package gosound

import (
    "errors"
)

type User struct {
    *resource
    Avatar_url    string
    Permalink     string
    Username      string
    Uri           string
    Permalink_url string

    // full struct
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
    r := newResource()
    u := &User{}
    u.resource = r
    return u
}

func (u *User) Update(s *SoundcloudApi) (bool, error) {
    if u.IsNew() {
        return false, errors.New("User is new, cannot be updated!")
    }

    // send put request to soundcloud
    return true, nil
}