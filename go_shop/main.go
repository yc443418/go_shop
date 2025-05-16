package main

import (
	"fmt"
	"go_shop/brands"
	"go_shop/database"
	"go_shop/products"
	"go_shop/shppingcar"
	"go_shop/tools"
	"go_shop/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	r := gin.Default()

	// 自动迁移
	database.Gdb.AutoMigrate(&shppingcar.Carts{})

	r.StaticFS("static", http.Dir("./static"))

	r.LoadHTMLGlob("./tpl/*")

	// 全局中间件
	var user users.Users
	r.Use(func(ctx *gin.Context) {
		// 验证用户是否登录
		val := tools.GetSession(ctx, "user")
		user, _ = val.(users.Users)
		// 流转
		ctx.Next()
	})

	r.GET("/", func(ctx *gin.Context) {
		indexProducts := products.GetIndexProducts()
		ctx.HTML(200, "index.html", gin.H{
			"indexProducts": indexProducts,
			"user":          user,
		})
	})

	// 2.特价功能
	r.GET("/special", func(ctx *gin.Context) {
		page := ctx.DefaultQuery("page", "1")
		pageInt, _ := strconv.Atoi(page)
		// 获取分页
		pageInfo := tools.GetPageCount()
		specicalProducts := products.GetProducts(pageInt, pageInfo)
		ctx.HTML(200, "tejia.html", gin.H{
			"specicalProducts": specicalProducts,
			"user":             user,
			"router":           "special",
			"page":             tools.GetPage(products.GetCount(), pageInt),
		})
	})

	// 限时抢购
	r.GET("/time", func(ctx *gin.Context) {
		page := ctx.DefaultQuery("page", "1")
		pageInt, _ := strconv.Atoi(page)
		// 获取分页
		pageInfo := tools.GetPageCount()
		timeProducts := products.GetTimeProducts(pageInt, pageInfo)
		ctx.HTML(200, "time.html", gin.H{
			"timeProducts": timeProducts,
			"user":         user,
			"router":       "time",
			"page":         tools.GetPage(products.GetCount(), pageInt),
		})
	})

	// 男士腕表
	r.GET("/male", func(ctx *gin.Context) {
		page := ctx.DefaultQuery("page", "1")
		pageInt, _ := strconv.Atoi(page)
		// 获取分页
		pageInfo := tools.GetPageCount()
		maleProducts := products.GetGendeProducts("男", pageInt, pageInfo)
		getBrands := brands.GetBrands()
		// 新品！！
		maleRecommend := products.GetNewProducts("男")
		ctx.HTML(200, "boys.html", gin.H{
			"maleProducts":  maleProducts,
			"getBrands":     getBrands,
			"maleRecommend": maleRecommend,
			"user":          user,
			"router":        "male",
			"page":          tools.GetPage(products.GetGenderCount("男"), pageInt),
		})
	})

	// 女士腕表
	r.GET("/woman", func(ctx *gin.Context) {
		page := ctx.DefaultQuery("page", "1")
		pageInt, _ := strconv.Atoi(page)
		pageInfo := tools.GetPageCount()
		womanProducts := products.GetGendeProducts("女", pageInt, pageInfo)
		womanRecommend := products.GetNewProducts("女")
		getBrands := brands.GetBrands()
		ctx.HTML(200, "girl.html", gin.H{
			"womanProducts":  womanProducts,
			"getBrands":      getBrands,
			"womanRecommend": womanRecommend,
			"user":           user,
			"router":         "woman",
			"page":           tools.GetPage(products.GetGenderCount("女"), pageInt),
		})
	})

	// 登录
	r.GET("/login", func(ctx *gin.Context) {
		if user.User_id != 0 {
			ctx.HTML(200, "redirect.html", gin.H{
				"mag":    "你已经登录过！",
				"target": "/",
			})
			return
		}
		// 获取上一个页面路由
		referer := ctx.Request.Header.Get("referer")

		ctx.HTML(200, "login.html", gin.H{
			"user":    user,
			"referer": referer, // 借助login.html将数据提取过去
		})
	})

	r.POST("/checkLogin", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		referer := ctx.PostForm("referer")

		user, err := users.CheckLogin(username, password)

		if err != nil {
			ctx.HTML(200, "redirect.html", gin.H{
				"msg":    err.Error(),
				"target": "/login",
			})
			return
		}

		// users存入session
		if err := tools.SetSession(ctx, "user", user); err != nil {
			// 处理 session 设置失败的情况
			ctx.HTML(200, "redirect.html", gin.H{
				"msg":    "存储会话信息失败",
				"target": "/",
			})
			return
		}

		ctx.HTML(200, "redirect.html", gin.H{
			"msg":    "登录成功！",
			"target": referer,
		})
	})

	// 退出登录状态
	r.GET("/quitLogin", func(ctx *gin.Context) {

		tools.DelSession(ctx, "user")

		ctx.HTML(200, "redirect.html", gin.H{
			"msg":    "退出成功！",
			"target": "/",
		})
	})

	//查看商品详情
	r.GET("/detail/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		detail := products.GetDetail(id)

		if detail.Brand_id == 0 {
			ctx.HTML(200, "redirect.html", gin.H{
				"msg":    "您访问的商品不存在",
				"target": "/",
				"user":   user,
			})
			return
		}

		hotProducts := products.GetBrandProducts(detail.Brand_id)

		ctx.HTML(200, "xiangxi.html", gin.H{
			"product":     detail,
			"hotProducts": hotProducts,
			"target":      "/",
			"user":        user,
		})
	})

	r.GET("/register", func(ctx *gin.Context) {
		ctx.HTML(200, "reg.html", gin.H{
			"captcha": tools.GetCaptcha(ctx),
		})
	})

	r.POST("/checkRegister", func(ctx *gin.Context) {
		var regInfo users.RegInfo

		if err := ctx.ShouldBindWith(&regInfo, binding.Form); err != nil {
			ctx.HTML(200, "redirect.html", gin.H{
				"msg":    err.Error(),
				"target": "/register",
			})
			return
		}

		if !tools.CheckCaptcha(ctx, regInfo.Captcha) {
			ctx.HTML(200, "redirect.html", gin.H{
				"msg":    "验证码错误",
				"target": "/register",
			})
			return
		}

		if err := users.CheckRegisterInfo(regInfo); err != nil {
			ctx.HTML(200, "redirect.html", gin.H{
				"msg":    err.Error(),
				"target": "/register",
			})
			return
		}

		ctx.HTML(200, "redirect.html", gin.H{
			"msg":    "注册成功",
			"target": "/login",
		})

	})

	// 获取验证码
	r.GET("/getCaptcha", func(ctx *gin.Context) {
		ctx.HTML(200, "reg.html", gin.H{
			"captcha": tools.GetCaptcha(ctx),
		})
	})

	// 购物车
	r.GET("/car", func(ctx *gin.Context) {
		shppingCarProducts := products.GetShoppingCarProducts(user.User_id)
		totalPrice := products.GetCarTotalPrice(user.User_id)
		ctx.HTML(200, "car.html", gin.H{
			"user":               user,
			"shppingCarProducts": shppingCarProducts,
			"totalPrice":         totalPrice,
		})
	})

	// 提交订单
	r.GET("/tijiao", func(ctx *gin.Context) {
		shppingCarProducts := products.GetShoppingCarProducts(user.User_id)
		totalPrice := products.GetCarTotalPrice(user.User_id)
		ctx.HTML(200, "tijiao.html", gin.H{
			"user":               user,
			"shppingCarProducts": shppingCarProducts,
			"totalPrice":         totalPrice,
		})
	})

	// 订单提交成功
	r.GET("/chenggong", func(ctx *gin.Context) {
		ctx.HTML(200, "chenggong.html", gin.H{})
	})

	// 商品加入购物车
	r.POST("/checkShoppingCar", func(ctx *gin.Context) {
		// 假设用户 ID 是固定的，实际应用中应从用户会话或 JWT 中获取
		userID := user.User_id                // 用户 ID
		productID := ctx.PostForm("goods_id") // 获取商品 ID
		// 从表单数据中获取数量和商品ID
		quantity := ctx.PostForm("goods_number")

		// 将商品数量转换为整数
		goodsNumber, err := strconv.Atoi(quantity)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的商品数量"})
			return
		}
		// 将商品 ID 转换为整数
		productIDInt, err := strconv.Atoi(productID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的商品 ID"})
			return
		}

		// 将商品加入购物车
		cartItem := shppingcar.Carts{
			User_id:    userID,
			Product_id: productIDInt,
			Quantity:   goodsNumber, // 这里是 int 类型
		}

		// 保存到数据库
		if err := database.Gdb.Create(&cartItem).Error; err != nil {
			fmt.Println("数据库错误:", err) // 打印错误信息
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "无法加入购物车"})
			return
		}

		// 返回响应
		ctx.HTML(200, "redirect.html", gin.H{
			"msg":    "商品已添加进购物车",
			"target": "/car",
		})
	})

	// 寄卖路由
	r.GET("/consignmentShop", func(ctx *gin.Context) {
		ctx.HTML(200, "jimai.html", gin.H{})
	})

	// 会员功能
	r.GET("/huiyuan", func(ctx *gin.Context) {
		shppingCarProducts := products.GetShoppingCarProducts(user.User_id)
		totalPrice := products.GetCarTotalPrice(user.User_id)
		ctx.HTML(200, "huiyuan.html", gin.H{
			"user":               user,
			"shppingCarProducts": shppingCarProducts,
			"totalPrice":         totalPrice,
		})
	})

	r.Run(":80")
}
