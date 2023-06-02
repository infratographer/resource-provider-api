package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"

	"go.infratographer.com/resource-provider-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// FindOrganizationalUnitByID is the resolver for the findOrganizationalUnitByID field.
func (r *entityResolver) FindOrganizationalUnitByID(ctx context.Context, id gidx.PrefixedID) (*OrganizationalUnit, error) {
	return &OrganizationalUnit{ID: id}, nil
}

// FindResourceProviderByID is the resolver for the findResourceProviderByID field.
func (r *entityResolver) FindResourceProviderByID(ctx context.Context, id gidx.PrefixedID) (*generated.ResourceProvider, error) {
	return r.client.ResourceProvider.Get(ctx, id)
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }