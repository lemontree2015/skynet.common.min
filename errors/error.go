package errors

import (
	"errors"
)

// errors
var (
	NormalError               = errors.New("Normal")
	Ignore                    = errors.New("Ignore")
	InvalidPacketError        = errors.New("Invalid Packet")
	UnknownPacketTypeError    = errors.New("Unknown Packet Type")
	LogoutNormalError         = errors.New("Logout Normal")
	NoUserError               = errors.New("No User")
	NoSessionError            = errors.New("No Session")
	NotFound                  = errors.New("Not Found")
	SessionClosedError        = errors.New("Session Closed")
	PasswordError             = errors.New("Password Error")
	ReplacedByNewSessionError = errors.New("Replaced By New Session")
	BlockingError             = errors.New("Blocking Error")
	SendToCloseError          = errors.New("Send To Close Error")
	InvalidArguments          = errors.New("Invalid Arguments")
	DBOpsError                = errors.New("DB Ops Error")
	InvalidFormat             = errors.New("Invalid Format")
	InvalidOps                = errors.New("Invalid Ops")
	InvalidMessageType        = errors.New("Invalid Message Type")
	InvalidJson               = errors.New("Invalid Json")
	MaxRetryError             = errors.New("Max Retry Error")
	LimitError                = errors.New("Limit Error")
	OutOfIndex                = errors.New("Out Of Index")
	DBError                   = errors.New("DB Error")
	AuthError                 = errors.New("Auth Error")
	InvalidLogic              = errors.New("Invalid Logic")
)
