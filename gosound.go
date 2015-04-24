package gosound

type Saver interface {
    Save()
}

type Updater interface {
    Update()
}

type Deleter interface {
    Delete()
}

type Resourcer interface {
    Id()    uint64
    Kind()  string
    IsNew() bool
}

type resource struct {
    id      uint64
    kind    string
}

func (r *resource) Id() uint64 {
    return r.id
}

func (r *resource) Kind() string {
    return r.kind
}

func (r *resource) IsNew() bool {
    if r.id > 0 {
        return false
    }
    return true
}

type Comment struct {
    resource
    user_id         uint64
    track_id        uint64
    timestamp       uint64
    created_at      string
    body            string "Html comment"
    uri             string
    User            *User
}

func (c *Comment) UserId() uint64 {
    return c.user_id
}

func (c *Comment) TrackId() uint64 {
    return c.track_id
}

func (c *Comment) SetTrackId(id uint64) bool {
    if !c.IsNew() {
        return false
    }
    c.track_id = id
    return true
}

func (c *Comment) Timestamp() uint64 {
    return c.timestamp
}

func (c *Comment) SetTimestamp(t uint64) bool {
    if !c.IsNew() {
        return false
    }
    c.timestamp = t
    return true
}

func (c *Comment) CreatedAt() string {
    return c.created_at
}

func (c *Comment) Body() string {
    return c.body
}

func (c *Comment) SetBody(b string) bool {
    if !c.IsNew() {
        return false
    }
    c.body = b
    return true
}

func (c *Comment) Uri() string {
    return c.uri
}

type User struct {
    *resource
    Avatar_url              string
    Permalink               string
    Username                string
    Last_modified           string
    Uri                     string
    Permalink_url           string

    // full struct
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

func newResource() *resource {
    return &resource{}
}

func NewComment() *Comment {
    return new(Comment)
}

func NewUser() *User {
    return new(User)
}