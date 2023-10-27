package models

import (
	"fmt"
	"time"
)

type Frozen struct {
	TypeID   int
	BatchID  int
	Name     string
	Quantity int
	TimeStamp
}

func CreateFrozen(name string, num int) Frozen {
	Counter++
	Counter2++
	return Frozen{
		TypeID:   Counter,
		BatchID:  Counter2,
		Name:     name,
		Quantity: num,
		TimeStamp: TimeStamp{
			FirstStock:    time.Now(),
			RecentStock:   time.Now(),
			NearestExpiry: time.Now().Add(time.Hour * 24 * 730),
		},
	}
}

// func (item Frozen) Add(number int) {
// 	item.Quantity += number
// 	item.RecentStock = time.Now()
// }
// func (item Frozen) Take(number int) {
// 	item.Quantity -= number
// 	item.RecentStock = time.Now()
// }

func (item *Frozen) Add(number int) {
	item.Quantity += number
	item.RecentStock = time.Now()
}
func (item *Frozen) Take(number int) {
	item.Quantity -= number
	item.RecentStock = time.Now()
}

func test() {
	frozenItem := CreateFrozen("Ice Cubes", 100)
	frozenItem.Add(1)

}

func (item Frozen) Print() {
	fmt.Println("Type ID : ", item.TypeID)
	fmt.Println("Batch ID : ", item.BatchID)
	fmt.Println("Name :", item.Name)
	fmt.Println("Quantity: ", item.Quantity)
	fmt.Printf("Last restock: %d-%d-%d %d:%d:%d\n",
		item.RecentStock.Year(),
		item.RecentStock.Month(),
		item.RecentStock.Day(),
		item.RecentStock.Hour(),
		item.RecentStock.Hour(),
		item.RecentStock.Second())
	println("====================================")
}
