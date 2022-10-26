// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package pgsessionstore

import (
	"context"
	"database/sql"
	"encoding/base32"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/pkg/oidc/pgsessionstore/queries"
)

const (
	defaultMaxAge          = 60 * 60 * 24 * 30     // 30 days
	defaultExpiresAge      = time.Minute * 60 * 24 // 1 day
	defaultRandomKeyLength = 32
)

// PGStore represents the currently configured session store.
type PGStore struct {
	logger  *zap.Logger
	Codecs  []securecookie.Codec
	Options *sessions.Options
	Path    string
	Querier queries.Queries
}

// PGSession type
type PGSession struct {
	ID        int64
	Key       string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresOn time.Time
}

// New creates a new PGStore instance.
func New(logger *zap.Logger, db *sql.DB, keyPairs ...[]byte) (*PGStore, error) {
	querier, err := queries.Prepare(context.Background(), db)
	if err != nil {
		return nil, err
	}

	return &PGStore{
		logger: logger,
		Codecs: securecookie.CodecsFromPairs(keyPairs...),
		Options: &sessions.Options{
			Path:   "/",
			MaxAge: defaultMaxAge,
		},
		Querier: *querier,
	}, nil
}

// Close closes the database connection.
func (db *PGStore) Close() error {
	return db.Querier.Close()
}

// Get Fetches a session for a given name after it has been added to the
// registry.
func (db *PGStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	return sessions.GetRegistry(r).Get(db, name)
}

// New returns a new session for the given name without adding it to the registry.
func (db *PGStore) New(r *http.Request, name string) (*sessions.Session, error) {
	session := sessions.NewSession(db, name)

	opts := *db.Options
	session.Options = &(opts)
	session.IsNew = true

	var err error
	if c, errCookie := r.Cookie(name); errCookie == nil {
		err = securecookie.DecodeMulti(name, c.Value, &session.ID, db.Codecs...)
		if err == nil {
			err = db.load(session)
			if err == nil {
				session.IsNew = false
			} else if errors.Is(err, sql.ErrNoRows) {
				err = nil
			}
		}
	}

	db.MaxAge(db.Options.MaxAge)

	return session, err
}

// Save saves the given session into the database and deletes cookies if needed
func (db *PGStore) Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	// Set delete if max-age is < 0
	if session.Options.MaxAge < 0 {
		if err := db.Querier.DeleteSession(context.Background(), []byte(session.ID)); err != nil {
			return err
		}

		http.SetCookie(w, sessions.NewCookie(session.Name(), "", session.Options))

		return nil
	}

	if session.ID == "" {
		// Generate a random session ID key suitable for storage in the DB
		session.ID = strings.TrimRight(
			base32.StdEncoding.EncodeToString(
				securecookie.GenerateRandomKey(defaultRandomKeyLength),
			), "=")
	}

	if err := db.save(session); err != nil {
		return err
	}

	// Keep the session ID key in a cookie so it can be looked up in DB later.
	encoded, err := securecookie.EncodeMulti(session.Name(), session.ID, db.Codecs...)
	if err != nil {
		return err
	}

	http.SetCookie(w, sessions.NewCookie(session.Name(), encoded, session.Options))

	return nil
}

// MaxLength restricts the maximum length of new sessions to l.
// If l is 0 there is no limit to the size of a session, use with caution.
// The default for a new PGStore is 4096. PostgreSQL allows for max
// value sizes of up to 1GB (http://www.postgresql.org/docs/current/interactive/datatype-character.html)
func (db *PGStore) MaxLength(l int) {
	for _, c := range db.Codecs {
		if codec, ok := c.(*securecookie.SecureCookie); ok {
			codec.MaxLength(l)
		}
	}
}

// MaxAge sets the maximum age for the store and the underlying cookie
// implementation. Individual sessions can be deleted by setting Options.MaxAge
// = -1 for that session.
func (db *PGStore) MaxAge(age int) {
	db.Options.MaxAge = age

	// Set the maxAge for each securecookie instance.
	for _, codec := range db.Codecs {
		if sc, ok := codec.(*securecookie.SecureCookie); ok {
			sc.MaxAge(age)
		}
	}
}

// load fetches a session by ID from the database and decodes its content
// into session.Values.
func (db *PGStore) load(session *sessions.Session) error {
	dbSession, err := db.Querier.GetSession(context.Background(), []byte(session.ID))
	if err != nil {
		return err
	}

	return securecookie.DecodeMulti(session.Name(), string(dbSession.Data), &session.Values, db.Codecs...)
}

// save writes encoded session.Values to a database record.
// writes to http_sessions table by default.
func (db *PGStore) save(session *sessions.Session) error {
	encoded, err := securecookie.EncodeMulti(session.Name(), session.Values, db.Codecs...)
	if err != nil {
		return err
	}

	crOn := session.Values["created_on"]
	exOn := session.Values["expires_on"]

	createdAt, ok := crOn.(time.Time)
	if !ok {
		createdAt = time.Now()
	}

	// default expiresOn is
	dbExpiresOn := sql.NullTime{
		Time:  time.Now().Add(defaultExpiresAge),
		Valid: true,
	}

	if session.Options.MaxAge != 0 {
		if exOn == nil {
			dbExpiresOn = sql.NullTime{
				Time:  time.Now().Add(time.Second * time.Duration(session.Options.MaxAge)),
				Valid: true,
			}
		} else {
			expiresOn := exOn.(time.Time)
			if expiresOn.Sub(time.Now().Add(time.Second*time.Duration(session.Options.MaxAge))) < 0 {
				expiresOn = time.Now().Add(time.Second * time.Duration(session.Options.MaxAge))
			}

			dbExpiresOn = sql.NullTime{
				Time:  expiresOn,
				Valid: true,
			}
		}
	}

	if session.IsNew {
		return db.Querier.CreateSession(context.Background(), &queries.CreateSessionParams{
			Key:       []byte(session.ID),
			Data:      []byte(encoded),
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			ExpiresOn: dbExpiresOn,
		})
	}

	return db.Querier.UpdateSession(context.Background(), &queries.UpdateSessionParams{
		Key:  []byte(session.ID),
		Data: []byte(encoded),
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ExpiresOn: dbExpiresOn,
	})
}
