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

type Webp struct {
	Enabled bool `json:"enabled"`
}

type ResourcePattern struct {
	Pattern string `json:"pattern"`
	Rank    int    `json:"rank"`
}

type GlobalDuration struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type Directive struct {
	ResourcePatterns        []ResourcePattern `json:"resourcePatterns"`
	DefaultResourcePatterns []string          `json:"defaultResourcePatterns"`
	GlobalDuration          GlobalDuration    `json:"globalDuration"`
}

type Cache struct {
	Enabled   bool      `json:"enabled"`
	Webp      Webp      `json:"webp"`
	Directive Directive `json:"directive"`
}

type CrsThematics struct {
	ScannerDetection    bool `json:"scannerDetection"`
	ProtocolEnforcement bool `json:"protocolEnforcement"`
	ProtocolAttack      bool `json:"protocolAttack"`
	Lfi                 bool `json:"lfi"`
	Rfi                 bool `json:"rfi"`
	Rce                 bool `json:"rce"`
	PhpInjection        bool `json:"phpInjection"`
	Xss                 bool `json:"xss"`
	Sqli                bool `json:"sqli"`
	SessionFixation     bool `json:"sessionFixation"`
	GeneralDataLeakages bool `json:"generalDataLeakages"`
	SqlDataLeakages     bool `json:"sqlDataLeakages"`
	JavaDataLeakages    bool `json:"javaDataLeakages"`
	PhpDataLeakages     bool `json:"phpDataLeakages"`
	IisDataLeakages     bool `json:"iisDataLeakages"`
}

type Waf struct {
	Enabled             bool         `json:"enabled"`
	DetectionOnly       bool         `json:"detectionOnly"`
	CrsSensitivityLevel bool         `json:"crsSensitivityLevel"`
	CrsThematics        CrsThematics `json:"crsThematics"`
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

func (c *Client) getWithNamespace(namespace string, url string, data interface{}) error {
	res, err := c.requestWithNamespace(namespace).Get(url)

	if err != nil {
		return fmt.Errorf("error retrieving %w: %w", url, err)
	}

	if !res.IsSuccess() {
		return fmt.Errorf("error retrieving %w: "+res.Status, url)
	}

	res.UnmarshalJson(data)

	return nil
}

func (c *Client) GetOrigin(namespace string) (*Origin, error) {
	var origin Origin
	err := c.getWithNamespace(namespace, "/api/configs/origin", &origin)

	if err != nil {
		return nil, err
	}

	return &origin, nil
}

func (c *Client) GetCache(namespace string) (*Cache, error) {
	var cache Cache
	err := c.getWithNamespace(namespace, "/api/configs/cache", &cache)

	if err != nil {
		return nil, err
	}

	return &cache, nil
}

func (c *Client) GetWaf(namespace string) (*Waf, error) {
	var waf Waf
	err := c.getWithNamespace(namespace, "/api/configs/waf", &waf)

	if err != nil {
		return nil, err
	}

	return &waf, nil
}

func (c *Client) GetCustomStaticRules(namespace string) ([]CustomStaticRule, error) {
	var customStaticRules []CustomStaticRule
	err := c.getWithNamespace(namespace, "/api/configs/custom-static-rules", &customStaticRules)

	if err != nil {
		return nil, err
	}

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
