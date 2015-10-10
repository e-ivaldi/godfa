package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
)

func main() {
	testDfa()
}

func testDfa() {
	fmt.Printf("DFA STUFF\n")

	s1 := NewState("s1")
	s2 := NewState("s2")
	states := mapset.NewSet()
	states.Add(s1)
	states.Add(s2)
	acceptStates := mapset.NewSet()
	startState := s1
	alphabet := mapset.NewSet()
	alphabet.Add(NewElement("0"))
	alphabet.Add(NewElement("1"))

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
	}else {

		fmt.Printf("Accept 10: %t\n", dfa.AcceptInput("10"))
		fmt.Printf("Accept 100: %t\n", dfa.AcceptInput("100"))
		fmt.Printf("Accept 100: %t\n", dfa.AcceptInput("1000"))
		fmt.Printf("Accept 100: %t\n", dfa.AcceptInput("10000"))
	}
}
