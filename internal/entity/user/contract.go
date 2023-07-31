package user

import "net/http"

type InterfaceUserService interface {
	Save() error
	Get() InterfaceUserService

	Create(w http.ResponseWriter, r *http.Request)
}
