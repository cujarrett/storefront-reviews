package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"

	"github.com/cujarrett/storefront-reviews/graph/model"
)

type Resolver struct{}

// FindRecordByID is the resolver for the findRecordByID field.
func (r *entityResolver) FindRecordByID(ctx context.Context, id string) (*model.Record, error) {
	// Another team owns Record. This subgraph only adds its reviews.
	found := []*model.Review{}
	for _, rev := range reviews {
		if rev.RecordID == id {
			found = append(found, rev)
		}
	}
	return &model.Record{ID: id, Reviews: found}, nil
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type Resolver struct{}
*/
