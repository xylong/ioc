package injector

import (
	"reflect"
)

var Factory *factory

func init() {
	Factory = NewFactory()
}

type factory struct {
	mapper Mapper
}

func NewFactory() *factory {
	return &factory{
		mapper: make(Mapper),
	}
}

func (f factory) Set(list ...interface{}) {
	if list == nil || len(list) == 0 {
		return
	}
	for _, item := range list {
		f.mapper.add(item)
	}
}

func (f factory) Get(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	value := f.mapper.get(v)
	if value.IsValid() {
		return value.Interface()
	}
	return nil
}

// Apply 处理依赖注入
func (f factory) Apply(any interface{}) {
	if any == nil {
		return
	}
	v := reflect.ValueOf(any)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}
	// 在容器里匹配
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		// 判断是否可以赋值
		if v.Field(i).CanSet() && field.Tag.Get("inject") != "" {
			if value := f.Get(field.Type); value != nil {
				v.Field(i).Set(reflect.ValueOf(value))
			}
		}
	}
}
