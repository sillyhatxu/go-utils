package statefsm

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"sync"
)

type FSMState string
type FSMEvent struct {
	EventName  string
	EventState FSMState
}
type FSMHandler func() error

//Finite State Machine
type FSM struct {
	mu          sync.Mutex
	state       FSMState
	flowDiagram map[FSMState][]FSMEvent //State Machine Map
}

func (f *FSM) setState(newState FSMState) {
	f.state = newState
}

func NewFSM(initState FSMState) *FSM {
	return &FSM{
		state:       initState,
		flowDiagram: make(map[FSMState][]FSMEvent),
	}
}

func (f *FSM) AddHandler(state FSMState, events []FSMEvent) *FSM {
	if _, ok := f.flowDiagram[state]; !ok {
		f.flowDiagram[state] = make([]FSMEvent, len(events))
	}
	if _, ok := f.flowDiagram[state]; ok {
		log.Warnf("The state (%s) event (%s) has been defined.", state, events)
	}
	f.flowDiagram[state] = events
	return f
}

func (f *FSM) Call(event FSMEvent, fsmHandler FSMHandler) error {
	if f == nil || f.state == "" {
		return errors.New("FSM data error.")
	}
	if event.EventState == "" {
		return errors.New("EventState is nil.")
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	events := f.flowDiagram[f.state]
	if events == nil || len(events) == 0 {
		return errors.New("Event undefined. State : " + string(f.state))
	}
	for _, e := range events {
		if e.EventName == event.EventName {
			log.Infof("State changed from %s to %s", f.state, event.EventState)
			f.setState(event.EventState)
			return fsmHandler()
		}
	}
	return errors.New("State transition error." + string(f.state) + " -> " + string(event.EventState))
}
