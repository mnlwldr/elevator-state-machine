package main

import (
	"log"
	"reflect"
)

type IElevatorState interface {
	Open() Open
	Close() Close
	Move() Move
	Stop() Stop
}

type ElevatorState struct{}

func (es ElevatorState) Open() Open   { panic("open not allowed") }
func (es ElevatorState) Close() Close { panic("close not allowed") }
func (es ElevatorState) Move() Move   { panic("move not allowed") }
func (es ElevatorState) Stop() Stop   { panic("stop not allowed") }

// Move
// only "Move" and "Stop" are allowed
type Move struct {
	ElevatorState
}

func (m Move) Move() Move { return Move{} }
func (m Move) Stop() Stop { return Stop{} }

// Open
// only "Close" are allowed
type Open struct {
	ElevatorState
}

func (o Open) Close() Close {
	return Close{}
}

// Stop
// only "Open" and "Move" are allowed
type Stop struct {
	ElevatorState
}

func (s Stop) Open() Open { return Open{} }
func (s Stop) Move() Move { return Move{} }

// Close
// only "Open" and "Move" are allowed
type Close struct {
	ElevatorState
}

func (c Close) Open() Open { return Open{} }
func (c Close) Move() Move { return Move{} }

type Elevator struct {
	State IElevatorState
}

func (e *Elevator) Open() {
	log.Printf("%s from %s", "open", reflect.TypeOf(e.State))
	e.State = e.State.Open()
}
func (e *Elevator) Close() {
	log.Printf("%s from %s", "close", reflect.TypeOf(e.State))
	e.State = e.State.Close()
}
func (e *Elevator) Move() {
	log.Printf("%s from %s", "move", reflect.TypeOf(e.State))
	e.State = e.State.Move()
}
func (e *Elevator) Stop() {
	log.Printf("%s from %s", "stop", reflect.TypeOf(e.State))
	e.State = e.State.Stop()
}

func (e *Elevator) IsOpen() bool {
	return reflect.TypeOf(e.State) == reflect.TypeOf(Open{})
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("-> %s <-", r)
		}
	}()

	e := Elevator{}
	e.State = Open{}
	e.Close()
	e.Open()
	e.Close()
	log.Printf("IsOpen: %v", e.IsOpen()) // <- false
	e.Move()
	e.Stop()
	e.Open()
	log.Printf("IsOpen: %v", e.IsOpen()) // <- true
	e.Close()
	e.Move()
	e.Move()
	e.Move()
	e.Move()
	e.Open() // <- panic

}
