package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"test_work/database"
	"test_work/handlers"
	"time"
)

// Отправная точка для запуска проекта
func main() {
	app := fiber.New(fiber.Config{Prefork: true})

	if err := database.InitDB("postgres://local:password@localhost:5432/tasks_db?sslmode=disable"); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app.Post("/tasks", handlers.CreateTask)
	app.Get("/tasks", handlers.GetTasks)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)

	// Запуск сервера в горутине
	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Создать канал для пинга сигнала
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Блокировка до получения сигнала
	<-quit
	log.Println("Shutting down server...")

	// Контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Попытка нормально завершить работу сервера
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped")
}
