package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/mkokoulin/c6er-wallet.git/internal/config"
	custom_error "github.com/mkokoulin/c6er-wallet.git/internal/errors"
	"github.com/mkokoulin/c6er-wallet.git/internal/helpers"
	"github.com/mkokoulin/c6er-wallet.git/internal/jwt"
	"github.com/mkokoulin/c6er-wallet.git/internal/models"
	"github.com/rs/zerolog"
)

type Repository interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	CheckUserPassword(ctx context.Context, user models.User) (string, error)
}

type ContextType string

const UserIDCtx ContextType = "ctxUserId"


type Handlers struct {
	repo   Repository
	cfg    config.Config
	logger *zerolog.Logger
}

func New(repo Repository, cfg *config.Config, logger *zerolog.Logger) *Handlers {
	return &Handlers {
		repo: repo,
		cfg: *cfg,
		logger: logger,
	}
}

func (h *Handlers) Registration(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "application/json; charset=utf-8")

	user := models.User{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	if len(body) == 0 {
		http.Error(w, "the body is missing", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = uuid.NewString()

	newUser, err := h.repo.CreateUser(r.Context(), user)
	var dbErr *custom_error.ErrorWithDB

	if errors.As(err, &dbErr) && dbErr.Title == "UniqConstraint" {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	token, err := jwt.CreateToken(newUser.ID, &h.cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token.AccessToken)
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "application/json; charset=utf-8")

	user := models.User{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	if len(body) == 0 {
		http.Error(w, "the body is missing", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID, err := h.repo.CheckUserPassword(r.Context(), user)
	var dbErr *custom_error.ErrorWithDB

	if errors.As(err, &dbErr) && dbErr.Title == "UserNotFound" {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := jwt.CreateToken(userID, &h.cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	atc, rtc := CreateAccessRefreshCookies(token)

	SetCookies(w, atc, rtc)

	w.Header().Set("Authorization", "Bearer "+token.AccessToken)
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CreateAccessRefreshCookies(token *jwt.TokenDetails)(*http.Cookie, *http.Cookie) {
	atc := helpers.CreateCookie("access_token", token.AccessToken, false, false);
	rtc := helpers.CreateCookie("refresh_token", token.RefreshToken, true, true);

	return atc, rtc
}

func SetCookies(w http.ResponseWriter, atc, rtc *http.Cookie) {
	http.SetCookie(w, atc)
	http.SetCookie(w, rtc)
}