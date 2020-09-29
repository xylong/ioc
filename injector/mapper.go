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

func (m Mapper) get(key interface{}) reflect.Value {
	t := reflect.TypeOf(key)
	if value, ok := m[t]; ok {
		return value
	}
	return reflect.Value{}
}
