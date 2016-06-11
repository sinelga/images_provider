package main

import (
	"github.com/sinelga/images_provider/domains"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net/url"
	"time"
	//	"fmt"
	"encoding/xml"
	"github.com/sinelga/images_provider/startones"
//	"strconv"
	"strings"

	"github.com/sinelga/images_provider/sitemap_maker/getLinks"
)

var siteFlag = flag.String("site", "", "must be test.com www.test.com")
var limitFlag = flag.Int("limit", 0, "if not will be 0")

func main() {
	flag.Parse() // Scan the arguments list

	site := *siteFlag
	limit := *limitFlag

	golog, _ := startones.Start()

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}
	defer c.Close()

	characters := getLinks.GetAllLinks(golog, c, site)

	var Url *url.URL

	docList := new(domains.Pages)
	docList.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

	for i := 0; i < limit; i++ {

		Url, err = url.Parse("http://" + site)
		if err != nil {
			golog.Crit(err.Error())			
		}
		
		permlink :=strings.Split(characters[i].Moto," ")
//		Url.Path += "/"+ strconv.Itoa(characters[i].Id) + "/" +permlink[0]+"-"+permlink[1]+".html"
		Url.Path += "/"+ characters[i].Id + "/" +permlink[0]+"-"+permlink[1]+".html"
		
		doc := new(domains.Page)
		doc.Loc = Url.String()
		doc.Lastmod =characters[i].Created_at.Format(time.RFC3339)
		doc.Changefreq = "weekly"

		docList.Pages = append(docList.Pages, doc)
		
	}

	resultXml, err := xml.MarshalIndent(docList, "", "  ")
	if err != nil {

		golog.Crit(err.Error())
	}

	fmt.Println(string(resultXml))

}
