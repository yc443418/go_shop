package tools

import (
	"encoding/gob"
	"go_shop/users"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte("recomend-32bytes-at-least"))

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   0,
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	gob.Register(users.Users{})
}

func SetSession(ctx *gin.Context, key string, val interface{}) error {
	session, err := store.Get(ctx.Request, "SHOPID")

	if err != nil {
		return err
	}
	session.Values[key] = val

	return session.Save(ctx.Request, ctx.Writer)
}

// 查询数据
func GetSession(ctx *gin.Context, key string) interface{} {
	session, err := store.Get(ctx.Request, "SHOPID")

	if err != nil {
		return err
	}

	return session.Values[key]
}

// 删除session值
func DelSession(ctx *gin.Context, key string) error {
	session, err := store.Get(ctx.Request, "SHOPID")

	if err != nil {
		return err
	}

	delete(session.Values, key)

	return session.Save(ctx.Request, ctx.Writer)
}
