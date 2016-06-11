package sitemap

import (
	"log"
	"net/http/httptest"
	"startones"
	"testing"
	//	"log/syslog"
	//	"github.com/garyburd/redigo/redis"
	"net/http"
)

func TestCheckGenerateSlug(t *testing.T) {

	golog, _ := startones.Start()
	site := "www.test.com"

	_, err := http.NewRequest("GET", "http://www.example.com/", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp := httptest.NewRecorder()

	CheckGenerateSlug(golog, resp, site)

}
