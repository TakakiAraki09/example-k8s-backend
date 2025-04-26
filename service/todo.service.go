package service

import (
	"context"
	"database/sql"

	"github.com/TakakiAraki09/example-k8s-backend/internal/database"
	"github.com/TakakiAraki09/example-k8s-backend/internal/model"
)

func (r *Service) GetTodoByUserId(ctx context.Context, id uint32) ([]*model.Todo, error) {
	data, err := r.Queries.GetTodosGetByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	// Convert the data to the model
	var todos []*model.Todo
	for _, d := range data {
		todos = append(todos, &model.Todo{
			ID:     d.ID,
			Text:   d.Text.String,
			Done:   d.Done.Valid, // BooleanになってないからSQLから訂正を行う
			UserID: d.UserID,
		})
	}
	return todos, nil
}

func (r *Service) GetTodoAll(ctx context.Context) ([]*model.Todo, error) {
	data, err := r.Queries.GetTodosGetAll(ctx)
	if err != nil {
		return nil, err
	}
	// Convert the data to the model
	var todos []*model.Todo
	for _, d := range data {
		todos = append(todos, &model.Todo{
			ID:     d.ID,
			Text:   d.Text.String,
			Done:   d.Done.Valid, // BooleanになってないからSQLから訂正を行う
			UserID: d.UserID,
		})
	}
	return todos, nil
}

func (r *Service) CreateTodo(ctx context.Context, args model.NewTodo) (*model.Todo, error) {
	data, err := r.Queries.CreateTodo(ctx, database.CreateTodoParams{
		Text:   sql.NullString{String: args.Text, Valid: true},
		UserID: args.UserID,
	})
	if err != nil {
		return nil, err
	}

	id, err := data.LastInsertId()
	if err != nil {
		return nil, err
	}
	todo, err := r.Queries.GetTodoById(ctx, uint32(id))
	if err != nil {
		return nil, err
	}

	return &model.Todo{
		ID:     todo.ID,
		Text:   todo.Text.String,
		Done:   todo.Done.Valid, // BooleanになってないからSQLから訂正を行う
		UserID: todo.UserID,
	}, nil
}
