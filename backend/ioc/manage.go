package ioc

import "github.com/gin-gonic/gin"

var nc = &NamespaceContainer{
	ns: map[string]*Container{
		"api":        NewContainer(),
		"controller": NewContainer(),
		"config":     NewContainer(),
		"default":    NewContainer(),
	},
}

func Init() error {
	return nc.Init()
}

func Destroy() error {
	return nc.Destroy()
}

// Controller
// ioc.Controller().Restiry()
// ioc.Controller().Get()
func Controller() *Container {
	return nc.ns["controller"]
}

// Api
// ioc.Api().Restiry
// ioc.Api().Get()
func Api() *Container {
	return nc.ns["api"]
}

func RegistryGinApi(rr gin.IRouter) {
	nc.RegistryGinApi(rr)
}

// NamespaceContainer Build a multiple container
type NamespaceContainer struct {
	// {}/[]
	ns map[string]*Container
}

func (c *NamespaceContainer) Init() error {
	// Traverse Namespace
	for key := range c.ns {
		c := c.ns[key]
		// Traverse Namespace object
		for objectName := range c.storage {
			if err := c.storage[objectName].Init(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *NamespaceContainer) Destroy() error {
	// Traverse Namespace
	for key := range c.ns {
		c := c.ns[key]
		// Traverse Namespace object
		for objectName := range c.storage {
			if err := c.storage[objectName].Destroy(); err != nil {
				return err
			}
		}
	}
	return nil
}

type GinApi interface {
	// Object Basic constraints
	Object
	// Registry additional constraints
	// ioc.Api().Get(token.AppName).(*api.TokenApiHandler).Registry(rr)
	Registry(rr gin.IRouter)
}

// RegistryGinApi Traverse all objects and complete Api registration
// Provided by the ioc caller object Registry Method, lets module register to gin root router
func (c *NamespaceContainer) RegistryGinApi(rr gin.IRouter) {
	// Traverse Namespace
	for key := range c.ns {
		c := c.ns[key]
		// Traverse Namespace object
		for objectName := range c.storage {
			obj := c.storage[objectName]
			// determine whether an object is a GinApi object (constraint)
			// Determine whether the object has Registry(rr gin.IRouter)
			// Assert that the object satisfies the GinApi interface and implements the Registry function
			if v, ok := obj.(GinApi); ok {
				v.Registry(rr)
			}
		}
	}
}
