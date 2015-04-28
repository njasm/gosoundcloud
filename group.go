package gosound

type group struct {
    Id                  uint64
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
