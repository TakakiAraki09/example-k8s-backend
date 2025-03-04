package resolver

import "github.com/TakakiAraki09/k8s-lesson/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.Service
}
