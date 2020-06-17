package credit

import "math"

//ФУНКЦИЯ РАСЧЕТА АННУИТЕТА
//Принимаемые параметры - сумма кредита в копейках, срок кредита в месяцах, ставка по кредиту в % годовых
//Возвращаемые значения - ежемесячный аннуитетный платеж, переплата по кредиту, общая выплата по кредиту (все в копейках)
func Calculate (loanAmount int64, loanPeriod, loanPercent int) (monthlyPayment, overPayment, totalPayment int64) {

	var percentageMonthly, annuityRate float64

	//Расчет месячной процентной ставки
	percentageMonthly = float64(loanPercent) / 12 / 100
	percentageMonthly = math.Round(percentageMonthly*10000) / 10000 //Округление месячной ставки до 3-го знака

	//Расчет коэффициента аннуитета
	annuityRate = percentageMonthly * math.Pow((percentageMonthly + 1), float64(loanPeriod)) / (math.Pow((percentageMonthly + 1), float64(loanPeriod)) - 1)
	annuityRate = math.Round(annuityRate * 1_000_000) / 1_000_000 //Округление коэффиента до 6 знака

	//Расчет суммы ежемесячного аннуитетного платежа
	monthlyPayment = int64(float64(loanAmount) * annuityRate)

	//Расчет общей выплаты по кредиту
	totalPayment = int64(loanPeriod) * monthlyPayment

	//Расчет переплаты по кредиту
	overPayment = totalPayment - loanAmount

	return
}