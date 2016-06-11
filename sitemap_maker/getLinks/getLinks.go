package getLinks

import (
	"github.com/sinelga/images_provider/domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
//	"fmt"
	//	"strconv"
)

func GetAllLinks(golog syslog.Writer, c redis.Conn, site string) []domains.CharacterRedis {

	//	limitstr :=strconv.Itoa(limit)

	var charactersRedis []domains.CharacterRedis

	if bcharactersRedis, err := redis.MultiBulk(c.Do("HVALS", site)); err != nil {

		golog.Crit(err.Error())

	} else {

//		fmt.Println(len(bcharactersRedis),site)
		
		for _, bcharacter := range bcharactersRedis {

			var v, ok = bcharacter.([]byte)

			if ok {

				var character domains.CharacterRedis

				if err := json.Unmarshal(v, &character); err != nil {
					golog.Crit(err.Error())
				} else {

					if character.Sex == "female" {

						charactersRedis = append(charactersRedis, character)

					}

				}
			}

		}

	}

	//	sqlstr := "select Id,Moto from characters where topic='sex' and sex='female' order by Created_at desc limit "+limitstr
	//
	//	rows, err := db.Query(sqlstr)
	//	if err != nil {
	//		golog.Err(err.Error())
	//	}
	//	defer rows.Close()
	//
	//	var characters []domains.Character
	//
	//	for rows.Next() {
	//
	//		var ch domains.Character
	//
	//		err := rows.Scan(&ch.Id, &ch.Moto)
	//		if err != nil {
	//			golog.Err(err.Error())
	//		}
	//
	//		characters = append(characters, ch)
	//
	//	}

	return charactersRedis
}
