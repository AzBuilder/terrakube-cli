package client

import (
	"azb/client/models"
	"fmt"
	"net/http"
)

type SecretClient struct {
	Client *Client
}

func (c *SecretClient) List(organizationId string, workspaceId string, filter string) ([]*models.Secret, error) {
	req, err := c.Client.newRequest(http.MethodGet, fmt.Sprintf("organization/%v/workspace/%v/secret", organizationId, workspaceId), nil)
	if err != nil {
		return nil, err
	}
	var resp models.GetBodySecret
	_, err = c.Client.do(req, &resp)
	return resp.Data, err
}

func (c *SecretClient) Create(organizationId string, workspaceId string, secret models.Secret) (*models.Secret, error) {
	reqBody := models.PostBodySecret{
		Data: &secret,
	}

	req, err := c.Client.newRequest(http.MethodPost, fmt.Sprintf("organization/%v/workspace/%v/secret", organizationId, workspaceId), reqBody)
	if err != nil {
		return nil, err
	}
	var resp models.PostBodySecret
	_, err = c.Client.do(req, &resp)
	return resp.Data, err
}

func (c *SecretClient) Delete(organizationId string, workspaceId string, secretId string) error {
	req, err := c.Client.newRequest(http.MethodDelete, fmt.Sprintf("organization/%v/workspace/%v/secret/%v", organizationId, workspaceId, secretId), nil)
	if err != nil {
		return err
	}

	_, err = c.Client.do(req, nil)
	return err
}

func (c *SecretClient) Update(organizationId string, workspaceId string, secret models.Secret) error {
	reqBody := models.PostBodySecret{
		Data: &secret,
	}

	req, err := c.Client.newRequest(http.MethodPatch, fmt.Sprintf("organization/%v/workspace/%v/secret/%v", organizationId, workspaceId, secret.ID), reqBody)
	if err != nil {
		return err
	}

	_, err = c.Client.do(req, nil)
	return err
}
