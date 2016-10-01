
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

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


	
