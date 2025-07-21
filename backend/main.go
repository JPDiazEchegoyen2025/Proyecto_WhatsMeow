package main

import (
	"github.com/gin-gonic/gin"
	"context"
)

func main() {
	r := gin.Default()

	// Middleware CORS (debe ir antes de las rutas)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/logout", func(c *gin.Context) {
		if waClient == nil {
			c.JSON(400, gin.H{"error": "No hay sesión activa"})
			return
		}
		err := waClient.Logout(context.Background())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Sesión cerrada correctamente"})
	})

	// Endpoint para verificar sesión activa
	r.GET("/session", func(c *gin.Context) {
		if waClient != nil && waClient.IsLoggedIn() {
			// Puedes personalizar los datos según tu modelo
			c.JSON(200, gin.H{
				"active": true,
				"user": waClient.Store.ID.String(),
				"phone": waClient.Store.ID.String(), // Ajusta si tienes el teléfono
				"status": "Activa",
			})
		} else {
			c.JSON(200, gin.H{"active": false})
		}
	})

	// Middleware CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/login", func(c *gin.Context) {
	qr, err := StartLogin()
	if err != nil {
		// Imprime el error en la consola para depuración
		println("Error en /login:", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"qr": qr})
	})

	r.Run() // por defecto en :8080
}
