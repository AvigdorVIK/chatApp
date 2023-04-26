package auth

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
            c.Abort()
            return
        }

        // Проверяем, что токен действительный и получаем информацию о пользователе
        user, err := parseToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
            c.Abort()
            return
        }

        // Добавляем информацию о пользователе в контекст
        c.Set("user", user)

        c.Next()
    }
}
r := gin.Default()

// Регистрация и вход
r.POST("/register", auth.Register)
r.POST("/login", auth.Login)

// Защищенные маршруты
private := r.Group("/private")
private.Use(auth.AuthMiddleware())
{
    private.GET("/profile", func(c *gin.Context) {
        // Извлекаем информацию о пользователе из контекста
        user, exists := c.Get("user")
        if !exists {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "User information not found"})
            return
        }

        // Используем информацию о пользователе для получения профиля
        // и возвращаем его в качестве ответа
        profile := getProfile(user.(auth.User))
        c.JSON(http.StatusOK, profile)
    })

    // другие защищенные маршруты
}

func getProfile(user auth.User) map[string]interface{} {
    // Возвращаем профиль пользователя в виде карты ключ-значение
    return map[string]interface{}{
        "email": user.Email,
        // Другие поля профиля
    }
}
