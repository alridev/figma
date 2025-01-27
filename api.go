package figma

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type FigmaAPI struct {
	// Token - токен для доступа к API Figma
	Token string

	// BaseURL - базовый URL API Figma (default: https://www.figma.com)
	BaseURL string

	// UserAgent - User-Agent по умолчанию для запросов к API (default: Figma-Go-Client)
	UserAgent string

	// AdditionalHeaders - дополнительные заголовки для запросов к API
	AdditionalHeaders map[string]string
	AdditionalQueries map[string]string

	// Timeout - таймаут для запросов к API (default: 10 seconds)
	Timeout time.Duration
}

func NewFigmaAPI(token string) *FigmaAPI {
	return &FigmaAPI{
		Token:     token,
		BaseURL:   BaseURL,
		UserAgent: UserAgent,
		Timeout:   Timeout,
	}
}

func (api *FigmaAPI) SetAdditionalHeaders(headers map[string]string) {
	api.AdditionalHeaders = headers
}

func (api *FigmaAPI) SetAdditionalQueries(queries map[string]string) {
	api.AdditionalQueries = queries
}

func (api *FigmaAPI) SetTimeout(timeout time.Duration) {
	api.Timeout = timeout
}

func (api *FigmaAPI) SetBaseURL(baseURL string) {
	api.BaseURL = baseURL
}

func (api *FigmaAPI) SetUserAgent(userAgent string) {
	api.UserAgent = userAgent
}

func (api *FigmaAPI) GetClient() *http.Client {
	return &http.Client{
		Timeout: api.Timeout,
	}
}

func (api *FigmaAPI) ActivateRequest(request *http.Request) {

	// add additional headers
	api.AddHeaders(request, api.AdditionalHeaders)

	// add additional queries
	for key, value := range api.AdditionalQueries {
		api.AddQuery(request, key, value)
	}

	// add X-FIGMA-TOKEN
	api.AddHeader(request, "X-FIGMA-TOKEN", api.Token)

	// add User-Agent
	api.AddHeader(request, "User-Agent", api.UserAgent)

}

func (api *FigmaAPI) GetResponse(request *http.Request) (*http.Response, error) {
	client := api.GetClient()
	api.ActivateRequest(request)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"figma api error: %d %s",
			response.StatusCode,
			FigmaErrorDescriptions[FigmaCodeError(response.StatusCode)],
		)
	}
	return response, nil
}

func (api *FigmaAPI) GetResponseNoActivate(request *http.Request) (*http.Response, error) {
	client := api.GetClient()
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (api *FigmaAPI) GetResponseBody(response *http.Response) ([]byte, error) {
	return GetResponseBody(response)
}

func GetResponseBody(response *http.Response) ([]byte, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (api *FigmaAPI) ParseJsonResponse(response *http.Response, to interface{}) error {
	return ParseJsonResponse(response, to)
}

func ParseJsonResponse(response *http.Response, to interface{}) error {
	body, err := GetResponseBody(response)
	if err != nil {
		return err
	}
	return ParseJsonFromBytes(body, to)
}

func ParseJsonFromBytes(body []byte, to interface{}) error {

	// Check if this is a oneOf/anyOf schema with GetActualInstance
	if actualObj, ok := to.(interface{ GetActualInstance() interface{} }); ok {
		// Verify it has UnmarshalJSON defined
		if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON(data []byte) error }); ok {
			return unmarshalObj.UnmarshalJSON(body)
		}
		return fmt.Errorf("no UnmarshalJSON method found")
	}

	err := json.Unmarshal(body, &to)
	if err != nil {
		return err
	}
	return nil
}

func (api *FigmaAPI) AddQuery(request *http.Request, key string, value string) {
	q := request.URL.Query()
	q.Add(key, value)
	request.URL.RawQuery = q.Encode()
}

func (api *FigmaAPI) AddQueries(request *http.Request, queries map[string]string) {
	for key, value := range queries {
		api.AddQuery(request, key, value)
	}
}

func (api *FigmaAPI) AddHeader(request *http.Request, key string, value string) {
	request.Header.Add(key, value)
}

func (api *FigmaAPI) AddHeaders(request *http.Request, headers map[string]string) {
	for key, value := range headers {
		api.AddHeader(request, key, value)
	}
}
