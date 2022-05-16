package injector

var (
	BeanFactory *BeanFactoryImpl
)

func init() {
	BeanFactory = NewBeanFactory()
}

type BeanFactoryImpl struct {
	beanMapper BeanMapper
}

func NewBeanFactory() *BeanFactoryImpl {
	return &BeanFactoryImpl{beanMapper: make(BeanMapper)}
}

func (i BeanFactoryImpl) Set(v ...interface{}) {
	if v == nil || len(v) == 0 {
		return
	}

	for _, item := range v {
		i.beanMapper.add(item)
	}
}

func (i BeanFactoryImpl) Get(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	_v := i.beanMapper.get(v)
	if _v.IsValid() {
		return _v.Interface()
	}

	return nil
}
