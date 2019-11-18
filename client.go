package payswitch

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
)

type ApiClient struct {
	baseUrl         string
	httpc           *resty.Client
	common          service
	PaymentService  *PaymentService
	CheckingService *CheckingService
}

/*

{
  "transaction_id":"000000000000",
  "status": "approved",
  "code": "000",
  "reason": "Transaction successful!"
}

*/
type ApiResult struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Code          string `json:"code"`
	Reason        string `json:"reason"`
}

type service struct {
	client *ApiClient
}

type PaymentService service
type CheckingService service

func NewApiClient(username string, key string, live bool) *ApiClient {

	h := resty.New().EnableTrace().SetBasicAuth(username, key)

	var baseUrl = "https://prod.theteller.net "

	if !live {
		baseUrl = "https://test.theteller.net"
	}

	c := &ApiClient{}
	c.httpc = h
	c.baseUrl = baseUrl
	c.common.client = c
	c.CheckingService = (*CheckingService)(&c.common)
	c.PaymentService = (*PaymentService)(&c.common)

	return c
}

func (p *PaymentService) ProcessCardPayment(m *CardPaymentRequest) (*ApiResult, error) {

	err := m.Validate()
	if err != nil {
		return nil, err
	}
	var client = p.client
	resp, err := client.httpc.R().SetBody(m).Post(client.baseUrl + "/v1.1/transaction/process")
	if err != nil {
		return nil, err
	}

	res := ApiResult{}

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, errors.New(string(resp.Body()))
	}

	return &res, nil
}

func (p *PaymentService) ProcessMobileMoneyPayment(m *MobileMoneyPaymentRequest) (*ApiResult, error) {

	err := m.Validate()
	if err != nil {
		return nil, err
	}

	var client = p.client
	resp, err := client.httpc.R().SetBody(m).Post(client.baseUrl + "/v1.1/transaction/process")
	if err != nil {
		return nil, err
	}

	res := ApiResult{}

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, errors.New(string(resp.Body()))
	}

	return &res, nil
}
