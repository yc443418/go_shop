package tools

import (
	"math/rand/v2"
	"strings"

	"github.com/gin-gonic/gin"
)

// 负责验证码功能
func GetCaptcha(ctx *gin.Context) string {
	var captcha string

	for i := 0; i < 5; i++ {
		key := rand.IntN(3)

		var ascii int
		if key == 0 {
			ascii = rand.IntN(9) + 49
		} else if key == 1 {
			ascii = rand.IntN(26) + 65
		} else {
			ascii = rand.IntN(26) + 97
		}

		captcha += string(ascii)
	}

	// 验证验证码到session
	SetSession(ctx, "captcha", captcha)
	return captcha
}

func CheckCaptcha(ctx *gin.Context, captcha string) bool {
	sessionCaptcha, _ := GetSession(ctx, "captcha").(string)

	return strings.EqualFold(sessionCaptcha, captcha)
}
