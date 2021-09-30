package client

import (
	"terrakube/client/models"
	"fmt"
	"net/http"
)

type VariableClient struct {
	Client *Client
}

func (c *VariableClient) List(organizationId string, workspaceId string, filter string) ([]*models.Variable, error) {
	req, err := c.Client.newRequest(http.MethodGet, fmt.Sprintf("organization/%v/workspace/%v/variable", organizationId, workspaceId), nil)
	if err != nil {
		return nil, err
	}
	var resp models.GetBodyVariable
	_, err = c.Client.do(req, &resp)
	return resp.Data, err
}

func (c *VariableClient) Create(organizationId string, workspaceId string, variable models.Variable) (*models.Variable, error) {
	reqBody := models.PostBodyVariable{
		Data: &variable,
	}

	req, err := c.Client.newRequest(http.MethodPost, fmt.Sprintf("organization/%v/workspace/%v/variable", organizationId, workspaceId), reqBody)
	if err != nil {
		return nil, err
	}
	var resp models.PostBodyVariable
	_, err = c.Client.do(req, &resp)
	return resp.Data, err
}

func (c *VariableClient) Delete(organizationId string, workspaceId string, variableId string) error {
	req, err := c.Client.newRequest(http.MethodDelete, fmt.Sprintf("organization/%v/workspace/%v/variable/%v", organizationId, workspaceId, variableId), nil)
	if err != nil {
		return err
	}

	_, err = c.Client.do(req, nil)
	return err
}

func (c *VariableClient) Update(organizationId string, workspaceId string, variable models.Variable) error {
	reqBody := models.PostBodyVariable{
		Data: &variable,
	}

	req, err := c.Client.newRequest(http.MethodPatch, fmt.Sprintf("organization/%v/workspace/%v/variable/%v", organizationId, workspaceId, variable.ID), reqBody)
	if err != nil {
		return err
	}

	_, err = c.Client.do(req, nil)
	return err
}
