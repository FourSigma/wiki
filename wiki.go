package main

import (
	"fmt"
	"github.com/jeffail/gabs"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//Takes in url string and returns JSON of wiki summary
func getSummary(url string) map[string]interface{} {

	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	//Get the body of the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	//Loads JSON into the gabs Parser
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		panic(err)
	}

	children, err := jsonParsed.Path("query.pages").Children()
	if err != nil {
		panic(err)
	}
	// This child contains relevant information (title, summary)
	child := children[0].Data()

	//Cast into a map of interfaces
	data := child.(map[string]interface{})

	return data
}

func main() {

	var query string

	//Checks for right number of command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: wiki <query> ")
		fmt.Println("Example: wiki python")
		return
	} else if len(os.Args) == 2 {
		query = os.Args[1]
	} else {
		args := os.Args[1:]
		query = strings.Join(args, "_")
	}

	//API Calls
	url := "https://en.wikipedia.org/w/api.php?continue=&action=query&titles=placeholder&prop=extracts&exintro=&explaintext=&format=json&redirects"

	url = strings.Replace(url, "placeholder", query, -1)

	data := getSummary(url)

	//Output to terminal
	fmt.Println("\n", data["title"], "\n\n", data["extract"], "\n")
}
