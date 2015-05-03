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

#### Naive Example

```go
package main

import (
    "fmt"
    "os"
    
    "github.com/njasm/gosoundcloud"
)

func main() {
    //  callback url is optional - nil in example
    s, _ := gosoundcloud.NewSoundcloudApi("client_id", "client_secret", nil)

    // request password credentials token - what soundcloud calls user credentials authentication
    if err = s.PasswordCredentialsToken("your_email@something.com", "your_password"); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    
    // get group id 3 data
    var group_id uint64 = 3
    group, _ := s.GetGroup(group_id)
    fmt.Println(group.Description)
    
    // get group members, that have "great" in they username, description, etc
    params := gosoundcloud.NewUrlParams()
    params.Set("q", "great");
    members, _ := s.GetGroupMembers(group, params)
    //members, _ := s.GetGroupMembers(group, nil) // or get all members

    for member := range members {
        fmt.Println(member.Username)
    }
}
```