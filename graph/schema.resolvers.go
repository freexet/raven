package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/freexet/raven/auth"
	"github.com/freexet/raven/graph/generated"
	"github.com/freexet/raven/graph/model"
	"github.com/gin-gonic/gin"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, params model.NewUser) (*auth.User, error) {
	gc := ctx.Value(ContextKey{Name: "ginCtx"}).(*gin.Context)
	a, _ := gc.Get("auth")
	return a.(auth.Service).Register(params.Username, params.Password)
}

func (r *mutationResolver) Login(ctx context.Context, params model.Login) (*auth.User, error) {
	gc := ctx.Value(ContextKey{Name: "ginCtx"}).(*gin.Context)
	a, _ := gc.Get("auth")
	return a.(auth.Service).Login(params.Username, params.Password)
}

func (r *queryResolver) Users(ctx context.Context) ([]*auth.User, error) {
	obj, err := Authenticate(ctx, func(ctx *gin.Context, user *auth.User) (interface{}, error) {
		return []*auth.User{user}, nil
	})

	if err != nil {
		return nil, err
	}

	return obj.([]*auth.User), err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
