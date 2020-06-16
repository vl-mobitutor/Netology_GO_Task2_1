package main
import (
	"fmt"
	"math"
)

func main() {
	annuityCalc(1_000_000_00, 36, 20)
}


//ФУНКЦИЯ РАСЧЕТА АННУИТЕТА
//Принимаемые параметры - сумма кредита в копейках, срок кредита в месяцах, ставка по кредиту в % годовых
//Возвращаемые значения - ежемесячный аннуитетный платеж, переплата по кредиту, общая выплата по кредиту (все в копейках)
func annuityCalc (loanSum, period, percent int64) int64 {

	var (
		percentageMonthly, annuityRate float64
		monthlyPayment, totalPayment, overPayment int64
	)

	//Расчет месячной процентной ставки
	percentageMonthly = float64(percent) / 12 / 100
	percentageMonthly = math.Round(percentageMonthly*10000) / 10000 //Округление месячной ставки до 3-го знака

	//Расчет коэффициента аннуитета
	annuityRate = percentageMonthly * math.Pow((percentageMonthly + 1), float64(period)) / (math.Pow((percentageMonthly + 1), float64(period)) - 1)
	annuityRate = math.Round(annuityRate * 1_000_000) / 1_000_000 //Округление коэффиента до 6 знака

	//Расчет суммы ежемесячного аннуитетного платежа
	monthlyPayment = int64(float64(loanSum) * annuityRate)

	//Расчет общей выплаты по кредиту
	totalPayment = period * monthlyPayment

	//Расчет переплаты по кредиту
	overPayment = totalPayment - loanSum

	fmt.Println(percentageMonthly, annuityRate, monthlyPayment, overPayment, totalPayment)

	return 1
}
