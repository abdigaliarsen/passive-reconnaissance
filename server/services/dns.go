package services

type Dns interface {
}

type dns struct {
}

func NewDns() Dns {
	return &dns{}
}
