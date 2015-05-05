package gosoundcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
    "io/ioutil"
    "strconv"

    "golang.org/x/oauth2"
)

var BaseApiURL = "https://api.soundcloud.com"
var AuthURL = "https://soundcloud.com/connect"
var TokenURL = "https://api.soundcloud.com/oauth2/token"

type UrlParams struct {
	url.Values
}

func NewUrlParams() *UrlParams {
	return &UrlParams{url.Values{}}
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

// Get Makes a get request to the specified url resource, p is adicional url params
func (s *SoundcloudApi) Get(url string, p *UrlParams) (*http.Response, error) {
	url = buildUrlWithParams(url, p)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return s.do(req)
}

// Post Makes a post request to the speciefied url resource, data interface will be encoded into json
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

// Put Makes a put request to the specified url resource, data interface will be encoded into json
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

// Delete Makes a delete request to the specified url resource
func (s *SoundcloudApi) Delete(url string) (*http.Response, error) {
	url = buildUrl(url)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return s.do(req)
}

// Resolve Resolves a soundcloud url and redirects automatically if found
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

// GetMe Requests the User resource of the authenticated user
func (s *SoundcloudApi) GetMe() (*User, error) {
    resp, err := s.Get("/me", nil)
    u := NewUser()
    if err = processAndUnmarshalResponses(resp, err, u); err != nil {
        return nil, err
    }
    return u, nil
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

func buildUrlWithParams(url string, p *UrlParams) string {
    url = buildUrl(url)
    if p != nil && len(p.Values) > 0 {
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

// cleanUrlPrefix adds a slash prefix if non-existent. only relative paths should call this function
func cleanUrlPrefix(url string) string {
    if url[:1] != "/" {
        url = "/" + url
    }
    return url
}

// prefixBaseUrlApi prefixes the soundcloud's api base url
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

func processAndUnmarshalResponses(resp *http.Response, err error, holder interface{}) (error) {
    if err != nil {
        return err
    }

    //TODO: check if StatusCode is 40x/50x if so set the body as the error and return

    data, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        return err
    }
    if err = json.Unmarshal(data, holder); err != nil {
        return err
    }
    return err
}

/********************
* PLAYLISTS Methods *
*********************/

func (s *SoundcloudApi) GetPlaylist(id uint64) (*Playlist, error) {
    url := "/playlists/" + strconv.FormatUint(id, 10)
    resp, err := s.Get(url, nil)
    p := NewPlaylist()
    if err = processAndUnmarshalResponses(resp, err, p); err != nil {
        return nil, err
    }
    return p, err
}

func (s *SoundcloudApi) GetPlaylists(p *UrlParams) ([]*Playlist, error) {
    return getPlaylists(s, p)
}

/****************
* USERS Methods *
*****************/

func (s *SoundcloudApi) GetUser(id uint64) (*User, error) {
    url := "/users/" + strconv.FormatUint(id, 10)
    resp, err := s.Get(url, nil)
    u := NewUser()
    if err = processAndUnmarshalResponses(resp, err, u); err != nil {
        return nil, err
    }
    return u, err
}

func (s *SoundcloudApi) GetUsers(p *UrlParams) ([]*User, error) {
    return getUsers(s, p)
}

func (s *SoundcloudApi) GetUserTracks(u *User, p *UrlParams) ([]*Track, error) {
    return u.getTracks(s, p)
}

func (s *SoundcloudApi) GetUserFollowings(u *User, p *UrlParams) ([]*User, error) {
    return u.getFollowings(s, p)
}

func (s *SoundcloudApi) GetUserFollowers(u *User, p *UrlParams) ([]*User, error) {
    return u.getFollowers(s, p)
}

func (s *SoundcloudApi) GetUserComments(u *User, p *UrlParams) ([]*Comment, error) {
    return u.getComments(s, p)
}

func (s *SoundcloudApi) GetUserFavorites(u *User, p *UrlParams) ([]*Track, error) {
    return u.getFavorites(s, p)
}

func (s *SoundcloudApi) GetUserPlaylists(u *User, p *UrlParams) ([]*Playlist, error) {
    return u.getPlaylists(s, p)
}

func (s *SoundcloudApi) GetUserGroups(u *User, p *UrlParams) ([]*Group, error) {
    return u.getGroups(s, p)
}

func (s *SoundcloudApi) GetUserWebProfiles(u *User, p *UrlParams) ([]*WebProfile, error) {
    return u.getWebProfiles(s, p)
}

/*******************
* COMMENTS Methods *
*******************/

func (s *SoundcloudApi) GetComment(id uint64) (*Comment, error) {
    url := "/comments/" + strconv.FormatUint(id, 10)
    resp, err := s.Get(url, nil)
    c := NewComment()
    if err = processAndUnmarshalResponses(resp, err, c); err != nil {
        return nil, err
    }
    return c, err
}

func (s *SoundcloudApi) GetComments(p *UrlParams) ([]*Comment, error) {
    return getComments(s, p)
}

/*****************
* GROUPS Methods *
******************/

func (s *SoundcloudApi) GetGroup(id uint64) (*Group, error) {
    url := "/groups/" + strconv.FormatUint(id, 10)
    resp, err := s.Get(url, nil)
    g := NewGroup()
    if err = processAndUnmarshalResponses(resp, err, g); err != nil {
        return nil, err
    }
    return g, err
}

func (s *SoundcloudApi) GetGroups(p *UrlParams) ([]*Group, error){
    return getGroups(s, p)
}

func (s *SoundcloudApi) GetGroupModerators(g *Group, p *UrlParams) ([]*User, error) {
    return g.getModerators(s, p)
}

func (s *SoundcloudApi) GetGroupMembers(g *Group, p *UrlParams) ([]*User, error) {
    return g.getMembers(s, p)
}

func (s *SoundcloudApi) GetGroupContributors(g *Group, p *UrlParams) ([]*User, error) {
    return g.getContributors(s, p)
}

func (s *SoundcloudApi) GetGroupUsers(g *Group, p *UrlParams) ([]*User, error) {
    return g.getUsers(s, p)
}

func (s *SoundcloudApi) GetGroupTracks(g *Group, p *UrlParams) ([]*Track, error) {
    return g.getTracks(s, p)
}

func (s *SoundcloudApi) GetGroupPendingTracks(g *Group, p *UrlParams) ([]*Track, error) {
    return g.getPendingTracks(s, p)
}

// should be redundant with GetTrack unless the track resource have added data here - to confirm
//func (s *SoundcloudApi) GetGroupPendingTrack(g *Group, id uint64) (*Track, error) {
//    return g.GetPendingTrack(s, uint64())
//}

func (s *SoundcloudApi) GetGroupContributions(g *Group, p *UrlParams) ([]*Track, error) {
    return g.getContributions(s, p)
}

// should be redundant with GetTrack unless the track resource have added data here - to confir
//func (s *SoundcloudApi) GetGroupContributionsTrack(g *Group) ([]*Track, error) {
//    return g.GetContributionsTrack(s)
//}
