package client

import (
	"terrakube/client/models"
	"fmt"
	"net/http"
)

type EnvironmentClient struct {
	Client *Client
}

func (c *EnvironmentClient) List(organizationId string, workspaceId string, filter string) ([]*models.Environment, error) {
	req, err := c.Client.newRequest(http.MethodGet, fmt.Sprintf("organization/%v/workspace/%v/environment", organizationId, workspaceId), nil)
	if err != nil {
		return nil, err
	}
	var resp models.GetBodyEnvironment
	_, err = c.Client.do(req, &resp)
	return resp.Data, err
}

func (c *EnvironmentClient) Create(organizationId string, workspaceId string, environment models.Environment) (*models.Environment, error) {
	reqBody := models.PostBodyEnvironment{
		Data: &environment,
	}

	req, err := c.Client.newRequest(http.MethodPost, fmt.Sprintf("organization/%v/workspace/%v/environment", organizationId, workspaceId), reqBody)
	if err != nil {
		return nil, err
	}
	var resp models.PostBodyEnvironment
	_, err = c.Client.do(req, &resp)
	return resp.Data, err
}

func (c *EnvironmentClient) Delete(organizationId string, workspaceId string, environmentId string) error {
	req, err := c.Client.newRequest(http.MethodDelete, fmt.Sprintf("organization/%v/workspace/%v/environment/%v", organizationId, workspaceId, environmentId), nil)
	if err != nil {
		return err
	}

	_, err = c.Client.do(req, nil)
	return err
}

func (c *EnvironmentClient) Update(organizationId string, workspaceId string, environment models.Environment) error {
	reqBody := models.PostBodyEnvironment{
		Data: &environment,
	}

	req, err := c.Client.newRequest(http.MethodPatch, fmt.Sprintf("organization/%v/workspace/%v/environment/%v", organizationId, workspaceId, environment.ID), reqBody)
	if err != nil {
		return err
	}

	_, err = c.Client.do(req, nil)
	return err
}
