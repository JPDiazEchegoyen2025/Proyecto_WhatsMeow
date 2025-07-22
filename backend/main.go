package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	_ "github.com/mattn/go-sqlite3"
)


// NewWhatsAppClientStable crea un cliente WhatsMeow con configuración SQLite optimizada y logging adecuado.
// Si existe un device en la base de datos, lo reutiliza (restaura sesión). Si no, permite crear uno nuevo.
func NewWhatsAppClientStable(container *sqlstore.Container) (*whatsmeow.Client, error) {
	// Mostrar información general del cliente
	clientLogger := waLog.Stdout("Client", "INFO", true)
	var deviceStore *store.Device
	devices, err := container.GetAllDevices(context.Background())
	if err != nil {
		return nil, err
	}
	if len(devices) > 0 {
		// Reutilizar el primer device encontrado (restaurar sesión)
		deviceStore = devices[0]
	} else {
		// No hay sesión previa, crear nuevo device
		deviceStore = container.NewDevice()
	}
	client := whatsmeow.NewClient(deviceStore, clientLogger)
	client.AddEventHandler(func(evt interface{}) {
		if _, ok := evt.(*events.Connected); ok {
			log.Println("🔗 WhatsApp conectado")
			log.Println("ℹ️  Los warnings de 'FOREIGN KEY constraint failed' durante sync son normales y no afectan la funcionalidad")
		}
	})
	return client, nil
}



var waClient *whatsmeow.Client
var waContainer *sqlstore.Container

func main() {
	// Inicializar el contenedor global (ajusta la ruta de la base de datos si es necesario)
	var err error
	dsn := "whatsmeow.db?_foreign_keys=on&_journal_mode=WAL&_synchronous=NORMAL&_busy_timeout=30000&_cache_size=20000&_temp_store=memory"
	waContainer, err = sqlstore.New(
		context.Background(),
		"sqlite3",
		dsn,
		waLog.Stdout("SQLStore", "ERROR", true),
	)
	if err != nil || waContainer == nil {
		log.Fatal("No se pudo inicializar el store de WhatsMeow: ", err)
	}

	// Intentar restaurar sesión existente al iniciar el backend
	waClient, err = NewWhatsAppClientStable(waContainer)
	if err != nil {
		log.Println("No se pudo restaurar sesión WhatsApp al iniciar:", err)
		waClient = nil
	}
	// Si hay sesión restaurada, reconectar automáticamente
	if waClient != nil && !waClient.IsConnected() {
		go func() {
			err := waClient.Connect()
			if err != nil {
				log.Println("Error al reconectar sesión WhatsApp:", err)
			}
		}()
	}

	r := gin.Default()

	// Middleware CORS (antes de las rutas)
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

	// Endpoint: login (genera QR solo si NO hay sesión activa)
	r.GET("/login", func(c *gin.Context) {
		if waClient != nil && waClient.IsLoggedIn() {
			c.JSON(400, gin.H{"error": "Ya hay una sesión activa. Cierra sesión antes de volver a vincular."})
			return
		}
		// Si hay un cliente previo, limpiar
		if waClient != nil {
			_ = waClient.Logout(context.Background())
			waClient.Disconnect()
			waClient = nil
		}
		client, err := NewWhatsAppClientStable(waContainer)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error inicializando cliente: " + err.Error()})
			return
		}
		waClient = client
		qrChan, err := waClient.GetQRChannel(context.Background())
		if err != nil {
			c.JSON(500, gin.H{"error": "Error obteniendo canal QR: " + err.Error()})
			return
		}
		err = waClient.Connect()
		if err != nil {
			c.JSON(500, gin.H{"error": "Error conectando cliente: " + err.Error()})
			return
		}
		select {
		case qr := <-qrChan:
			if qr.Event == "code" {
				c.JSON(200, gin.H{"qr": qr.Code})
				return
			} else {
				c.JSON(500, gin.H{"error": "Evento QR inesperado: " + qr.Event})
				return
			}
		case <-time.After(30 * time.Second):
			c.JSON(500, gin.H{"error": "Timeout esperando QR"})
			return
		}
	})

	// Endpoint: logout
	r.POST("/logout", func(c *gin.Context) {
		if waClient == nil || !waClient.IsLoggedIn() {
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

	// Endpoint: verificar sesión activa
	r.GET("/session", func(c *gin.Context) {
		if waClient != nil && waClient.IsLoggedIn() && waClient.Store != nil && waClient.Store.ID != nil {
			c.JSON(200, gin.H{
				"active": true,
				"user": waClient.Store.ID.String(),
				"phone": waClient.Store.ID.String(),
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

	// Endpoint: ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	log.Println("🚀 Backend WhatsApp listo en :8080")
	r.Run() // por defecto en :8080
}
