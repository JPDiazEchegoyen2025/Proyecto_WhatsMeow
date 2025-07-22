
package main

import (
    "context"
    "database/sql"
    "log"
    "os"

    "go.mau.fi/whatsmeow"
    "go.mau.fi/whatsmeow/store/sqlstore"
    "go.mau.fi/whatsmeow/types/events"
    waLog "go.mau.fi/whatsmeow/util/log"

    _ "github.com/mattn/go-sqlite3"
)

// Gestor principal que maneja las dos fases
type WhatsAppManager struct {
    client         *whatsmeow.Client
    isInitialSetup bool
}

// Constructor que detecta automáticamente el modo
func NewWhatsAppManager() (*WhatsAppManager, error) {
    dbPath := "whatsmeow.db"
    isFirst := isFirstRun(dbPath)
    var client *whatsmeow.Client
    var err error
    if isFirst {
        log.Println("[WhatsAppManager] Primera ejecución: inicializando sin foreign keys...")
        client, err = initFirstTime(dbPath)
        if err != nil {
            return nil, err
        }
        return &WhatsAppManager{client: client, isInitialSetup: true}, nil
    } else {
        log.Println("[WhatsAppManager] DB existente: usando configuración con foreign keys...")
        client, err = initNormal(dbPath)
        if err != nil {
            return nil, err
        }
        return &WhatsAppManager{client: client, isInitialSetup: false}, nil
    }
}

// Inicialización para primera vez (sin FK)
func initFirstTime(dbPath string) (*whatsmeow.Client, error) {
    dsn := dbPath + "?_journal_mode=WAL&_synchronous=NORMAL"
    container, err := sqlstore.New(
        context.Background(),
        "sqlite3",
        dsn,
        waLog.Stdout("SQLStore", "DEBUG", true),
    )
    if err != nil {
        return nil, err
    }
    deviceStore := container.NewDevice()
    client := whatsmeow.NewClient(deviceStore, waLog.Stdout("Client", "INFO", true))
    // Handler para saber cuándo termina el sync inicial
    client.AddEventHandler(func(evt interface{}) {
        if e, ok := evt.(*events.HistorySyncFinished); ok {
            log.Println("⚠️  SYNC INICIAL COMPLETADO")
            log.Println("🔄 REINICIA LA APLICACIÓN para habilitar foreign keys")
        }
    })
    return client, nil
}

// Inicialización normal (con FK)
func initNormal(dbPath string) (*whatsmeow.Client, error) {
    dsn := dbPath + "?_foreign_keys=on&_journal_mode=WAL&_synchronous=NORMAL&_busy_timeout=30000"
    container, err := sqlstore.New(
        context.Background(),
        "sqlite3",
        dsn,
        waLog.Stdout("SQLStore", "DEBUG", true),
    )
    if err != nil {
        return nil, err
    }
    deviceStore := container.NewDevice()
    client := whatsmeow.NewClient(deviceStore, waLog.Stdout("Client", "INFO", true))
    return client, nil
}

// Función auxiliar para detectar primera ejecución
func isFirstRun(dbPath string) bool {
    // Si no existe el archivo, es primera vez
    if _, err := os.Stat(dbPath); os.IsNotExist(err) {
        return true
    }
    // Si existe, verificar si hay dispositivos registrados
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        log.Printf("[WhatsAppManager] Error abriendo DB para detección: %v", err)
        // Si la DB está corrupta o no se puede abrir, tratamos como primera vez
        return true
    }
    defer db.Close()
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM whatsmeow_device WHERE registration_id IS NOT NULL").Scan(&count)
    if err != nil {
        log.Printf("[WhatsAppManager] Error consultando dispositivos: %v", err)
        // Si la tabla no existe o hay error, tratamos como primera vez
        return true
    }
    return count == 0
}

// Getters útiles
func (w *WhatsAppManager) GetClient() *whatsmeow.Client {
    return w.client
}

func (w *WhatsAppManager) IsInitialSetup() bool {
    return w.isInitialSetup
}
