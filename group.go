package gosoundcloud

import (
    "errors"
    "strconv"
)

// For what i understand you cannot create nor delete a group via api. - to confirm
type Group struct {
    Id                  uint64
    Kind                string // to confirm if soundcloud respond with a kind for group types.
    Created_at          string "2009/06/18 15:46:46 +0000"
    Permalink           string "made-with-ableton-live"
    Name                string "Made with Ableton Live!"
    Short_description   string "tracks produced with Ableton Live! music software, no DJ mixes!"
    Description         string "send your tracks, no DJ mixes please!"
    Uri                 string "http://api.soundcloud.com/groups/3"
    Artwork_url         string "http://i1.sndcdn.com/artworks-000000481400-f0ynk3-large.jpg?142a848"
    Permalink_url       string "http://soundcloud.com/groups/made-with-ableton-live"
    Creator             User
    /* user contains based on api docs
    "creator": {
        "id": 1433,
        "permalink": "matas",
        "username": "matas",
        "uri": "http://api.soundcloud.com/users/1433",
        "permalink_url": "http://soundcloud.com/matas",
        "avatar_url": "http://i1.sndcdn.com/avatars-000001548772-zay6dy-large.jpg?142a848"
    }*/
}

func (g Group) GetId() uint64 {
    return g.Id
}

func (g Group) GetKind() string {
    return g.Kind
}

func (g Group) IsNew() bool {
    if g.Id > 0 {
        return false
    }
    return true
}

func (g *Group) Save(s *SoundcloudApi) error {
    if !g.IsNew() {
        return errors.New("Group is not new, cannot be saved!")
    }
    // save group
    return nil
}

func (g *Group) Delete(s *SoundcloudApi) error {
    if g.IsNew() {
        return errors.New("Group is new, cannot be deleted!")
    }

    url := "/groups/" + strconv.FormatUint(g.Id, 10)
    _, err := s.Delete(url)
    return err
}
