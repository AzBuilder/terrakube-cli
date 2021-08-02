package models

type GetBodySecret struct {

	// data
	Data []*Secret `json:"data"`
}

type PostBodySecret struct {

	// data
	Data *Secret `json:"data"`
}

type Secret struct {

	// attributes
	Attributes *SecretAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// relationships
	Relationships *SecretRelationships `json:"relationships,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

type SecretAttributes struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

type SecretRelationships struct {

	// workspace
	Workspace *SecretRelationshipsWorkspace `json:"workspace,omitempty"`
}

type SecretRelationshipsWorkspace struct {

	// id
	ID string `json:"id,omitempty"`

	// type
	// Enum: [workspace]
	Type string `json:"type,omitempty"`
}
