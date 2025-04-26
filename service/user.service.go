package service

import (
	"context"

	"github.com/TakakiAraki09/example-k8s-backend/internal/model"
)

func (r *Service) GetUserByTodo(ctx context.Context, id uint32) (*model.User, error) {
	data, err := r.Queries.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:   data.ID,
		Name: data.Name.String,
	}, nil
}
