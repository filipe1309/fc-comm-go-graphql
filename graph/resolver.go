package graph

import "go-graphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Categories []*model.Category // In memory, could be a db repository, or an API...
	Courses    []*model.Course
	Chapters   []*model.Chapter
}
