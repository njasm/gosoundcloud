package gosoundcloud

type WebProfile struct{
    Id        uint64
    Kind      string
}

func NewWebProfile() *WebProfile {
    return &WebProfile{
        Kind: "web-profile",
    }
}