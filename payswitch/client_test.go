package payswitch

import (
	"github.com/joho/godotenv"
	"github.com/mambisi/payswitch-client-go/payswitch/helper"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestPaymentService_ProcessCardPayment(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	pUsername := os.Getenv("PAYSWITCH_USERNAME")
	pApiKey := os.Getenv("PAYSWITCH_API_KEY")
	psc := NewApiClient(pUsername, pApiKey, false)

	req := CardPaymentRequest{}
	req.RSwitch = "VIS"
	req.ProcessingCode = "000000"
	req.TransactionID = generateTransactionID()
	req.MerchantID = "TTM-00000865"
	req.Amount = "000000000100"
	req.Pan = "4448366600675430"
	req.Cvv = "330"
	req.ExpMonth = "02"
	req.ExpYear = "21"
	req.CardHolder = "Mambisi Zempare"
	req.CustomerEmail = "test@email.com"
	req.Desc = "Test Payment"
	req.Currency = "GHS"
	req.RedirectUrl = "https://hubitcloud.com"

	res, err := psc.PaymentService.ProcessCardPayment(&req)

	if err != nil {
		t.Fatal(err)
	}

	if res.Code != "000" {
		t.Fatal(res)
	}

}

func TestPaymentService_ProcessMobileMoneyPayment(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	pUsername := os.Getenv("PAYSWITCH_USERNAME")
	pApiKey := os.Getenv("PAYSWITCH_API_KEY")
	pMID := os.Getenv("PAYSWITCH_M_ID")

	psc := NewApiClient(pUsername, pApiKey, true)

	req := MobileMoneyPaymentRequest{}
	req.RSwitch = "MTN"
	req.ProcessingCode = "000200"
	req.TransactionID = generateTransactionID()
	req.MerchantID = pMID
	req.Amount = "000000000200"
	req.Desc = "Test Payment"
	req.SubscriberNumber = "024000001"

	res, err := psc.PaymentService.ProcessMobileMoneyPayment(&req)

	if err != nil {
		t.Fatal(err)
	}

	if res.Code != "000" {
		t.Fatal(res)
	}
}

func TestVerificationService_VerifyTransaction(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	pUsername := os.Getenv("PAYSWITCH_USERNAME")
	pApiKey := os.Getenv("PAYSWITCH_API_KEY")
	pMID := os.Getenv("PAYSWITCH_M_ID")

	psc := NewApiClient(pUsername, pApiKey, false)

	amount, err := helper.ConvT12DigitAmount(3.25)

	if err != nil {
		t.Fatal(err)
	}

	req := CardPaymentRequest{}
	req.RSwitch = "VIS"
	req.ProcessingCode = "000000"
	req.TransactionID = generateTransactionID()
	req.MerchantID = pMID
	req.Amount = amount
	req.Pan = "4448366600675430"
	req.Cvv = "330"
	req.ExpMonth = "02"
	req.ExpYear = "21"
	req.CardHolder = "Mambisi Zempare"
	req.CustomerEmail = "test@email.com"
	req.Desc = "Test Payment"
	req.Currency = "GHS"
	req.RedirectUrl = "https://hubitcloud.com"

	res, err := psc.PaymentService.ProcessCardPayment(&req)

	if err != nil {
		t.Fatal(err)
	}

	if res.Code != "000" {
		t.Fatal(res)
	}

	p, err := psc.VerificationService.VerifyTransaction(req.TransactionID, req.MerchantID)

	if err != nil {
		t.Fatal(err)
	}

	if p.Code != "000" {
		t.Fatal(res)
	}

	if p.TransactionID != res.TransactionID && p.RSwitch != req.RSwitch {
		t.Fatal(res)
	}

}

func generateTransactionID() string {
	rand.Seed(time.Now().UnixNano())
	k := rand.Intn(9999)
	t, _ := helper.ConvT12DigitTransactionID(k)
	return t
}
