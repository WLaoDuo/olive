package usrgrp

import (
	"context"
	"fmt"
	"net/http"

	v1Web "github.com/WLaoDuo/olive/business/web/v1"
	"github.com/WLaoDuo/olive/business/web/v1/mid"
	"github.com/WLaoDuo/olive/engine/kernel"
	"github.com/WLaoDuo/olive/foundation/web"
	"go.uber.org/zap"
)

// Handlers manages the set of check enpoints.
type Handlers struct {
	Log *zap.SugaredLogger
	K   *kernel.Kernel
}

// Login handler is for User logins.
func (h Handlers) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var req LoginRequest
	web.Decode(r, &req)

	if !h.K.IsValidPortalUser(req.Username, req.Password) {
		return v1Web.NewRequestError(fmt.Errorf("invalid Username or Password"), http.StatusBadRequest)
	}

	status := struct {
		Permissions []string `json:"permissions"`
	}{
		Permissions: []string{"*.*.*"},
	}

	return mid.Respond(ctx, w, status, http.StatusOK)
}

// Logout handler is for User logouts.
func (h Handlers) Logout(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return mid.Respond(ctx, w, nil, http.StatusOK)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
