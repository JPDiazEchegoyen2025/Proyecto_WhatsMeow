package main

import (
	   "context"
	   "net/http"
	   "log"
	   "github.com/gin-gonic/gin"
)

// Estructura de respuesta para el frontend
 type ContactResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	LastMessage     string `json:"lastMessage"`
	UnreadCount     int    `json:"unreadCount"`
	LastMessageTime string `json:"lastMessageTime"`
}

func ContactsHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	   if waClient == nil || !waClient.IsConnected() {
			   c.JSON(http.StatusServiceUnavailable, gin.H{"error": "WhatsApp client not connected"})
			   return
	   }

	   contactsMap, err := waClient.Store.Contacts.GetAllContacts(context.Background())
	   if err != nil {
			   log.Printf("Error obteniendo contactos: %v", err)
			   c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo contactos"})
			   return
	   }

	   var contacts []ContactResponse
	   for jid, contact := range contactsMap {
			   if contact.FullName == "" {
					   continue // Solo contactos con full_name
			   }
			   name := contact.FullName
			   contacts = append(contacts, ContactResponse{
					   ID:              jid.String(),
					   Name:            name,
					   Avatar:          "",  // El frontend generará iniciales
					   LastMessage:     "",  // Se actualizará con eventos
					   UnreadCount:     0,   // Se actualizará con eventos
					   LastMessageTime: "",  // Se actualizará con eventos
			   })
	   }
	   c.JSON(http.StatusOK, contacts)
}
