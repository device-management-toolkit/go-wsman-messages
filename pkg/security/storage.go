package security

// NewStorage function to create a new Storage instance with a keyring interface (for testing).
func NewStorage(serviceName string, kr Keyring) Storage {
	return Storage{
		ServiceName: serviceName,
		Keyring:     kr,
	}
}

// NewKeyRingStorage function to create a new Storage instance with RealKeyring.
func NewKeyRingStorage(serviceName string) Storage {
	return Storage{
		ServiceName: serviceName,
		Keyring:     RealKeyring{},
	}
}

// SetKeyValue method to set a key-value pair in the keyring.
func (s Storage) SetKeyValue(key, value string) error {
	return s.Keyring.Set(s.ServiceName, key, value)
}

// GetKeyValue method to get a value from the keyring by key.
func (s Storage) GetKeyValue(key string) (string, error) {
	return s.Keyring.Get(s.ServiceName, key)
}

// DeleteKeyValue method to delete a key-value pair from the keyring.
func (s Storage) DeleteKeyValue(key string) error {
	return s.Keyring.Delete(s.ServiceName, key)
}
