package main

import (
    "errors"
    "flag"
    "github.com/mrjones/oauth"
//    "fmt"
)

type Dropbox struct {}

func (d *Dropbox) Authorize() (string, error) {
    
    flags := flag.NewFlagSet("authorize dropbox", flag.ContinueOnError)
    key := flags.String("key", "", "")
    secret := flags.String("secret", "", "")
    flags.Parse(subCommandFlags)

    if len(*key) == 0 || len(*secret) == 0 {
        return "", errors.New("Necessary flags -key or -secret undefined!")
    }

    consumer := oauth.NewConsumer(*key, *secret, oauth.ServiceProvider{
            RequestTokenUrl:   "https://api.dropbox.com/1/oauth/request_token",
			AuthorizeTokenUrl: "https://www.dropbox.com/1/oauth/authorize",
			AccessTokenUrl:    "https://api.dropbox.com/1/oauth/access_token",
	})

    _, url, err := consumer.GetRequestTokenAndUrl("oob")
	if err != nil {
		return "", err
	}
        
    return url, nil
}
