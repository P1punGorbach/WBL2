package main

import "fmt"

// Command interface
type Command interface {
	Execute()
}

// ConcreteCommand struct
type ConcreteCommand struct {
	receiver *Receiver
}

// NewConcreteCommand returns a new instance of ConcreteCommand
func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{receiver: receiver}
}

// Execute method for ConcreteCommand
func (cc *ConcreteCommand) Execute() {
	cc.receiver.Action()
}

// Receiver struct
type Receiver struct{}

// Action method for Receiver
func (r *Receiver) Action() {
	fmt.Println("Receiver executing action")
}

// Invoker struct
type Invoker struct {
	command Command
}

// SetCommand method for Invoker
func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

// ExecuteCommand method for Invoker
func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	receiver := &Receiver{}
	command := NewConcreteCommand(receiver)
	invoker := &Invoker{}

	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}