package main

import (
	"ergo.services/ergo/act"
	"ergo.services/ergo/gen"
)

func factoryMyActor() gen.ProcessBehavior {
	return &myActor{}
}

type myActor struct {
	act.Actor

	value int
}

func (a *myActor) Init(args ...any) error {
	a.value = 1
	a.Send(a.PID(), "hello")
	a.Log().Debug("actor started %s %d", a.PID(), a.value)
	a.Spawn(factoryMyActor, gen.ProcessOptions{}, 1)
	a.Spawn(factoryMyActor, gen.ProcessOptions{})
	return nil
}

func (a *myActor) HandleMessage(from gen.PID, message any) error {
	switch message {
	case "increase":
		a.value += 8
		a.Log().Info("increased value by %d", 8)
		a.Send(gen.Atom("abc"), a.value)

		x, _ := a.CallPID(gen.PID{}, 12345, 1)
		a.Send(gen.Atom("def"), x)
		return nil
	}
	return gen.TerminateReasonNormal
}
