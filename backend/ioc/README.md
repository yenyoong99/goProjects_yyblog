## IOC Design

### Function:

- Object Hosted
  - Registration of objects
  - Get Object

- Unified manage of objects:
  - Configuration
  - initialization
  - destroy

### Implement container based on Map

#### Single container design

```go
container = map[string]any
```

```go
type IocObject interface {
    Init() error
    Destory() error
}


type TestStruct struct {

}

func (t *TestStruct) Init() error {

}

func (t *TestStruct) Destory() error {

}

// container = map[string]IocObject
ioc.Registry("service name", &TestStruct{})

// When starting, complete the unified management of objects, 
// cycle all objects in the container, and call the Init method
ioc.Init()
```

1 map, Duplicate names are not allowed
+ TokenServiceImpl: Controller
+ TokenApiHandler: Api


```go
// Register controller
ioc.Controller.Registry("module name", &TokenServiceImpl{})
ioc.Api.Registry("module name", &TOkenApiHandler{})
```

According to the program design and the responsibility constraints on these objects, the container is partitioned:
+ Api: Responsible for registering objects of Api implementation types
+ Controller: Responsible for registering the object of the service implementation class
+ Config: configuration object, db, kafak, redis
+ Default: reserved area

#### multiple container design
```go
api_contailer = map[string]IocObject
controller_container = map[string]IocObject
```

### Implement IOC

Encapsulation Container
```go
func TestContainerGetAndRegistry(t *testing.T) {
	c := ioc.NewContainer()
	c.Registry("TestStruct", &TestStruct{})
	t.Log(c.Get("TestStruct"))

	// Assert using
	c.Get("TestStruct").(*TestStruct).XXX()
}
```


Encapsulation Manager
```go
func TestManageGetAndRegistry(t *testing.T) {
	ioc.Controller().Registry("TestStruct", &TestStruct{})
	t.Log(ioc.Controller().Get("TestStruct"))

	// Assert using
	ioc.Controller().Get("TestStruct").(*TestStruct).XXX()
}
```
