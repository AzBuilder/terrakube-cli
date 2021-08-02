package models

type GetBodyEnvironment struct {

	// data
	Data []*Environment `json:"data"`
}

type PostBodyEnvironment struct {

	// data
	Data *Environment `json:"data"`
}

type Environment struct {

	// attributes
	Attributes *EnvironmentAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// relationships
	Relationships *EnvironmentRelationships `json:"relationships,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

type EnvironmentAttributes struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

type EnvironmentRelationships struct {

	// workspace
	Workspace *EnvironmentRelationshipsWorkspace `json:"workspace,omitempty"`
}

type EnvironmentRelationshipsWorkspace struct {

	// id
	ID string `json:"id,omitempty"`

	// type
	// Enum: [workspace]
	Type string `json:"type,omitempty"`
}
