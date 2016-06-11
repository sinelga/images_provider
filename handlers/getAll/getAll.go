package getAll

import (
	"github.com/sinelga/images_provider/domains"
//	"fmt"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
	"encoding/json"
	"math/rand"
	"time"
	//		"strings"
)

func shuffleSlice(slice []interface{}) []interface{} {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func GetAll(golog syslog.Writer, c redis.Conn, site string) ([]domains.CharacterRedis,bool) {

	var charactersRedis []domains.CharacterRedis
	var exist bool = false

	if existint, err := redis.Int(c.Do("EXISTS", site)); err != nil {

		golog.Crit("bcharacterRedis " + err.Error() + " " + site)

	} else {

		if existint == 1 {

			exist = true
			if allfilds, err := redis.Strings(c.Do("HKEYS", site)); err != nil {

				golog.Crit(err.Error())

			} else {

				list := make([]interface{}, len(allfilds))

				for i, fild := range allfilds {

					list[i] = fild

				}

				shuffleSlice(list)

				var args []interface{}

				args = append(args, site)

				for i := 0; i < 15; i++ {

					args = append(args, list[i].(string))

				}

				if bcharactersRedis, err := redis.MultiBulk(c.Do("HMGET", args...)); err != nil {

					golog.Crit(err.Error())

				} else {

					for _, x := range bcharactersRedis {
						var v, ok = x.([]byte)
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

			}

		}
	}

	return charactersRedis,exist

}
