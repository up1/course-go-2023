package main

import "testing"

type MockDatabase struct{}

func (d *MockDatabase) connect() error {
	return nil
}
func (d *MockDatabase) insertData() error {
	return nil
}

type MockEmailService struct{}

func (e *MockEmailService) connect() error {
	return nil
}

func (e *MockEmailService) send() error {
	return nil
}

func TestDI(t *testing.T) {
	db := &MockDatabase{}
	email := &MockEmailService{}
	service := NewMyService(db, email)
	service.Process()
}
