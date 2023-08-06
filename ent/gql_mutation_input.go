// Code generated by ent, DO NOT EDIT.

package ent

// CreateExpertInput represents a mutation input for creating experts.
type CreateExpertInput struct {
	Name string
}

// Mutate applies the CreateExpertInput on the ExpertMutation builder.
func (i *CreateExpertInput) Mutate(m *ExpertMutation) {
	m.SetName(i.Name)
}

// SetInput applies the change-set in the CreateExpertInput on the ExpertCreate builder.
func (c *ExpertCreate) SetInput(i CreateExpertInput) *ExpertCreate {
	i.Mutate(c.Mutation())
	return c
}
