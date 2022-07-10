package injector

import "reflect"

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
