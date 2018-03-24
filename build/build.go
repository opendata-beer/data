package main

import (
	"log"
	"fmt"
	"strings"
	"os"
	"io"
	"io/ioutil"
	"path/filepath"
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
    Country string	`json:"country"`
    Brewery string	`json:"brewery"`
    Style string	`json:"style"`
    Hops []string	`json:"hops"`
    Links []string	`json:"links"`
}

type Params struct {
	Path string
	Country string
	Brewery string
}

func cleanup() (string, string) {
	dirJson := "output/json/"
	dirHtml := "output/html/"

	// first cleanup
	os.RemoveAll(dirJson)
	os.RemoveAll(dirHtml)

	return dirJson, dirHtml
}

func process(params Params) {
    file, err := os.Open(params.Path)
    if err != nil {
        log.Fatalf("failed opening directory: %s", err)
    }
    defer file.Close()
 
    files,_ := file.Readdir(0) // 0 to read all files and folders
    for _,file := range files {
        path := params.Path + "/" + file.Name()
        if (file.IsDir()) {
        	newParams := Params{Path: path, Country: params.Country, Brewery: params.Brewery}
        	if (params.Country == "") {
				fmt.Println("country ", filepath.Base(path))
				newParams.Country = filepath.Base(path)
        	} else if (params.Brewery == "") {
				fmt.Println("brewery ", filepath.Base(path))
        		newParams.Brewery = filepath.Base(path)
        	}
       		process(newParams)
        } else {
        	if (strings.HasSuffix(path,".beer")) {
   		        beer := new(Beer)
   		        content,_ := ioutil.ReadFile(path)
   		        err := yaml.Unmarshal(content, &beer)
        		if (err != nil) {
                	log.Fatalf("problem when unmarshaling path: %s", err)
        		}
        		beer.Country = params.Country
        		beer.Brewery = params.Brewery

        		jsonBeer,err := json.MarshalIndent(beer, "", "  ");
        		WriteJson("countries/" + beer.Country + "/breweries/" + beer.Brewery + "/beers/" + beer.Id, string(jsonBeer))
        	}
        	if (strings.HasSuffix(path,".brewery")) {

        	}
    	} 
    }
}

func WriteJson(file string, content string) error {
    
    path := "output/" + file + ".json"
    os.MkdirAll(filepath.Dir(path), os.ModePerm);
    fo, err := os.Create(path)
    if err != nil {
        return err
    }
    defer fo.Close()
    _, err = io.Copy(fo, strings.NewReader(content))
    if err != nil {
        return err
    }

    return nil
}


func main() {
	process(Params{Path: "data"})
}