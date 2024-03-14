package domain

const (
	SignerAccount = "tx-services-account"
)

type Transaction struct {
	ToAddress string
	Amount    int64
	Memo      string
}
