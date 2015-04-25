package gosound

import (
    "testing"
    "golang.org/x/oauth2"
    "io/ioutil"
    "fmt"
    "net/url"
    "os"
    "bufio"
    "strings"
    "io"
)

var client_id string
var client_secret string
var username string
var passwd string

func init() {
    f, err := os.Open(".env")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
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

func TestNewSoundcloudApi(t *testing.T) {
    s, err := NewSoundcloudApi(client_id, client_secret, "")
    _, err = s.PasswordCredentialsToken(username, passwd)
    if err != nil {
        t.Error(err)
    }
    getParams := url.Values{}
    getParams.Set("q", "travis")
    r, err := s.Get("/tracks", getParams)
    if err != nil {
        t.Error(err)
    }
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    fmt.Println(string(body))

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

func TestBuildUrlParams(t *testing.T) {
    v := "/me"
    p := make(map[string][]string)
    result := buildUrlParams(v, p)
    if result != v {
        t.Errorf("Expected: %q, got: %q", v, result)
    }

    p["q"] = []string{"testString"}
    result = buildUrlParams(v, p)
    if result != "/me?q=testString" {
        t.Errorf("Expected: %q, got: %q", "/me?q=testString", result)
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
    prefixBaseUrlApi(&value)
    if expected != value {
        t.Errorf("Expected value: %q, got: %q", expected, value)
    }
}