package session

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/management-api/models"
	mock_session "go.nlx.io/nlx/management-api/session/mock"
)

func TestSessionAuthorizer_Authorize(t *testing.T) {
	type args struct {
		r *http.Request
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	readonlySession := mock_session.NewMockSession(mockCtrl)
	readonlySession.EXPECT().Account().Return(&models.Account{Role: "readonly"}, nil).AnyTimes()

	adminSession := mock_session.NewMockSession(mockCtrl)
	adminSession.EXPECT().Account().Return(&models.Account{Role: "admin"}, nil).AnyTimes()

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "GET anonymous",
			args: args{r: httptest.NewRequest("GET", "/anonymous", nil)},
			want: false,
		},
		{
			name: "GET readonly",
			args: args{r: httptest.NewRequest("GET", "/readonly", nil).WithContext(context.WithValue(context.Background(), contextKey, readonlySession))},
			want: true,
		},
		{
			name: "GET admin",
			args: args{r: httptest.NewRequest("GET", "/admin", nil).WithContext(context.WithValue(context.Background(), contextKey, adminSession))},
			want: true,
		},
		{
			name: "POST readonly",
			args: args{r: httptest.NewRequest("POST", "/readonly", nil).WithContext(context.WithValue(context.Background(), contextKey, readonlySession))},
			want: false,
		},
		{
			name: "POST admin",
			args: args{r: httptest.NewRequest("POST", "/admin", nil).WithContext(context.WithValue(context.Background(), contextKey, adminSession))},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			authorizer := Authorizer{}
			assert.Equal(t, tt.want, authorizer.Authorize(tt.args.r))
		})
	}
}
