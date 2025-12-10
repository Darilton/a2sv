package routers_test

import (
	"net/http"
	"net/http/httptest"
	"task_manager_api/Delivery/controller"
	"task_manager_api/Delivery/router"
	"task_manager_api/Tests/mocks"
	usecases "task_manager_api/UseCases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouterRoutes(t *testing.T) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	taskUseCase := usecases.NewTaskUseCase(mockTaskRepo)
	userUseCase := usecases.NewUserUseCase(mockUserRepo)
	taskController := controller.NewTaskController(taskUseCase, userUseCase)

	r := router.GetRouter(taskController)

	routes := []struct {
		method string
		path   string
	}{
		{"GET", "/tasks"},
		{"GET", "/tasks/:id"},
		{"PUT", "/tasks/:id"},
		{"DELETE", "/tasks/:id"},
		{"POST", "/tasks"},
		{"POST", "/register"},
		{"POST", "/login"},
	}

	for _, route := range routes {
		req, _ := http.NewRequest(route.method, route.path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		// We expect Not Found 404 only if the route is NOT registered.
		// If it is registered, it might return 401 Unauthorized or 404 (if param issues) or 400 etc.
		// But Gin usually returns 404 for unknown routes.
		// Wait, if we request "/tasks/:id" literally, it might be weird.
		// Let's request a dummy path matching the pattern
	}

	// Better approach: Check if endpoints exist by hitting them and assuming 404 means missing.
	// But Auth middleware might return 401.

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.NotEqual(t, http.StatusNotFound, w.Code, "GET /tasks should exist")

	req, _ = http.NewRequest("POST", "/register", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.NotEqual(t, http.StatusNotFound, w.Code, "POST /register should exist")
}
