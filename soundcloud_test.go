package gosound

import (
    "testing"
    "net/http"
    "golang.org/x/oauth2"
    "fmt"
    "os"
    "bufio"
    "strings"
    "io"
    "reflect"
    "time"
)

var client_id string
var client_secret string
var username string
var passwd string

func init() {
    f, err := os.Open(".env2")
    if err != nil {
        //fmt.Println(err)
        //os.Exit(1)
        // example data
        client_id = "123456789"
        client_secret = "987654321"
        username = "john@doe.com"
        passwd = "whynot"
    } else {
        defer f.Close()
        reader := bufio.NewReader(f)
        for lino := 0; true; lino++ {
            line, _, err := reader.ReadLine()
            if err == io.EOF {
                err = nil
                break
            }
            parts := strings.Split(string(line), "=")
            switch parts[0] {
                case "CLIENT_ID":
                client_id = parts[1]
                case "CLIENT_SECRET":
                client_secret = strings.TrimSpace(parts[1])
                case "USERNAME":
                username = strings.TrimSpace(parts[1])
                case "PASSWD":
                passwd = strings.TrimSpace(parts[1])
                default:
                fmt.Println("ERROR!!")
            }
        }
    }

    go runHttpTestServer()
    fmt.Println("Sleeping 2 seconds..")
    time.Sleep(2 * time.Second)
}

func runHttpTestServer() {
    // http test's should be asserted here
    // soundcloud delete body response: {"status":"200 - OK"}
    http.HandleFunc("/oauth2/token", func(w http.ResponseWriter, req *http.Request) {
        jresponse := "{\"expires_in\":86400, \"refresh_token\":\"18927670-R-2v8e8ycA419RaaVWY9Xz4APp\", \"access_token\":\"18926970-A-nMnSHDqg8Fsunm6Qx1cF1APp\"}"
        w.Header().Set("content-type", "application/json")
        fmt.Fprint(w, jresponse)
    })
    http.HandleFunc("/tracks", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("content-type", "application/json")
        fmt.Fprint(w, "{[]}")
    })
    http.HandleFunc("/me", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("content-type", "application/json")
        fmt.Fprint(w, "{[]}")
    })
    http.HandleFunc("/playlists", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("content-type", "application/json")
        fmt.Fprint(w, "{[]}")
    })
    http.HandleFunc("/playlists/101716717", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("content-type", "application/json")
        fmt.Fprint(w, "{[]}")
    })
    http.HandleFunc("/resolve", func(w http.ResponseWriter, req *http.Request) {
        string := req.URL.Query().Encode()
        w.Header().Set("content-type", "application/json")
        if string != "url=https%3A%2F%2Fsoundcloud.com%2Fhybrid-species" {
            fmt.Fprint(w, "{\"result\": \"fail\"}")
        } else {
            fmt.Fprint(w, "{\"result\": \"sucess\"}")
        }
    })

    http.ListenAndServe(":8282", nil)
}

func getMockedSoundcloudApi() *SoundcloudApi {
    conf := &oauth2.Config{
        ClientID:     client_id, //"CLIENT ID",
        ClientSecret: client_secret, //"CLIENT SECRET",
        RedirectURL:  "", //"YOUR_REDIRECT_URL",
        Scopes: []string{"non-expiring"},
        Endpoint: oauth2.Endpoint{
            AuthURL: "http://127.0.0.1:8282/auth",
            TokenURL: "http://127.0.0.1:8282/oauth2/token",
        },
    }
    s := &SoundcloudApi{conf: conf}
    bt := reflect.ValueOf(&BaseApiURL)
    v := bt.Elem()
    v.SetString("http://127.0.0.1:8282")

    return s
}

func TestNewSoundcloudApi(t *testing.T) {
    s := getMockedSoundcloudApi()
    _, err := s.PasswordCredentialsToken(username, passwd)
    if err != nil {
        t.Error(err)
    }
}

func TestGet(t *testing.T) {
    s := getMockedSoundcloudApi()
    _, err := s.PasswordCredentialsToken(username, passwd)
    if err != nil {
        t.Error(err)
    }

    r, err := s.Get("/me", UrlParams{})
    if err != nil {
        t.Error(err)
    }
    r.Body.Close()
}

func TestPost(t *testing.T) {
    s := getMockedSoundcloudApi()
    _, err := s.PasswordCredentialsToken(username, passwd)
    if err != nil {
        t.Error(err)
    }
    playlist := map[string]map[string]string{
        "playlist":{
            "title": "golang yay!",
            "sharing": "public",
        },
    }

    r, err := s.Post("/playlists", playlist)
    r.Body.Close()
    if err != nil {
        t.Error(err)
    }
}

func TestPut(t *testing.T) {
    s := getMockedSoundcloudApi()
    _, err := s.PasswordCredentialsToken(username, passwd)
    if err != nil {
        t.Error(err)
    }
    // {"playlist":{"tracks":[{"id":29720509},{"id":26057359}]}}
    tracks := map[string]map[string][]map[string]uint64{
        "playlist":{
            "tracks": {
                {"id": 29720509},
                {"id": 26057359},
            },
        },
    }
    // playlists/101716717
    r, err := s.Put("/playlists/101716717", tracks)
    r.Body.Close()
    if err != nil {
        t.Error(err)
    }
}

func TestDelete(t *testing.T) {
    s := getMockedSoundcloudApi()
    _, err := s.PasswordCredentialsToken(username, passwd)
    if err != nil {
        t.Error(err)
    }

    // playlists/101716717
    r, err := s.Delete("/playlists/101716717")
    r.Body.Close()
    if err != nil {
        t.Error(err)
    }
}

func TestResolve(t *testing.T) {
    s := getMockedSoundcloudApi()
    _, err := s.PasswordCredentialsToken(username, passwd)
    if err != nil {
        t.Error(err)
    }

    r, err := s.Resolve("https://soundcloud.com/hybrid-species")
    r.Body.Close()
    if err != nil {
        t.Error(err)
    }
}

func TestDefaultTokenType(t *testing.T) {
    tok := oauth2.Token{}
    tok.TokenType = ""
    defaultTokenType(&tok)
    if tok.TokenType != "OAuth" {
        t.Error("TokenType should be OAuth")
    }
    tok.TokenType = "test"
    defaultTokenType(&tok)
    if tok.TokenType != "test" {
        t.Errorf("%q", "Token changed.")
    }
}

func TestCleanUrlPrefix(t *testing.T) {
    v := "/with"
    expected := "/with"
    r := cleanUrlPrefix(v)
    if r != expected {
        t.Errorf("Expected: %q, got: %q", expected, r)
    }
    v = "without"
    expected = "/without"
    r = cleanUrlPrefix(v)
    if r != expected {
        t.Errorf("Expected: %q, got: %q", expected, r)
    }
}

func TestPrefixBaseUrlApi(t *testing.T) {
    value := "/me"
    expected := BaseApiURL + value
    value = prefixBaseUrlApi(value)
    if expected != value {
        t.Errorf("Expected value: %q, got: %q", expected, value)
    }
}