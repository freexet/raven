package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/freexet/raven/auth"
	"github.com/freexet/raven/graph/generated"
	"github.com/freexet/raven/graph/model"
	"github.com/gin-gonic/gin"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, params model.NewUser) (*auth.User, error) {
	gc := ctx.Value("ginCtx").(*gin.Context)
	a, _ := gc.Get("auth")
	return a.(auth.Service).Register(params.Username, params.Password)
}

func (r *queryResolver) Users(ctx context.Context) ([]*auth.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
