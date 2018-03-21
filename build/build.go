package main

import (
	"log"
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/ghodss/yaml"	
)

type Brewery struct {
    Id string		`json:"id"`
    Name string 	`json:"name"`
    Address string	`json:"artist"`
    Links []string	`json:"links"`
}

type Beer struct {
    Id string		`json:"id"`
    Name string	    `json:"name"`
    Brewery string	`json:"brewery"`
    Style string	`json:"style"`
    Hops []string	`json:"hops"`
    Links []string	`json:"links"`
}

func cleanup() (string, string) {
	dirJson := "output/json/"
	dirHtml := "output/html/"

	// first cleanup
	os.RemoveAll(dirJson)
	os.RemoveAll(dirHtml)

	return dirJson, dirHtml
}

func process(dir string) {
    file, err := os.Open(dir)
    if err != nil {
        log.Fatalf("failed opening directory: %s", err)
    }
    defer file.Close()
 
    files,_ := file.Readdir(0) // 0 to read all files and folders
    for _,file := range files {
        filepath := dir + "/" + file.Name()
        if (file.IsDir()) {
        	process(filepath)
        } else {
        	if (strings.HasSuffix(filepath,".beer")) {
   		        beer := new(Beer)
   		        content,_ := ioutil.ReadFile(filepath)
   		        err := yaml.Unmarshal(content, &beer)
        		if (err != nil) {
                	log.Fatalf("problem when unmarshaling filepath: %s", err)
        		}
        		jsonBeer,err := json.MarshalIndent(beer, "", "  ");
        		fmt.Println(string(jsonBeer))	
        	}
        	if (strings.HasSuffix(filepath,".brewery")) {

        	}

        	
        	
    	} 
    }
}

func main() {
	process(".")
}