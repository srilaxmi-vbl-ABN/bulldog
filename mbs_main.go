
package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type loan struct {
	loanAccNo int
	loanAmt int
	loanRate int
	loanTerm int
	propertyNo string
	borrowerName string
	borrowerBSN string
	borrowerCreditRating int
}

type tranche struct {
	trancheId int
	trancheRating int
	trancheRate int	
	loans []loan //array of loans

}

type trancheRatingList struct {
	list [3]string //harcoded list of values
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var loan A1, A2 //loans 
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	//int i: = 1
	//for i<=2
	//{
	//	i = i + 1	
	//}
	
	A1 = loan{loanAccNo: 101, loanAmt: 100000, loanRate: 5, loanTerm: 10, propertyNo "HSR1001", borrowerName "R. Anndurai", borrowerBSN: "000001",	borrowerCreditRating: 8}
	A2 = loan{loanAccNo: 102, loanAmt: 200000, loanRate: 4, loanTerm: 10, propertyNo "HSR1002", borrowerName "R. Murthy", borrowerBSN: "000002",	borrowerCreditRating: 5}
	
	fmt.Printf("loanA1No = %d, loanA2No = %d\n", A1.loanAccNo, A2.loanAccNo)

	// Write the state to the ledger
	err = stub.PutState(A1, []byte(strconv.Itoa(A1.loanAccNo)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(A2, []byte(strconv.Itoa(A2.loanAccNo)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


	
