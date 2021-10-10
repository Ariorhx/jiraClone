package CannotResolve

import (
	"sync"
)

type data struct {
	users    map[string]int
	priority map[string]int
}

var once sync.Once
var d *data

func GetUsers() map[string]int {
	once.Do(initData)
	return d.users
}

func GetPriority() map[string]int {
	once.Do(initData)
	return d.priority
}

func initData() {
	d = &data{GetDatabase().GetUsers(), GetDatabase().GetPriority()}
}
