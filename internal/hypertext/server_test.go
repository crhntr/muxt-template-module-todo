package hypertext_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/crhntr/muxt-template-module-todo/internal/database"
	"github.com/crhntr/muxt-template-module-todo/internal/fake"
	"github.com/crhntr/muxt-template-module-todo/internal/hypertext"
)

func TestServer_Index(t *testing.T) {
	t.Run("it returns a name", func(t *testing.T) {
		queries := new(fake.Querier)
		queries.SelectListsReturns([]database.List{
			{ID: 1, Name: "to eat"},
			{ID: 2, Name: "to drink"},
			{ID: 3, Name: "to watch"},
			{ID: 3, Name: "to play"},
		}, nil)
		conn := new(fake.DBConnection)
		server := hypertext.Server{
			DBQuery: queries,
			DBConn:  conn,
		}
		ctx := context.Background()
		data := server.Index(ctx)
		require.Len(t, data.Lists, 4)
		if assert.Equal(t, 1, queries.SelectListsCallCount()) {
			ctx, conn := queries.SelectListsArgsForCall(0)
			require.NotNil(t, ctx)
			require.NotNil(t, conn)
		}
	})
}
