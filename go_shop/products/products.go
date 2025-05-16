package products

import (
	"fmt"
	"go_shop/database"
)

type IndexProducts struct {
	Product_id       int     `gorm:"primarykey"`
	English_name     string  `gorm:"size:100"`
	Chinese_name     string  `gorm:"size:100"`
	Series_name      string  `gorm:"size:100"`
	Price            float64 `gorm:""`
	Ordered_num      int     `gorm:""`
	Is_special_offer bool    `gorm:""`
	Special_price    float64 `gorm:""`
	Image            string  `gorm:"size:100"`
	Gender           string  `gorm:"size:10"`
	Grade            string  `gorm:"size:2"`
	Style            string  `gorm:"size:30"`
	Quantity         int     `gorm:"default:1"` // 默认值为 1
	Brand_id         int     `gorm:""`
	Brand_story      string  `gorm:"size:50"`
}

func GetIndexProducts() []IndexProducts {
	var indexProducts []IndexProducts

	database.Gdb.Raw("select p.Product_id,p.Price,p.Ordered_num,b.English_name,b.Chinese_name,p.Is_special_offer,s.Series_name,p.Special_price,p.Image,p.Gender,b.Grade from products p left join series s on p.Series_id = s.Series_id left join brands b on s.Brand_id = b.Brand_id;").Scan(&indexProducts)

	return indexProducts
}

// 获取所有男士商品
func GetGendeProducts(gender string, page int, pageCount int) []IndexProducts {
	var products []IndexProducts

	start := (page - 1) * pageCount
	fmt.Println(page, start, pageCount)
	database.Gdb.Raw(fmt.Sprintf("select p.Product_id,p.Image,p.Price,p.Is_special_offer,p.Special_price,p.Gender,p.style,p.Ordered_num,s.series_name,b.chinese_name,b.english_name,b.grade from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id where p.gender = '%s' order by p.price desc limit %d,%d", gender, start, pageCount)).Scan(&products)

	return products
}

// 获取商品详情
type Products struct {
	Product_id       int     `gorm:"primary"`
	Series_id        int     `gorm:""`
	Series_name      string  `gorm:"size:100"`
	Case_material    string  `gorm:"size:10"`
	Case_back        string  `gorm:"size:10"`
	Strap_color      string  `gorm:"size:10"`
	Functions        string  `gorm:""`
	Style            string  `gorm:"size:10"`
	Size             string  `gorm:"size:20"`
	Watch_glass      string  `gorm:"size:20"`
	Watch_buckle     string  `gorm:"size:10"`
	Launch_year      string  `gorm:"size:10"`
	Movement         string  `gorm:"size:10"`
	Thickness        float64 `gorm:""`
	Dial             string  `gorm:"size:10"`
	Water_resistance string  `gorm:"size:50"`
	Price            float64 `gorm:""`
	Is_new           bool    `gorm:""`
	Is_special_offer bool    `gorm:""`
	Special_price    float64 `gorm:""`
	Image            string  `gorm:"size:100"`
	Gender           string  `gorm:"size:10"`
	Ordered_num      int     `gorm:""`

	// 品牌信息
	Brand_id     int    `gorm:""`
	Chinese_name string `gorm:"size:100"`
	English_name string `gorm:"size:100"`
	Grade        string `gorm:"size:2"`
	Brand_story  string `gorm:"size:50"`
}

func GetDetail(id string) Products {
	var products Products

	database.Gdb.Raw("select p.*,s.series_name,b.Brand_id,b.Chinese_name,b.English_name,b.grade from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id where p.product_id = " + id).Scan(&products)

	return products
}

// 获取指定商品的热销商品（降序），前3条
func GetBrandProducts(brand_id int) []Products {
	var products []Products

	database.Gdb.Raw("select p.product_id,p.image,p.price,p.is_special_offer,p.special_price,p.gender,p.style,p.ordered_num,s.series_name,b.chinese_name,b.grade from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id where b.brand_id = ? order by p.ordered_num desc limit 3", brand_id).Scan(&products)

	return products
}

