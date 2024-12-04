package handler

import "instagram/pkg/service"

type Handler interface {
	AuthHandler
}

type handler struct {
	s service.Service
}

func New(s service.Service) *handler {
	return &handler{
		s: s,
	}
}