package handlers

import (
    "testing"
    "log"
	"log/syslog"
//	"github.com/garyburd/redigo/redis"
	"net/http"
	"net/http/httptest"
	"github.com/zenazn/goji/web"
)

func TestMhandleAll(t *testing.T) {
	
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
				
	}
	
//	startparameters :=[]string{"tcp",":6379","5000"}
	
	
	req, err := http.NewRequest("GET", "http://www.example.com/", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp := httptest.NewRecorder()
	
	var c  web.C
	
	m := make(map[string]string)
	
	m["id"]="1"
	
	c.URLParams = m
	
	MhandleAll(c,resp,req)

}

