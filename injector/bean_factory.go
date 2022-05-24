package injector

import (
	"github.com/shenyisyn/goft-expr/src/expr"
	"log"
	"reflect"
)

var (
	BeanFactory *BeanFactoryImpl
)

func init() {
	BeanFactory = NewBeanFactory()
}

type BeanFactoryImpl struct {
	beanMapper BeanMapper
	ExprMap    map[string]interface{}
}

func NewBeanFactory() *BeanFactoryImpl {
	return &BeanFactoryImpl{
		beanMapper: make(BeanMapper),
		ExprMap:    make(map[string]interface{}),
	}
}

func (b BeanFactoryImpl) Set(v ...interface{}) {
	if v == nil || len(v) == 0 {
		return
	}

	for _, item := range v {
		b.beanMapper.add(item)
	}
}

func (b BeanFactoryImpl) Get(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	_v := b.beanMapper.get(v)
	if _v.IsValid() {
		return _v.Interface()
	}

	return nil
}

// Apply 注入
func (b BeanFactoryImpl) Apply(bean interface{}) {
	if bean == nil {
		return
	}

	v := reflect.ValueOf(bean)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.CanSet() && v.Type().Field(i).Tag == "inject" {
			if v.Type().Field(i).Tag.Get("inject") == "" {
				if val := b.Get(f.Type()); val != nil {
					f.Set(reflect.ValueOf(val))
				}
			} else {
				log.Println("aaa")
				r := expr.BeanExpr(v.Type().Field(i).Tag.Get("inject"), b.ExprMap)
				if r != nil && !r.IsEmpty() {
					v.Field(i).Set(reflect.ValueOf(r[0]))
				}
			}
		}
	}
}
