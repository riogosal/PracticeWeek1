package models

import (
	"fmt"
	"time"
)

type Catalogue struct {
	ID       int
	Name     string
	Quantity int
	TimeStamp
}

func CreateCatalogue(name string, num int) Catalogue {
	Counter++
	return Catalogue{
		ID:       Counter,
		Name:     name,
		Quantity: num,
		TimeStamp: TimeStamp{
			FirstStock:    time.Now(),
			RecentStock:   time.Now(),
			NearestExpiry: time.Now().Add(time.Hour * 24 * 365),
		},
	}
}

func (item Catalogue) Print() {
	fmt.Println("ID : ", item.ID)
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

func (item *Catalogue) Add(number int) {
	item.Quantity += number
	item.RecentStock = time.Now()
}
func (item *Catalogue) Take(number int) {
	item.Quantity -= number
	item.RecentStock = time.Now()
}
