### IOC容器
轻量级IoC容器。支持单例、多例、运行时获取bean，依赖支持结构体指针和接口。

- **Unwrap** 自动设置依赖，支持***单例***和***多例***注入
- **Set** 手动设置依赖，支持***单例***注入
- **Get** 获取依赖
- **Apply** 注入依赖

***

### 安装
```
go get -u github.com/xylong/ioc
```

***

创建依赖配置
```go
type Service struct {
}

func NewService() *Service {
    return &Service{}
}

// Order 创建订单service
func (s *Service) Order() *service.OrderService {
    return service.NewOrderService()
}

// Login 创建登录service
func (s *Service) Login() *service.LoginService {
    return service.NewLoginService()
}
```

***

使用
```go
ioc.Factory.Unwrap(config.NewService()) // 自动展开依赖配置方法
ioc.Factory.Set(service.NewOrderService())    // 手动加入依赖配置

fmt.Println(ioc.Factory.Get((*service.OrderService)(nil)))  // 获取依赖

user := service.NewUserService()
ioc.Factory.Apply(user) // 为UserService注入依赖
```

***

单例注入
```go
type UserService struct {
    Order *OrderService `inject:"-"`
}
```
struct字段需设置**`inject:"-"`**标签，依赖每次直接从容器中取，支持***Set、Unwrap***两种方式的依赖

多例注入
```go
type UserService struct {
    Order *OrderService `inject:"Service.Order()"`
}
```
struct字段需设置**`inject:"结构体名.方法名()"`** ，多例就是类名+方法名执行，每次都会创建一个新实例(***只支持Unwrap方式的依赖配置***)