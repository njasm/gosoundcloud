package gosoundcloud

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

type User struct {
	Id            uint64
	Avatar_url    string
	Permalink     string
	Username      string
	Uri           string
	Permalink_url string

	// full struct
	Kind                    string
	Last_modified           string
	First_name              string
	Last_name               string
	Description             string
	Country                 string
	City                    string
	Playlist_count          uint64
	Following_count         uint64
	Followers_count         uint64
	Upload_seconds_left     uint64
	Private_tracks_count    uint64
	Public_favorites_count  uint64
	Private_playlists_count uint64
	Track_count             uint64
	Full_name               string
	Myspace_name            string
	Discogs_name            string
	Website                 string
	Website_title           string
	Online                  bool
	Plan                    string
	Quota                   map[string]interface{}
	Primary_email_confirmed bool
}

type UserConnections struct {
	Id            uint64
	Uri           string
	Type          string
	Service       string
	Created_at    string
	Display_name  string
	Post_favorite bool
	Post_publish  bool
	/*
	   {
	   "created_at": "2010/12/05 16:46:34 +0000",
	   "display_name": "a facebook artist page",
	   "id": 313104,
	   "post_favorite": false,
	   "post_publish": false,
	   "service": "facebook_page",
	   "type": "facebook_page",
	   "uri": "https://api.soundcloud.com/connections/313104"
	   }
	*/
}

func NewUser() *User {
	return &User{Kind: "user"}
}

func (u User) GetId() uint64 {
	return u.Id
}

func (u User) GetKind() string {
	return u.Kind
}

func (u User) IsNew() bool {
	if u.Id > 0 {
		return false
	}
	return true
}

func (u *User) Update(s *SoundcloudApi) error {
	if u.IsNew() {
		return errors.New("User is new, cannot be updated!")
	}
	if u.Uri == "" {
		return errors.New("Empty Uri, cannot be updated!")
	}
	_, err := s.Put(u.Uri, u)
	return err
}

func getUsers(s *SoundcloudApi, p *UrlParams) ([]*User, error) {
	resp, err := s.Get("/users", p)
	var slice []*User
	if err = processAndUnmarshalResponses(resp, err, slice); err != nil {
		return nil, err
	}
	return slice, err
}

func (u *User) getTracks(s *SoundcloudApi, p *UrlParams) ([]*Track, error) {
	url := u.Uri + "/tracks"
	resp, err := s.Get(url, p)
	var slice []*Track
	if err = processAndUnmarshalResponses(resp, err, slice); err != nil {
		return nil, err
	}
	return slice, err
}

func (u *User) getPlaylists(s *SoundcloudApi, p *UrlParams) ([]*Playlist, error) {
	url := u.Uri + "/playlists"
	resp, err := s.Get(url, p)
	var slice []*Playlist
	if err = processAndUnmarshalResponses(resp, err, slice); err != nil {
		return nil, err
	}
	return slice, err
}

func (u *User) getFollowings(s *SoundcloudApi, p *UrlParams) ([]*User, error) {
	url := u.Uri + "/followings"
	resp, err := s.Get(url, p)
	var slice []*User
	if err = processAndUnmarshalResponses(resp, err, slice); err != nil {
		return nil, err
	}
	return slice, err
}

// should be redundant with GetUser - to confirm
//func (u *User) getFollowing(s *SoundcloudApi, id uint64) (*User, error) {
//}

func (u *User) addFollowing(s *SoundcloudApi, followed *User) (*User, error) {
	url := u.Uri + "/followings/" + strconv.FormatUint(followed.Id, 10)
	resp, err := s.Put(url, followed)
	if err = processAndUnmarshalResponses(resp, err, followed); err != nil {
		return nil, err
	}
	return followed, err
}

func (u *User) deleteFollowing(s *SoundcloudApi, other *User) error {
	url := u.Uri + "/followings/" + strconv.FormatUint(other.Id, 10)
	resp, err := s.Delete(url)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		return nil
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	return errors.New(string(bytes))
}

func (u *User) getFollowers(s *SoundcloudApi, p *UrlParams) ([]*User, error) {
	url := u.Uri + "/followers"
	resp, err := s.Get(url, p)
	var slice []*User
	if err = processAndUnmarshalResponses(resp, err, slice); err != nil {
		return nil, err
	}
	return slice, err
}

// should be redundant with GetUser - to confirm
//func (u *User) getFollower(s *SoundcloudApi, id uint64) (*User, error) {
//}

func (u *User) getComments(s *SoundcloudApi, p *UrlParams) ([]*Comment, error) {
	url := u.Uri + "/comments"
	resp, err := s.Get(url, p)
	var slice []*Comment
	if err = processAndUnmarshalResponses(resp, err, slice); err != nil {
		return nil, err
	}
	return slice, err
}

func (u *User) getFavorites(s *SoundcloudApi, p *UrlParams) ([]*Track, error) {
	url := u.Uri + "/favorites"
	resp, err := s.Get(url, p)
	var slice []*Track
	if err = processAndUnmarshalResponses(resp, err, slice); err != nil {
		return nil, err
	}
	return slice, err
}

// should be redundant with GetTrack - to confirm
//func (u *User) getFavorite(s *SoundcloudApi, id uint64) (*User, error) {
//}

func (u *User) addFavorite(s *SoundcloudApi, t *Track) (*Track, error) {
	url := u.Uri + "/favorites/" + strconv.FormatUint(t.Id, 10)
	resp, err := s.Put(url, t)
	if err = processAndUnmarshalResponses(resp, err, t); err != nil {
		return nil, err
	}
	return t, err
}

func (u *User) deleteFavorite(s *SoundcloudApi, t *Track) error {
	url := u.Uri + "/favorites/" + strconv.FormatUint(t.Id, 10)
	resp, err := s.Delete(url)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		return nil
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	return errors.New(string(bytes))
}

func (u *User) getGroups(s *SoundcloudApi, p *UrlParams) ([]*Group, error) {
	url := u.Uri + "/groups"
	resp, err := s.Get(url, p)
	var slice []*Group
	if err = processAndUnmarshalResponses(resp, err, slice); err != nil {
		return nil, err
	}
	return slice, err
}

func (u *User) getWebProfiles(s *SoundcloudApi, p *UrlParams) ([]*WebProfile, error) {
	url := u.Uri + "/web-profiles"
	resp, err := s.Get(url, p)
	var slice []*WebProfile
	if err = processAndUnmarshalResponses(resp, err, slice); err != nil {
		return nil, err
	}
	return slice, err
}

func (u *User) getWebProfile(s *SoundcloudApi, id uint64) (*WebProfile, error) {
	url := u.Uri + "/web-profiles/" + strconv.FormatUint(id, 10)
	resp, err := s.Get(url, nil)
	p := NewWebProfile()
	if err = processAndUnmarshalResponses(resp, err, p); err != nil {
		return nil, err
	}
	return p, err
}

// use Exported DeleteResource
//func (u *User) deleteWebProfile(s *SoundcloudApi, *WebProfile) error {
//
//}

func (u User) MarshalJSON() ([]byte, error) {
	j := map[string]map[string]interface{}{
		"user": {
			"city":        u.City,
			"country":     u.Country,
			"description": u.Description,
			"first_name":  u.First_name,
			"last_name":   u.Last_name,
			"permalink":   u.Permalink,
			"username":    u.Username,
		},
	}

	return json.Marshal(j)
}
