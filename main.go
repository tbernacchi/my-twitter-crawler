package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"./model"
	"github.com/dghubble/oauth1"
	elastic "gopkg.in/olivere/elastic.v7"
)

//Config to load secrets from file.
type Config struct {
	keys struct {
		Consumerkey    string `json:"CONSUMERKEY"`
		Consumersecret string `json:"CONSUMERSECRET"`
		Acesstoken     string `json:"ACESSTOKEN"`
		Accesssecret   string `json:"ACESSSECRET"`
	} `json:"keys"`
	Consumerkey    string `json:"CONSUMERKEY"`
	Consumersecret string `json:"CONSUMERSECRET"`
	Acesstoken     string `json:"ACESSTOKEN"`
	Accesssecret   string `json:"ACESSSECRET"`
}

//LoadConfiguration to load the environments varibles.
func LoadConfiguration(filename string) (Config, error) {

	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

//GetESClient to set an elastic client connection.
func GetESClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	return client, err

}

func main() {
	config, _ := LoadConfiguration("secrets.json")

	key := oauth1.NewConfig(config.Consumerkey, config.Consumersecret)
	token := oauth1.NewToken(config.Acesstoken, config.Accesssecret)

	//httpClient will automatically authorize http.Request's
	httpClient := key.Client(oauth1.NoContext, token)

	//tags
	tags := []string{"openbanking", "remediation", "devops", "sre", "microservices", "observability", "oauth", "metrics", "logmonitoring", "opentracing"}

	//Search tags
	for _, tag := range tags {
		path := "https://api.twitter.com/1.1/search/tweets.json?q=" + tag + "&result_type=mixed&count=100"
		response, _ := httpClient.Get(path)
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] Erro trying to search tweets: ", err.Error())
			return
		}
		defer response.Body.Close()

		searchResponse := model.SearchResponse{}
		err = json.Unmarshal([]byte(body), &searchResponse)
		if err != nil {
			fmt.Println("Error Unmarshall: ", err)
		}

		//Elastic
		ctx := context.Background()
		esclient, err := GetESClient()
		if err != nil {
			fmt.Println("Error initializing elastic: ", err)
		}

		fmt.Println("Inserting post with hashtag " + tag + " into Elasticsearch...")

		for _, src := range searchResponse.Statuses {
			dataJSON, err := json.Marshal(src)
			if err != nil {
				fmt.Println("Error marshall: ", err)
			}
			js := string(dataJSON)
			_, err = esclient.Index().Index("twitter-hashtags").BodyJson(js).Do(ctx)
			if err != nil {
				fmt.Println("Error sending data to ES: ", err)
			}
		}
	}
	fmt.Println("Done!")
}
