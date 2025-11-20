package hypertext

//go:generate go run github.com/typelate/muxt/cmd/muxt generate --receiver-type=Server --routes-func=TemplateRoutes
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o=../fake/server.go        --fake-name=Server       . RoutesReceiver

//counterfeiter:generate -o=../fake/querier.go       --fake-name=Querier      ../database Querier
//counterfeiter:generate -o=../fake/db_connection.go --fake-name=DBConnection ../database Connection
//counterfeiter:generate -o=../fake/tx.go            --fake-name=Tx           github.com/jackc/pgx/v5.Tx
