package helper

import (
	"context"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-protos-go-apiv2/gateway"
	"google.golang.org/grpc/status"
	"log"
)

func LedgerErrorHandler(contract *client.Contract, err error) error {
	switch err := err.(type) {
	case *client.EndorseError:
		log.Printf("Endorse error with gRPC status %v: %s\n", status.Code(err), err)
		return convertError(err)
	case *client.SubmitError:
		log.Printf("Submit error with gRPC status %v: %s\n", status.Code(err), err)
		return convertError(err)
	case *client.CommitStatusError:
		if errors.Is(err, context.DeadlineExceeded) {
			log.Printf("Timeout waiting for transaction %s commit status: %s\n", err.TransactionID, err)
			return convertError(err)
		} else {
			log.Printf("Error obtaining commit status with gRPC status %v: %s\n", status.Code(err), err)
			return convertError(err)
		}
	case *client.CommitError:
		log.Printf("Transaction %s failed to commit with status %d: %s\n", err.TransactionID, int32(err.Code), err)
		return convertError(err)
	}

	return nil
}

// Any error that originates from a peer or orderer node external to the gateway will have its details
// embedded within the gRPC status error. The following code shows how to extract that.
func convertError(err error) error {
	statusErr := status.Convert(err)
	for _, detail := range statusErr.Details() {
		switch detail := detail.(type) {
		case *gateway.ErrorDetail:
			log.Printf("Error from endpoint: %s, mspId: %s, message: %s\n", detail.Address, detail.MspId, detail.Message)
			return fmt.Errorf(detail.Message)
		}
	}
	return nil
}
