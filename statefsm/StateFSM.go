package statefsm

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"sync"
)

type FSMState string

type FSMEvent struct {
	EventName  string
	EventState string
}

type FSMHandler func() error

//Finite State Machine
type FSM struct {
	mu          sync.Mutex
	state       FSMState
	flowDiagram map[FSMState]map[FSMEvent]FSMState //State Machine Map
}

func (f *FSM) setState(newState FSMState) {
	f.state = newState
}

func NewFSM(initState FSMState) *FSM {
	return &FSM{
		state:       initState,
		flowDiagram: make(map[FSMState]map[FSMEvent]FSMState),
	}
}

func (f *FSM) AddHandler(state FSMState, event FSMEvent, eventState FSMState) *FSM {
	if _, ok := f.flowDiagram[state]; !ok {
		f.flowDiagram[state] = make(map[FSMEvent]FSMState)
	}
	if _, ok := f.flowDiagram[state][event]; ok {
		log.Warnf("The state (%s) event (%s) has been defined.", state, event)
	}
	f.flowDiagram[state][event] = eventState
	return f
}

func (f *FSM) Call(event FSMEvent, fsmHandler FSMHandler) error {
	if f == nil || f.state == "" {
		return errors.New("FSM data error.")
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	events := f.flowDiagram[f.state]
	if events == nil || len(events) == 0 {
		return errors.New("Event undefined. State : " + string(f.state))
	}
	for _, e := range events {
		if e == event {
			f.setState(event)
			return fsmHandler()
		}
	}

	if fn, ok := events[event]; ok {
		oldState := f.getState()
		f.setState(fn.Execute())
		newState := f.getState()
		log.Infof("State changed from %s to %s", oldState, newState)
	}
	return f.getState()
}
