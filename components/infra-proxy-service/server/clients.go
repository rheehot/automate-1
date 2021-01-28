package server

import (
	"context"
	"fmt"

	chef "github.com/go-chef/chef"

	"github.com/chef/automate/api/interservice/infra_proxy/request"
	"github.com/chef/automate/api/interservice/infra_proxy/response"
)

// Client represents api client struct.
type Client struct {
	Name       string `json:"name"`
	ClientName string `json:"clientname"`
	OrgName    string `json:"orgname"`
}

// ClientResult list result from Search API
type ClientResult struct {
	Total int       `json:"total"`
	Start int       `json:"start"`
	Rows  []*Client `json:"rows"`
}

// GetClients get clients list
func (s *Server) GetClients(ctx context.Context, req *request.Clients) (*response.Clients, error) {
	c, err := s.createClient(ctx, req.OrgId, req.ServerId)
	if err != nil {
		return nil, err
	}

	res, err := c.SearchClients(req.SearchQuery)
	if err != nil {
		return nil, ParseAPIError(err)
	}

	return &response.Clients{
		Clients: fromAPIToListClients(res.Rows),
		Start:   int32(res.Start),
		Total:   int32(res.Total),
	}, nil
}

// SearchClients gets client list from Chef Infra Server search API.
func (c *ChefClient) SearchClients(searchQuery *request.SearchQuery) (ClientResult, error) {
	var result ClientResult
	var searchAll bool
	inc := 1000
	var query chef.SearchQuery

	if searchQuery == nil || searchQuery.Q == "" {
		searchAll = true
		query = chef.SearchQuery{
			Index: "client",
			Query: "*:*",
			Start: 0,
			Rows:  inc,
		}
	} else {
		query = chef.SearchQuery{
			Index: "client",
			Query: searchQuery.GetQ(),
			Start: int(searchQuery.GetStart()),
			Rows:  int(searchQuery.GetRows()),
		}
	}

	fullURL := fmt.Sprintf("search/%s", query)
	newReq, err := c.client.NewRequest("GET", fullURL, nil)

	if err != nil {
		return result, ParseAPIError(err)
	}

	res, err := c.client.Do(newReq, &result)
	if err != nil {
		return result, ParseAPIError(err)
	}

	defer res.Body.Close() // nolint: errcheck

	if searchAll {
		var searchResult ClientResult
		start := result.Start
		// the total rows available for this query across all pages
		total := result.Total
		for start+inc <= total {
			query.Start = query.Start + inc
			fullURL = fmt.Sprintf("search/%s", query)

			res1, err := c.client.Do(newReq, &searchResult)
			if err != nil {
				return result, ParseAPIError(err)
			}

			defer res1.Body.Close() // nolint: errcheck

			// add this page of results to the primary SearchResult instance
			result.Rows = append(result.Rows, searchResult.Rows...)
		}
	}
	return result, nil
}

// GetClient get client
func (s *Server) GetClient(ctx context.Context, req *request.Client) (*response.Client, error) {
	c, err := s.createClient(ctx, req.OrgId, req.ServerId)
	if err != nil {
		return nil, err
	}

	ic, err := c.client.Clients.Get(req.Name)
	if err != nil {
		return nil, ParseAPIError(err)
	}

	return &response.Client{
		Name:       ic.Name,
		ClientName: ic.ClientName,
		OrgName:    ic.OrgName,
		Validator:  ic.Validator,
		JsonClass:  ic.JsonClass,
		ChefType:   ic.ChefType,
	}, nil

}

// fromAPIToListClients a response.Clients from a struct of ClientList
func fromAPIToListClients(al []*Client) []*response.ClientListItem {
	cl := make([]*response.ClientListItem, len(al))
	for index, c := range al {
		cl[index] = &response.ClientListItem{
			Name: c.Name,
		}
		index++
	}
	return cl
}
