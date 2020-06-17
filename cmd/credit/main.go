package main

import (
	"fmt"
	"github.com/vl-mobitutor/Netology_GO_Task2_1/pkg/credit"
)

func main() {
	//Исходные данные для расчета параметров аннуитета
	myLoanAmount := int64(1_000_000_00) //Сумма кредита в копейках
	myLoanPeriod := 36 //Срок кредита в месяцах
	myLoanPercent := 20 //Процентная ставка по кредиту в % годовых


	myMonthlyPayment, myOverPayment, myTotalPayment := credit.Calculate(myLoanAmount, myLoanPeriod, myLoanPercent)

	//Вывод результатов расчета параметров аннуитета
	fmt.Printf("Ежемесячныйй аннуитетный платеж: %d копеек \n", myMonthlyPayment)
	fmt.Printf("Переплата по кредиту: %d копеек \n", myOverPayment)
	fmt.Printf("Полная выплата по кредиту: %d копеек \n", myTotalPayment)

}

