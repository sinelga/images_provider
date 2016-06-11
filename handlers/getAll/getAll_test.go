package getAll

import (
	"github.com/garyburd/redigo/redis"
	"startones"
	"testing"
	"fmt"
)

func TestGetAll(t *testing.T) {

	golog, _ := startones.Start()

	rds, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}
	defer rds.Close()

	site,exits := GetAll(golog, rds, "www.test.com")
	
	if exits {
		
		fmt.Println(site)
		
	} else {
		
		fmt.Println("NOT EXITS www.test.com")
		
	}

}
