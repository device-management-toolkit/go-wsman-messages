package base

import (
	"encoding/xml"
	"reflect"
	"strings"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type WSManService[T any] struct {
	Base message.Base
}

func NewService[T any](creator *message.WSManMessageCreator, resourceURI string, client client.WSMan) WSManService[T] {
	return WSManService[T]{
		Base: message.NewBaseWithClient(creator, resourceURI, client),
	}
}

func (s WSManService[T]) Get() (T, error) {
	return s.getBySelector(nil)
}

func (s WSManService[T]) getBySelector(selector *message.Selector) (T, error) {
	var out T

	msg := &client.Message{XMLInput: s.Base.Get(selector)}

	injectMessage(&out, msg)

	if err := s.Base.Execute(msg); err != nil {
		return out, err
	}

	if err := xml.Unmarshal([]byte(msg.XMLOutput), &out); err != nil {
		return out, err
	}

	injectMessage(&out, msg)

	return out, nil
}

func (s WSManService[T]) GetByName(name string) (T, error) {
	selector := &message.Selector{
		Name:  "Name",
		Value: name,
	}

	return s.getBySelector(selector)
}

func (s WSManService[T]) GetByInstanceID(name string) (T, error) {
	selector := &message.Selector{
		Name:  "InstanceID",
		Value: name,
	}

	return s.getBySelector(selector)
}

func (s WSManService[T]) Enumerate() (T, error) {
	var out T

	msg := &client.Message{XMLInput: s.Base.Enumerate()}

	injectMessage(&out, msg)

	if err := s.Base.Execute(msg); err != nil {
		return out, err
	}

	if err := xml.Unmarshal([]byte(msg.XMLOutput), &out); err != nil {
		return out, err
	}

	injectMessage(&out, msg)

	return out, nil
}

func (s WSManService[T]) Pull(ctx string) (T, error) {
	var out T

	msg := &client.Message{XMLInput: s.Base.Pull(ctx)}

	injectMessage(&out, msg)

	if err := s.Base.Execute(msg); err != nil {
		return out, err
	}

	if err := xml.Unmarshal([]byte(msg.XMLOutput), &out); err != nil {
		return out, err
	}

	injectMessage(&out, msg)

	return out, nil
}

func (s WSManService[T]) Put(request any) (T, error) {
	var out T

	injectNamespace(request, s.Base.ClassName)

	msg := &client.Message{XMLInput: s.Base.Put(request, false, nil)}

	injectMessage(&out, msg)

	if err := s.Base.Execute(msg); err != nil {
		return out, err
	}

	if err := xml.Unmarshal([]byte(msg.XMLOutput), &out); err != nil {
		return out, err
	}

	injectMessage(&out, msg)

	return out, nil
}

// Inject *client.Message into out.Message using reflection.
func injectMessage[T any](v *T, msg *client.Message) {
	rv := reflect.ValueOf(v).Elem()
	msgField := rv.FieldByName("Message")

	if msgField.IsValid() && msgField.CanSet() {
		msgField.Set(reflect.ValueOf(msg))
	}
}

func injectNamespace(request any, resourceURI string) {
	rv := reflect.ValueOf(request)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	field := rv.FieldByName("H")
	if !field.IsValid() || !field.CanSet() || field.Kind() != reflect.String {
		return
	}

	var schema string

	switch {
	case strings.HasPrefix(resourceURI, "IPS_"):
		schema = message.IPSSchema
	case strings.HasPrefix(resourceURI, "AMT_"):
		schema = message.AMTSchema
	case strings.HasPrefix(resourceURI, "CIM_"):
		schema = message.CIMSchema
	default:
		return
	}

	field.SetString(schema + resourceURI)
}
