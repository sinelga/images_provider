package getOne

import (
	"github.com/sinelga/images_provider/domains"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"log/syslog"
)

func GetById(golog syslog.Writer, c redis.Conn, site string, id string) (domains.Character,bool) {

	var character domains.Character
	var exist bool = false

	golog.Info("GetById "+ id+" "+site)

	if existint, err := redis.Int(c.Do("HEXISTS", site, id)); err != nil {

		golog.Crit("bcharacterRedis " + err.Error())
	} else {

		if existint == 1 {
			
			golog.Info("Exist "+site+" "+ id)
			exist = true

			if bcharacterRedis, err := redis.Bytes(c.Do("HGET", site, id)); err != nil {

				golog.Crit("bcharacterRedis " + err.Error())

			} else {

				if err := json.Unmarshal(bcharacterRedis, &character); err != nil {
					golog.Crit(err.Error())
				}

			}

		} 

	}

	return character,exist

}
