package getLinks

import (
	"github.com/garyburd/redigo/redis"
	"testing"
	"startones"
	"fmt"
)

func TestGetAllLinks(t *testing.T) {

	golog, _ := startones.Start()
	
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}
	defer c.Close()
	
	site :="www.test.com"
	
	
	res := GetAllLinks(golog,c,site)
	
	
	for _,character := range res {
		
		fmt.Println(character.Moto)
		
		
	}

}
