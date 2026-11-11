package graph

import "github.com/cujarrett/storefront-reviews/graph/model"

var reviews = []*model.Review{
	{ID: "1", RecordID: "1", Rating: 5, Body: "Essential."},
	{ID: "2", RecordID: "1", Rating: 4, Body: "Cool as it gets."},
	{ID: "3", RecordID: "2", Rating: 5, Body: "Spiritual."},
}
