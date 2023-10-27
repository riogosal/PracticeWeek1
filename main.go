package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"project1/models"
)

type Item interface {
	Add(quantityAdded int)
	Take(quantityTaken int)
	Print()
}

func ViewStorage(storage map[int]Item) {
	for _, value := range storage {
		value.Print()
	}
}

func HomePage() int {
	fmt.Print("Input action (1-view storage, 2-add item (Catalogue), 3-add item (Frozen), 4-restock item, 5-take item, 6-exit): ")
	var temp int
	fmt.Scan(&temp)
	return temp
}

func main() {
	fmt.Printf("Struct playground\n\n")
	scanner := bufio.NewScanner(os.Stdin)

	// bread1 := models.CreateCatalogue("Croissant", 10)
	bread2 := models.CreateCatalogue("Rotiboy", 8)
	// fish := models.CreateCatalogue("Salmon", 10)
	// vegs := models.CreateCatalogue("Selada", 30)
	frozenSalmon := models.CreateFrozen("Frozen Salmon", 20)

	var storage = map[int]Item{
		// bread1.ID: bread1,
		bread2.ID: &bread2,
		// fish.ID:   fish,
		// vegs.ID:   vegs,
		frozenSalmon.TypeID: &frozenSalmon,
	}

	condition := HomePage()
	var temp1, temp2 int
	for condition != 5 {
		if condition == 1 {
			ViewStorage(storage)
		} else if condition == 2 {
			fmt.Print("Item name: ")
			scanner.Scan()
			if err := scanner.Err(); err != nil {
				panic(err)
			}
			str := scanner.Text()
			fmt.Print("Item quantity: ")
			fmt.Scan(&temp1)
			new := models.CreateCatalogue(str, temp2)
			storage[new.ID] = &new
		} else if condition == 3 {
			fmt.Print("Input item ID: ")
			fmt.Scan(&temp1)
			fmt.Print("Number of items: ")
			fmt.Scan(&temp2)
			val := storage[temp1]
			val.Add(temp2)
			storage[temp1] = val
		} else if condition == 4 {
			fmt.Print("Input item ID: ")
			fmt.Scan(&temp1)
			fmt.Print("Number of items: ")
			fmt.Scan(&temp2)
			val := storage[temp1]
			val.Take(temp2)
			storage[temp1] = val
		}
		condition = HomePage()
	}
	jsonEncoded, err := json.Marshal(storage)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile("database.json", jsonEncoded, 0644); err != nil {
		panic(err)
	}
}

//Previous playground, kept for reference
/*
func Closure(testCase, num1, num2 int) func() {
	num := num1 * num2
	return func() { fmt.Printf("Closure test %d = %d\n", testCase, num) }
}*/
/*fmt.Printf("\nFunctions Playground\n\n")
init := func(pi, radius float32) float32 {
	area := pi * radius * radius
	return area
}(3.14, 5.00)

somethinSKetchy := func(num1 int) int {
	out := num1 * 4
	out++
	return out
}(44253)

fmt.Println("Test circle area = ", init)
fmt.Println("Function basic test = ", baseMath.Add(3, 4))
Closure(1, 222, 2)()
funcVar := Closure(2, 222, 3)
funcVar()
fmt.Println("千万不要搜索! ", somethinSKetchy)

fmt.Printf("\nArrays vs Maps 部分/Playground\n\n")
arr := [6]int{8, 5, 6, 3, 1, 2} //array size needs to be studied more: declaring a function that can receive an array without prior knowledge of the array size
fmt.Println("Unsorted array : ", arr)
fmt.Println("Sorted array : ", sorting.Selection(arr[:]))

var mp map[string]int = map[string]int{
	"完了":   3,
	"神经病":  666,
	"啥意思?": 9,
	"我爱你":  520,
	"我姓石, 无论何时, 与你相识我都值": 1,
	"海公牛": 8,
}
fmt.Println("Map print : ", mp)
search, okOrNotOk := mp["完了"]
fmt.Println("Checking for map variable availability: ", search, okOrNotOk)
search, okOrNotOk = mp["haha"]
fmt.Println("If doesn't exist : ", search, okOrNotOk)*/
