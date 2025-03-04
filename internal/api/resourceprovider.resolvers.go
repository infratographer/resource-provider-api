package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"

	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/resource-provider-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// ResourceProviderCreate is the resolver for the resourceProviderCreate field.
func (r *mutationResolver) ResourceProviderCreate(ctx context.Context, input generated.CreateResourceProviderInput) (*ResourceProviderCreatePayload, error) {
	if err := permissions.CheckAccess(ctx, input.OwnerID, actionResourceProviderCreate); err != nil {
		return nil, err
	}

	t, err := r.client.ResourceProvider.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &ResourceProviderCreatePayload{ResourceProvider: t}, err
}

// ResourceProviderUpdate is the resolver for the resourceProviderUpdate field.
func (r *mutationResolver) ResourceProviderUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateResourceProviderInput) (*ResourceProviderUpdatePayload, error) {
	if err := permissions.CheckAccess(ctx, id, actionResourceProviderUpdate); err != nil {
		return nil, err
	}

	t, err := r.client.ResourceProvider.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	t, err = t.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &ResourceProviderUpdatePayload{ResourceProvider: t}, err
}

// ResourceProviderDelete is the resolver for the resourceProviderDelete field.
func (r *mutationResolver) ResourceProviderDelete(ctx context.Context, id gidx.PrefixedID) (*ResourceProviderDeletePayload, error) {
	if err := permissions.CheckAccess(ctx, id, actionResourceProviderDelete); err != nil {
		return nil, err
	}

	if err := r.client.ResourceProvider.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, err
	}

	return &ResourceProviderDeletePayload{DeletedID: id}, nil
}

// ResourceProvider is the resolver for the resourceProvider field.
func (r *queryResolver) ResourceProvider(ctx context.Context, id gidx.PrefixedID) (*generated.ResourceProvider, error) {
	if err := permissions.CheckAccess(ctx, id, actionResourceProviderGet); err != nil {
		return nil, err
	}

	return r.client.ResourceProvider.Get(ctx, id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
