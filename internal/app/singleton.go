package app

import (
	"sync"
)

var lock = &sync.Mutex{}

type server struct {
	App *App
}

var instance *server

func Server() *server {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(server)
	}
	return instance
}

func (ca server) Set(app *App) {
	ca.App = app
}

func (ca server) Get() *App {
	return ca.App
}
