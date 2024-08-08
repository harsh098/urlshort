package internal

import (
	"log"
	"net/http"
	"os"
)

func MapHandler(path_to_urls map[string]string, fallback http.Handler) http.HandlerFunc {
	mux := http.NewServeMux()
	for k, v := range path_to_urls {
		log.Printf("Register %v as %v\n", k, v)
		url_pattern := k
		redirect_url := v
		mux.HandleFunc(url_pattern, func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Calling For %v : %v\n", url_pattern, redirect_url)
			http.Redirect(w, r, redirect_url, http.StatusMovedPermanently)
		})
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok := path_to_urls[r.URL.Path]; ok {
			mux.ServeHTTP(w, r)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func YAMLHandler(fallback http.Handler) (http.HandlerFunc, error) {
	var err error
	var config *Config
	config, err = GetConfig()
	if err != nil {
		log.Fatalf("Invalid or Missing Config:\n %v\n", err.Error())
		os.Exit(-1)
	}
	paths := config.Paths
	var pathMap = make(map[string]string)
	for _, path := range paths {
		pathMap[path.Path] = path.Target
	}
	return MapHandler(pathMap, fallback), err
}
