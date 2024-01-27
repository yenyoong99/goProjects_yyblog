package ioc_test

import (
	"fmt"
	"github.com/yenyoong99/goProjects_yyblog/ioc"
	"testing"
)

func TestContainerGetAndRegistry(t *testing.T) {
	c := ioc.NewContainer()
	c.Registry("TestStruct", &TestStruct{})
	t.Log(c.Get("TestStruct"))

	err := c.Get("TestStruct").(*TestStruct).XXX()
	if err != nil {
		return
	}
}

type TestStruct struct {
}

func (t *TestStruct) Init() error {
	return nil
}

func (t *TestStruct) Destroy() error {
	return nil
}

func (t *TestStruct) XXX() error {
	fmt.Println("xxx log")
	return nil
}
