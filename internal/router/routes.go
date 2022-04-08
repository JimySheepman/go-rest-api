package router

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/JimySheepman/go-rest-api/internal/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

type db *mongo.Database

var routes = []route{
	newRoute("POST", "/api/v1/fetch-data", handler.GetFetchDataHandler()),
	newRoute("POST", "/api/v1/in-memory", handler.PostInMemeoryDataHandler()),
	newRoute("GET", "/api/v1/in-memory", handler.GetInMemeoryDataHandler()),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	fmt.Println(route{method, regexp.MustCompile("^" + pattern + "$"), handler})
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type ctxKey struct{}

func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		fmt.Println(matches)
		fmt.Println(len(matches))
		fmt.Println(route.method)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			route.handler(w, r)
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}
