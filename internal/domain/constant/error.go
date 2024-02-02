package constant

import "errors"

var (
	ErrQueryRequest         = errors.New("error on performing query request")
	ErrTransactionalRequest = errors.New("error on performing transactional request")
	ErrAlreadyEnabledWallet = errors.New("Already Enabled")
)
