package behaviour

//GenericInterface represents a generic active characteristic of an entity
type GenericInterface interface {
	getIteratedOperations() []func(operands ...(*int))
}

//Behaviour is the reified object that stores data pertaining to the action of system subcomponents and their actionable systems themselves
type Behaviour struct {
	GenericInterface
}
