package hypertext

import (
	"context"
	"embed"
	"html/template"

	"github.com/crhntr/muxt-template-module-todo/internal/database"
)

//go:embed templates/*.gohtml
var templateFiles embed.FS

var templates = template.Must(template.ParseFS(templateFiles, "templates/*.gohtml"))

type Server struct {
	DBQuery database.Querier
	DBConn  database.Connection
}

func (srv *Server) PatchTaskComplete(ctx context.Context, id int64) (database.Task, error) {
	return srv.DBQuery.UpdateTaskSetCompletedAtNow(ctx, srv.DBConn, id)
}

func (srv *Server) DeleteList(ctx context.Context, id int64) error {
	return srv.DBQuery.DeleteList(ctx, srv.DBConn, id)
}

type PostTaskValues struct {
	ListID      int64
	Description string
}

type PostTaskResult struct {
	List struct {
		ID int64
	}
	Task database.Task
}

func (srv *Server) PostTask(ctx context.Context, values PostTaskValues) (PostTaskResult, error) {
	row, err := srv.DBQuery.InsertTask(ctx, srv.DBConn, database.InsertTaskParams(values))
	if err != nil {
		return PostTaskResult{}, err
	}
	return PostTaskResult{
		List: struct {
			ID int64
		}{ID: values.ListID},
		Task: row,
	}, nil
}

type PostListValues struct {
	Name        string
	Description string
}

type PostListResult struct {
	List   database.List
	Values PostListValues
	Error  error
	Tasks  []database.Task
}

func (srv *Server) PostList(ctx context.Context, params PostListValues) PostListResult {
	row, err := srv.DBQuery.InsertList(ctx, srv.DBConn, database.InsertListParams(params))
	if err != nil {
		return PostListResult{
			Values: params,
			Error:  err,
		}
	}
	return PostListResult{
		List:   row,
		Values: params,
		Error:  nil,
	}
}

type ListData struct {
	List  database.List
	Tasks []database.Task
}

func (srv *Server) GetList(ctx context.Context, id int64) (ListData, error) {
	list, err := srv.DBQuery.SelectList(ctx, srv.DBConn, id)
	if err != nil {
		return ListData{}, err
	}
	tasks, err := srv.DBQuery.SelectTasksForList(ctx, srv.DBConn, id)
	if err != nil {
		return ListData{}, err
	}
	return ListData{
		List:  list,
		Tasks: tasks,
	}, err
}

type IndexData struct {
	Lists []database.List
	Error error
}

func (srv *Server) Index(ctx context.Context) IndexData {
	list, err := srv.DBQuery.SelectLists(ctx, srv.DBConn)
	return IndexData{
		Lists: list,
		Error: err,
	}
}
