package injector

import (
	"github.com/shenyisyn/goft-expr/src/expr"
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
	Expr map[string]interface{}
}

func NewMapperFactory() *MapperFactory {
	return &MapperFactory{
		mapper: make(map[reflect.Type]reflect.Value),
		Expr:   make(map[string]interface{}),
	}
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
// 注入的实体必须是struct的指针，并且被注入的字段需设置inject标签，如`inject:"-"`或`inject:"Service.Order()"`
// tag为-时，主动调用Set函数传入实体
// tag为表达式时，需设置Expr并且Set传入Expr对象的值，且tag的表达式和Expr的键保持一致
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
		tag := field.Tag.Get(injectTag)

		if v.Field(i).CanSet() && tag != "" {
			if tag == "" || tag == "-" {
				// 直接注入
				if val := f.Get(field.Type); val != nil {
					v.Field(i).Set(reflect.ValueOf(val))
				}
			} else {
				// 表达式注入
				if result := expr.BeanExpr(tag, f.Expr); result != nil && !result.IsEmpty() {
					v.Field(i).Set(reflect.ValueOf(result[0]))
				}
			}
		}
	}
}
