package service

type EventChanType chan Event

type ObserverFunc func(EventChanType)
type ServiceFunc func(EventChanType)

type Manager struct {
	EventChan EventChanType
	Service   ServiceFunc
	Observer  ObserverFunc
}

func (Manager) Create(service ServiceFunc, observer ObserverFunc) Manager {
	return Manager{
		EventChan: make(EventChanType),
		Service:   service,
		Observer:  observer,
	}
}

func (m *Manager) Run() {
	go m.Observer(m.EventChan)
	m.Service(m.EventChan)
}
