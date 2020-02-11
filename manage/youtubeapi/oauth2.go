package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getClient(scope, clientSecretFile, tokenFile string) *http.Client {
	b, err := ioutil.ReadFile(clientSecretFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, scope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		tok, err = getTokenFromPrompt(config, authURL)
		saveToken(tokenFile, tok)

	}
	return config.Client(context.TODO(), tok)
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

func getTokenFromPrompt(config *oauth2.Config, authURL string) (*oauth2.Token, error) {
	var code string
	fmt.Printf("Go to the following link in your browser. After completing "+
		"the authorization flow, enter the authorization code on the command "+
		"line: \n%v\n", authURL)

	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}
	fmt.Println(authURL)
	return exchangeToken(config, code)
}

func exchangeToken(config *oauth2.Config, code string) (*oauth2.Token, error) {
	tok, err := config.Exchange(context.TODO(), code)
	if err != nil {
		log.Fatalf("Unable to retrieve token %v", err)
	}
	return tok, nil
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Println("trying to save token")
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
