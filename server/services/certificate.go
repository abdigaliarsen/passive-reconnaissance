package services

import "errors"

type CertificateRequest struct {
	Host      *string `json:",omitempty"`
	IpAddress *string `json:",omitempty"`
}

type CertificateResponse struct {
}

func (c *CertificateRequest) GetHost() (string, error) {
	if c.Host == nil || len(*c.Host) == 0 {
		return "", errors.New("host is required")
	}

	return *c.Host, nil
}

func (c *CertificateRequest) GetIpAddress() (string, error) {
	if c.IpAddress == nil || len(*c.IpAddress) == 0 {
		return "", errors.New("ipAddress is required")
	}

	return *c.IpAddress, nil
}

type Certificate interface {
	Exec()
}

type certificate struct {
}

func NewCertificate() Certificate {
	return &certificate{}
}
