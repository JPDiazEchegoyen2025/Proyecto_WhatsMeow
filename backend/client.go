

package main

import (
    "context"
    "errors"
    "sync"
    "time"

    _ "github.com/mattn/go-sqlite3"

    "go.mau.fi/whatsmeow/store/sqlstore"
    waLog "go.mau.fi/whatsmeow/util/log"
    "go.mau.fi/whatsmeow"
)

var (
    waClient     *whatsmeow.Client
    waClientOnce sync.Once
    waClientErr  error
)

// StartLogin inicializa el cliente y devuelve el código QR en base64
func StartLogin() (string, error) {
    waClientOnce.Do(func() {
        dbPath := "file:whatsmeow.db?_foreign_keys=on"
        container, err := sqlstore.New(context.Background(), "sqlite3", dbPath, waLog.Stdout("SQLStore", "INFO", true))
        if err != nil {
            waClientErr = err
            return
        }
        device, err := container.GetFirstDevice(context.Background())
        if err != nil {
            waClientErr = err
            return
        }
        waClient = whatsmeow.NewClient(device, waLog.Stdout("WA", "INFO", true))
    })
    if waClientErr != nil {
        return "", waClientErr
    }

    if waClient.IsLoggedIn() {
        return "", errors.New("ya está logueado")
    }

    qrChan, _ := waClient.GetQRChannel(context.Background())
    go func() {
        _ = waClient.Connect()
    }()

    select {
    case evt := <-qrChan:
        if evt.Event == "code" {
            return evt.Code, nil
        } else if evt.Event == "timeout" {
            return "", errors.New("timeout esperando QR")
        } else if evt.Event == "success" {
            return "", errors.New("ya logueado")
        }
    case <-time.After(60 * time.Second):
        return "", errors.New("timeout esperando QR")
    }
    return "", errors.New("no se pudo obtener el QR")
}
