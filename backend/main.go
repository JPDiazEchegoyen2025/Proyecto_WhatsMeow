
package main

import (
   "github.com/gin-gonic/gin"
   "context"
   "go.mau.fi/whatsmeow"
   "go.mau.fi/whatsmeow/store/sqlstore"
   waLog "go.mau.fi/whatsmeow/util/log"
   "time"
)

var waClient *whatsmeow.Client
var waContainer *sqlstore.Container

func main() {
   r := gin.Default()

   // Inicializar el container global (ajusta la ruta de la base de datos si es necesario)
   var err error
   waContainer, err = sqlstore.New(
	   context.Background(),
	   "sqlite3",
	   "whatsmeow.db?_foreign_keys=on",
	   waLog.Stdout("SQLStore", "DEBUG", true),
   )
   if err != nil || waContainer == nil {
	   panic("No se pudo inicializar el store de WhatsMeow: " + err.Error())
   }

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
	   waClient.Disconnect()
	   waClient = nil
	   c.JSON(200, gin.H{"message": "Sesión cerrada correctamente"})
	})

	// Endpoint para verificar sesión activa
	r.GET("/session", func(c *gin.Context) {
		if waClient != nil && waClient.IsLoggedIn() && waClient.Store != nil && waClient.Store.ID != nil {
			c.JSON(200, gin.H{
			   "active": true,
			   "user": waClient.Store.ID.String(),
			   "phone": waClient.Store.ID.String(), // Ajusta si tienes el teléfono
			   "status": "Activa",
			})
		} else {
			c.JSON(200, gin.H{
			   "active": false,
			   "user": "",
			   "phone": "",
			   "status": "",
			})
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
   // Si hay un cliente existente, cerrar sesión, desconectarlo y limpiarlo
   if waClient != nil {
	   _ = waClient.Logout(context.Background()) // Ignorar error, solo limpiar
	   waClient.Disconnect()
	   waClient = nil
   }
   // Crear nuevo device store y cliente
   deviceStore := waContainer.NewDevice()
   client := whatsmeow.NewClient(deviceStore, waLog.Stdout("Client", "DEBUG", true))
   if client.IsConnected() || client.IsLoggedIn() {
	   println("[LOGIN] Cliente aún conectado, espere un momento")
	   c.JSON(500, gin.H{"error": "cliente aún conectado, espere un momento"})
	   return
   }
   waClient = client
   qrChan, err := waClient.GetQRChannel(context.Background())
   if err != nil {
	   println("[LOGIN] Error al obtener canal QR:", err.Error())
	   c.JSON(500, gin.H{"error": err.Error()})
	   return
   }
   err = waClient.Connect()
   if err != nil {
	   println("[LOGIN] Error al conectar cliente:", err.Error())
	   c.JSON(500, gin.H{"error": err.Error()})
	   return
   }
   select {
   case qr := <-qrChan:
	   println("[LOGIN] Evento QR recibido:", qr.Event)
	   if qr.Event == "code" {
		   println("[LOGIN] Código QR generado:", qr.Code)
		   c.JSON(200, gin.H{"qr": qr.Code})
		   return
	   } else {
		   println("[LOGIN] Evento QR inesperado:", qr.Event)
		   c.JSON(500, gin.H{"error": "evento QR inesperado: " + qr.Event})
		   return
	   }
   case <-time.After(30 * time.Second):
	   println("[LOGIN] Timeout esperando QR")
	   c.JSON(500, gin.H{"error": "timeout esperando QR"})
	   return
   }
   println("[LOGIN] No se pudo generar QR")
   c.JSON(500, gin.H{"error": "no se pudo generar QR"})
	})

	r.Run() // por defecto en :8080
}
