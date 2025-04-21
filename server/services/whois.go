package services

type Whois interface {
}

type whois struct {
}

func NewWhois() Whois {
	return &whois{}
}
