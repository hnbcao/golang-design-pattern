package observer

import "github.com/sirupsen/logrus"

// 主题，被观察对象
type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s.context)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

// 观察者接口
type Observer interface {
	Update(context string)
}

// 观察者接口具体实现
type Watcher struct {
	name string
}

func NewWatcher(name string) Observer {
	return &Watcher{
		name: name,
	}
}

func (r Watcher) Update(context string) {
	logrus.Infoln(r.name, "received notify:", context)
}
