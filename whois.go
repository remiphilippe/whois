package whois

import (
	"errors"
	"strings"

	"github.com/domainr/go-whois/servers"
	. "github.com/domainr/go-whois/types"
)

// Whois queries a whois server for q and returns the result.
func Whois(q string) (string, error) {
	req, err := Resolve(q)
	if err != nil {
		return "", err
	}

	res, err := req.Fetch()
	if err != nil {
		return "", err
	}

	return string(res.Body), nil
}

// Resolve finds a whois server for q and prepares a Request.
func Resolve(q string) (*Request, error) {
	req := NewRequest(q)

	labels := strings.Split(q, ".")
	zone := labels[len(labels)-1]

	req.Host = zones[zone]
	if req.Host == "" {
		return req, errors.New("No whois server found for " + q)
	}

	srv, ok := servers.Servers[req.Host]
	if !ok {
		srv = servers.Default
	}
	srv.Resolve(req)

	return req, nil
}
