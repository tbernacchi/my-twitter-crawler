# twitter-crawler
Simple crawler using search tags from Twitter API using Go.

## Requirements 
* Go: https://golang.org/
* A developer account on Twitter: https://developer.twitter.com/en/apply-for-access
* Docker, ELK: you can use this one: https://github.com/tbernacchi/meu-elk-compose
## How to use 

With your account approved on Twitter, you must create an app with tokens to access the API;

Edit the ```secrets.json``` with your token keys and then:

```
$ go mod init github.com/tbernacchi/twitter-crawler
$ go run main.go 
```

Voilá! 
