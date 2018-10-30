package main

import (
	"fmt"
)

func main() {
	var dt DataTable
	// dt.dataValues = make([0]DataItem)
	fmt.Println("Plot Generator")
	dt.AddValue("Miecz", 1)
	dt.AddValue("Topór", 1)
	dt.AddValue("Sztylet", 3)
	dt.AddValue("Pałka", 1)
	dt.ReadFromFile("weapons.rdf")
	fmt.Println(dt.GetSize())
	dt.DumpData()
	fmt.Println(dt.GetElementAtValue(1))

}
