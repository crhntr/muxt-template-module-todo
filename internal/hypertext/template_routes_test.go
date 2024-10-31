package hypertext_test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/crhntr/dom/domtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/crhntr/muxt-template-module-todo/internal/database"
	"github.com/crhntr/muxt-template-module-todo/internal/fake"
	"github.com/crhntr/muxt-template-module-todo/internal/hypertext"
)

func TestTemplates(t *testing.T) {
	for _, tt := range []struct {
		Name     string
		Request  func(rsv *fake.Server) *http.Request
		Response func(rsv *fake.Server, res *http.Response)
	}{
		{
			Name: "the header has the name",
			Request: func(srv *fake.Server) *http.Request {
				srv.IndexReturns(hypertext.IndexData{
					Lists: []database.List{
						{ID: 1, Name: "to eat"},
						{ID: 2, Name: "to drink"},
						{ID: 3, Name: "to watch"},
						{ID: 3, Name: "to play"},
					},
				})
				return httptest.NewRequest(http.MethodGet, "/", nil)
			},
			Response: func(srv *fake.Server, res *http.Response) {
				if assert.Equal(t, 1, srv.IndexCallCount()) {
					ctx := srv.IndexArgsForCall(0)
					require.NotNil(t, ctx)
				}
				assert.Equal(t, http.StatusOK, res.StatusCode)
				doc := domtest.Response(t, res)
				if links := doc.QuerySelectorAll(`ul#lists li a`); assert.Equal(t, 4, links.Length()) {
					for el := range doc.QuerySelectorEach(`ul#lists li a`) {
						assert.NotZero(t, el.TextContent())
						assert.Regexp(t, regexp.MustCompile(`/list/\d+`), el.GetAttribute("href"))
					}
				}
			},
		},
		{
			Name: "the about page is routable",
			Request: func(srv *fake.Server) *http.Request {
				return httptest.NewRequest(http.MethodGet, "/about", nil)
			},
			Response: func(rsv *fake.Server, res *http.Response) {
				assert.Equal(t, http.StatusOK, res.StatusCode)
				doc := domtest.Response(t, res)
				if el := doc.QuerySelector(`h1`); assert.NotNil(t, el) {
					assert.Equal(t, "About", strings.TrimSpace(el.TextContent()))
				}
			},
		},
	} {
		t.Run(tt.Name, func(t *testing.T) {
			srv := new(fake.Server)
			mux := http.NewServeMux()
			hypertext.TemplateRoutes(mux, srv)
			rec := httptest.NewRecorder()
			req := tt.Request(srv)
			mux.ServeHTTP(rec, req)
			tt.Response(srv, rec.Result())
		})
	}
}
