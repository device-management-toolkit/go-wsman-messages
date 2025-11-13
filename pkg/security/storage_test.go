package security_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/security"
)

// MockKeyring to mock the keyring interface for unit testing.
type MockKeyring struct {
	mock.Mock
}

func (m *MockKeyring) Set(serviceName, key, value string) error {
	args := m.Called(serviceName, key, value)

	return args.Error(0)
}

func (m *MockKeyring) Get(serviceName, key string) (string, error) {
	args := m.Called(serviceName, key)

	return args.String(0), args.Error(1)
}

func (m *MockKeyring) Delete(serviceName, key string) error {
	args := m.Called(serviceName, key)

	return args.Error(0)
}

func TestSetKeyValue(t *testing.T) {
	mockKeyring := new(MockKeyring)
	storage := security.NewStorage("testService", mockKeyring)

	mockKeyring.On("Set", "testService", "testKey", "testValue").Return(nil)

	err := storage.SetKeyValue("testKey", "testValue")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	mockKeyring.AssertExpectations(t)
}

func TestGetKeyValue(t *testing.T) {
	mockKeyring := new(MockKeyring)
	storage := security.NewStorage("testService", mockKeyring)

	mockKeyring.On("Get", "testService", "testKey").Return("testValue", nil)

	value, err := storage.GetKeyValue("testKey")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != "testValue" {
		t.Errorf("Expected value 'testValue', got %v", value)
	}

	mockKeyring.AssertExpectations(t)
}

func TestDeleteKeyValue(t *testing.T) {
	mockKeyring := new(MockKeyring)
	storage := security.NewStorage("testService", mockKeyring)

	mockKeyring.On("Delete", "testService", "testKey").Return(nil)

	err := storage.DeleteKeyValue("testKey")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	mockKeyring.AssertExpectations(t)
}

func TestGetSecretValue(t *testing.T) {
	mockKeyring := new(MockKeyring)
	storage := security.NewStorage("testService", mockKeyring)

	mockKeyring.On("Get", "testService", "path/key").Return("testValue", nil)

	value, err := storage.GetSecretValue(context.Background(), "path", "key")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != "testValue" {
		t.Errorf("Expected value 'testValue', got %v", value)
	}

	mockKeyring.AssertExpectations(t)
}

func TestSetSecretValue(t *testing.T) {
	mockKeyring := new(MockKeyring)
	storage := security.NewStorage("testService", mockKeyring)

	mockKeyring.On("Set", "testService", "path/key", "testValue").Return(nil)

	err := storage.SetSecretValue(context.Background(), "path", "key", "testValue")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	mockKeyring.AssertExpectations(t)
}

func TestDeleteSecret(t *testing.T) {
	mockKeyring := new(MockKeyring)
	storage := security.NewStorage("testService", mockKeyring)

	mockKeyring.On("Delete", "testService", "testPath").Return(nil)

	err := storage.DeleteSecret(context.Background(), "testPath")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	mockKeyring.AssertExpectations(t)
}

func TestGetSecret_NotSupported(t *testing.T) {
	mockKeyring := new(MockKeyring)
	storage := security.NewStorage("testService", mockKeyring)

	_, err := storage.GetSecret(context.Background(), "testPath")
	if !errors.Is(err, security.ErrNotSupportedByKeyring) {
		t.Errorf("Expected ErrNotSupportedByKeyring, got %v", err)
	}
}

func TestSetSecret_NotSupported(t *testing.T) {
	mockKeyring := new(MockKeyring)
	storage := security.NewStorage("testService", mockKeyring)

	data := map[string]interface{}{"key": "value"}
	err := storage.SetSecret(context.Background(), "testPath", data)
	if !errors.Is(err, security.ErrNotSupportedByKeyring) {
		t.Errorf("Expected ErrNotSupportedByKeyring, got %v", err)
	}
}
