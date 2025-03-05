package routes

import (
	"objective-service/api/business"

	api "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-api"
	"github.com/danielgtaylor/huma/v2"
)

type Router struct {
	*api.Router
	Business business.IBusiness
}

func NewRouter(business business.Logic, r1 *api.Router) *Router {
	r := &Router{Router: r1, Business: business}
	Setup(r)
	return r
}

func Setup(r *Router) {
	// Version
	huma.Register(r.Api, GetVersionOperation, r.GetVersion)

	// Objectives
	huma.Register(r.Api, GetAllObjectivesOperation, r.GetAllObjectives)
}
