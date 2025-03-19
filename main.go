package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    origLoanAmtPtr := flag.Float64("origLoanAmt", 0, "Original Loan Amount, required")
    interestPtr := flag.Float64("interest", 0, "Loan Interest Rate, required")
    origLoanTermPtr := flag.Int("loanTerm", 30, "Loan Term in years")
    lumpSumPtr := flag.Float64("lumpSum", 0, "An additional lump sum paid")
    monthlyExtraPtr := flag.Float64("monthlyExtra", 0, "Extra monthly principle payment")

    flag.Parse()

    fmt.Println("Mortgage Payoff Calculations coming soon...")
    
    err := validateInput(*origLoanAmtPtr, *interestPtr)
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

func validateInput(origLoanAmt float64, interest float64) error {
    if origLoanAmt <= 0 {
        return fmt.Errorf("Invalid original loan amount: %v, must be greater than 0", origLoanAmt) 
    }

    if interest < 0 {
        return fmt.Errorf("Invalid interest rate: %v, must be 0 or greater", interest)
    }
    
    return nil
}