// 限时抢购分页数据(暂时输出全部商品)
func GetTimeProducts(page int, pageCount int) []IndexProducts {
	var products []IndexProducts

	start := (page - 1) * pageCount
	fmt.Println(page, start, pageCount)
	database.Gdb.Raw(fmt.Sprintf("select b.brand_story,p.product_id,p.image,p.price,p.is_special_offer,p.special_price,p.gender,p.style,p.ordered_num,s.series_name,b.chinese_name,b.grade from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id order by p.price asc limit %d,%d", start, pageCount)).Scan(&products)

	return products
}

type Count struct {
	Count int `gorm:""`
}

// 获取男士总记录数
func GetGenderCount(gender string) int {
	var count Count

	database.Gdb.Raw("select count(*) count from products where gender = '" + gender + "'").Find(&count)

	return count.Count
}

// 获取限时抢购总记录数（价格升序前12个商品）
func GetTimeCount() int {
	var count Count

	database.Gdb.Raw("select count(*) count from products p order by p.price asc limit 12").Find(&count)

	return count.Count
}

// 获取特价总记录数
func GetCount() int {
	var count Count

	database.Gdb.Raw("select count(*) count from products p where p.Is_special_offer = true").Find(&count)

	return count.Count
}

// 获取所有特价分页商品
func GetProducts(page int, pageCount int) []IndexProducts {
	var products []IndexProducts

	start := (page - 1) * pageCount
	fmt.Println(page, start, pageCount)
	database.Gdb.Raw(fmt.Sprintf("select p.Product_id,p.Image,p.Price,p.Is_special_offer,p.Special_price,p.Gender,p.style,p.Ordered_num,s.series_name,b.chinese_name,b.english_name,b.grade from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id order by p.price desc limit %d,%d", start, pageCount)).Scan(&products)

	return products
}

// 获取所有男士新品商品
func GetNewProducts(gender string) []IndexProducts {
	var products []IndexProducts

	database.Gdb.Raw("select p.product_id,p.image,p.price,p.is_special_offer,p.special_price,p.gender,p.style,p.ordered_num,s.series_name,b.chinese_name,b.grade from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id where p.gender = '" + gender + "' order by p.price desc").Scan(&products)

	return products
}

// 获取购物车信息
func GetShoppingCarProducts(user_id int) []IndexProducts {
	var shoppingCarProducts []IndexProducts

	database.Gdb.Raw("select c.Quantity, p.Product_id,p.Price,p.Ordered_num,b.English_name,b.Chinese_name,p.Is_special_offer,s.Series_name,p.Special_price,p.Image,p.Gender,b.Grade from carts c left join products p on c.product_id = p.Product_id left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id left join users u on u.user_id = c.user_id where c.user_id = ?", user_id).Scan(&shoppingCarProducts)

	return shoppingCarProducts
}

type Carts struct {
	Cart_id    int `gorm:"primaryKey;autoIncrement"` // 主键，自增
	User_id    int `gorm:"not null"`                 // 不允许为空
	Product_id int `gorm:"not null"`                 // 不允许为空
	Quantity   int `gorm:"default:1"`                // 默认值为 1
}

// 获取产品价格
func GetProductPrice(productId int) (float64, error) {
	var product Products
	// 从数据库中查询产品
	result := database.Gdb.First(&product, productId)
	if result.Error != nil {
		return 0, result.Error // 返回错误
	}
	return product.Price, nil // 返回产品价格
}

func GetCarTotalPrice(user_id int) int {
	var carts []Carts
	var totalPrice float64

	// 获取所有购物车项
	database.Gdb.Where("User_id = ?", user_id).Find(&carts)

	// 计算总价格
	for _, cart := range carts {
		productPrice, err := GetProductPrice(cart.Product_id) // 获取产品价格
		if err != nil {
			// 处理错误，例如记录日志或返回错误
			continue // 跳过此项
		}
		totalPrice += productPrice * float64(cart.Quantity) // 计算总价格
	}

	return int(totalPrice) // 返回总价格
}
