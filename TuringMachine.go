package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type (
	turingMachine struct {
		transitions []transition
		finalStates []string
		tapeContent []string
		initialState string
	}

	transition struct {
		currentState string
		nextState string
		inputCharacter string
		replaceCharacter string
		rightOrLeft bool // true is right false is left
	}
)

func main() {
	printSample()

	machine := &turingMachine{}

	getTransitions(machine)
	getFinalStates(machine)
	getInitialState(machine)
	getTapeContent(machine)
	checkTheInput(machine)

	var end string
	fmt.Scanln(&end)
	if end == "end" {
		return
	}
}

func printSample() {
	fmt.Println("Sample input:")
	fmt.Println("Transition Rules:")
	fmt.Println("q0,a=q1,b,R")
	fmt.Println("q1,a=q1,b,R")
	fmt.Println("q1,b=q2,_,L")
	fmt.Println("end")
	fmt.Println("Final States:")
	fmt.Println("q2")
	fmt.Println("end")
	fmt.Println("Initial State:")
	fmt.Println("q0")
	fmt.Println("Tape Content:")
	fmt.Println("aab")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func getTransitions(machine *turingMachine) {
	fmt.Println("Enter the transition rules line by line: (enter 'end' when finished)")

	reader := bufio.NewReader(os.Stdin)
	transitionString, _ := reader.ReadString('\n')
	transitionString = strings.TrimSpace(transitionString)
	transitionString = strings.ReplaceAll(transitionString, " ", "")

	for transitionString != "end" {
		transition := transition{}
		transition.currentState = string(transitionString[1])
		transition.inputCharacter = string(transitionString[3])
		transition.nextState = string(transitionString[6])
		transition.replaceCharacter = string(transitionString[8])
		if string(transitionString[10]) == "R" {
			transition.rightOrLeft = true
		} else {
			transition.rightOrLeft = false
		}

		machine.transitions = append(machine.transitions, transition)

		fmt.Scanln(&transitionString)
		transitionString = strings.TrimSpace(transitionString)
	}
}

func getFinalStates(machine *turingMachine) {
	fmt.Println("Enter the final states: (enter 'end' when finished)")

	var f string
	fmt.Scanln(&f)

	for f != "end" {
		f = string(f[1])

		machine.finalStates = append(machine.finalStates, f)

		fmt.Scanln(&f)
	}
}

func getTapeContent(machine *turingMachine) {
	for i := 0; i < 100; i++ {
		machine.tapeContent = append(machine.tapeContent, "_")
	}

	fmt.Println("Enter tape content:")
	var t string
	fmt.Scanln(&t)
	c := strings.Split(t, "")
	for i := 0; i < len(c); i++ {
		machine.tapeContent = append(machine.tapeContent, c[i])
	}

	for i := 0; i < 100; i++ {
		machine.tapeContent = append(machine.tapeContent, "_")
	}
}

func getInitialState(machine *turingMachine) {
	fmt.Println("Enter the initial state:")
	var i string
	fmt.Scanln(&i)
	machine.initialState = string(i[1])
}

func checkTheInput(machine *turingMachine) {
	currentState := machine.initialState
	currentInputCharacter := machine.tapeContent[100]
	currentHead := 100

	foundTransition := false
	for true {
		for i := 0; i < len(machine.transitions); i++ {
			if machine.transitions[i].currentState == currentState && machine.transitions[i].inputCharacter == currentInputCharacter {
				currentState = machine.transitions[i].nextState
				machine.tapeContent[currentHead] = machine.transitions[i].replaceCharacter
				if machine.transitions[i].rightOrLeft {
					currentHead++
				} else {
					currentHead--
				}

				foundTransition = true
				break
			}
		}

		if !foundTransition {
			for j := 0; j < len(machine.finalStates); j++ {
				if currentState == machine.finalStates[j] {
					fmt.Println("Accepted. Machine halted in a final state.")
					return
				}
			}

			fmt.Println("NOT Accepted. Machine halted in a non-final state.")
			return
		}

		foundTransition = false
		currentInputCharacter = machine.tapeContent[currentHead]
	}
}