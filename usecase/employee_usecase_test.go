package usecase

import (
	"Test_Mek/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockEmployeeRepo is a mock implementation of the IEmployee repository interface
type MockEmployeeRepo struct {
	mock.Mock
}

// FindAll mocks the FindAll method of the IEmployee repository interface
func (m *MockEmployeeRepo) FindAll(req *model.Employee) ([]*model.Employee, error) {
	args := m.Called(req)
	return args.Get(0).([]*model.Employee), args.Error(1)
}

// CreateOne mocks the CreateOne method of the IEmployee repository interface
func (m *MockEmployeeRepo) CreateOne(employeeData *model.Employee) error {
	args := m.Called(employeeData)
	return args.Error(0)
}

// UpdateOne mocks the UpdateOne method of the IEmployee repository interface
func (m *MockEmployeeRepo) UpdateOne(employeeData *model.Employee) error {
	args := m.Called(employeeData)
	return args.Error(0)
}

// DeleteOne mocks the DeleteOne method of the IEmployee repository interface
func (m *MockEmployeeRepo) DeleteOne(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// FindOneFirst mocks the FindOneFirst method of the IEmployee repository interface
func (m *MockEmployeeRepo) FindOneFirst(id string) (*model.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Employee), args.Error(1)
}

func TestListEmployeeUseCase(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockEmployeeRepo)
	// Inject the mock repository into the use case
	usecase := NewEmployeeUsecase(nil, mockRepo)

	// Define expected data and behavior for the mock repository
	expectedData := []*model.Employee{{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com"}}
	mockRepo.On("FindAll", mock.AnythingOfType("*model.Employee")).Return(expectedData, nil)

	// Call the use case method
	resp := usecase.ListEmployeeUseCase(&model.Employee{})

	// Assert that the response is as expected
	assert.Equal(t, expectedData, resp.Data)
}

// Similar tests for other use case methods...
