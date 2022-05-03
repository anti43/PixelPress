package system

import (
	"fmt"
	"reflect"
)

type Registry struct {
	types map[string]reflect.Type
}

func (s *Registry) RegisterControllerType(name string, t BaseController) {
	a := reflect.TypeOf(t)
	s.types["controller."+name] = a
}

func (s *Registry) NewController(name string) (*BaseController, error) {
	t, ok := s.types["controller."+name]
	if !ok {
		return nil, fmt.Errorf("unrecognized type: %s in %s", name, s.types)
	}
	var v = BaseController(reflect.New(t)) //fixme this breaks
	return &v, nil
}

func (s *Registry) RegisterType(name string, t interface{}) {
	a := reflect.TypeOf(t)
	s.types[name] = a
}

func (s *Registry) New(name string) (interface{}, error) {
	t, ok := s.types[name]
	if !ok {
		return nil, fmt.Errorf("unrecognized type: %s in %s", name, s.types)
	}
	return reflect.New(t).Interface(), nil
}

func NewRegistry() *Registry {
	return &Registry{types: map[string]reflect.Type{}}
}
