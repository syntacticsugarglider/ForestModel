package world

import (
	"ForestModel/entity"
	"ForestModel/util"
	"errors"
)

//World represents the reified datastructure of the world itself
type World struct {
	Entities []*entity.Entity
	Cells    map[util.Point]*Cell
	size     util.Rect
}

//Cell represents a containing subsection of forest
type Cell struct {
	util.Rect
	Strata    []*Stratum
	Substrate *entity.Soil
	Contents  []entity.GenericInterface
	WPointer  *World
	Position  util.Point
}

//Stratum represents the properties of one of the stratified layers of biomedia within the modelled biome
type Stratum struct {
	Humidity    float32
	Temperature float32
	LightLevel  float32
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
			world.Cells = map[util.Point]*Cell{}
			world.Cells[util.Point{int32(x), int32(y)}] = &Cell{
				Substrate: entity.NewSoil(),
				Strata:    []*Stratum{},
				WPointer:  &world,
			}
		}
	}
	return world, nil
}

//SetHumidity assigns all cell substrates/strata to the given moisture/humidity value
func (w *World) SetHumidity(value float32) {
	for cell := range w.Cells {
		w.Cells[cell].Substrate.Humidity = value
		for stratum := range w.Cells[cell].Strata {
			w.Cells[cell].Strata[stratum].Humidity = value
		}
	}
}

//SetTemperature assigns all cell substrates/strata to the given temperature value
func (w *World) SetTemperature(value float32) {
	for cell := range w.Cells {
		w.Cells[cell].Substrate.Temperature = value
		for stratum := range w.Cells[cell].Strata {
			w.Cells[cell].Strata[stratum].Temperature = value
		}
	}
}

//Simulate models a certain number of seconds of world time
func (w *World) Simulate(ticks int) {
	temp := (func() map[util.Point]Cell {
		_map := map[util.Point]Cell{}
		for k, v := range w.Cells {
			_map[k] = *v
		}
		return _map
	})()
	for i := 0; i <= ticks; i++ {
		for cell := range temp {
			for entity := range temp[cell].Contents {
				temp[cell].Contents[entity].Simulate(w)
			}
		}
	}
	w.Cells = (func() map[util.Point]*Cell {
		_map := map[util.Point]*Cell{}
		for k, v := range w.Cells {
			_map[k] = v
		}
		return _map
	})()
}

//GetNeighbors returns the neighboring cells
func (c *Cell) GetNeighbors() {
	neighbors := map[util.Point]*Cell{}
}
