package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
)

type State struct {
	value string
}

func NewState(value string) State {
	return State{value}
}

func (this *State) equals(any interface{}) bool {
	switch any.(type) {
	case string:
		str := any.(string)
		return this.value == str
	case State:
		str := any.(State).value
		return this.value == str
	default:
		return false
	}
}

type Element struct {
	value string
}

func NewElement(value string) Element {
	return Element{value}
}

func (this *Element) equals(any interface{}) bool {
	switch any.(type) {
	case string:
		str := any.(string)
		return this.value == str
	case Element:
		str := any.(Element).value
		return this.value == str
	default:
		return false
	}
}

type Dfa struct {
	states       mapset.Set
	alphabet     mapset.Set
	transition   func(State, Element) State
	startState   State
	acceptStates mapset.Set
}

func (dfa *Dfa) assertInvariants() error {
	if dfa.states.Contains(dfa.startState) &&
	dfa.acceptStates.IsSubset(dfa.states){
		return nil
	}
	return fmt.Errorf("failing invariants")
}

func (dfa *Dfa) isFinalState(state State) bool {
	return dfa.acceptStates.Contains(state)
}

func (dfa *Dfa) AcceptInput(input string) bool {
	state := dfa.startState
	for _, c := range input {
		alphEle := NewElement(string(c))
		state = dfa.transition(state, alphEle)
	}
	return dfa.isFinalState(state)
}

func NewDfa(states, acceptStates mapset.Set, startState State, alphabet mapset.Set, transition func(State, Element) State) (Dfa, error) {
	dfa := Dfa{states, alphabet, transition, startState, acceptStates}

	err := dfa.assertInvariants()

	if err != nil {
		return dfa, err
	}

	return dfa, err

}
