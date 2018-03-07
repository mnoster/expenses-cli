package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type MonthlyExpenses struct{
	rent int
	carInsurance int
	rentersInsurance int
	gifts int
	travel int
	healthInsurance int
}

//type Expenses interface{
//	loadExpensePrompts()
//	getMonthlyExpenses()
//	getWeeklyExpenses()
//	getDailyExpenses()
//}

func getMonthlyExpenses() [6]int {

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

	s := [6]int{rent, carInsurance, rentersInsurance, gifts, travel, healthInsurance}
	fmt.Println(s)
	return s

}


type weeklyExpenses struct{
	gas int
	drinks int
	groceries int
}

func getWeeklyExpenses() [3]int{
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("-----Weekly Expenses----")
	fmt.Println("")

	fmt.Println("Enter Gas Expense: ")
	input1, _  := reader.ReadString('\n')
	format1 := strings.Split(input1, "\n")[0]
	gas, _ := strconv.Atoi(format1)

	fmt.Println("Enter Drinks Expense: ")
	input2, _  := reader.ReadString('\n')
	format2 := strings.Split(input2, "\n")[0]
	drinks, _ := strconv.Atoi(format2)

	fmt.Println("Enter Groceries Expense: ")
	input3, _  := reader.ReadString('\n')
	format3 := strings.Split(input3, "\n")[0]
	groceries, _ := strconv.Atoi(format3)

	// Assign data back to pointed reference of expense struct

	s := [3]int{gas, drinks, groceries}
	fmt.Println(s)
	return s

}

type dailyExpenses struct{
	lunch int
	dinner int
}

func getDailyExpenses() [3]int{

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("-----Daily Expenses----")
	fmt.Println("")

	fmt.Println("Enter Lunch Expense: ")
	input1, _  := reader.ReadString('\n')
	format1 := strings.Split(input1, "\n")[0]
	lunch, _ := strconv.Atoi(format1)

	fmt.Println("Enter Dinner Expense: ")
	input2, _  := reader.ReadString('\n')
	format2 := strings.Split(input2, "\n")[0]
	dinner, _ := strconv.Atoi(format2)

	// Assign data back to pointed reference of expense struct

	s := [3]int{lunch, dinner}
	fmt.Println(s)
	return s

}

func calculateAverage(){
// Monthly * 12
	monthly := getMonthlyExpenses()
	m := 0
	for i := 0; i < len(monthly) ; i++ {
		m += monthly[i]

	}
	fmt.Println("TOTAL MONTHLY: " , m)
// Weekly * 52
	weekly := getWeeklyExpenses()
	w := 0
	for t := 0; t < len(weekly) ; t++ {
		w += weekly[t] *  3

	}
	fmt.Println("TOTAL WEEKLY: " ,w)

// Daily * 365
	daily := getDailyExpenses()
	d := 0
	for n := 0; n < len(daily) ; n++ {
		d += daily[n] * 30

	}
	fmt.Println("TOTAL DAILY: ", d) 

	total := (m + w + d) / 30
	fmt.Println("Total: ", total )
// Total / 365
}

func main(){
	calculateAverage()
}


