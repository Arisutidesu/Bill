package main 

import (
	"fmt"
	"os"
)

type bill struct {
	name string
	items map[string]float64
	tip float64
}

func newBill(name string) bill{
	a := bill{
		name : name,
		items : map[string]float64{},
		tip : 0,
	}

	return a
}

func (a bill) format() string {
	b := "Bill breakdown : \n"
	var total float64 = 0

	for k, v := range a.items{
		b += fmt.Sprintf("%v  %v K \n", k+":", v)
		total += v
	}

	b += fmt.Sprintf("%v  %v K", "Tip:", a.tip)

	b += fmt.Sprintf("%v  %v K", "\nTotal:", a.tip + total)
	return b
}

func (a *bill) upTip(tip float64){
	a.tip = tip
}

func (a *bill) addItem(name string, price float64){
	a.items[name] = price
}

func (a *bill) save(){
	data := []byte(a.format())

	err := os.WriteFile("bill/"+a.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved")
}
