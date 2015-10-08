package main

import (
	"fmt"
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
		str, _ := any.(string)
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
		str, _ := any.(string)
		return this.value == str
	case Element:
		str := any.(Element).value
		return this.value == str
	default:
		return false
	}
}

type Dfa struct {
	states       []State
	alphabet     []Element
	transition   func(State, Element) State
	startState   State
	acceptStates []State
}

func (dfa *Dfa) assertInvariants() error {
	//TODO: implement this
	return nil
}

func (dfa *Dfa) isFinalState(state State) bool {
	for _, acceptState := range dfa.acceptStates {
		if acceptState.equals(state) {
			return true
		}
	}
	return false
}

func (dfa *Dfa) AcceptInput(input string) bool {
	state := dfa.startState
	for _, c := range input {
		alphEle := NewElement(string(c))
		state = dfa.transition(state, alphEle)
	}
	return dfa.isFinalState(state)
}

func NewDfa(states, acceptStates []State, startState State, alphabet []Element, transition func(State, Element) State) (Dfa, error) {
	dfa := Dfa{states, alphabet, transition, startState, acceptStates}

	err := dfa.assertInvariants()

	if err != nil {
		return dfa, fmt.Errorf("failing invariants")
	}

	return dfa, err

}

func main() {
	fmt.Printf("DFA STUFF\n")

	s1 := NewState("s1")
	s2 := NewState("s2")
	states := []State{s1, s2}
	acceptStates := []State{s1}
	startState := s1
	alphabet := []Element{NewElement("0"), NewElement("1")}

	trx := func(state State, alphEle Element) State {
		if state.equals(s1) && alphEle.equals("0") {
			return s2
		} else if state.equals(s1) && alphEle.equals("1") {
			return s1
		} else if state.equals(s2) && alphEle.equals("0") {
			return s1
		} else if state.equals(s2) && alphEle.equals("1") {
			return s2
		} else {
			return NewState("non existent")
		}
	}

	dfa, err := NewDfa(states, acceptStates, startState, alphabet, trx)

	if err != nil {
		fmt.Printf("Error initializing the dfa %s", err.Error())
	}

	fmt.Printf("Accept 10: %t\n", dfa.AcceptInput("10"))
	fmt.Printf("Accept 100: %t\n", dfa.AcceptInput("100"))
	fmt.Printf("Accept 100: %t\n", dfa.AcceptInput("1000"))
	fmt.Printf("Accept 100: %t\n", dfa.AcceptInput("10000"))
}
