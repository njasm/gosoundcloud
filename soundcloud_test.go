package gosound

import (
    "testing"
    "golang.org/x/oauth2"
)

const (
    CLIENT_ID = "8427643cbe50e5302f955814f98dccfe"
    CLIENT_SECRET = "c844cf0f21296d2643a717b34c145556"
    USERNAME = "blindedspecie@gmail.com"
    PASSWD = "fullon!2014"
)

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