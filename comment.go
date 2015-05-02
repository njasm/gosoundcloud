package gosound

import (
    "errors"
    "strconv"
    "encoding/json"
)

type Comment struct {
    Id          uint64
    Kind        string
    User_id     uint64
    Track_id    uint64
    Timestamp   uint64
    Created_at  string
    Body        string "Html comment"
    Uri         string
    Creator     User
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

func NewComment() *Comment {
    return &Comment{Kind:"comment"}
}

func (c Comment) GetId() uint64 {
    return c.Id
}

func (c Comment) GetKind() string {
    return c.Kind
}

func (c Comment) IsNew() bool {
    if c.Id > 0 {
        return false
    }
    return true
}

func (c Comment) GetUserId() uint64 {
    return c.User_id
}

func (c Comment) GetTrackId() uint64 {
    return c.Track_id
}

func (c *Comment) SetTrackId(id uint64) bool {
    if !c.IsNew() {
        return false
    }
    c.Track_id = id
    return true
}

func (c *Comment) GetTimeStamp() uint64 {
    return c.Timestamp
}

func (c *Comment) SetTimeStamp(t uint64) bool {
    if !c.IsNew() {
        return false
    }
    c.Timestamp = t
    return true
}

func (c *Comment) GetCreatedAt() string {
    return c.Created_at
}

func (c *Comment) GetBody() string {
    return c.Body
}

func (c *Comment) SetBody(b string) bool {
    if !c.IsNew() {
        return false
    }
    c.Body = b
    return true
}

func (c *Comment) GetUri() string {
    return c.Uri
}

func (c *Comment) Save(s *SoundcloudApi) error {
    if !c.IsNew() {
        return errors.New("Comment is not new, cannot be saved!")
    }
    if c.Track_id == 0 {
        return errors.New("Track ID missing, cannot be saved!")
    }

    url := "/tracks/" + strconv.FormatUint(c.Track_id, 10) + "/comments"
    _, err := s.Post(url, *c)
    return err
}

func (c *Comment) Delete(s *SoundcloudApi) error {
    if c.IsNew() {
        return errors.New("Comment is new, cannot be deleted!")
    }

    url := "/comments/" + strconv.FormatUint(c.Id, 10)
    _, err := s.Delete(url)
    return err
}

func (c Comment) MarshalJSON() ([]byte, error) {
    j := map[string]map[string]interface{}{
        "comment": {
            "body": c.Body,
            "timestamp": c.Timestamp,
        },
    }
    return json.Marshal(j)
}