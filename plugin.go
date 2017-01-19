package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"net/url"
	"path"
)

// Plugin defines the ansibrest plugin parameters.
type Plugin struct {
	EndPoint  string
	PlayBook  string
	Inventory string
	ExtraVars string
}

// Exec runs the plugin
func (p *Plugin) Exec() error {
	endpoint, _ := url.Parse(p.EndPoint)
	endpoint.Path = path.Join(endpoint.Path, fmt.Sprintf("/api/playbooks/%s", p.PlayBook))

	req, err := http.NewRequest("POST", endpoint.String(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	q := req.URL.Query()
	q.Add("inventory", p.Inventory)
	q.Add("extraVars", p.ExtraVars)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if res.StatusCode != 200 {
		log.Infof("res: %+v", res)
		return fmt.Errorf("api response is not 200: %d", res.StatusCode)
	}
	defer res.Body.Close()
	return nil
}
