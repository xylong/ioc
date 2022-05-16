package injector

import "reflect"

type BeanMapper map[reflect.Type]reflect.Value

func (m BeanMapper) add(bean interface{}) {
	t := reflect.TypeOf(bean)

	if t.Kind() != reflect.Ptr {
		panic("required ptr")
	}

	m[t] = reflect.ValueOf(bean)
}

func (m BeanMapper) get(bean interface{}) reflect.Value {
	t := reflect.TypeOf(bean)

	if v, ok := m[t]; ok {
		return v
	}

	return reflect.Value{}
}
