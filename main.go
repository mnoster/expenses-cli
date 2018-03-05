package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){

	const income uint = 65000

	type monthlyExpenses struct{
		rent int
		carInsurance int
		rentersInsurance int
		gifts int
		travel int
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
	fmt.Println("-----Monthly Expenses----")

	fmt.Print("Enter Rent: ")
	input, _  := reader.ReadString('\n')
	format := strings.Split(input, "\n")[0]
	rent, _ := strconv.Atoi(format)

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



	s := monthlyExpenses{rent, carInsurance, rentersInsurance, gifts, travel}
	sp := &s
	sp.rent = rent
	fmt.Println(s)

}


