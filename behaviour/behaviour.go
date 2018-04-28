package behaviour

//GenericInterface represents a generic active characteristic of an entity
type GenericInterface interface {
	getIteratedOperations() []func(operands ...(*int))
}

//Behaviour is the reified object that stores data pertaining to the action of system subcomponents and their actionable systems themselves
type Behaviour struct {
	Probability float32
	Action      func(*interface{})
}

//MoistureDispersal is a behaviour that controls equalization of moisture between lateral environments
var MoistureDispersal = &Behaviour{
	Probability: 1,
	Action: func(w *interface{}) {

	},
}
