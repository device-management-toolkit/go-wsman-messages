package security_test

import (
	"testing"

	"github.com/99designs/keyring"
	"github.com/stretchr/testify/mock"

	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/security"
)

// MockKeyring to mock the keyring interface for unit testing.
type MockKeyring struct {
	mock.Mock
}

func (m *MockKeyring) Set(serviceName string, item keyring.Item) error {
	args := m.Called(serviceName, item)

	return args.Error(0)
}

func (m *MockKeyring) Get(serviceName, key string) (keyring.Item, error) {
	args := m.Called(serviceName, key)

	return args.Get(0).(keyring.Item), args.Error(1)
}

func TestSetKeyValue(t *testing.T) {
	mockKeyring := new(MockKeyring)
	storage := security.NewStorage("testService", mockKeyring)

	mockKeyring.On("Set", "testService", keyring.Item{Key: "testKey", Data: []byte("testValue")}).Return(nil)

	err := storage.SetKeyValue("testKey", "testValue")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	mockKeyring.AssertExpectations(t)
}

func TestGetKeyValue(t *testing.T) {
	mockKeyring := new(MockKeyring)
	storage := security.NewStorage("testService", mockKeyring)

	mockKeyring.On("Get", "testService", "testKey").Return(keyring.Item{Key: "testKey", Data: []byte("testValue")}, nil)

	value, err := storage.GetKeyValue("testKey")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != "testValue" {
		t.Errorf("Expected value 'testValue', got %v", value)
	}

	mockKeyring.AssertExpectations(t)
}
