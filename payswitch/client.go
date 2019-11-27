package payswitch

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type ApiClient struct {
	baseUrl             string
	httpc               *resty.Client
	common              service
	PaymentService      *PaymentService
	VerificationService *VerificationService
	FundTransferService *FundTransferService
}

type ApiResult struct {
	TransactionID    string  `json:"transaction_id"`
	Status           string  `json:"status"`
	Code             string  `json:"code"`
	Reason           string  `json:"reason"`
	RSwitch          string  `json:"r_switch"`
	SubscriberNumber string  `json:"subscriber_number"`
	Amount           float64 `json:"amount"`
}

type VerifyApiResult struct {
	TransactionID    string `json:"transaction_id"`
	Status           string `json:"status"`
	Code             string `json:"code"`
	Reason           string `json:"reason"`
	RSwitch          string `json:"r_switch"`
	SubscriberNumber string `json:"subscriber_number"`
	Amount           string `json:"amount"`
}

type service struct {
	client *ApiClient
}

type PaymentService service
type VerificationService service
type FundTransferService service

func NewApiClient(username string, key string, live bool) *ApiClient {

	h := resty.New().EnableTrace().SetBasicAuth(username, key)

	var baseUrl = "https://prod.theteller.net"

	if !live {
		baseUrl = "https://test.theteller.net"
	}

	c := &ApiClient{}
	c.httpc = h
	c.baseUrl = baseUrl
	c.common.client = c
	c.VerificationService = (*VerificationService)(&c.common)
	c.PaymentService = (*PaymentService)(&c.common)
	c.FundTransferService = (*FundTransferService)(&c.common)

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

func (s *PaymentService) ProcessMobileMoneyPayment(m *MobileMoneyPaymentRequest) (*ApiResult, error) {

	err := m.Validate()
	if err != nil {
		return nil, err
	}

	var client = s.client
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

func (s *VerificationService) VerifyTransaction(id string, merchant string) (*VerifyApiResult, error) {
	if len(id) != 12 {
		return nil, errors.New("invalid id len < 12")
	}

	var endpoint = fmt.Sprintf("/v1.1/users/transactions/%s/status", id)
	var client = s.client
	resp, err := client.httpc.R().SetHeader("Merchant-Id", merchant).Get(client.baseUrl + endpoint)
	if err != nil {
		return nil, err
	}

	res := VerifyApiResult{}

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, errors.New(string(resp.Body()))
	}

	return &res, nil
}
