package test

import (
	"fmt"
	"github.com/BoltApp/sleet"
	"github.com/BoltApp/sleet/gateways/authorize_net"
	"testing"
)

func TestAuthNet(t *testing.T) {
	client := authorize_net.NewClient(getEnv("AUTH_NET_LOGIN_ID"), getEnv("AUTH_NET_TXN_KEY"))
	authRequest := baseAuthorizationRequest()
	resp, err := client.Authorize(authRequest)
	fmt.Printf("resp: [%+v] err [%s]\n", resp, err)

	capResp, err := client.Capture(&sleet.CaptureRequest{
		Amount:               authRequest.Amount,
		TransactionReference: resp.TransactionReference,
	})
	fmt.Printf("capResp: [%+v] err [%s]\n", capResp, err)

	lastFour := authRequest.CreditCard.Number[len(authRequest.CreditCard.Number)-4:]
	options := make(map[string]interface{})
	options["credit_card"] = lastFour
	refundResp, err := client.Refund(&sleet.RefundRequest{
		Amount:               authRequest.Amount,
		TransactionReference: resp.TransactionReference,
		Options:              options,
	})
	fmt.Printf("refundResp: [%+v] err [%s]\n", refundResp, err)
}