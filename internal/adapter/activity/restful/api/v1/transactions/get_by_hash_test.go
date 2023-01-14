package transactions

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/blackhorseya/ethscan/internal/pkg/errorx"
	"github.com/stretchr/testify/mock"
)

func (s *suiteTester) Test_impl_GetByHash() {
	type args struct {
		hash string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "get by hash then 500",
			args: args{hash: "hash", mock: func() {
				s.biz.On("GetByHash", mock.Anything, "hash").Return(nil, errorx.ErrGetTx).Once()
			}},
			wantCode: 500,
		},
		{
			name: "get by hash then 200",
			args: args{hash: "hash", mock: func() {
				s.biz.On("GetByHash", mock.Anything, "hash").Return(nil, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri, _ := url.Parse(fmt.Sprintf("/api/v1/transactions/%s", tt.args.hash))
			req := httptest.NewRequest(http.MethodGet, uri.String(), nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			if got.StatusCode != tt.wantCode {
				t.Errorf("GetByHash() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			}

			s.assertExpectation(t)

		})
	}
}
