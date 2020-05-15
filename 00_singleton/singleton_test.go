package singleton

import (
	"sync"
	"testing"
)

func TestInstance(t *testing.T) {
	inst00 := Instance()
	inst01 := Instance()
	if inst00 == inst01 {
		t.Log("this two object are equal, and singleton object name is", inst00.name)
	}
}

const parCount = 100

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instances[index] = Instance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		} else {
			t.Log("instance is equal")
		}
	}
}
