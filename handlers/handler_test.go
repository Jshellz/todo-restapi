package handlers

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"testing"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Exec(ctx context.Context, query string, args ...interface{}) (interface{}, error) {
	callArgs := m.Called(ctx, query, args)
	return callArgs.Get(0), callArgs.Error(1)
}

func TestCreateTask(t *testing.T) {
	app := fiber.New()
	mockDB := new(MockDB)

	app.Post("/tasks", CreateTask)

	tests := []struct {
		name           string
		body           string
		mockReturn     int64
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Successful Task Creation",
			body:           `{"title": "Test Task", "description": "Test Description", "status": "pending"}`,
			mockReturn:     1,
			mockError:      nil,
			expectedStatus: fiber.StatusCreated,
			expectedBody:   `{"title":"Test Task","description":"Test Description","status":"pending"}`,
		},
		{
			name:           "Database Error",
			body:           `{"title": "Test Task", "description": "Test Description", "status": "pending"}`,
			mockReturn:     0,
			mockError:      errors.New("database error"),
			expectedStatus: fiber.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB.On("Exec", mock.Anything, "INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)", mock.Anything).
				Return(tt.mockReturn, tt.mockError)

			req := fiber.Request{}
			req.SetRequestURI("/tasks")
			req.Header.SetMethod("POST")
			req.Header.Set("Content-Type", "application/json")
			req.SetBody([]byte(tt.body))

			resp, err := app.Test(&http.Request{})
			assert.NoError(t, err)
			defer resp.Body.Close().Error()

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expectedBody, string(body))
		})
	}
}
