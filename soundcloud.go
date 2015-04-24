package gosound

import (
    "fmt"
    "golang.org/x/oauth2"
    "net/http"
    "net/url"
)

const (
    BaseApiURL = "https://api.soundcloud.com"
)

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
    err := validateNotEmptyString(v)
    if err != nil {
        return nil, err
    }
    conf := &oauth2.Config{
        ClientID:     c, //"CLIENT ID",
        ClientSecret: cs, //"CLIENT SECRET",
        RedirectURL:  callback, //"YOUR_REDIRECT_URL",
        Scopes: []string{"non-expiring"},
        Endpoint: oauth2.Endpoint{
            AuthURL:"https://soundcloud.com/connect",
            TokenURL: "https://api.soundcloud.com/oauth2/token",
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
    err := validateNotEmptyString(v)
    if err != nil {
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
        url = buildGetParams(url, p)
    }
    prefixBaseUrlApi(&url)
    var err error
    s.response, err = s.httpClient.Get(url)
    if err != nil {
        return nil, err
    }
    return s.response, nil
}

func (s *SoundcloudApi) Post(url string, p map[string][]string) (*http.Response, error) {
    url = cleanUrlPrefix(url)
    if len(p) != 0 {
        url = buildGetParams(url, p)
    }
    prefixBaseUrlApi(&url)
    var err error
    s.response, err = s.httpClient.Get(url)
    if err != nil {
        return nil, err
    }
    return s.response, nil
}

// prefix the soundcloud baseUrlApi
func prefixBaseUrlApi(url *string) {
    *url = BaseApiURL + *url
}

// build queryParams for a GET Request
func buildGetParams(uri string, p map[string][]string) string {
    if len(p) == 0 {
        return uri
    }
    values := new(url.Values)
    for k, v := range p {
        for _, v := range v {
            values.Add(k, v)
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
