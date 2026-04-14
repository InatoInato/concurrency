package main

import (
	"bytes"
	"concurrency/internal/db"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 1. Define the Mock
type MockQueries struct {
	// We use fields to tell the mock what to return
	UserToReturn db.User
	ErrToReturn  error
}

// 2. Implement the Interface for the Mock
func (m *MockQueries) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return m.UserToReturn, m.ErrToReturn
}
func (m *MockQueries) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return m.UserToReturn, m.ErrToReturn
}
func (m *MockQueries) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return m.UserToReturn, m.ErrToReturn
}
func (m *MockQueries) DeleteUser(ctx context.Context, id int32) error {
	return m.ErrToReturn
}

// 3. The Actual Test
func TestCreateUser(t *testing.T) {
	// Setup the mock data
	mockUser := db.User{ID: 1, Name: "David", Age: 25}
	mockQueries := &MockQueries{
		UserToReturn: mockUser,
	}

	server := &Server{queries: mockQueries}

	// Create a dummy request body
	userParams := db.CreateUserParams{Name: "David", Age: 25}
	body, _ := json.Marshal(userParams)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	// Execute
	handler := http.HandlerFunc(server.createUser)
	handler.ServeHTTP(rr, req)

	// Assertions
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	var returnedUser db.User
	json.NewDecoder(rr.Body).Decode(&returnedUser)

	if returnedUser.Name != "David" {
		t.Errorf("expected name David, got %s", returnedUser.Name)
	}
}

