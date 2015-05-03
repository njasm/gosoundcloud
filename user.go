package gosoundcloud

import (
    "errors"
    "encoding/json"
)

type User struct {
    Id                      uint64
    Avatar_url              string
    Permalink               string
    Username                string
    Uri                     string
    Permalink_url           string

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

func (u User) MarshalJSON() ([]byte, error) {
    j := map[string]map[string]interface{}{
        "user": {
            "city": u.City,
            "country": u.Country,
            "description": u.Description,
            "first_name": u.First_name,
            "last_name": u.Last_name,
            "permalink": u.Permalink,
            "username": u.Username,
        },
    }

    return json.Marshal(j)
}