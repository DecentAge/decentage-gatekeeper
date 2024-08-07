package authorization

import (
	"strconv"
)

type AuthzDecision int

const (
	UndefinedAuthz AuthzDecision = iota
	AllowedAuthz   AuthzDecision = iota
	DeniedAuthz    AuthzDecision = iota
)

func (decision AuthzDecision) String() string {
	switch decision {
	case AllowedAuthz:
		return strconv.Itoa(int(AllowedAuthz))
	case DeniedAuthz:
		return strconv.Itoa(int(DeniedAuthz))
	case UndefinedAuthz:
		return ""
	}
	return strconv.Itoa(int(DeniedAuthz))
}

type Provider interface {
	Authorize() (AuthzDecision, error)
}
