package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"test_work/database"
	"test_work/models"
)

// CreateTask создание таска
func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err := database.DB.Exec(context.Background(),
		"INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)",
		task.Title, task.Description, task.Status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(task)
}

// GetTasks получние всех тасков
func GetTasks(c *fiber.Ctx) error {
	rows, err := database.DB.Query(context.Background(),
		"SELECT * FROM tasks")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		tasks = append(tasks, task)
	}

	//return c.Status(http.StatusOK).JSON(tasks) можно использовать для большего пояснения со статусом
	return c.JSON(tasks)
}

// UpdateTask обновление таска в базе данных
func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task := new(models.Task)
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result, err := database.DB.Exec(context.Background(),
		"UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = now() WHERE id = $4",
		task.Title, task.Description, task.Status, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if result.RowsAffected() == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Todo updated successfully",
	})
}

// DeleteTask удаление таска с базы даных
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := database.DB.Exec(context.Background(),
		"DELETE FROM tasks WHERE id = $1",
		id,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if result.RowsAffected() == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Todo deleted successfully",
	})
}
