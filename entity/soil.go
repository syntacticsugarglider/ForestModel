package entity

import "ForestModel/behaviour"

//Soil represents the soil entity with nutrient levels, moisture, temp, etc.
type Soil struct {
	Levels      map[string]float32
	Temperature float32
	Humidity    float32
	Entity
}

//NewSoil creates a new configured soil entity
func NewSoil() *Soil {
	soil := Soil{
		Levels:      make(map[string]float32),
		Temperature: 0,
		Humidity:    0,
	}
	soil.Behaviours = append(soil.Behaviours, behaviour.MoistureDispersal)
	return &soil
}
