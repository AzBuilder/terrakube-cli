package models

type GetBodyVariable struct {

	// data
	Data []*Variable `json:"data"`
}

type PostBodyVariable struct {

	// data
	Data *Variable `json:"data"`
}
type Variable struct {

	// attributes
	Attributes *VariableAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// relationships
	Relationships *VariableRelationships `json:"relationships,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

type VariableAttributes struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// category
	Category string `json:"category,omitempty"`

	// sensitive
	Sensitive bool `json:"sensitive,omitempty"`

	// hcl
	Hcl bool `json:"hcl,omitempty"`
}

type VariableRelationships struct {

	// workspace
	Workspace *VariableRelationshipsWorkspace `json:"workspace,omitempty"`
}

type VariableRelationshipsWorkspace struct {

	// id
	ID string `json:"id,omitempty"`

	// type
	// Enum: [workspace]
	Type string `json:"type,omitempty"`
}
