package main

import (
	"fmt"
	"os"
)

// in go -> struct is used
// struct is a blueprint that descibes data
// WE dont use class in GO

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	// create a bill structure
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// receive bill into this function
// limiting to call only by bill object
// myBill.format() ✅
// format() ❌

// receiver function -> also is a copy of obj(bill)
// basically methodsss
func (b *bill) format() string {
	fs := "Your Bill Breakdown: \n\n"
	var total float64 = 0

	for k, v := range b.items {
		// -25 leaves 25 char spaces on the right
		fs += fmt.Sprintf("%-25v ...%v Rupees\n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("\n%-25v ...%0.2f Rupees", "Tip:", b.tip)
	fs += fmt.Sprintf("\n%-25v ...%0.2f Rupees", "Total:", total+b.tip)
	return fs

}

// we are passing pointer because we want it to update
// this takes less space as well
func (b *bill) updateTip(tip float64) {
	// (*b).tip = tip
	b.tip = tip // automatically deferences struct
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// this saves the file
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
}
