package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var response ProjectFromNotionDTO
	url := "https://api.notion.com/v1/databases/8fe53600c574469db93481966f2898ab/query"
	//fmt.Println("URL:>", url)
	jsonStr := []byte(`{
		"filter": {
			"property": "Client",
			"relation": {
				"contains": "933a9d42-a694-477b-a618-bd8322a21920"
			}
		}
	}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Notion-Version", "2022-02-22")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(res)
	err := json.Unmarshal(body, &response)
	if err == nil {
		errors.New("response is not in the same JSON format")
	}

	if len(response.Results) == 0 {
		errors.New("no projects found")
	}
	//listProjects := []Project{}
	fmt.Println(response)
	fmt.Println(string(body))
}
