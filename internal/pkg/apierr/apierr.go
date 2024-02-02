package apierr

import "net/http"

type (
	ApiErr struct {
		Message          string
		StatusCode       int
		ExposeStacktrace bool
		Stacktrace       interface{}
	}

	ErrOpt func(o *ApiErr)
)

func New(message string, opts ...ErrOpt) *ApiErr {
	apiErr := &ApiErr{
		Message:          message,
		Stacktrace:       message,
		StatusCode:       http.StatusInternalServerError,
		ExposeStacktrace: false,
	}

	for _, o := range opts {
		o(apiErr)
	}

	return apiErr
}

func (e *ApiErr) Error() string {
	return e.Message
}

func WithStacktrace(stacktrace interface{}) ErrOpt {
	return func(o *ApiErr) {
		o.Stacktrace = stacktrace
	}
}

func WithStatusCode(statusCode int) ErrOpt {
	return func(o *ApiErr) {
		o.StatusCode = statusCode
	}
}
