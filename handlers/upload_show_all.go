package handlers

import (
	"database/sql"
	"github.com/sinelga/images_provider/domains"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zenazn/goji/web"
	"encoding/json"
	"net/http"
	"github.com/sinelga/images_provider/startones"
)

func ShowAll(c web.C, w http.ResponseWriter, r *http.Request) {

	golog, config = startones.Start()

	db, err := sql.Open("mysql", config.Database.ConStr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer db.Close()

//	sqlstr := "select ch.Id,Name,Age,Moto,ph.Phone,Description,Region_id,City,Adv_phone_id,Img_orient,Topic,Sex,ch.Created_at,ch.Updated_at,Img_file_name,Img_content_type,Img_file_size,Img_updated_at from characters as ch,adv_phones as ph,regions as re where re.id=ch.region_id and ph.id=ch.adv_phone_id and topic='sex' and sex='female' order by ch.Created_at desc limit 30"
	sqlstr := "select Id,Age,Created_at,Img_file_name,Img_content_type from characters where topic='sex' and sex='female' order by Created_at desc limit 30"

	rows, err := db.Query(sqlstr)
	if err != nil {
		golog.Err(err.Error())
	}
	defer rows.Close()

	var characters []domains.Character

	for rows.Next() {

		var ch domains.Character

		err := rows.Scan(&ch.Id, &ch.Age, &ch.Created_at, &ch.Img_file_name, &ch.Img_content_type)
		if err != nil {
			golog.Err(err.Error())
		}

		characters = append(characters, ch)

	}

	bytes, e := json.Marshal(characters)
	if e != nil {

		golog.Err(e.Error())

	}

	w.Write(bytes)

}
