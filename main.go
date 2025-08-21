package main

import "fmt"

type Sensor struct {
	Name        string  //""
	Id          string  //""
	Temperature float32 //0
}

// struct receiver function
func (s *Sensor) convertToCelsius() {
	//convert to celsius
	celcius := (s.Temperature - 32) * 5 / 9
	fmt.Printf("Sensor Name: %s, ftoc: %f\n", s.Name, celcius)
}

func (s *Sensor) displayInfo() {
	fmt.Printf("Sensor Name: %s, Id: %s, Temperature: %f\n", s.Name, s.Id, s.Temperature)
}

func main() {
	//	sensor := new(Sensor)
	//	fmt.Printf("%v", sensor)
	// sensor.Name = "Temperature Sensor"
	// sensor.Id = "TS001"
	// sensor.Flow = 100
	sensor := &Sensor{
		Name:        "Temperature Sensor",
		Id:          "TS001",
		Temperature: 90, //farenheit, celcius
	}
	//pointer * value & address
	//sensor.convertToCelsius() //call the method on the struct
	sensor.displayInfo() // call the method to display info

}

// return,  void
//func convertToCelsiusReturn(sensor *Sensor) float32 {
//	return (sensor.Temperature - 32) * 5 / 9
//}
//
//func convertToCelsius(sensor *Sensor) {
//	celcius := (sensor.Temperature - 32) * 5 / 9
//
//	fmt.Printf("Sensor Name: %s\n, ftoc: %f", sensor.Name, celcius)
//
//}
