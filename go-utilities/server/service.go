package server

// ServiceType server
type ServiceType string

// Service server dependency
type Service interface {
	GetModules() []Module
	Name() ServiceType
}
