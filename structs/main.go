package main

import "log"

type location struct {
	longitude float64
	latitude  float64
}

func main() {
	newYork := location{
		latitude:  40.73,
		longitude: -73.93,
	}

	newYork.changeLatitude()

	log.Println(newYork)
}

func (lo *location) changeLatitude() {
	lo.latitude = 41.0
}
