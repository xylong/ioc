package injector

import "reflect"

type Mapper map[reflect.Type]reflect.Value

func (m Mapper) add(v interface{}) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Ptr {
		panic("require ptr")
	}
	m[t] = reflect.ValueOf(v)
}

func (m Mapper) get(any interface{}) reflect.Value {
	var key reflect.Type
	if t, ok := any.(reflect.Type); ok {
		key = t
	} else {
		key = reflect.TypeOf(any)
	}
	if value, ok := m[key]; ok {
		return value
	}
	return reflect.Value{}
}
