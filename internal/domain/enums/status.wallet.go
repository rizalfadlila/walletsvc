package enums

type StatusWallet string

const (
	Enabled  StatusWallet = "enabled"
	Disabled StatusWallet = "disabled"
)

func (s StatusWallet) Against() StatusWallet {
	if s == Enabled {
		return Disabled
	}
	return Enabled
}
