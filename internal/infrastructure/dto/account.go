package dto

import "github.com/julo/walletsvc/internal/pkg/datatype"

type (
	RequestAccessAccount struct {
		CustomerID string `form:"customer_xid"`
	}

	AccessAccount struct {
		Token string `json:"token"`
	}
)

func (c *RequestAccessAccount) Validate() datatype.Fields {
	if c.CustomerID == "" {
		return datatype.Fields{
			"customer_xid": []string{
				"Missing data for required field.",
			},
		}
	}

	return nil
}
