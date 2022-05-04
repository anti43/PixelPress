package system

import (
	"fmt"
	"reflect"
	"strings"
)

type ControllerRegistry struct {
	types map[string]*BaseControllerWithActions
}

func (s *ControllerRegistry) NewController(name string) (*BaseControllerWithActions, error) {
	t, ok := s.types["controller."+strings.Title(name)+"Controller"]
	if !ok {
		return nil, fmt.Errorf("unrecognized type: %s in %s", name, s.types)
	}
	return t, nil
}
func (s *ControllerRegistry) RegisterNamedType(name string, t BaseControllerWithActions) {
	s.types[name] = &t
}

func (s *ControllerRegistry) RegisterType(t BaseControllerWithActions) {
	name := reflect.TypeOf(t).String()
	s.types[name] = &t
}

func NewControllerRegistry() *ControllerRegistry {
	return &ControllerRegistry{types: map[string]*BaseControllerWithActions{}}
}
