package injector

import "reflect"

type mapper map[reflect.Type]reflect.Value

func (m mapper) set(any interface{}) {
	t := reflect.TypeOf(any)

	if t.Kind() != reflect.Ptr {
		panic("required ptr")
	}

	m[t] = reflect.ValueOf(any)
}

func (m mapper) get(any interface{}) reflect.Value {
	t := reflect.TypeOf(any)

	if v, ok := m[t]; ok {
		return v
	}

	return reflect.Value{}
}
