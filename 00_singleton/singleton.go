package singleton

import "sync"

var singleton *Singleton
var once sync.Once

type Singleton struct {
	name string
}

func Instance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{
			name: "Singleton",
		}
	})
	return singleton
}
