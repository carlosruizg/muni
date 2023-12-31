package muni

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"fmt"

	"github.com/carlosruizg/muni/ent"
)

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, input ent.CreateLabellingProjectInput) (*ent.LabellingProject, error) {
	return r.client.LabellingProject.Create().SetInput(input).Save(ctx)
}

// UpdateProject is the resolver for the updateProject field.
func (r *mutationResolver) UpdateProject(ctx context.Context, id int, input ent.UpdateLabellingProjectInput) (*ent.LabellingProject, error) {
	return r.client.LabellingProject.UpdateOneID(id).SetInput(input).Save(ctx)
}

// CreateExpert is the resolver for the createExpert field.
func (r *mutationResolver) CreateExpert(ctx context.Context, input ent.CreateExpertInput) (*ent.Expert, error) {
	panic(fmt.Errorf("not implemented: CreateExpert - createExpert"))
}

// UpdateExpert is the resolver for the updateExpert field.
func (r *mutationResolver) UpdateExpert(ctx context.Context, id int, input ent.UpdateExpertInput) (*ent.Expert, error) {
	panic(fmt.Errorf("not implemented: UpdateExpert - updateExpert"))
}

// CreateLabellingTask is the resolver for the createLabellingTask field.
func (r *mutationResolver) CreateLabellingTask(ctx context.Context, input ent.CreateLabellingTaskInput) (*ent.LabellingTask, error) {
	panic(fmt.Errorf("not implemented: CreateLabellingTask - createLabellingTask"))
}

// UpdateLabellingTask is the resolver for the updateLabellingTask field.
func (r *mutationResolver) UpdateLabellingTask(ctx context.Context, id int, input ent.UpdateLabellingTaskInput) (*ent.LabellingTask, error) {
	panic(fmt.Errorf("not implemented: UpdateLabellingTask - updateLabellingTask"))
}

// CreateLabellingTaskResponse is the resolver for the createLabellingTaskResponse field.
func (r *mutationResolver) CreateLabellingTaskResponse(ctx context.Context, input ent.CreateLabellingTaskResponseInput) (*ent.LabellingTaskResponse, error) {
	panic(fmt.Errorf("not implemented: CreateLabellingTaskResponse - createLabellingTaskResponse"))
}

// UpdateLabellingTaskResponse is the resolver for the updateLabellingTaskResponse field.
func (r *mutationResolver) UpdateLabellingTaskResponse(ctx context.Context, id int, input ent.UpdateLabellingTaskResponseInput) (*ent.LabellingTaskResponse, error) {
	panic(fmt.Errorf("not implemented: UpdateLabellingTaskResponse - updateLabellingTaskResponse"))
}

// CreateQualification is the resolver for the createQualification field.
func (r *mutationResolver) CreateQualification(ctx context.Context, input ent.CreateQualificationInput) (*ent.LabellingTaskResponse, error) {
	panic(fmt.Errorf("not implemented: CreateQualification - createQualification"))
}

// UpdateQualification is the resolver for the updateQualification field.
func (r *mutationResolver) UpdateQualification(ctx context.Context, id int, input ent.UpdateQualificationInput) (*ent.LabellingTaskResponse, error) {
	panic(fmt.Errorf("not implemented: UpdateQualification - updateQualification"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
