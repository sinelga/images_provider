package handlers

import (

	"github.com/zenazn/goji/web"
//	"image"
	"image/jpeg"
//	"image/png"
	"net/http"
	"github.com/disintegration/imaging"
	"runtime"
	"github.com/sinelga/images_provider/startones"
	"strconv"
	
)

func ImageShow(c web.C, w http.ResponseWriter, r *http.Request) {
	
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	golog, config = startones.Start()
	width := 100;
	height := 100;
	
	id := c.URLParams["id"]
	imgfile := c.URLParams["imgfile"]
//	mime := c.URLParams["mime"]
	
	if c.URLParams["width"] !="" {
		
		widthstr :=c.URLParams["width"]
		width,_ = strconv.Atoi(widthstr)
			
	}
	
	if c.URLParams["height"] !="" {
		
		heightstr :=c.URLParams["height"]
		height,_ = strconv.Atoi(heightstr)
			
	}	

	w.Header().Set("Content-Type", "image/jpeg")

	filestr := config.Store.StoreDir + id + "/original/" + imgfile

	file, err := imaging.Open(filestr)
//	defer file.Close()
	if err != nil {
		golog.Err(err.Error()+" "+filestr)
		return
	}

	m := imaging.Thumbnail(file,width ,height , imaging.CatmullRom)
	
	jpeg.Encode(w, m, nil)

}
