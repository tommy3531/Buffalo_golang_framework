package actions

import (

	"github.com/gobuffalo/buffalo"
	"net/http"
	"os"
	"encoding/json"
	"github.com/tommarler/bottle2/model"
	"log"

)

func PropublicaHandler(c buffalo.Context) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.propublica.org/congress/v1/115/senate/members.json", nil)
	if err != nil {
		os.Exit(1)
	}
	req.Header.Add("X-API-KEY", "SpzjlPZlkMlPKKGCLQS1OqZtCN96lPl7sszOTKra")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	var responseObject model.Response
	if err := json.NewDecoder(resp.Body).Decode(&responseObject); err != nil {
		log.Println(err)
	}

	return c.Render(200, r.JSON(responseObject))
}
