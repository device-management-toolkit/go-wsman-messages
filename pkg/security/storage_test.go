package security_test

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/security"
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
