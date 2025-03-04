package service

import "github.com/TakakiAraki09/k8s-lesson/internal/database"

type Service struct {
	Queries *database.Queries
}
