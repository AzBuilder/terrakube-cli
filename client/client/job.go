package client

import (
	"azb/client/models"
	"fmt"
	"net/http"
)

type JobClient struct {
	Client *Client
}

func (c *JobClient) List(organizationId string, filter string) ([]*models.Job, error) {
	req, err := c.Client.newRequest(http.MethodGet, basePath+fmt.Sprintf("organization/%v/job", organizationId), nil)
	if err != nil {
		return nil, err
	}
	var resp models.GetBodyJob
	_, err = c.Client.do(req, &resp)
	return resp.Data, err
}

func (c *JobClient) Create(organizationId string, job models.Job) (*models.Job, error) {
	reqBody := models.PostBodyJob{
		Data: &job,
	}

	req, err := c.Client.newRequest(http.MethodPost, basePath+fmt.Sprintf("organization/%v/job", organizationId), reqBody)
	if err != nil {
		return nil, err
	}
	var resp models.PostBodyJob
	_, err = c.Client.do(req, &resp)
	return resp.Data, err
}

func (c *JobClient) Delete(organizationID string, moduleId string) error {
	req, err := c.Client.newRequest(http.MethodDelete, fmt.Sprintf(basePath+"organization/%v/module/%v", organizationID, moduleId), nil)
	if err != nil {
		return err
	}

	_, err = c.Client.do(req, nil)
	return err
}

func (c *JobClient) Update(organizationId string, module models.Module) error {
	reqBody := models.PostBodyModule{
		Data: &module,
	}

	req, err := c.Client.newRequest(http.MethodPatch, fmt.Sprintf(basePath+"organization/%v/module/%v", organizationId, module.ID), reqBody)
	if err != nil {
		return err
	}

	_, err = c.Client.do(req, nil)
	return err
}
