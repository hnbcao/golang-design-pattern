package observer

import "testing"

func TestSubject_UpdateContext(t *testing.T) {
	subject := NewSubject()
	watcher00 := NewWatcher("watcher00")
	watcher01 := NewWatcher("watcher01")
	watcher02 := NewWatcher("watcher02")
	subject.Attach(watcher00)
	subject.Attach(watcher01)
	subject.Attach(watcher02)

	subject.UpdateContext("observer mode")
}
