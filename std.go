package presenter

import (
	"encoding/json"
)

type Std struct {
	Ok               bool        `json:"ok"`
	Result           interface{} `json:"result,omitempty"`
	ErrorCode        int         `json:"error_code,omitempty"`
	ErrorDescription string      `json:"description,omitempty"`
}

func (s *Std) String() string {
	marshaled, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(marshaled)
}
