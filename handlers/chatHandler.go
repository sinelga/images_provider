package handlers

import (
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
	"net/url"
	"github.com/sinelga/images_provider/startones"
	"io/ioutil"
//	"encoding/json"
//	"domains"
)

func GetChatAnswer(c web.C, w http.ResponseWriter, r *http.Request) {

	uuid := c.URLParams["uuid"]
	phone := c.URLParams["phone"]
	ask := c.URLParams["ask"]
	fmt.Println("id", uuid, phone, ask)
	golog, _ := startones.Start()

	var Url *url.URL
	Url, err := url.Parse("http://79.125.25.179:8000")
	if err != nil {

		golog.Err(err.Error())
	}

	Url.Path += "/bot_answer"

	parameters := url.Values{}
	parameters.Add("uuid", uuid)
	parameters.Add("phone", phone)
	parameters.Add("say", ask)

	Url.RawQuery = parameters.Encode()
	
	fmt.Printf("Encoded URL is %q\n", Url.String())
	
	if res, err := http.Get(Url.String());err != nil {
		
		golog.Err(err.Error())
		
	} else {
		
		defer res.Body.Close()
		
		if body, err := ioutil.ReadAll(res.Body);err != nil {
			golog.Err(err.Error())
			
		} else {
			
//			var chat domains.Chat
//			
//			json.Unmarshal(body, &chat)
	
			w.Write(body)
			
						
		}
		
		
	}
	
	

	//	urlstr := "http://79.125.25.179:8000/bot_answer/?uuid="+uuid+"

	//	http://79.125.25.179:8000/bot_answer/?uuid='+uuid+'&phone='+this.selectedCharacter.phone+'&say='+$scope.mAsk+'

}
