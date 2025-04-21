package server

import "passive-reconnaissance/server/services"

type Service interface {
	services.Scanner
}

type service struct {
}

func NewService() Service {
	return &service{}
}
