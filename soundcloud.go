package gosoundcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
)

var BaseApiURL = "https://api.soundcloud.com"
var AuthURL = "https://soundcloud.com/connect"
var TokenURL = "https://api.soundcloud.com/oauth2/token"

type UrlParams struct {
	url.Values
}

func NewUrlParams() UrlParams {
	return UrlParams{url.Values{}}
}

type SoundcloudApi struct {
	conf       *oauth2.Config
	httpClient *http.Client
	response   *http.Response
}

func validateNotEmptyString(m map[string]string) error {
    for field, message := range m {
        if field == "" {
            return fmt.Errorf(message)
        }
    }
	return nil
}

func NewSoundcloudApi(c string, cs string, callback *string) (*SoundcloudApi, error) {
	v := map[string]string{
		c: "Client Id cannot be Blank",
		cs: "Client Secret cannot be Blank",
	}
	if err := validateNotEmptyString(v); err != nil {
		return nil, err
	}
    if callback == nil {
        empty := ""
        callback = &empty
    }
	conf := &oauth2.Config{
		ClientID:     c,        //"CLIENT ID",
		ClientSecret: cs,       //"CLIENT SECRET",
		RedirectURL:  *callback, //"YOUR_REDIRECT_URL",
		Scopes:       []string{"non-expiring"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthURL,
			TokenURL: TokenURL,
		},
	}
	return &SoundcloudApi{conf: conf}, nil
}

// user credentials authorization
func (s *SoundcloudApi) PasswordCredentialsToken(u string, p string) error {
	v := map[string]string{
		u: "Username/Email cannot be Blank",
		p: "Password cannot be Blank",
	}
	if err := validateNotEmptyString(v); err != nil {
		return err
	}
	tok, err := s.conf.PasswordCredentialsToken(oauth2.NoContext, u, p)
	if err != nil {
		return err
	}
	defaultTokenType(tok)
	s.httpClient = s.conf.Client(oauth2.NoContext, tok)
	return err
}

// make a get request, p map are the url params
func (s *SoundcloudApi) Get(url string, p UrlParams) (*http.Response, error) {
	if len(p.Values) > 0 {
		url = buildUrlWithParams(url, p)
	} else {
		url = buildUrl(url)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return s.do(req)
}

// make a post request, data interface will be encoded into json
func (s *SoundcloudApi) Post(url string, data interface{}) (*http.Response, error) {
	url = buildUrl(url)
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(body)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}

	req.ContentLength = int64(reader.Len())
	req.Header.Set("content-type", "application/json")

	return s.do(req)
}

// make a put request, data interface will be encoded into json
func (s *SoundcloudApi) Put(url string, data interface{}) (*http.Response, error) {
	url = buildUrl(url)
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(body)
	req, err := http.NewRequest("PUT", url, reader)
	if err != nil {
		return nil, err
	}

	req.ContentLength = int64(reader.Len())
	req.Header.Set("content-type", "application/json")

	return s.do(req)
}

// make a delete request
func (s *SoundcloudApi) Delete(url string) (*http.Response, error) {
	url = buildUrl(url)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return s.do(req)
}

// resolves a soundcloud url and redirects automatically if found
func (s *SoundcloudApi) Resolve(searchUrl string) (*http.Response, error) {
	p := NewUrlParams()
	p.Set("url", searchUrl)
	url := buildUrlWithParams("/resolve", p)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return s.do(req)
}

func buildUrlWithParams(url string, p UrlParams) string {
	url = buildUrl(url)
	if len(p.Values) > 0 {
		url = url + "?" + p.Values.Encode()
	}
	return url
}

func buildUrl(url string) string {
    if len(url) >= 4 && url[:4] == "http" {
        return url
    }
	url = cleanUrlPrefix(url)
	url = prefixBaseUrlApi(url)
	return url
}

// adds a slash prefix if non-existent
func cleanUrlPrefix(url string) string {
	if url[:1] != "/" {
		url = "/" + url
	}
	return url
}

// prefix the soundcloud baseUrlApi
func prefixBaseUrlApi(url string) string {
	return BaseApiURL + url
}

// work-around for Soundcloud OAuth2 implementation,
// header must be OAuth instead of bearer
func defaultTokenType(t *oauth2.Token) {
	if t.TokenType == "" {
		t.TokenType = "OAuth"
	}
}

// send http request
func (s *SoundcloudApi) do(req *http.Request) (*http.Response, error) {
	var err error
	s.response, err = s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return s.response, nil
}

func (s *SoundcloudApi) GetLastResponse() (*http.Response) {
    return s.response
}

func (s *SoundcloudApi) SaveResource(r Saver) error {
    return r.Save(s)
}

func (s *SoundcloudApi) UpdateResource(r Updater) error {
    return r.Update(s)
}

func (s *SoundcloudApi) DeleteResource(r Deleter) error {
    return r.Delete(s)
}
