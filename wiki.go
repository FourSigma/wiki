package main

import (
	"fmt"
	"github.com/jeffail/gabs"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	//Checks for right number of command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: wiki <query> ")
		fmt.Println("Example: wiki python")
		return
	}

	query := os.Args[1]

	//API Calls
	url := "https://en.wikipedia.org/w/api.php?continue=&action=query&titles=placeholder&prop=extracts&exintro=&explaintext=&format=json&redirects"
	url = strings.Replace(url, "placeholder", query, -1)
	r, _ := http.Get(url)
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	//Loads JSON into the gabs Parser
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		panic(err)
	}
	children, _ := jsonParsed.Path("query.pages").Children()
	child := children[0].Data()            // This child contains relevant information (title, summary)
	data := child.(map[string]interface{}) //Cast into a map of interfaces

	//Output to terminal
	fmt.Println("\n", data["title"], "\n\n", data["extract"], "\n")
}
