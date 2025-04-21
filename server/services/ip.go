package services

import (
	"errors"
	"net"
	"passive-reconnaissance/utils"
)

type IpField struct {
	Value *string `json:",omitempty"`
	Error *string `json:",omitempty"`
}

type IpResponse struct {
	IP          []IpField `json:"ip"`
	DefaultMask []IpField `json:"default_mask"`
}

type IpRequest struct {
	Host      *string `json:",omitempty"`
	IpAddress *string `json:",omitempty"`
}

func (r *IpRequest) GetHost() (string, error) {
	if r.Host == nil || len(*r.Host) == 0 {
		return "", errors.New("host is empty")
	}

	return *r.Host, nil
}

func (r *IpRequest) GetIpAddress() (string, error) {
	if r.IpAddress == nil || len(*r.IpAddress) == 0 {
		return "", errors.New("ip address is empty")
	}

	return *r.IpAddress, nil
}

type Ip interface {
	Exec(*IpRequest) *IpResponse
}

type scanner struct {
}

func NewIp() Ip {
	return &scanner{}
}

func (s *scanner) Exec(request *IpRequest) *IpResponse {
	rt := &IpResponse{}

	if host, err := request.GetHost(); err != nil {
		ips, err := net.LookupIP(host)
		if err != nil {
			e := err.Error()
			rt.IP = append(rt.IP, IpField{Error: &e})
		} else {
			for _, ip := range ips {
				v := ip.String()
				rt.IP = append(rt.IP, IpField{
					Value: &v,
				})
			}
		}
	}

	if ipAddress, err := request.GetIpAddress(); err != nil {
		if !utils.Any(utils.Map(rt.IP, func(ip IpField) string {
			if ip.Value == nil {
				return ""
			}
			return *ip.Value
		}), ipAddress) {
			errMsg := "requested ip address is not in set of ip addresses of the host"
			rt.DefaultMask = []IpField{{Error: &errMsg}}
		} else {
			m := net.ParseIP(ipAddress).DefaultMask().String()
			rt.DefaultMask = []IpField{{Value: &m}}
		}
	} else {
		rt.DefaultMask = make([]IpField, 0, len(rt.IP))
		for _, ip := range rt.IP {
			if ip.Value == nil {
				msg := "requested ip address is not in set of ip addresses"
				rt.DefaultMask = append(rt.DefaultMask, IpField{Error: &msg})
				continue
			}

			m := net.ParseIP(*ip.Value).DefaultMask().String()
			rt.DefaultMask = append(rt.DefaultMask, IpField{Value: &m})
		}
	}

	return rt
}
