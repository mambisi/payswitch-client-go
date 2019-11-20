# GO Client for [TheTeller](https://www.theteller.net) by PaySwitch

Official Document of the api can be found [here](https://www.theteller.net/documentation)


## Usage
```go
import payswitch "mambisi/payswitch-client-go"
```

```go
client  :=  payswitch.NewApiClient("API_USERNAME","API_KEY", false /* set to true if in live mode */)
``` 

**Example card payment**


```go
pUsername := "API_USERNAME"
pApiKey := "API_KEY"
psc := payswitch.NewApiClient(pUsername,pApiKey,false)

req := payswitch.CardPaymentRequest{}
req.RSwitch = "VIS"
req.ProcessingCode = "000000"
req.TransactionID = "000000000035"
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
	log.Fatal(err)
}

if res.Code != "000"{
	log.Fatal(res)
}
```

**Example Mobile Wallet payment** 

```go
pUsername := "API_USERNAME"
pApiKey := "API_KEY"
psc := payswitch.NewApiClient(pUsername,pApiKey,false)

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
	log.Fatal(err)
}
if res.Code != "000" {
	log.Fatal(res)
}
```