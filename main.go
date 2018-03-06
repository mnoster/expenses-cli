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

func getMonthlyExpenses() []int {

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

	s []int = [rent, carInsurance, rentersInsurance, gifts, travel, healthInsurance]
	fmt.Println(s)
	return s

}


type weeklyExpenses struct{
	gas int
	drinks int
	groceries int
}

func getWeeklyExpenses() weeklyExpenses{
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

	s := weeklyExpenses{gas, drinks, groceries}
// 	sp := &s
	fmt.Println(s)
	return s

}

type dailyExpenses struct{
	lunch int
	dinner int
}

func getDailyExpenses() dailyExpenses{

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

	s := dailyExpenses{lunch, dinner}
	fmt.Println(s)
	return s

}

func calculateAverage(){
// Monthly * 12
	monthly := getMonthlyExpenses()
	j := 0
	for i := 0; i < 6 ; i++ {
		j += monthly[i]

	}
	fmt.Println(j)
// Weekly * 52
// Daily * 365

// Total / 365
}

func main(){
	calculateAverage()
	//	getMonthlyExpenses()
	getWeeklyExpenses()
	getDailyExpenses()

}


