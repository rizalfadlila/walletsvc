package enums

type StatusTransaction string

const (
	Deposit    StatusTransaction = "deposit"
	Withdrawal StatusTransaction = "withdrawal"
)

func (s StatusTransaction) String() string {
	return string(s)
}
