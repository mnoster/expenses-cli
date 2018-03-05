package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){

	type monthlyExpenses struct{
		rent int
		carInsurance int
		rentersInsurance int
		gifts int
		travel int
		healthInsurance int
	}
	type WeeklyExpenses struct{
		gas int
		drinks int
		groceries int
	}
	type DailyExpenses struct{
		lunch int
		dinner int
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter yearly income after taxes: ")
	input, _  := reader.ReadString('\n')
	format := strings.Split(input, "\n")[0]
	income, _ := strconv.Atoi(format)
	fmt.Println("Income set: ", income)
	fmt.Println("")

	fmt.Println("-----Monthly Expenses----")
	fmt.Println("")

	fmt.Println("Enter Rent: ")
	input1, _  := reader.ReadString('\n')
	format1 := strings.Split(input1, "\n")[0]
	rent, _ := strconv.Atoi(format1)

	fmt.Println("Enter Car Insurance: ")
	input2, _  := reader.ReadString('\n')
	format2 := strings.Split(input2, "\n")[0]
	carInsurance, _ := strconv.Atoi(format2)

	fmt.Println("Enter Renters Insurance: ")
	input3, _  := reader.ReadString('\n')
	format3 := strings.Split(input3, "\n")[0]
	rentersInsurance, _ := strconv.Atoi(format3)

	fmt.Println("Enter Gifts:  ")
	input4, _  := reader.ReadString('\n')
	format4 := strings.Split(input4, "\n")[0]
	gifts, _ := strconv.Atoi(format4)

	fmt.Println("Enter Travel:  ")
	input5, _  := reader.ReadString('\n')
	format5 := strings.Split(input5, "\n")[0]
	travel, _ := strconv.Atoi(format5)
	
	fmt.Println("Enter Health Insurance: ")
	input6, _  := reader.ReadString('\n')
	format6 := strings.Split(input6, "\n")[0]
	healthInsurance, _ := strconv.Atoi(format6)



	// Assign data back to pointed reference of expense struct

	s := monthlyExpenses{rent, carInsurance, rentersInsurance, gifts, travel, healthInsurance}
	sp := &s
	sp.rent = rent
	fmt.Println(s)

}


