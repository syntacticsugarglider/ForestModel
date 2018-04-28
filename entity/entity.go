package entity

import "ForestModel/behaviour"

//GenericInterface is a generic representational interface with function/type members
type GenericInterface interface {
	GetBehaviours() map[string]behaviour.Behaviour
}

//Entity is the reified object with GenericInterface-described members
type Entity struct {
	behaviours map[string]behaviour.Behaviour
	GenericInterface
}

//GetBehaviours returns the behaviours performed by the parent entity
func (e *Entity) GetBehaviours() map[string]behaviour.Behaviour {
	return e.behaviours
}

//Soil represents the soil entity with nutrient levels, moisture, temp, etc.
type Soil struct {
	Levels      map[string]float32
	Temperature float32
	Moisture    float32
}
