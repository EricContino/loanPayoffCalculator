package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    origLoanAmtPtr := flag.Float64("origLoanAmt", 0, "Original Loan Amount, required")
    interestPtr := flag.Float64("interest", 0, "Loan Interest Rate, required")
    origLoanTermPtr := flag.Int("loanTerm", 360, "Loan Term in months")
    lumpSumPtr := flag.Float64("lumpSum", 0, "An additional lump sum paid")
    monthlyExtraPtr := flag.Float64("monthlyExtra", 0, "Extra monthly principle payment")
    startMonthPtr := flag.Int("startMonth", 1, "Month of first payment(1-12)")
    startYearPtr := flag.Int("startYear", 2025, "Year of first payment")

    flag.Parse()

    fmt.Println("Mortgage Payoff Calculations coming soon...")
    
    err := validateInput(*origLoanAmtPtr, *interestPtr, *origLoanTermPtr, *lumpSumPtr, *monthlyExtraPtr, *startMonthPtr)
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }


    fmt.Println(*origLoanAmtPtr)
    fmt.Println(*interestPtr)
    fmt.Println(*origLoanTermPtr)
    fmt.Println(*lumpSumPtr)
    fmt.Println(*monthlyExtraPtr)
}

func validateInput(origLoanAmt float64, interest float64, origLoanTerm int, lumpSum float64, monthlyExtra float64, startMonth int) error {
    if origLoanAmt <= 0 {
        return fmt.Errorf("Invalid original loan amount: %v, must be greater than 0", origLoanAmt) 
    }

    if interest < 0 {
        return fmt.Errorf("Invalid interest rate: %v, must be 0 or greater", interest)
    }

    if origLoanTerm < 1 {
        return fmt.Errorf("Invalid loan term length: %v, must be 1 or more months", origLoanTerm)
    }

    if lumpSum < 0 {
        return fmt.Errorf("Invalid lump sum payment: %v, must be a positive value", lumpSum)
    }

    if monthlyExtra < 0 {
        return fmt.Errorf("Invalid monthly extra principle: %v, must be a positive value", monthlyExtra)
    }

    if startMonth < 1 || startMonth > 12 {
        return fmt.Errorf("Invalid month: %v, must be between 1-12")
    }
    
    return nil
}
