package baleen_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	req "github.com/imroc/req/v3"
)

type Client struct {
	c *req.Client
}

type Namespace struct {
	ID   string
	Name string
}

type Account struct {
	ID          string    `json:"id"`
	Activated   bool      `json:"activated"`
	CreatedDate time.Time `json:"createdDate"`
	Namespaces  []Namespace
}

type ErrorPages struct {
	Custom404Page bool `json:"custom404Page"`
	Custom50xPage bool `json:"custom50xPage"`
}

type Origin struct {
	ErrorPages ErrorPages `json:"errorPages"`
	URL        string     `json:"url"`
}

type Condition struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Operator string `json:"operator"`
}

type CustomStaticRule struct {
	ID          string        `json:"id"`
	Category    string        `json:"category"`
	TrackingID  string        `json:"trackingId"`
	Enabled     bool          `json:"enabled"`
	Description string        `json:"description"`
	Conditions  [][]Condition `json:"conditions"`
	Labels      []string      `json:"labels"`
}

type ClientOption func(c *Client)

func New(options ...ClientOption) *Client {
	c := &Client{
		c: req.C(),
	}
	c.c.SetBaseURL("https://console.baleen.cloud")

	for _, o := range options {
		o(c)
	}

	return c
}

func (c *Client) GetAccount() (*Account, error) {
	res, err := c.r().Get("/api/account")

	if err != nil {
		return nil, fmt.Errorf("error retrieving account: %w", err)
	}

	if !res.IsSuccess() {
		return nil, errors.New("error retrieving account: " + res.Status)
	}

	account := new(Account)

	res.UnmarshalJson(account)
	var result map[string]interface{}
	bytes, err := res.ToBytes()
	json.Unmarshal(bytes, &result)

	namespacesObject := result["namespaces"].(map[string]interface{})

	namespaces := []Namespace{}

	for id, name := range namespacesObject {
		namespace := Namespace{ID: id, Name: name.(string)}
		namespaces = append(namespaces, namespace)
	}
	account.Namespaces = namespaces

	return account, nil
}

func (c *Client) requestWithNamespace(namespace string) *req.Request {
	return c.r().SetCookies(
		&http.Cookie{
			Name:  "baleen-namespace",
			Value: namespace,
		})
}

func (c *Client) GetOrigin(namespace string) (*Origin, error) {
	res, err := c.requestWithNamespace(namespace).Get("/api/configs/origin")

	if err != nil {
		return nil, fmt.Errorf("error retrieving origin: %w", err)
	}

	if !res.IsSuccess() {
		return nil, errors.New("error retrieving origin: " + res.Status)
	}

	origin := new(Origin)

	res.UnmarshalJson(origin)

	return origin, nil
}

func (c *Client) GetCustomStaticRules(namespace string) ([]CustomStaticRule, error) {
	res, err := c.requestWithNamespace(namespace).Get("/api/configs/custom-static-rules")

	if err != nil {
		return nil, fmt.Errorf("error retrieving custom-static-rules: %w", err)
	}

	if !res.IsSuccess() {
		return nil, errors.New("error retrieving custom-static-rules: " + res.Status)
	}

	customStaticRules := []CustomStaticRule{}

	res.UnmarshalJson(&customStaticRules)

	return customStaticRules, nil
}

func (c *Client) r() *req.Request {
	return c.c.R()
}

func (c *Client) setToken(token string) *Client {
	c.c.SetCommonHeader("X-Api-Key", token)
	return c
}

func WithToken(apiKey string) ClientOption {
	return func(c *Client) {
		c.setToken(apiKey)
	}
}
