[![Build Status](https://travis-ci.org/njasm/gosoundcloud.svg?branch=master)](https://travis-ci.org/njasm/gosoundcloud) 
[![Coverage Status](https://coveralls.io/repos/njasm/gosoundcloud/badge.svg?branch=master)](https://coveralls.io/r/njasm/gosoundcloud?branch=master)

## Soundcloud.com API for GO

Package is already usable, but still under heavy development, API might change!
Still missing complete map of soundcloud resources to structs, helper functions, etc.

#### Implemented features 

* User Credentials Flow Authentication (Password Credentials)
* Access to all GET, PUT, POST and DELETE Resources

#### Soon to come

* User Authorization/Authentication
* Media File Download/Upload

#### Naive Low-level Example

```go
package main

import (
    "github.com/njasm/gosoundcloud"
    "fmt"
    "io/ioutil"
)

func main() {
    //  callback url is optional - nil in example
    s, err := gosoundcloud.NewSoundcloudApi("client_id", "client_secret", nil)
    err = s.PasswordCredentialsToken("your_email@something.com", "your_password")
    if err != nil {
        fmt.Println(err)
    }
    getParams := gosoundcloud.NewUrlParams()
    getParams.Set("q", "HybridSpecies")
    r, err := s.Get("/tracks", getParams)
    if err != nil {
        fmt.Println(err)
    }
    defer r.Body.Close()
    data, err = ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(data))
}
```