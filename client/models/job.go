package models

type GetBodyJob struct {

	// data
	Data []*Job `json:"data"`
}

type PostBodyJob struct {

	// data
	Data *Job `json:"data"`
}

type Job struct {

	// attributes
	Attributes *JobAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// relationships
	Relationships *JobRelationships `json:"relationships,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

type JobAttributes struct {

	// command
	// Enum: [plan apply destroy]
	Command string `json:"command,omitempty"`

	// output
	Output string `json:"output,omitempty"`

	// status
	// Enum: [pending queue running completed]
	Status string `json:"status,omitempty"`
}

type JobRelationships struct {

	// organization
	Organization *JobRelationshipsOrganization `json:"organization,omitempty"`

	// workspace
	Workspace *JobRelationshipsWorkspace `json:"workspace,omitempty"`
}

type JobRelationshipsOrganization struct {

	// id
	ID string `json:"id,omitempty"`

	// type
	// Enum: [organization]
	Type string `json:"type,omitempty"`
}

type JobRelationshipsWorkspace struct {

	// id
	ID string `json:"id,omitempty"`

	// type
	// Enum: [workspace]
	Type string `json:"type,omitempty"`
}
