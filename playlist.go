package gosoundcloud

type Playlist struct {
    Id              uint64
    Kind            string
    Created_at      string
    User_id         uint64
    Duration        uint64
    Sharing         string
    Tag_list        string
    Permalink       string
    Track_count     uint64
    Streamable      bool
    Downloadable    bool
    Embeddable_by   string
    Purchase_url    *string
    Label_id        *uint64
    Type            string
    Playlist_type   string
    Ean             string
    Description     string
    Genre           string
    Release         string
    Purchase_title  *string
    Label_name      string
    Title           string
    Release_year    *uint32
    Release_month   *string
    Release_day     *uint
    License         string
    Uri             string
    Permalink_url   string
    Artwork_url     string
    User            *User
    tracks          []*Track
/*
{
  "kind": "playlist",
  "id": 405726,
  "created_at": "2010/11/02 09:24:50 +0000",
  "user_id": 3207,
  "duration": 154516,
  "sharing": "public",
  "tag_list": "",
  "permalink": "field-recordings",
  "track_count": 5,
  "streamable": true,
  "downloadable": true,
  "embeddable_by": "me",
  "purchase_url": null,
  "label_id": null,
  "type": "other",
  "playlist_type": "other",
  "ean": "",
  "description": "a couple of field recordings to test http://soundiverse.com.\r\n\r\nrecorded with the fire recorder: http://soundcloud.com/apps/fire",
  "genre": "",
  "release": "",
  "purchase_title": null,
  "label_name": "",
  "title": "Field Recordings",
  "release_year": null,
  "release_month": null,
  "release_day": null,
  "license": "all-rights-reserved",
  "uri": "http://api.soundcloud.com/playlists/405726",
  "permalink_url": "http://soundcloud.com/jwagener/sets/field-recordings",
  "artwork_url": "http://i1.sndcdn.com/artworks-000025801802-1msl1i-large.jpg?5e64f12",
  "user": {
    "id": 3207,
    "kind": "user",
    "permalink": "jwagener",
    "username": "Johannes Wagener",
    "uri": "http://api.soundcloud.com/users/3207",
    "permalink_url": "http://soundcloud.com/jwagener",
    "avatar_url": "http://i1.sndcdn.com/avatars-000014428549-3at7qc-large.jpg?5e64f12"
  },
  "tracks": [
    {
      "kind": "track",
      "id": 6621631,
      "created_at": "2010/11/02 09:08:43 +0000",
      "user_id": 3207,
      "duration": 27099,
      "commentable": true,
      "state": "finished",
      "original_content_size": 2382624,
      "sharing": "public",
      "tag_list": "Fieldrecording geo:lat=52.527544 geo:lon=13.402905",
      "permalink": "coffee-machine",
      "streamable": true,
      "embeddable_by": "all",
      "downloadable": false,
      "purchase_url": null,
      "label_id": null,
      "purchase_title": null,
      "genre": "",
      "title": "coffee machine",
      "description": "",
      "label_name": "",
      "release": "",
      "track_type": "",
      "key_signature": "",
      "isrc": "",
      "video_url": null,
      "bpm": null,
      "release_year": null,
      "release_month": null,
      "release_day": null,
      "original_format": "wav",
      "license": "cc-by",
      "uri": "http://api.soundcloud.com/tracks/6621631",
      "user": {
        "id": 3207,
        "kind": "user",
        "permalink": "jwagener",
        "username": "Johannes Wagener",
        "uri": "http://api.soundcloud.com/users/3207",
        "permalink_url": "http://soundcloud.com/jwagener",
        "avatar_url": "http://i1.sndcdn.com/avatars-000014428549-3at7qc-large.jpg?5e64f12"
      },
      "created_with": {
        "id": 64,
        "kind": "app",
        "name": "FiRe - Field Recorder",
        "uri": "http://api.soundcloud.com/apps/64",
        "permalink_url": "http://soundcloud.com/apps/fire",
        "external_url": "http://itunes.apple.com/us/app/fire-2-field-recorder/id436241643?mt=8"
      },
      "permalink_url": "http://soundcloud.com/jwagener/coffee-machine",
      "artwork_url": "http://i1.sndcdn.com/artworks-000002863219-4zpxc0-large.jpg?5e64f12",
      "waveform_url": "http://w1.sndcdn.com/Yva1Qimi7TVd_m.png",
      "stream_url": "http://api.soundcloud.com/tracks/6621631/stream",
      "playback_count": 1249,
      "download_count": 114,
      "favoritings_count": 14,
      "comment_count": 11,
      "attachments_uri": "http://api.soundcloud.com/tracks/6621631/attachments"
    },
    {
      "kind": "track",
      "id": 6621549,
      "created_at": "2010/11/02 09:00:23 +0000",
      "user_id": 3207,
      "duration": 65618,
      "commentable": true,
      "state": "finished",
      "original_content_size": 5780256,
      "sharing": "public",
      "tag_list": "Fieldrecording geo:lat=52.528181 geo:lon=13.412658",
      "permalink": "tram-in-berlin",
      "streamable": true,
      "embeddable_by": "all",
      "downloadable": true,
      "purchase_url": null,
      "label_id": null,
      "purchase_title": null,
      "genre": "",
      "title": "tram in berlin",
      "description": "",
      "label_name": "",
      "release": "",
      "track_type": "recording",
      "key_signature": "",
      "isrc": "",
      "video_url": null,
      "bpm": null,
      "release_year": null,
      "release_month": null,
      "release_day": null,
      "original_format": "wav",
      "license": "cc-by",
      "uri": "http://api.soundcloud.com/tracks/6621549",
      "user": {
        "id": 3207,
        "kind": "user",
        "permalink": "jwagener",
        "username": "Johannes Wagener",
        "uri": "http://api.soundcloud.com/users/3207",
        "permalink_url": "http://soundcloud.com/jwagener",
        "avatar_url": "http://i1.sndcdn.com/avatars-000014428549-3at7qc-large.jpg?5e64f12"
      },
      "created_with": {
        "id": 64,
        "kind": "app",
        "name": "FiRe - Field Recorder",
        "uri": "http://api.soundcloud.com/apps/64",
        "permalink_url": "http://soundcloud.com/apps/fire",
        "external_url": "http://itunes.apple.com/us/app/fire-2-field-recorder/id436241643?mt=8"
      },
      "permalink_url": "http://soundcloud.com/jwagener/tram-in-berlin",
      "artwork_url": "http://i1.sndcdn.com/artworks-000002863163-6f2aqe-large.jpg?5e64f12",
      "waveform_url": "http://w1.sndcdn.com/u04ibjx6FYdM_m.png",
      "stream_url": "http://api.soundcloud.com/tracks/6621549/stream",
      "download_url": "http://api.soundcloud.com/tracks/6621549/download",
      "playback_count": 578,
      "download_count": 93,
      "favoritings_count": 4,
      "comment_count": 3,
      "attachments_uri": "http://api.soundcloud.com/tracks/6621549/attachments"
    },
    {
      "kind": "track",
      "id": 6668072,
      "created_at": "2010/11/03 19:47:11 +0000",
      "user_id": 3207,
      "duration": 21871,
      "commentable": true,
      "state": "finished",
      "original_content_size": 1921800,
      "sharing": "public",
      "tag_list": "geo:lat=52.527529 geo:lon=13.402961",
      "permalink": "alex-playing-drums",
      "streamable": true,
      "embeddable_by": "all",
      "downloadable": true,
      "purchase_url": null,
      "label_id": null,
      "purchase_title": null,
      "genre": null,
      "title": "alex playing drums",
      "description": "",
      "label_name": "",
      "release": "",
      "track_type": "recording",
      "key_signature": "",
      "isrc": "",
      "video_url": null,
      "bpm": null,
      "release_year": null,
      "release_month": null,
      "release_day": null,
      "original_format": "wav",
      "license": "cc-by",
      "uri": "http://api.soundcloud.com/tracks/6668072",
      "user": {
        "id": 3207,
        "kind": "user",
        "permalink": "jwagener",
        "username": "Johannes Wagener",
        "uri": "http://api.soundcloud.com/users/3207",
        "permalink_url": "http://soundcloud.com/jwagener",
        "avatar_url": "http://i1.sndcdn.com/avatars-000014428549-3at7qc-large.jpg?5e64f12"
      },
      "created_with": {
        "id": 64,
        "kind": "app",
        "name": "FiRe - Field Recorder",
        "uri": "http://api.soundcloud.com/apps/64",
        "permalink_url": "http://soundcloud.com/apps/fire",
        "external_url": "http://itunes.apple.com/us/app/fire-2-field-recorder/id436241643?mt=8"
      },
      "permalink_url": "http://soundcloud.com/jwagener/alex-playing-drums",
      "artwork_url": "http://i1.sndcdn.com/artworks-000002888918-takbu6-large.jpg?5e64f12",
      "waveform_url": "http://w1.sndcdn.com/MQnxWxIH94ai_m.png",
      "stream_url": "http://api.soundcloud.com/tracks/6668072/stream",
      "download_url": "http://api.soundcloud.com/tracks/6668072/download",
      "playback_count": 400,
      "download_count": 84,
      "favoritings_count": 2,
      "comment_count": 1,
      "attachments_uri": "http://api.soundcloud.com/tracks/6668072/attachments"
    },
    {
      "kind": "track",
      "id": 6698933,
      "created_at": "2010/11/04 19:09:32 +0000",
      "user_id": 3207,
      "duration": 12726,
      "commentable": true,
      "state": "finished",
      "original_content_size": 1116936,
      "sharing": "public",
      "tag_list": "geo:lat=52.528450 geo:lon=13.404099",
      "permalink": "typing",
      "streamable": true,
      "embeddable_by": "all",
      "downloadable": false,
      "purchase_url": null,
      "label_id": null,
      "purchase_title": null,
      "genre": null,
      "title": "typing",
      "description": "",
      "label_name": "",
      "release": "",
      "track_type": "recording",
      "key_signature": "",
      "isrc": "",
      "video_url": null,
      "bpm": null,
      "release_year": null,
      "release_month": null,
      "release_day": null,
      "original_format": "wav",
      "license": "cc-by",
      "uri": "http://api.soundcloud.com/tracks/6698933",
      "user": {
        "id": 3207,
        "kind": "user",
        "permalink": "jwagener",
        "username": "Johannes Wagener",
        "uri": "http://api.soundcloud.com/users/3207",
        "permalink_url": "http://soundcloud.com/jwagener",
        "avatar_url": "http://i1.sndcdn.com/avatars-000014428549-3at7qc-large.jpg?5e64f12"
      },
      "created_with": {
        "id": 64,
        "kind": "app",
        "name": "FiRe - Field Recorder",
        "uri": "http://api.soundcloud.com/apps/64",
        "permalink_url": "http://soundcloud.com/apps/fire",
        "external_url": "http://itunes.apple.com/us/app/fire-2-field-recorder/id436241643?mt=8"
      },
      "permalink_url": "http://soundcloud.com/jwagener/typing",
      "artwork_url": "http://i1.sndcdn.com/artworks-000002903990-le6t7d-large.jpg?5e64f12",
      "waveform_url": "http://w1.sndcdn.com/ZSil4IqhP6Hh_m.png",
      "stream_url": "http://api.soundcloud.com/tracks/6698933/stream",
      "playback_count": 1151,
      "download_count": 102,
      "favoritings_count": 1,
      "comment_count": 0,
      "attachments_uri": "http://api.soundcloud.com/tracks/6698933/attachments"
    },
    {
      "kind": "track",
      "id": 6770077,
      "created_at": "2010/11/07 02:45:11 +0000",
      "user_id": 3207,
      "duration": 27202,
      "commentable": true,
      "state": "finished",
      "original_content_size": 2392840,
      "sharing": "public",
      "tag_list": "geo:lat=52.531203 geo:lon=13.412165",
      "permalink": "bassy",
      "streamable": true,
      "embeddable_by": "all",
      "downloadable": true,
      "purchase_url": null,
      "label_id": null,
      "purchase_title": null,
      "genre": "",
      "title": "bassy",
      "description": "",
      "label_name": "",
      "release": "",
      "track_type": "recording",
      "key_signature": "",
      "isrc": "",
      "video_url": null,
      "bpm": null,
      "release_year": null,
      "release_month": null,
      "release_day": null,
      "original_format": "wav",
      "license": "cc-by",
      "uri": "http://api.soundcloud.com/tracks/6770077",
      "user": {
        "id": 3207,
        "kind": "user",
        "permalink": "jwagener",
        "username": "Johannes Wagener",
        "uri": "http://api.soundcloud.com/users/3207",
        "permalink_url": "http://soundcloud.com/jwagener",
        "avatar_url": "http://i1.sndcdn.com/avatars-000014428549-3at7qc-large.jpg?5e64f12"
      },
      "created_with": {
        "id": 64,
        "kind": "app",
        "name": "FiRe - Field Recorder",
        "uri": "http://api.soundcloud.com/apps/64",
        "permalink_url": "http://soundcloud.com/apps/fire",
        "external_url": "http://itunes.apple.com/us/app/fire-2-field-recorder/id436241643?mt=8"
      },
      "permalink_url": "http://soundcloud.com/jwagener/bassy",
      "artwork_url": "http://i1.sndcdn.com/artworks-000002938592-960ejc-large.jpg?5e64f12",
      "waveform_url": "http://w1.sndcdn.com/bxaiyNJt3vWK_m.png",
      "stream_url": "http://api.soundcloud.com/tracks/6770077/stream",
      "download_url": "http://api.soundcloud.com/tracks/6770077/download",
      "playback_count": 335,
      "download_count": 78,
      "favoritings_count": 1,
      "comment_count": 1,
      "attachments_uri": "http://api.soundcloud.com/tracks/6770077/attachments"
    }
  ]
}
*/
}

func NewPlaylist() *Playlist {
    return &Playlist{Kind: "playlist"}
}

func getPlaylists(s *SoundcloudApi, p *UrlParams) ([]*Playlist, error) {
    resp, err := s.Get("/playlists", p)
    var slice []*Playlist
    if err = processAndUnmarshalResponses(resp, err, slice); err != nil {
        return nil, err
    }
    return slice, err
}