package types

import "errors"

type Status string

var ErrTypeStatusInvalid = errors.New("Invalid Status")

const (
	SEND     = "Send"
	ACCEPTED = "Accepted"
	REFUSED  = "Refused"
)

func (st Status) IsValid() error {
	switch st {
	case SEND, ACCEPTED, REFUSED:
		return nil
	}
	return ErrTypeStatusInvalid
}
