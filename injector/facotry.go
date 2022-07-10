package injector

import (
	"reflect"
)

// injectTag 注入标签
const injectTag = "inject"

func init() {
	Factory = NewMapperFactory()
}

// Factory 工厂
var Factory *MapperFactory

type MapperFactory struct {
	mapper
}

func NewMapperFactory() *MapperFactory {
	return &MapperFactory{mapper: make(map[reflect.Type]reflect.Value)}
}

// Set 设置
func (f *MapperFactory) Set(item ...interface{}) {
	if item == nil || len(item) == 0 {
		return
	}

	for _, i := range item {
		f.mapper.set(i)
	}
}

// Get 获取
func (f *MapperFactory) Get(key interface{}) interface{} {
	if key != nil {
		if v := f.mapper.get(key); v.IsValid() {
			return v.Interface()
		}
	}

	return nil
}

// Apply 处理依赖注入
func (f *MapperFactory) Apply(obj interface{}) {
	if obj == nil {
		return
	}

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if v.Field(i).CanSet() && (field.Tag == injectTag || field.Tag.Get(injectTag) == "" || field.Tag.Get(injectTag) == "-") {
			if val := f.Get(field.Type); val != nil {
				v.Field(i).Set(reflect.ValueOf(val))
			}
		}
	}
}
