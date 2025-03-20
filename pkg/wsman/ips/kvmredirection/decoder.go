package kvmredirection

const (
	IPSKVMRedirectionSettingData string = "IPS_KVMRedirectionSettingData"
	TerminateSession             string = "TerminateSession"
	ValueNotFound                string = "Value not found in map"
)

const (
	ReturnValueSuccess ReturnValue = iota
	ReturnValueInternalError
)

// returnValueToString maps ReturnValue values to strings.
var returnValueToString = map[ReturnValue]string{
	ReturnValueSuccess:       "Success",
	ReturnValueInternalError: "InternalError",
}

// String returns a human-readable string representation.
func (r ReturnValue) String() string {
	if s, ok := returnValueToString[r]; ok {
		return s
	}

	return ValueNotFound
}
