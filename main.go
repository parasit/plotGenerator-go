package main

import (
	"fmt"

	rpgforge "github.com/parasit/rpgforge-go"
)

func main() {
	var dt DataTable
	fmt.Println("Plot Generator")
	dt.AddValue("Miecz", 1)
	dt.AddValue("Topór", 1)
	dt.AddValue("Sztylet", 3)
	dt.AddValue("Pałka", 1)
	fmt.Println(dt.GetSize())
	dt.DumpData()
	fmt.Println(dt.GetElementForValue(3))
	roll := rpgforge.ThrowDices("1d6").Sum()
	fmt.Println(roll)
	fmt.Println(dt.GetElementForValue(roll))

}
