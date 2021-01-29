package main

import (
	"golang.org/x/sync/errgroup"
	"path/filepath"
	"sync"
)

func main() {
	filepath.Walk()
	errgroup.Group{}
	sync.Pool{}
}
