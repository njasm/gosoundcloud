package gosound

import (
    "fmt"
    "golang.org/x/oauth2"
    "net/http"
    "net/url"
    "encoding/json"
    "bytes"
    //"os"
    //"strings"
)

var BaseApiURL = "https://api.soundcloud.com"
var AuthURL = "https://soundcloud.com/connect"
var TokenURL = "https://api.soundcloud.com/oauth2/token"

type SoundcloudApi struct {
    conf        *oauth2.Config
    httpClient  *http.Client
    response    *http.Response
}

func validateNotEmptyString(m []map[string]string) error {
    for _, mp := range m {
        for field, message := range mp {
            if field == "" {
                return fmt.Errorf(message)
            }
        }
    }
    return nil
}

func NewSoundcloudApi(c string, cs string, callback string) (*SoundcloudApi, error) {
    v := []map[string]string{
        {c: "Client Id cannot be Blank"},
        {cs: "Client Secret cannot be Blank"},
    }
    if err := validateNotEmptyString(v); err != nil {
        return nil, err
    }
    conf := &oauth2.Config{
        ClientID:     c, //"CLIENT ID",
        ClientSecret: cs, //"CLIENT SECRET",
        RedirectURL:  callback, //"YOUR_REDIRECT_URL",
        Scopes: []string{"non-expiring"},
        Endpoint: oauth2.Endpoint{
            AuthURL: AuthURL,
            TokenURL: TokenURL,
        },
    }
    return &SoundcloudApi{conf: conf}, nil
}

// user credentials authorization
func (s *SoundcloudApi) PasswordCredentialsToken(u string, p string) (bool, error) {
    v := []map[string]string{
        {u: "Username/Email cannot be Blank"},
        {p: "Password cannot be Blank"},
    }
    if err := validateNotEmptyString(v); err != nil {
        return false, err
    }
    tok, err := s.conf.PasswordCredentialsToken(oauth2.NoContext, u, p)
    if err != nil {
        return false, err
    }
    defaultTokenType(tok)
    s.httpClient = s.conf.Client(oauth2.NoContext, tok)
    return true, nil
}

func (s *SoundcloudApi) Get(url string, p map[string][]string) (*http.Response, error) {
    url = cleanUrlPrefix(url)
    if len(p) > 0 {
        url = buildUrlParams(url, p)
    }
    prefixBaseUrlApi(&url)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    return s.do(req)
}

// make a post request, data interface will be json encoded
func (s *SoundcloudApi) Post(url string, data interface{}) (*http.Response, error) {
    url = cleanUrlPrefix(url)
    prefixBaseUrlApi(&url)

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

func (s *SoundcloudApi) Put(url string, data interface{}) (*http.Response, error) {
    url = cleanUrlPrefix(url)
    prefixBaseUrlApi(&url)

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

func (s *SoundcloudApi) Delete(url string) (*http.Response, error) {
    url = cleanUrlPrefix(url)
    prefixBaseUrlApi(&url)

    req, err := http.NewRequest("DELETE", url, nil)
    if err != nil {
        return nil, err
    }

    return s.do(req)
}

// resolves a soundcloud url and redirects automatically if found
func (s *SoundcloudApi) Resolve(url string) (*http.Response, error) {
    p := map[string][]string{
        "url":{url},
    }
    url = buildUrlParams("/resolve", p)
    prefixBaseUrlApi(&url)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    return s.do(req)
}

// prefix the soundcloud baseUrlApi
func prefixBaseUrlApi(url *string) {
    *url = BaseApiURL + *url
}

// build queryParams for a GET Request
func buildUrlParams(uri string, p map[string][]string) string {
    if len(p) == 0 {
        return uri
    }
    values := url.Values{}
    oldKey := ""
    for k, v := range p {
        for _, v := range v {
            if oldKey != k {
                values.Set(k, v)
            } else {
                values.Add(k, v)
            }
            oldKey = k
        }
    }
    return uri + "?" + values.Encode()
}

// adds a slash prefix if non-existent
func cleanUrlPrefix(url string) string {
    if url[:1] != "/" {
        return "/" + url
    }
    return url
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
