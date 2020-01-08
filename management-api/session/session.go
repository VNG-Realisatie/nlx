package session

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	uuid "github.com/satori/go.uuid"

	"go.nlx.io/nlx/management-api/models"
)

type Session interface {
	IsAuthenticated() (bool, error)
	Account() (*models.Account, error)
	AccountByName(name string) (*models.Account, error)
	Login(w http.ResponseWriter, id fmt.Stringer) error
	Logout(w http.ResponseWriter) error
}

type HTTPSession struct {
	authenicationManager *AuthenticationManager
	session              *sessions.Session
	r                    *http.Request
}

// IsAuthenticated returns if a user is logged in
func (s *HTTPSession) IsAuthenticated() (bool, error) {
	if s.session.Values["account"] != nil {
		id := uuid.FromStringOrNil(s.session.Values["account"].(string))
		account, err := s.authenicationManager.accountRepo.GetByID(id)

		if err != nil {
			return false, err
		}

		if account != nil {
			return true, nil
		}
	}

	return false, nil
}

// Account returns the model of the current logged in account
func (s *HTTPSession) Account() (*models.Account, error) {
	rawID := s.session.Values["account"]
	if rawID == nil {
		return nil, nil
	}

	id := uuid.FromStringOrNil(rawID.(string))

	account, err := s.authenicationManager.accountRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// Login attaches the Account with the provided id to the current session
func (s *HTTPSession) Login(w http.ResponseWriter, id fmt.Stringer) error {
	if id == uuid.Nil {
		return errors.New("field id is nil")
	}

	s.session.Values["account"] = id.String()

	err := s.session.Save(s.r, w)
	if err != nil {
		return err
	}

	return nil
}

// Logout detaches the Account from the current session
func (s *HTTPSession) Logout(w http.ResponseWriter) error {
	delete(s.session.Values, "account")

	err := s.session.Save(s.r, w)
	if err != nil {
		return err
	}

	return nil
}

func (s HTTPSession) AccountByName(name string) (*models.Account, error) {
	return s.authenicationManager.accountRepo.GetByName(name)
}
