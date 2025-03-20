package money

import (
    "math"
)

func TotalMonthlyPayment(loanAmt float64, monthlyInterest float64, numPayments int) float64{
    i := math.Pow((1+monthlyInterest), float64(numPayments))
    totalMonthlyPayment := loanAmt * ((monthlyInterest*i)/(i-1))
    return totalMonthlyPayment
}

func PrinciplePayment(totalMonthlyPayment float64, outstandingBalance float64, monthlyInterest float64) float64 {
    interestPayment := outstandingBalance * monthlyInterest
    return totalMonthlyPayment - interestPayment
}

func Payoff(loanAmt float64, monthlyInt float64, totalMonthly float64, lumpSum float64, extraMonthly float64) (float64, int) {
    var numPayments int = 0
    var intPaid float64 = 0
    
    loanAmt -= lumpSum
    for loanAmt > 0 {
        principle := PrinciplePayment(totalMonthly, loanAmt, monthlyInt)
        intPaid += (totalMonthly - principle)
        if loanAmt > principle {
            loanAmt = loanAmt - principle

            // I'm assuming extra monthly principle payments happen after regular payment
            if loanAmt > extraMonthly {
                loanAmt -= extraMonthly
            } else {
                loanAmt = 0
            }
        } else {
            loanAmt = 0
        }
        numPayments++
    }

    return intPaid, numPayments
}
