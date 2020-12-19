<h1 align="">Twitter crawler using Golang 👋</h1>
<p>
</p>

> Simple crawler using search tags from Twitter API using Go. Insert the posts into a Elasticsearch.

![Ansible](/.github/assets/img/Golang-logo.png)

<div align=>
	<img align="right" width="200px" src=/.github/assets/img/elklogo.png>
</div> 

## Table of Contents

* **Golang**  
  * [Official Website](https://golang.org/)
  * [Official Docs](https://golang.org/doc/)
  * [Official Github](https://github.com/golang/go)
* **ELK**
  * [Offical Website](https://www.elastic.co)
  * [Offical Docs](https://www.elastic.co/guide/index.html)
  * [Official Github](https://github.com/elastic)

## Requirements 
* Golang;
* A developer account on Twitter: https://developer.twitter.com/en/apply-for-access
* Docker: https://www.docker.com/;
* ELK: you can use this one: https://github.com/tbernacchi/meu-elk-compose

## How to use 

With your account approved on Twitter, you must create an app with tokens to access the API;

Edit the ```secrets.json``` with your token keys and then:

```
$ go mod init github.com/tbernacchi/twitter-crawler
$ go run main.go 
```

Voilá! 

## Author

👤 **Tadeu Bernacchi**

* Website: http://www.tadeubernacchi.com.br/
* Twitter: [@tadeuuuuu](https://twitter.com/tadeuuuuu)
* Github: [@tbernacchi](https://github.com/tbernacchi)
* LinkedIn: [@tadeubernacchi](https://linkedin.com/in/tadeubernacchi)

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
