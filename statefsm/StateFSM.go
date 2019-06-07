package statefsm

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

type FSMState string
type FSMEvent string
type FSMHandler func() FSMState

//Finite State Machine
type FSM struct {
	mu       sync.Mutex
	state    FSMState
	handlers map[FSMState]map[FSMEvent]FSMHandler
}

func (f *FSM) getState() FSMState {
	return f.state
}

func (f *FSM) setState(newState FSMState) {
	f.state = newState
}

func NewFSM(initState FSMState) *FSM {
	return &FSM{
		state:    initState,
		handlers: make(map[FSMState]map[FSMEvent]FSMHandler),
	}
}

func (f *FSM) AddHandler(state FSMState, event FSMEvent, handler FSMHandler) *FSM {
	if _, ok := f.handlers[state]; !ok {
		f.handlers[state] = make(map[FSMEvent]FSMHandler)
	}
	if _, ok := f.handlers[state][event]; ok {
		log.Warnf("The state (%s) event (%s) has been defined.", state, event)
	}
	f.handlers[state][event] = handler
	return f
}

func (f *FSM) Call(event FSMEvent) FSMState {
	f.mu.Lock()
	defer f.mu.Unlock()
	events := f.handlers[f.getState()]
	if events == nil {
		return f.getState()
	}
	if fn, ok := events[event]; ok {
		oldState := f.getState()
		f.setState(fn())
		newState := f.getState()
		log.Infof("State changed from %s to %s", oldState, newState)
	}
	return f.getState()
}
