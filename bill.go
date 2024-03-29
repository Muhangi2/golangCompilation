package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill {
		name:  name,
		items: map[string]float64{"jabula": 566.99, "cake": 400},
		tip:   0,
	}
	return b
}

// format the types receriver functions
func (b bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0
	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}
	//adding tip
	fs += fmt.Sprintf("%-25v...$%v\n", "tip:", b.tip)
	//total
	fs += fmt.Sprintf("%v....$%0.2f", "total:", total+b.tip)
	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
//save bill
func (b *bill) save(){
data:=[]byte(b.format())

err:=os.WriteFile("billsss/"+b.name+".txt",data,0644)
if err != nil {
	panic(err)
}
fmt.Println("bill was saved on the file")
}