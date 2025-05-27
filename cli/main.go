package main

import (
	"5M/lib/config"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет, мир!")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Сервис работает")
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/health", healthCheckHandler)
	return mux
}

func main() {
	// Загрузка конфигурации
	configPath := filepath.Join(".", "config.yaml")
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Ошибка при загрузке конфигурации: %v", err)
	}

	// Инициализация маршрутов
	mux := setupRoutes()

	// Формирование адреса сервера
	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	// Вывод информации о запуске сервера
	log.Printf("Запуск сервера на %s", serverAddr)

	// Запуск HTTP-сервера
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
