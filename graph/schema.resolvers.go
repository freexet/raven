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

func (r *mutationResolver) GenerateOtp(ctx context.Context) (*model.Otp, error) {
	otp, err := Authenticate(ctx, func(ctx *gin.Context, user *auth.User) (interface{}, error) {
		a, _ := ctx.Get("auth")

		secret, imgData, err := a.(auth.Service).GenerateOTP(user)
		if err != nil {
			return nil, err
		}

		return &model.Otp{SecretKey: secret, ImgData: imgData}, nil
	})

	if err != nil {
		return nil, err
	}

	return otp.(*model.Otp), err
}

func (r *mutationResolver) ValidateOtp(ctx context.Context, code string) (*auth.User, error) {
	user, err := Authenticate(ctx, func(ctx *gin.Context, user *auth.User) (interface{}, error) {
		a, _ := ctx.Get("auth")

		err := a.(auth.Service).ValidateOTP(code, user)
		if err != nil {
			return nil, err
		}

		return user, nil
	})

	if err != nil {
		return nil, err
	}

	return user.(*auth.User), err
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
