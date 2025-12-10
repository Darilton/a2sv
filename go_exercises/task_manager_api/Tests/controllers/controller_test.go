package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"task_manager_api/Delivery/controller"
	"task_manager_api/Delivery/router"
	domain "task_manager_api/Domain"
	"task_manager_api/Infrastructure"
	"task_manager_api/Tests/mocks"
	usecases "task_manager_api/UseCases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter() (*mocks.MockTaskRepository, *mocks.MockUserRepository, http.Handler) {
	mockTaskRepo := new(mocks.MockTaskRepository)
	mockUserRepo := new(mocks.MockUserRepository)

	taskUseCase := usecases.NewTaskUseCase(mockTaskRepo)
	userUseCase := usecases.NewUserUseCase(mockUserRepo)

	taskController := controller.NewTaskController(taskUseCase, userUseCase)
	r := router.GetRouter(taskController)
	return mockTaskRepo, mockUserRepo, r
}

func TestRegisterUser_Success(t *testing.T) {
	_, mockUserRepo, r := setupRouter()

	user := domain.User{UserName: "test", Password: "password"}
	mockUserRepo.On("AddUser", mock.Anything).Return(nil)

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginUser_Success(t *testing.T) {
	_, mockUserRepo, r := setupRouter()

	password := "password"
	hashed, _ := Infrastructure.HashPassword(password)
	user := domain.User{UserName: "test", Password: hashed, UserRole: "regular"}

	mockUserRepo.On("GetUser", "test").Return(user, nil)

	loginData := domain.User{UserName: "test", Password: password}
	jsonValue, _ := json.Marshal(loginData)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotEmpty(t, response["token"])
}

func TestGetTasks_Authorized(t *testing.T) {
	mockTaskRepo, _, r := setupRouter()

	tasks := []domain.Task{{Id: "1", Title: "Task 1"}}
	mockTaskRepo.On("GetTasks").Return(tasks)

	token, _ := Infrastructure.GenerateJWT("test", "regular")
	req, _ := http.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTasks_Unauthorized(t *testing.T) {
	_, _, r := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCreateTask_Admin(t *testing.T) {
	mockTaskRepo, _, r := setupRouter()

	task := domain.Task{Id: "1", Title: "New Task"}
	mockTaskRepo.On("AddTask", mock.Anything).Return(nil)

	token, _ := Infrastructure.GenerateJWT("admin", "admin")
	jsonValue, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateTask_NotAdmin(t *testing.T) {
	_, _, r := setupRouter()

	task := domain.Task{Id: "1", Title: "New Task"}
	token, _ := Infrastructure.GenerateJWT("user", "regular")
	jsonValue, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestGetTask_NotFound(t *testing.T) {
	mockTaskRepo, _, r := setupRouter()

	mockTaskRepo.On("GetTask", "999").Return(domain.Task{}, errors.New("not found"))

	token, _ := Infrastructure.GenerateJWT("test", "regular")
	req, _ := http.NewRequest("GET", "/tasks/999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
