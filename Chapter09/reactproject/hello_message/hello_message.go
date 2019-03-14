// hello_message.go

package hellomessage

import (
	"fmt"

	"honnef.co/go/js/dom"
	"myitcv.io/react"
)

//go:generate reactGen

// Step 1
// Declare a type that has (at least) an anonymous embedded react.ComponentDef
// (it can have other fields); this type must have the suffix 'Def', which corresponds to
// 'Definition'
//
type HelloMessageDef struct {
	react.ComponentDef
}

// Step 2
// Optionally declare a props type; the naming convention is *Props
//
type HelloMessageProps struct {
	Message string
}

// Step 3
// Optionally declare a state type; the naming convention is *State
//
type HelloMessageState struct {
	CurrName string
	Names    []string
}

// hello_message.go continued....

// Step 4
// Declare a function to create instances of the component, i.e. an element. If
// your component requires props to be specified, add this to the function
// signature. If the props are optional, use a props pointer type.
//
// buildHelloMessageElem is code generated to wrap a call to react.CreateElement.
//
// Convention is that this function is given the name of the component, HelloMessage
// in this instance. Because this component has props, we also accept these as part
// of the constructor.
//
func HelloMessage(p HelloMessageProps) *HelloMessageElem {
	fmt.Println("Building element...")
	return buildHelloMessageElem(p)
}

// Step 5
// Define a Render method on the component's non-pointer type
//
func (r HelloMessageDef) Render() react.Element {
	InputName := react.Input(&react.InputProps{
		Type:        "text",
		Key:         "FirstName",
		Placeholder: "Mina",
		Value:       r.State().CurrName,
		OnChange:    r,
	}, nil)
	InputBtn := react.Input(&react.InputProps{
		Type:  "Submit",
		Value: "Submit",
	}, nil)
	Form := react.Form(&react.FormProps{
		OnSubmit: r,
	},
		react.S("Name: "),
		InputName,
		InputBtn)
	names := r.State().Names
	fmt.Println(names)
	entries := make([]react.RendersLi, len(names))
	for i, name := range names {
		entries[i] = react.Li(nil, react.S(r.Props().Message+" "+name))
	}
	return react.Div(nil,
		Form,
		react.S(r.Props().Message+" "+r.State().CurrName),
		react.Ul(nil, entries...),
	)
}

func (r HelloMessageDef) OnSubmit(e *react.SyntheticEvent) {
	//Prevent the default form submission action
	e.PreventDefault()
	//Add the new name to the list of names in the state object
	names := r.State().Names
	names = append(names, r.State().CurrName)
	/*
		Change the state so that the current name is now empty, and the new name gets added to the existing list of names
	*/
	r.SetState(HelloMessageState{CurrName: "", Names: names})
}

func (c HelloMessageState) Equals(v HelloMessageState) bool {
	if c.CurrName != v.CurrName {
		return false
	}

	if len(c.Names) != len(v.Names) {
		return false
	}

	for i := range v.Names {
		if v.Names[i] != c.Names[i] {
			return false
		}
	}
	return true
}

func (r HelloMessageDef) OnChange(e *react.SyntheticEvent) {
	//we need to import "honnef.co/go/js/dom" for this to work
	//get target: our input text component
	target := e.Target().(*dom.HTMLInputElement)
	//get current state
	currState := r.State()
	//change state to include new value in our input text component, as well as the existing history of names
	r.SetState(HelloMessageState{CurrName: target.Value, Names: currState.Names})
}
