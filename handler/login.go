package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fgunawan1995/lemonilo/model"
	"github.com/pkg/errors"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var p model.Login
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	token, err := h.usecase.Login(p)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, token)
}

func (h *Handler) TestToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(model.ContextUserIDKey).(string)
	if !ok {
		writeErrorResponse(w, http.StatusBadRequest, errors.New("user_id not exist"))
		return
	}
	writeSuccessResponse(w, fmt.Sprintf("your id is %s", userID))
}
