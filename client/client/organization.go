package client

import (
	"azb/client/models"
	"fmt"
	"net/http"
)

type OrganizationClient struct {
	Client *Client
}

func (c *OrganizationClient) List(filter string) ([]*models.Organization, error) {
	req, err := c.Client.newRequest(http.MethodGet, basePath+"organization", nil)
	if err != nil {
		return nil, err
	}
	var organizationResp models.GetBodyOrganization
	_, err = c.Client.do(req, &organizationResp)
	return organizationResp.Data, err
}

func (c *OrganizationClient) Create(organization models.Organization) (*models.Organization, error) {
	reqBody := models.PostBodyOrganization{
		Data: &organization,
	}

	req, err := c.Client.newRequest(http.MethodPost, basePath+"organization", reqBody)
	if err != nil {
		return nil, err
	}
	var organizationResp models.PostBodyOrganization
	_, err = c.Client.do(req, &organizationResp)
	return organizationResp.Data, err
}

func (c *OrganizationClient) Delete(organizationID string) error {
	req, err := c.Client.newRequest(http.MethodDelete, fmt.Sprintf(basePath+"organization/%v", organizationID), nil)
	if err != nil {
		return err
	}

	_, err = c.Client.do(req, nil)
	return err
}

func (c *OrganizationClient) Update(organization models.Organization) error {
	reqBody := models.PostBodyOrganization{
		Data: &organization,
	}

	req, err := c.Client.newRequest(http.MethodPatch, fmt.Sprintf(basePath+"organization/%v", organization.ID), reqBody)
	if err != nil {
		return err
	}

	_, err = c.Client.do(req, nil)
	return err
}
