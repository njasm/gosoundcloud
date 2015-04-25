[![Coverage Status](https://coveralls.io/repos/njasm/gosound/badge.svg?branch=master)](https://coveralls.io/r/njasm/gosound?branch=master)

## Soundcloud.com API for GO

Package is already usable, but still under heavy development, API might change!
Still missing complete map of soundcloud resources to structs, helper functions, etc.

#### Implemented features 

* User Credentials Flow Authentication (Password Credentials)
* Access to all GET, PUT, POST and DELETE Resources

#### Soon to come

* User Authorization/Authentication
* Media File Download/Upload

#### Naive Example

```go
// empty string is callback url (optional)
s, err := NewSoundcloudApi("client_id", "client_secret", "")
_, err = s.PasswordCredentialsToken("your_email@something.com", "your_password")
if err != nil {
    t.Error(err)
}
getParams := url.Values{}
getParams.Set("q", "HybridSpecies")
r, err := s.Get("/tracks", getParams)
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}
data, err = ioutil.ReadAll(r.Body)
r.Body.Close()
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}
fmt.Println(string(data))
```