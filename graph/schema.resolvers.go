package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/felipehirano/go-graphQL/graph/generated"
	"github.com/felipehirano/go-graphQL/graph/model"
)

// Courses is the resolver for the courses field.
func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	var courses []*model.Course

	for _, value := range r.Resolver.Courses {
		if value.Category.ID == obj.ID {
			courses = append(courses, value)
		}
	}

	return courses, nil
}

// Chapters is the resolver for the chapters field.
func (r *courseResolver) Chapters(ctx context.Context, obj *model.Course) ([]*model.Chapter, error) {
	var chapters []*model.Chapter

	for _, value := range r.Resolver.Chapters {
		if value.Course.ID == obj.ID {
			chapters = append(chapters, value)
		}
	}

	return chapters, nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input *model.NewCategory) (*model.Category, error) {
	category := model.Category{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Description: &input.Description,
	}

	r.Categories = append(r.Categories, &category)

	return &category, nil
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input *model.NewCourse) (*model.Course, error) {
	var category *model.Category

	for _, value := range r.Categories {
		if value.ID == input.CategoryID {
			category = value
			break
		}
	}

	course := model.Course{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Description: &input.Description,
		Category:    category,
	}

	r.Courses = append(r.Courses, &course)

	return &course, nil
}

// CreateChapter is the resolver for the createChapter field.
func (r *mutationResolver) CreateChapter(ctx context.Context, input *model.NewChapter) (*model.Chapter, error) {
	var course *model.Course

	for _, value := range r.Courses {
		if value.ID == input.CourseID {
			course = value
			break
		}
	}

	chapter := model.Chapter{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Name:   input.Name,
		Course: course,
	}

	r.Chapters = append(r.Chapters, &chapter)

	return &chapter, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.Resolver.Categories, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return r.Resolver.Courses, nil
}

// Chapters is the resolver for the chapters field.
func (r *queryResolver) Chapters(ctx context.Context) ([]*model.Chapter, error) {
	return r.Resolver.Chapters, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Course returns generated.CourseResolver implementation.
func (r *Resolver) Course() generated.CourseResolver { return &courseResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type courseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
