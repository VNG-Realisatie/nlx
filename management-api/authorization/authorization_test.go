package authorization

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_authorization "go.nlx.io/nlx/management-api/authorization/mock"
)

func TestAuthorization(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	type fields struct {
		authorizer Authorizer
	}

	type args struct {
		request *http.Request
	}

	forbiddenRequest := httptest.NewRequest("GET", "/forbidden", nil)
	allowedRequest := httptest.NewRequest("GET", "/allowed", nil)
	authorizer := mock_authorization.NewMockAuthorizer(controller)
	authorizer.EXPECT().Authorize(forbiddenRequest).Return(false).AnyTimes()
	authorizer.EXPECT().Authorize(allowedRequest).Return(true).AnyTimes()

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "allowed",
			fields: fields{authorizer: authorizer},
			args:   args{allowedRequest},
			want:   "OK",
		},
		{
			name:   "forbidden",
			fields: fields{authorizer: authorizer},
			args:   args{forbiddenRequest},
			want:   http.StatusText(http.StatusForbidden),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			a := Authorization{
				authorizer: tt.fields.authorizer,
			}

			middleware := a.Middleware(handlerFunc)
			rs := httptest.NewRecorder()
			middleware.ServeHTTP(rs, tt.args.request)

			got := strings.Trim(rs.Body.String(), "\n")
			assert.Equal(t, tt.want, got)
		})
	}
}
