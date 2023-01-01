package blocks

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/blackhorseya/ethscan/internal/pkg/errorx"
	bb "github.com/blackhorseya/ethscan/pkg/entity/domain/block/biz"
	"github.com/stretchr/testify/mock"
)

func (s *suiteTester) Test_impl_List() {
	type args struct {
		page string
		size string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "invalid page then 400",
			args:     args{page: "a", size: "10"},
			wantCode: 400,
		},
		{
			name:     "invalid size then 400",
			args:     args{page: "1", size: "a"},
			wantCode: 400,
		},
		{
			name: "list then 500",
			args: args{page: "2", size: "5", mock: func() {
				s.biz.On("List", mock.Anything, bb.ListCondition{Page: 2, Size: 5}).Return(nil, 0, errorx.ErrGetRecord).Once()
			}},
			wantCode: 500,
		},
		{
			name: "default query to list then 200",
			args: args{mock: func() {
				s.biz.On("List", mock.Anything, bb.ListCondition{Page: 1, Size: 10}).Return(nil, 10, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri, _ := url.Parse("/api/v1/blocks")
			query := uri.Query()
			if len(tt.args.page) != 0 {
				query.Add("page", tt.args.page)
			}
			if len(tt.args.size) != 0 {
				query.Add("size", tt.args.size)
			}
			uri.RawQuery = query.Encode()
			req := httptest.NewRequest(http.MethodGet, uri.String(), nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			if got.StatusCode != tt.wantCode {
				t.Errorf("List() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			}

			s.assertExpectation(t)
		})
	}
}
