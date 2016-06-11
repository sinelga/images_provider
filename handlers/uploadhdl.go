package handlers

import (
	"github.com/zenazn/goji/web"
	//	"image"
	//	"image/jpeg"
	//	"image/png"
	"database/sql"
	//	"fmt"
	"net/http"
	"path/filepath"
	//	"domains"
//		"fmt"
	"github.com/disintegration/imaging"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"mime"
	"os"
	"github.com/sinelga/images_provider/startones"
	"strconv"
	"time"
	"runtime"
)

func MakeUpload(c web.C, w http.ResponseWriter, r *http.Request) {


	runtime.GOMAXPROCS(runtime.NumCPU())
	golog, config = startones.Start()

	if r.Method == "OPTIONS" {

		golog.Info("Upload start OPTIONS")

	} else if r.Method == "POST" {

		golog.Info("Upload start POST")

		topic := "sex"
		sex := "female"

		store := config.Store.StoreDir

		err := r.ParseMultipartForm(10000000)
		if err != nil {

			golog.Err(err.Error())
		}

		files := r.MultipartForm.File["file"]

		golog.Info("age POST " + r.FormValue("age"))
		golog.Info("age POST " + r.FormValue("sex"))
		golog.Info("age POST " + r.FormValue("topic"))

		topic = r.FormValue("topic")
		sex = r.FormValue("sex")
		agestr := r.FormValue("age")
		age, _ := strconv.Atoi(agestr)

		for i, _ := range files {

			golog.Info("Upload start " + files[i].Filename)

			file, err := files[i].Open()
			defer file.Close()

			img_orient := "portrait"

			ext := filepath.Ext(files[i].Filename)
			mime := mime.TypeByExtension(ext)

			db, err := sql.Open("mysql", config.Database.ConStr)
			defer db.Close()

			timenow := time.Now()

			res, err := db.Exec("insert into characters (age,region_id,adv_phone_id,img_orient,topic,sex,created_at,updated_at,img_content_type,img_file_size,img_updated_at) values (?,?,?,?,?,?,?,?,?,?,?)", age, 0, 0, img_orient, topic, sex, timenow, timenow, mime, 0, timenow)
			if err != nil {
				//			return http.StatusInternalServerError, err.Error()
				golog.Err(err.Error())
			}

			id, err := res.LastInsertId()
			if err != nil {
				//				return http.StatusInternalServerError, err.Error()
				golog.Err(err.Error())

			}

			imgfilestr := strconv.FormatInt(id, 10) + "/original/" + files[i].Filename

			if err := os.MkdirAll(store+filepath.Dir(imgfilestr), 0777); err != nil {

				golog.Err(err.Error())

			}

			dst, err := os.Create(store + imgfilestr)
			defer dst.Close()

			if _, err := io.Copy(dst, file); err != nil {
				golog.Err(err.Error())
			}

			fi, _ := dst.Stat()
			size := fi.Size()

			dst.Close()
			file.Close()

			img, err := imaging.Open(store + imgfilestr)
			if err != nil {

				golog.Err(err.Error())
			}
			
			if img.Bounds().Dx() > img.Bounds().Dy() {
				img_orient = "landscape"
				
			}
						

			if _, err := db.Exec("update characters set img_file_name =?,img_file_size=?,img_orient=? where id =?", files[i].Filename, size, img_orient, id); err != nil {

				golog.Err(err.Error())

			}

		}

	}

}
