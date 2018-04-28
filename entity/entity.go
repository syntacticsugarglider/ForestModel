package entity

import "ForestModel/behaviour"

//GenericInterface is a generic representational interface with function/type members
type GenericInterface interface {
	GetBehaviours() map[string]*behaviour.Behaviour
	Simulate(interface{})
}

//Entity represents the default required contents of an entity
type Entity struct {
	Behaviours []*behaviour.Behaviour
}

//Simulate simulates a single tick of a specific entity
func (e *Entity) Simulate(w *interface{}) {
	for behaviour := range e.Behaviours {
		e.Behaviours[behaviour].Action(w)
	}
}

//Humid represents an entity that has a humidity value
type Humid interface {
	GetHumidity() float32
}
