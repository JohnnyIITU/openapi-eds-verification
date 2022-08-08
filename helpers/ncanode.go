package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type NcaNodeDefault struct {
	Host    string
	Version string
}

type NcaNodeRequestBody struct {
	Version string            `json:"version"`
	Method  string            `json:"method"`
	Params  map[string]string `json:"params"`
}

var ncaNodeDefault *NcaNodeDefault

func NcaNodeRequest(method string, params map[string]string) (*http.Response, error) {
	initData()

	requestBody, _ := json.Marshal(&NcaNodeRequestBody{
		Version: ncaNodeDefault.Version,
		Method:  method,
		Params:  params,
	})

	requestBytes := bytes.NewBuffer(requestBody)

	return http.Post(ncaNodeDefault.Host, "application/json", requestBytes)
}

func initData() {
	if ncaNodeDefault == nil {
		version := GetEnvDefault("NCA_NODE_VER", "1.0")
		host := GetEnvDefault("NCA_NODE_HOST", "https://dev-ncanode-integration-microservices.apps.dev.bcc.kz")
		ncaNodeDefault = &NcaNodeDefault{
			Version: version,
			Host:    host,
		}
	}
}
