package user

import (
	"Rest-api/internal/apperror"
	"Rest-api/internal/handlers"
	"Rest-api/pkg/logging"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))

}
func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {

	w.Write([]byte("this is list of users"))
	w.WriteHeader(200)
	return nil

}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this is create users"))
	w.WriteHeader(201)
	return nil

}
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this is user by uuid"))
	w.WriteHeader(200)
	return nil

}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this is fully update users"))
	w.WriteHeader(204)
	return nil

}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this is partially update user"))
	w.WriteHeader(204)
	return nil

}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this is delete user"))
	w.WriteHeader(204)

	return nil
}