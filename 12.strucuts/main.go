package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// create a function to take user input
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	// remove any white space
	input = strings.TrimSpace(input)
	return input, err
}

func createBill() bill {
	// take input
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create bill for: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill for -", b.name)

	return b
}

func addItem() {

}

func promptOptions(b *bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Your Choice: (a: add item) (s: save bill) (t: add tip) - ", reader)

	switch opt {
	case "a":
		fmt.Println("you have selected adding item")
		name, _ := getInput("Item Name - ", reader)
		price, _ := getInput("Item Price - ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("price is not a number..")
			promptOptions(b)
		}

		b.addItem(name, p)
		fmt.Println("Item added Successfully")
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you have selected saving bill")
	case "t":
		tip, _ := getInput("Enter tip - ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Tip must be a number..")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip updated")
		promptOptions(b)
	default:
		fmt.Println("not a valid option")
		promptOptions(b)

	}
}

func main() {
	// make new bills
	myBill := createBill()
	promptOptions(&myBill)
	fmt.Println(myBill)

}
