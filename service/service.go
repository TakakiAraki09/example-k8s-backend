package service

import "github.com/TakakiAraki09/example-k8s-backend/internal/database"

type Service struct {
	Queries *database.Queries
}
