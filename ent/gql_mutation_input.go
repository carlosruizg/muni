// Code generated by ent, DO NOT EDIT.

package ent

// CreateExpertInput represents a mutation input for creating experts.
type CreateExpertInput struct {
	Name             string
	TaskResponseIDs  []int
	QualificationIDs []int
}

// Mutate applies the CreateExpertInput on the ExpertMutation builder.
func (i *CreateExpertInput) Mutate(m *ExpertMutation) {
	m.SetName(i.Name)
	if v := i.TaskResponseIDs; len(v) > 0 {
		m.AddTaskResponseIDs(v...)
	}
	if v := i.QualificationIDs; len(v) > 0 {
		m.AddQualificationIDs(v...)
	}
}

// SetInput applies the change-set in the CreateExpertInput on the ExpertCreate builder.
func (c *ExpertCreate) SetInput(i CreateExpertInput) *ExpertCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateLabellingTaskInput represents a mutation input for creating labellingtasks.
type CreateLabellingTaskInput struct {
	Title                string
	Description          *string
	ResponseIDs          []int
	ExpertRequirementIDs []int
}

// Mutate applies the CreateLabellingTaskInput on the LabellingTaskMutation builder.
func (i *CreateLabellingTaskInput) Mutate(m *LabellingTaskMutation) {
	m.SetTitle(i.Title)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.ResponseIDs; len(v) > 0 {
		m.AddResponseIDs(v...)
	}
	if v := i.ExpertRequirementIDs; len(v) > 0 {
		m.AddExpertRequirementIDs(v...)
	}
}

// SetInput applies the change-set in the CreateLabellingTaskInput on the LabellingTaskCreate builder.
func (c *LabellingTaskCreate) SetInput(i CreateLabellingTaskInput) *LabellingTaskCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateLabellingTaskResponseInput represents a mutation input for creating labellingtaskresponses.
type CreateLabellingTaskResponseInput struct {
	Response string
	TaskID   *int
	ExpertID *int
}

// Mutate applies the CreateLabellingTaskResponseInput on the LabellingTaskResponseMutation builder.
func (i *CreateLabellingTaskResponseInput) Mutate(m *LabellingTaskResponseMutation) {
	m.SetResponse(i.Response)
	if v := i.TaskID; v != nil {
		m.SetTaskID(*v)
	}
	if v := i.ExpertID; v != nil {
		m.SetExpertID(*v)
	}
}

// SetInput applies the change-set in the CreateLabellingTaskResponseInput on the LabellingTaskResponseCreate builder.
func (c *LabellingTaskResponseCreate) SetInput(i CreateLabellingTaskResponseInput) *LabellingTaskResponseCreate {
	i.Mutate(c.Mutation())
	return c
}
