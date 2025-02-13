package keyring

import (
	"encoding/json"
	"errors"
	"reflect"
)

var (
	ErrInvalidKeyType            = errors.New("invalid key type")
	ErrSignatureValidationFailed = errors.New("signature validation failed")
)

var allowedKeyTypes = map[reflect.Type]struct{}{}

type completeKeyring struct {
	SharedKeys *SharedKeys `json:"sharedKeys,omitempty"`
	PKPKey     *PKPKey     `json:"pkpKey,omitempty"`
}

func init() {
	complete := completeKeyring{}
	fields := reflect.TypeOf(complete).NumField()
	for i := 0; i < fields; i++ {
		field := reflect.TypeOf(complete).Field(i)
		fieldType := field.Type
		allowedKeyTypes[fieldType] = struct{}{}
	}
}

type UseKeyFn interface{} // func(*KeyringType)

type Keyring interface {
	Try(...UseKeyFn) bool
	ForEach(func(key interface{}))
	Marshal() ([]byte, error)
}

type keyring struct {
	Keys map[reflect.Type]interface{}
}

func New(keys ...interface{}) Keyring {
	m := map[reflect.Type]interface{}{}
	for _, key := range keys {
		t := reflect.TypeOf(key)
		if _, ok := allowedKeyTypes[t]; !ok {
			panic(ErrInvalidKeyType)
		}
		m[t] = key
	}
	return &keyring{
		Keys: m,
	}
}

func (kr *keyring) Try(fns ...UseKeyFn) bool {
	found := false
	for _, fn := range fns {
		fnValue := reflect.ValueOf(fn)
		fnType := reflect.TypeOf(fn)
		if fnType.Kind() != reflect.Func {
			panic("invalid UseKeyFn")
		}
		if fnType.NumIn() != 1 {
			panic("invalid UseKeyFn (requires one parameter)")
		}
		argType := fnType.In(0)
		if f, ok := kr.Keys[argType]; ok {
			found = true
			fnValue.Call([]reflect.Value{reflect.ValueOf(f)})
		}
	}
	return found
}

func (kr *keyring) ForEach(fn func(key interface{})) {
	for _, k := range kr.Keys {
		fn(k)
	}
}

func (kr *keyring) Marshal() ([]byte, error) {
	complete := completeKeyring{}
	completePtr := reflect.ValueOf(&complete)
	fields := reflect.TypeOf(complete).NumField()
	for i := 0; i < fields; i++ {
		field := reflect.TypeOf(complete).Field(i)
		fieldType := field.Type
		if f, ok := kr.Keys[fieldType]; ok {
			completePtr.Elem().Field(i).Set(reflect.ValueOf(f))
		}
	}
	return json.Marshal(complete)
}

func Unmarshal(data []byte) (Keyring, error) {
	complete := completeKeyring{}
	if err := json.Unmarshal(data, &complete); err != nil {
		return nil, err
	}
	completePtr := reflect.ValueOf(&complete)
	fields := reflect.TypeOf(complete).NumField()
	values := []interface{}{}
	for i := 0; i < fields; i++ {
		field := completePtr.Elem().Field(i)
		if !field.IsNil() {
			values = append(values, field.Interface())
		}
	}
	return New(values...), nil
}
