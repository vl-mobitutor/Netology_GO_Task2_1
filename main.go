package main
import (
	"fmt"
	"math"
)
func main() {
	anuitetCalc(1_000_000_00, 36, 20)
}
func anuitetCalc (loanSum, period, percent int64) int64 {
	var percentageMonthly, anuitetRate float64
	percentageMonthly = float64(percent) / 12 / 100
	percentageMonthly = math.Round(percentageMonthly*10000) / 10000
	fmt.Println(percentageMonthly)
	anuitetRate = percentageMonthly * math.Pow((percentageMonthly + 1), float64(period)) / (math.Pow((percentageMonthly + 1), float64(period)) - 1)
	anuitetRate = math.Round(anuitetRate * 1_000_000) / 1_000_000
	fmt.Println(anuitetRate)
	return 1
}
