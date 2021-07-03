package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/fgunawan1995/lemonilo/model"
	"github.com/gorilla/mux"
)

func writeErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	log.Printf("error = %+v", err)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model.BuildAPIResponseError(statusCode, errors.Cause(err)))
}

func writeSuccessResponse(w http.ResponseWriter, data interface{}) {
	statusCode := http.StatusOK
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model.BuildAPIResponseSuccess(statusCode, data))
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	if userID == "" {
		writeErrorResponse(w, http.StatusBadRequest, errors.New(model.ErrorUserIDMustBeProvided))
		return
	}
	result, err := h.usecase.GetUserByID(userID)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, result)
}

func (h *Handler) InsertUser(w http.ResponseWriter, r *http.Request) {
	var p model.InsertUser
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	err = h.usecase.InsertUser(p)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, model.GeneralSuccessMessage)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var p model.UpdateUser
	vars := mux.Vars(r)
	userID := vars["id"]
	if userID == "" {
		writeErrorResponse(w, http.StatusBadRequest, errors.New(model.ErrorUserIDMustBeProvided))
		return
	}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	p.ID = userID
	err = h.usecase.UpdateUser(p)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, model.GeneralSuccessMessage)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	if userID == "" {
		writeErrorResponse(w, http.StatusBadRequest, errors.New(model.ErrorUserIDMustBeProvided))
		return
	}
	err := h.usecase.DeleteUser(userID)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, model.GeneralSuccessMessage)
}
