package payswitch

const (
	MoMoProcessingCode = "000200"
	CardProcessingCode = "000000"
)

const (
	ResponseCodeTransactionSuccessful                       = "000"
	ResponseCodeInsufficientFunds                           = "101"
	ResponseCodeTransactionNotPermittedToCardHolderOrFailed = "100"
	ResponseCodeNumberNotRegisteredForMoMo                  = "102"
	ResponseCodeWrongPin                                    = "103"
	ResponseCodeTransactionDeclined                         = "105"
	ResponseCodePaymentRequestSent                          = "111"
	ResponseCodeUSSDBusy                                    = "107"
	ResponseCodeInvalidVoucherCode                          = "114"
	ResponseCodeVBVRequired                                 = "200"
	ResponseCodeAccessDenied                                = "600"
	ResponseCodeAccessDeniedInvalidCredentials              = "979"
	ResponseCodeDuplicateTransactionID                      = "909"
	ResponseCodeAccessDeniedMerchantIDNotSet                = "999"
)
