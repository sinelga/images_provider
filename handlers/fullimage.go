package handlers

import (
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/sinelga/images_provider/startones"
)

func ImageFullShow(c web.C, w http.ResponseWriter, r *http.Request) {
	
	golog, config = startones.Start()
	
//	id := c.URLParams["id"]
//	imgfile := c.URLParams["imgfile"]
	
//	golog.Info(id+" "+imgfile)
	
	fs := http.FileServer(http.Dir("upload/img"))
	http.Handle("/fullimage/", http.StripPrefix("/fullimage/", fs))
		
}