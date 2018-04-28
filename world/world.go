package world

import (
	"ForestModel/entity"
	"ForestModel/util"
	"errors"
)

//GenericInterface represents the operable outlets/inlets of a world
type GenericInterface interface {
}

//World represents the reified datastructure of the world itself
type World struct {
	GenericInterface
	Entities []entity.Entity
	Cells    map[util.Point]Cell
	size     util.Rect
}

//Cell represents a containing subsection of forest
type Cell struct {
	util.Rect
	Strata    []Stratum
	Substrate entity.Soil
}

//Stratum represents the properties of one of the stratified layers of biomedia within the modelled biome
type Stratum struct {
	Humidity    int
	Temperature int
	LightLevel  int
}

//InitializeWorld allocates and generates a simulated world with the specified properties
func InitializeWorld(size util.Rect, cellSize int) (World, error) {
	world := World{}
	world.size = size
	if size.Width%cellSize != 0 || size.Height%cellSize != 0 {
		return World{}, errors.New("World size not evenly divisible by subcell size")
	}
	for x := 0; x <= size.Width; x += cellSize {
		for y := 0; y <= size.Height; y += cellSize {
			world.Cells[util.Point{int32(x), int32(y)}] = Cell{
				Substrate: entity.Soil{
					Temperature: 0,
					Moisture:    0,
					Levels:      map[string]float32{},
				},
				Strata: []Stratum{},
			}
		}
	}
	return world, nil
}
