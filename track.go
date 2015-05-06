package gosoundcloud

type Track struct {
    Id              uint64
    Kind            string
    Created_with    *ScApp
}

func NewTrack() *Track {
    return &Track{
        Kind: "track",
    }
}
