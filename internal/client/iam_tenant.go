package client

import "context"

type TenantClient struct {
	c *Client
}

func (i *IAM) Tenant() *TenantClient {
	return &TenantClient{i.c}
}

type Tenant struct {
	ID        string `terraform:"id"`
	Name      string `terraform:"name"`
	SNC       bool   `terraform:"snc"`
	CompanyID string `terraform:"company_id"`
}

func (t *TenantClient) List(ctx context.Context) ([]*Tenant, error) {
	r := t.c.newRequest("GET", "/api/iam/v2/tenants")
	resp, err := t.c.doRequest(ctx, r)
	if err != nil {
		return nil, err
	}
	defer closeResponseBody(resp)
	if err := requireOK(resp); err != nil {
		return nil, err
	}

	var out []*Tenant
	if err := decodeBody(resp, &out); err != nil {
		return nil, err
	}

	return out, nil
}
