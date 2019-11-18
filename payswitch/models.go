package payswitch

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

type PaymentMethod struct {
	Amount         string `json:"amount"`
	ProcessingCode string `json:"processing_code"`
	RSwitch        string `json:"r-switch"`
	Desc           string `json:"desc"`
	MerchantID     string `json:"merchant_id"`
	TransactionID  string `json:"transaction_id"`
}
type CardPaymentRequest struct {
	PaymentMethod
	Pan           string `json:"pan"`
	RedirectUrl   string `json:"3d_url_response"`
	ExpMonth      string `json:"exp_month"`
	ExpYear       string `json:"exp_year"`
	Cvv           string `json:"cvv"`
	Currency      string `json:"currency"`
	CardHolder    string `json:"card_holder"`
	CustomerEmail string `json:"customer_email"`
}

func (c *CardPaymentRequest) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Amount, validation.Required, validation.Length(12, 12)),
		validation.Field(&c.ProcessingCode, validation.Required),
		validation.Field(&c.RSwitch, validation.Required),
		validation.Field(&c.Desc, validation.Required),
		validation.Field(&c.MerchantID, validation.Required),
		validation.Field(&c.TransactionID, validation.Required, validation.Length(12, 12)),
		//CardPaymentRequest Specific Fields
		validation.Field(&c.Pan, validation.Required),
		validation.Field(&c.RedirectUrl, validation.Required),
		validation.Field(&c.ExpMonth, validation.Required),
		validation.Field(&c.ExpYear, validation.Required),
		validation.Field(&c.Cvv, validation.Required),
		validation.Field(&c.Currency, validation.Required),
		validation.Field(&c.CardHolder, validation.Required),
		validation.Field(&c.CustomerEmail, validation.Required),
	)
}

type MobileMoneyPaymentRequest struct {
	PaymentMethod
	SubscriberNumber string `json:"subscriber_number"`
	VoucherCode      string `json:"voucher_code"`
}

func (c *MobileMoneyPaymentRequest) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Amount, validation.Required),
		validation.Field(&c.ProcessingCode, validation.Required),
		validation.Field(&c.RSwitch, validation.Required),
		validation.Field(&c.Desc, validation.Required),
		validation.Field(&c.MerchantID, validation.Required),
		validation.Field(&c.TransactionID, validation.Required),
		//Mobile Payment Specific Fields
		validation.Field(&c.SubscriberNumber, validation.Required),
		validation.Field(&c.VoucherCode, validation.By(rSwitchEquals(c.RSwitch))),

	)
}

func rSwitchEquals(r string) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if r == "VDF" && s == "" {
			return errors.New("rswitch:VDF must have voucher_code")
		}

		return nil
	}
}
