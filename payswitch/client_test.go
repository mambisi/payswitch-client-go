package payswitch

import (
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)
/*
PAYSWITCH_USERNAME
PAYSWITCH_API_KEY
 */
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

	psc := NewApiClient(pUsername, pApiKey, false)

	req := MobileMoneyPaymentRequest{}
	req.RSwitch = "MTN"
	req.ProcessingCode = "000200"
	req.TransactionID = generateTransactionID()
	req.MerchantID = "TTM-00000865"
	req.Amount = "000000000200"
	req.Desc = "Test Payment"
	req.SubscriberNumber = "0550000002"

	res, err := psc.PaymentService.ProcessMobileMoneyPayment(&req)

	if err != nil {
		t.Fatal(err)
	}

	if res.Code != "000" {
		t.Fatal(res)
	}
}

func generateTransactionID() string {
	rand.Seed(time.Now().UnixNano())
	k := rand.Intn(9999)
	n := strconv.Itoa(k)

	z := 12 - len(n)

	t := ""

	for i := 0; i < z; i++ {
		t += "0"
	}

	t += n

	log.Print("Generated Transaction ID: ", t)

	return t
}
