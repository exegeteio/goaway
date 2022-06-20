package router

import (
	"net/http"
	"net/url"
	"path"

	"github.com/exegeteio/goaway/pkg/config"
)

type Router struct {
	Config config.Config
}

// Get a prefix match for the request, check for fixed length
// routes, then redirect.
func (router *Router) Handler(w http.ResponseWriter, r *http.Request) {
	request_path := r.RequestURI[1:] // Remove the leading slash.
	// Parse the prefix character to determine redirect domain.
	domain, destination := router.Config.Route(request_path[0:1])
	request_path = request_path[1:] // Remove the prefix character.
	// Parse the fixed length of the remainder of the path, to determine redirect destination.
	fixed_domain, fixed_destination, err := router.Config.FixedLength(len(request_path[1:]), request_path[0:1])
	if err == nil {
		request_path = request_path[1:] // Remove the prefix character.
		// TODO:  Is this idiomatic go?
		if fixed_domain != "" {
			domain = fixed_domain
		}
		if fixed_destination != "" {
			destination = fixed_destination
		}
	}
	// Append un-parsed remainder of path to destination.
	destination = destination + request_path
	http.Redirect(w, r, router.buildURL(domain, destination), http.StatusFound)
}

func (r *Router) buildURL(domain string, destination string) string {
	u, err := url.Parse(domain)
	if err != nil {
		return path.Join(domain, destination)
	}
	u.Path = destination
	return u.String()
}
