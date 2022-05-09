package baleen_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
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

type CrsThematicStatus struct {
	ID      string
	Enabled bool
}

type Waf struct {
	Enabled             bool `json:"enabled"`
	DetectionOnly       bool `json:"detectionOnly"`
	CrsSensitivityLevel bool `json:"crsSensitivityLevel"`
	CrsThematics        []CrsThematicStatus
}

type Headers struct {
	DenyFrameOptions bool `json:"denyFrameOptions"`
	NoSniffMimeType  bool `json:"noSniffMimeType"`
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

type RedirectRule struct {
	Source          string `json:"source"`
	Destination     string `json:"destination"`
	Type            int    `json:"type"`
	WithQueryString bool   `json:"withQueryString"`
}

type RewriteRule struct {
	Source          string `json:"source"`
	Destination     string `json:"destination"`
	WithQueryString bool   `json:"withQueryString"`
}

type UrlRules struct {
	RedirectRules []RedirectRule `json:"redirectRules"`
	RewriteRules  []RewriteRule  `json:"rewriteRules"`
}

type CrsThematic struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Group       string `json:"group"`
}

type AccessLog struct {
	Timestamp               time.Time `json:"timestamp"`
	Status                  string    `json:"status"`
	RemoteAddr              net.IP    `json:"remoteAddr"`
	Upstream                string    `json:"upstream"`
	Scheme                  string    `json:"scheme"`
	RequestFateAction       string    `json:"requestFateAction"`
	BodyBytesSent           string    `json:"bodyBytesSent"`
	BotCategory             string    `json:"botCategory"`
	TriggeredRuleId         string    `json:"triggeredRuleId"`
	TriggeredRuleTrackingID string    `json:"triggeredRuleTrackingId"`
	HttpHost                string    `json:"httpHost"`
	HttpUserAgent           string    `json:"httpUserAgent"`
	RemoteUser              string    `json:"remoteUser"`
	RequestTime             string    `json:"requestTime"`
	ClientIP                net.IP    `json:"clientIp"`
	HttpXForwardedFor       string    `json:"httpXForwardedFor"`
	HttpReferrer            string    `json:"httpReferrer"`
	UpstreamResponseTime    string    `json:"upstreamResponseTime"`
	RequestFate             string    `json:"requestFate"`
	RequastJa3              string    `json:"requestJa3"`
	SslProtocol             string    `json:"sslProtocol"`
	ServerProtocol          string    `json:"serverProtocol"`
	RequestUri              string    `json:"requestUri"`
	RequestMethod           string    `json:"requestMethod"`
	RequestIsp              string    `json:"requestIsp"`
	RequestCountry          string    `json:"requestCountry"`
	RequestAsn              string    `json:"requestAsn"`
	RequestConnectionType   string    `json:"requestConnectionType"`
	RequestIsAnonymousProxy string    `json:"requestIsAnonymousProxy"`
	ResponseContentType     string    `json:"responseContentType"`
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

func (c *Client) getWithNamespace(namespace string, url string, data interface{}) (*req.Response, error) {
	res, err := c.requestWithNamespace(namespace).Get(url)

	if err != nil {
		return res, fmt.Errorf("error retrieving %w: %w", url, err)
	}

	if !res.IsSuccess() {
		return res, fmt.Errorf("error retrieving %w: "+res.Status, url)
	}

	res.UnmarshalJson(data)

	return res, nil
}

func (c *Client) GetOrigin(namespace string) (*Origin, error) {
	var origin Origin
	_, err := c.getWithNamespace(namespace, "/api/configs/origin", &origin)

	if err != nil {
		return nil, err
	}

	return &origin, nil
}

func (c *Client) GetCache(namespace string) (*Cache, error) {
	var cache Cache
	_, err := c.getWithNamespace(namespace, "/api/configs/cache", &cache)

	if err != nil {
		return nil, err
	}

	return &cache, nil
}

func (c *Client) GetWaf(namespace string) (*Waf, error) {
	var waf Waf
	res, err := c.getWithNamespace(namespace, "/api/configs/waf", &waf)

	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	bytes, err := res.ToBytes()

	if err != nil {
		return nil, err
	}

	json.Unmarshal(bytes, &result)

	crsThematicsObject := result["crsThematics"].(map[string]interface{})

	crsThematicsStatuses := []CrsThematicStatus{}

	for id, enabled := range crsThematicsObject {
		crsThematic := CrsThematicStatus{ID: id, Enabled: enabled.(bool)}
		crsThematicsStatuses = append(crsThematicsStatuses, crsThematic)
	}
	waf.CrsThematics = crsThematicsStatuses

	return &waf, nil
}

func (c *Client) GetCrsThematics(namespace string) ([]CrsThematic, error) {
	var crsThematics []CrsThematic
	_, err := c.getWithNamespace(namespace, "/api/refs/waf/crs-thematics", &crsThematics)

	if err != nil {
		return nil, err
	}

	return crsThematics, nil
}

func (c *Client) GetHeaders(namespace string) (*Headers, error) {
	var headers Headers
	_, err := c.getWithNamespace(namespace, "/api/configs/headers", &headers)

	if err != nil {
		return nil, err
	}

	return &headers, nil
}

func (c *Client) GetCustomStaticRules(namespace string) ([]CustomStaticRule, error) {
	var customStaticRules []CustomStaticRule
	_, err := c.getWithNamespace(namespace, "/api/configs/custom-static-rules", &customStaticRules)

	if err != nil {
		return nil, err
	}

	return customStaticRules, nil
}

func (c *Client) GetUrlRules(namespace string) (*UrlRules, error) {
	var urlRules UrlRules
	_, err := c.getWithNamespace(namespace, "/api/configs/url-rules", &urlRules)

	if err != nil {
		return nil, err
	}

	return &urlRules, nil
}

type AccessLogParams struct {
	Start int
	End   int
	Size  int
	Page  int
}

type PaginationInfo struct {
	TotalCount int
	TotalHits  int
}

func (c *Client) GetAccessLogs(namespace string, params AccessLogParams) ([]AccessLog, *PaginationInfo, error) {
	var accessLogs []AccessLog
	res, err := c.requestWithNamespace(namespace).SetQueryParams(map[string]string{
		"start": strconv.Itoa(params.Start),
		"end":   strconv.Itoa(params.End),
		"size":  strconv.Itoa(params.Size),
		"page":  strconv.Itoa(params.Page),
	}).SetBodyJsonString(`{"filters":[]}`).Post("/api/logs/access-logs")

	if err != nil {
		return nil, nil, err
	}

	if !res.IsSuccess() {
		return nil, nil, fmt.Errorf("error retrieving %w: "+res.Status, "access-logs")
	}

	res.UnmarshalJson(&accessLogs)

	totalCount, _ := strconv.Atoi(res.Header.Get("x-total-count"))
	totalHits, _ := strconv.Atoi(res.Header.Get("x-total-hits"))
	pagination := PaginationInfo{
		TotalCount: totalCount,
		TotalHits:  totalHits,
	}

	return accessLogs, &pagination, nil
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
