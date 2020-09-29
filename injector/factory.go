package injector

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
