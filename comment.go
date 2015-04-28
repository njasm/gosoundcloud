package gosound

type Comment struct {
    *resource
    User_id    uint64
    Track_id   uint64
    Timestamp  uint64
    Created_at string
    Body       string "Html comment"
    Uri        string
    User       User
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
    r := newResource()
    c := &Comment{}
    c.resource = r
    return c
}

func (c *Comment) GetUserId() uint64 {
    return c.User_id
}

func (c *Comment) GetTrackId() uint64 {
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