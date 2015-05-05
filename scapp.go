package gosoundcloud

type ScApp struct {
    Id            uint64
    Uri           string
    Permalink_url string
    External_url  string
    Creator       string
}

func NewScApp() *ScApp {
    return &ScApp{}
}

func (a *ScApp) getScAppTracks(s *SoundcloudApi, p *UrlParams) ([]*Track, error) {
    url := a.Uri + "/tracks"
    resp, err := s.Get(url, p)
    var slice []*Track
    if err = processAndUnmarshalResponses(resp, err, slice); err != nil {
        return nil, err
    }
    return slice, err
}