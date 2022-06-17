package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"notionwithgolang/adapters"
	"notionwithgolang/models"
)

func main() {
	var response adapters.ProjectFromNotionDTO
	url := "https://api.notion.com/v1/databases/8fe53600c574469db93481966f2898ab/query"
	//fmt.Println("URL:>", url)
	jsonStr := []byte(`{
		"filter": {
			"and": [
				{
					"property": "Client",
					"relation": {
						"contains": "6d41ea58-543e-4512-80bc-11c864b73475"
					}
				},
				{
					"property": "Project status",
					"select": {
						"equals": "active"
					}
				}
			]
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
	listProjects := []models.Project{}
	reportingTo := ""
	for _, project := range response.Results {
		if len(project.Properties.ReportingTo.RichText) == 0 {
			reportingTo = ""
		} else {
			reportingTo = project.Properties.ReportingTo.RichText[0].PlainText
		}
		project := models.Project{
			ID:          project.ID,
			Name:        project.Properties.Name.Title[0].Text.Content,
			Member:      []models.Member{},
			ReportingTo: reportingTo,
			Status:      project.Properties.ProjectStatus.Select.Name,
		}
		listProjects = append(listProjects, project)
	}
	//fmt.Println(response.Results)
	fmt.Println(len(response.Results))
	//fmt.Println(string(body))
	fmt.Println(listProjects)
}
