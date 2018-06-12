package sportsdatabase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var (
	httpClient = &http.Client{Timeout: 10 * time.Second}
)

//SQDLResponse x
type SQDLResponse struct {
	Headers []string `json:"headers"`
	Groups  []struct {
		Sdql    string          `json:"sdql"`
		Columns [][]interface{} `json:"columns"`
	} `json:"groups"`
}

//RecordsCount returns bool if response contains records
func (response *SQDLResponse) RecordsCount() int {
	recordsCount := 0
	if len(response.Groups) > 0 {
		recordsCount = len(response.Groups[0].Columns[0])
	}
	return recordsCount
}

//QueryConfig x
type QueryConfig struct {
	APIKey                                      string
	Sport                                       string
	Teams, TeamColumns, GameColumns, Conditions []string
}

//Query x
func Query(qc QueryConfig) (*SQDLResponse, error) {
	x := make([]string, len(qc.GameColumns)+(len(qc.Teams)*len(qc.TeamColumns)))
	i := 0
	for _, gc := range qc.GameColumns {
		x[i] = gc
		i++
	}
	for _, t := range qc.Teams {
		for _, tc := range qc.TeamColumns {
			c := fmt.Sprintf("%s:%s", t, tc)
			x[i] = c
			i++
		}
	}

	sb := strings.Builder{}
	sb.WriteString(strings.Join(x, `,`))
	sb.WriteString(` @ `)
	sb.WriteString(strings.Join(qc.Conditions, ` and `))
	sqdl := sb.String()

	u := &url.URL{
		Scheme: "http",
		Host:   "api.sportsdatabase.com",
		Path:   fmt.Sprintf("%s/query.json", qc.Sport),
	}
	q := u.Query()
	q.Add("api_key", qc.APIKey)
	q.Add("output", "json")
	q.Add("sdql", sqdl)

	// ?sdql=date%2Cpoints%40team%3DBears%20and%20season%3D2011
	u.RawQuery = q.Encode()
	// log.Print(u)

	response, err := httpClient.Get(u.String())
	if err != nil {
		return nil, errors.Wrap(err, "can't get http response from sports database.")
	}
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

	// ioutil.WriteFile("beforeParsing.json", buf.Bytes(), 0666)

	body := buf.String()
	body = strings.TrimPrefix(body, `json_callback(`)
	body = strings.Replace(body, `);`, ``, -1)
	body = strings.Replace(body, `['`, `["`, -1)
	body = strings.Replace(body, `']`, `"]`, -1)
	body = strings.Replace(body, `','`, `","`, -1)
	body = strings.Replace(body, `,'`, `,"`, -1)
	body = strings.Replace(body, `', '`, `","`, -1)
	body = strings.Replace(body, `',`, `",`, -1)
	b := []byte(body)

	var sqdlResponse SQDLResponse
	// ioutil.WriteFile("currentParsing.json", b, 0666)
	err = json.Unmarshal(b, &sqdlResponse)
	if err != nil {
		ioutil.WriteFile("badjson.json", b, 0666)
		return nil, errors.Wrap(err, "can't unmarshall SDQL json")
	}

	return &sqdlResponse, nil
}
