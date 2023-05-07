package api

import (
	"csbackend/api/auth"
	"encoding/gob"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func AttachRoutes(r *gin.Engine) error {
	g := r.Group("/")

	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	g.Use(sessions.Sessions("auth-session", store))

	authEndpoints, err := auth.New()
	if err != nil {
		return err
	}

	g.GET("/auth/login", authEndpoints.Login)
	g.GET("/auth/callback", authEndpoints.Callback)
	g.GET("/auth/logout", authEndpoints.Logout)

	user := User{}
	g.GET("/user", user.GET)

	return nil
}
