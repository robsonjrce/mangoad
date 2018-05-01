package site

import (
	"log"
	"sync"
)

var (
	modulesMux sync.Mutex
	modules    = []Module{}
)

// Module define all modules available
type Module struct {
	ID string
	Fn ModuleInstantiate
}

// ModuleInstantiate is responsible to instantiate correct type of module
type ModuleInstantiate func(string) Site

// RegisterModule register modules to be supported
func RegisterModule(url string, fn ModuleInstantiate) {
	modulesMux.Lock()
	modules = append(modules, Module{ID: url, Fn: fn})
	log.Printf("=> registering module: %s \n", url)
	modulesMux.Unlock()
}
