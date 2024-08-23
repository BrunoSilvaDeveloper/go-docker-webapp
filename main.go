package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Olá, Mundo!")
}

func status(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"status": "ok"}
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.HandleFunc("/status", status)

    port := ":8080"
    env := os.Getenv("APP_ENV")
    if env == "" {
        env = "development"
    }
    fmt.Printf("Iniciando servidor no ambiente %s na porta %s\n", env, port)

    if err := http.ListenAndServe(port, nil); err != nil {
        fmt.Println("Erro ao iniciar o servidor:", err)
    }
}
