package usecase

import (
	"context"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
	"open-rev.com/helper"
)

type testUsecase struct {
	Contract *client.Contract
}

func (t testUsecase) TestVersion(context context.Context, contract client.Contract) (string, error) {

	log.Printf("Evaluate Transaction: TestVersion1")

	version, err := contract.EvaluateTransaction("TestVersion1")

	if err != nil {
		return "", helper.LedgerErrorHandler(&contract, err)
	}

	return string(version[:]), nil
}

type TestUsecase interface {
	TestVersion(context context.Context, contract client.Contract) (string, error)
}

func NewTestUsecase(contract *client.Contract) TestUsecase {
	return &testUsecase{Contract: contract}
}
