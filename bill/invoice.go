package main 

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
)

func getInput(prompt string, r *bufio.Reader)(string, error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createInvoice() bill{
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Create invoice name : ")
	// name, _ := reader.readString('\n')
	// name = strings.trimSpace(name)

	name, _ := getInput("Create invoice name : ", reader)

	a := newBill(name)
	fmt.Println("Create the invoice - ", a.name)

	return a
}

func promptOption(a bill){
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Option (a = add item, b = save bill, c = add tip) : ", reader)
	// fmt.Println(opt)

	// Switch case
	switch opt{
	case "a":
		name, _ := getInput("Item name : ", reader)
		price, _ := getInput("Item price : ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil{
			fmt.Println("The price must be a number")
			promptOption(a)
		}
		a.addItem(name, p)
		
		fmt.Println("Item added - ", name, price)
		promptOption(a)
	case "b":
		a.save()
		fmt.Println("Saved the file - ", a.name)
	case "c":
		tip, _ := getInput("Enter tip : ", reader)
		
		c, err := strconv.ParseFloat(tip, 64)
		if err != nil{
			fmt.Println("The price must be a number")
			promptOption(a)
		}
		a.upTip(c)

		fmt.Println("Tip added - ", tip)
		promptOption(a)
	default:
		fmt.Println("That's not Valid")
		promptOption(a)
	}
}

func main(){
	myInvoice := createInvoice()
	promptOption(myInvoice)


}