package api

import (
	"net/http"

	"github.com/Gory007/auth-service/internal/oauth"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	// Кнопка "Login with Google"
	r.GET("/auth/google/login", func(c *gin.Context) {
		url := oauth.GoogleOAuthConfig.AuthCodeURL(oauth.RandomState)
		c.Redirect(http.StatusTemporaryRedirect, url)
	})

	// Callback от Google
	r.GET("/auth/google/callback", func(c *gin.Context) {
		// Проверяем state (защита от CSRF)
		if c.Query("state") != oauth.RandomState {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid state"})
			return
		}

		// Получаем код от Google
		code := c.Query("code")
		token, err := oauth.GoogleOAuthConfig.Exchange(c, code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Здесь можно получить данные пользователя (email, имя и т.д.)
		c.JSON(http.StatusOK, gin.H{"token": token})
	})
}
