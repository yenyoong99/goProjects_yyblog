package ioc

func NewContainer() *Container {
	return &Container{
		storage: map[string]Object{},
	}
}

type Container struct {
	storage map[string]Object
}

func (c *Container) Registry(name string, obj Object) {
	c.storage[name] = obj
}

// Get ioc.Get("module name").(*TestService)
func (c *Container) Get(name string) any {
	return c.storage[name]
}

// // Provide an object traversal method,
// // The user passes a function and passes the traversed object as a parameter to this function (callback function)
// func (c *Container) ForEach(cb func(objectName string, objectValue Object)) {
// 	for key := range c.storage {
// 		cb(key, c.storage[key])
// 	}
// }
