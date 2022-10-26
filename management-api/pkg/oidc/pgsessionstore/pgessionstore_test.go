// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//go:build integration

package pgsessionstore_test

import (
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

// nolint:gosec // this is a test secret
const secret = "EyaC2BPcJtNqU3tjEHy+c+Wmqc1yihYIbUWEl/jk0Ga73kWBclmuSFd9HuJKwJw/Wdsh1XnjY2Bw1HBVph6WOw=="

type headerOnlyResponseWriter http.Header

func (ho headerOnlyResponseWriter) Header() http.Header {
	return http.Header(ho)
}

func (ho headerOnlyResponseWriter) Write([]byte) (int, error) {
	panic("NOIMPL")
}

func (ho headerOnlyResponseWriter) WriteHeader(int) {
	panic("NOIMPL")
}

// nolint:funlen,gocyclo // this is a test
func TestPGStore(t *testing.T) {
	t.Parallel()

	ss := New(t, secret)
	defer ss.Close()

	// Check that the cookie is being saved
	req, err := http.NewRequest("GET", "http://www.example.com", http.NoBody)
	assert.NoError(t, err, "could not create request")

	session, err := ss.Get(req, "mysess")
	if err != nil {
		t.Fatal("failed to get session", err.Error())
	}

	session.Values["counter"] = 1

	m := make(http.Header)
	if err = ss.Save(req, headerOnlyResponseWriter(m), session); err != nil {
		t.Fatal("Failed to save session:", err.Error())
	}

	if m["Set-Cookie"][0][0:6] != "mysess" {
		t.Fatal("Cookie wasn't set!")
	}

	// check that the cookie can be retrieved
	req, err = http.NewRequest("GET", "http://www.example.com", http.NoBody)
	assert.NoError(t, err, "could not create request")

	encoded, err := securecookie.EncodeMulti(session.Name(), session.ID, ss.Codecs...)
	if err != nil {
		t.Fatal("Failed to make cookie value", err)
	}

	req.AddCookie(sessions.NewCookie(session.Name(), encoded, session.Options))

	session, err = ss.Get(req, "mysess")
	assert.NoError(t, err, "could not retrieve session")

	if session.Values["counter"] != 1 {
		t.Fatal("Retrieved session had wrong value:", session.Values["counter"])
	}

	session.Values["counter"] = 9 // set new value for round 3
	if err = ss.Save(req, headerOnlyResponseWriter(m), session); err != nil {
		t.Fatal("Failed to save session:", err.Error())
	}

	// Check that the cookie has been updated
	req, err = http.NewRequest("GET", "http://www.example.com", http.NoBody)
	assert.NoError(t, err, "could not create request")

	req.AddCookie(sessions.NewCookie(session.Name(), encoded, session.Options))

	session, err = ss.Get(req, "mysess")
	assert.NoError(t, err, "could not retrieve session")

	assert.Equal(t, 9, session.Values["counter"])

	// Increase max length
	req, err = http.NewRequest("GET", "http://www.example.com", http.NoBody)
	assert.NoError(t, err, "could not create request")

	req.AddCookie(sessions.NewCookie(session.Name(), encoded, session.Options))

	session, err = ss.New(req, "my session")
	assert.NoError(t, err, "could not create session")

	session.Values["big"] = make([]byte, base64.StdEncoding.DecodedLen(4096*2))

	err = ss.Save(req, headerOnlyResponseWriter(m), session)
	assert.Error(t, err)

	ss.MaxLength(4096 * 3) // A bit more than the value size to account for encoding overhead.

	err = ss.Save(req, headerOnlyResponseWriter(m), session)
	assert.NoError(t, err, "could not save session")
}

func TestSessionOptionsAreUniquePerSession(t *testing.T) {
	t.Parallel()

	ss := New(t, secret)
	defer ss.Close()

	maxAge := 900
	ss.Options.MaxAge = maxAge

	req, err := http.NewRequest("GET", "http://www.example.com", http.NoBody)
	assert.NoError(t, err, "could not create request")

	session, err := ss.Get(req, "newsess")
	assert.NoError(t, err, "could not create session")

	session.Options.MaxAge = -1

	assert.Equal(t, maxAge, ss.Options.MaxAge)
}
