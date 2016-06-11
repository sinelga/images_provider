package notjsbots

import (
	"github.com/sinelga/images_provider/domains"
	"github.com/zenazn/goji/web"
	"log"
	"log/syslog"
	"net/http"
	"path"
	"html/template"
	"strings"
)

func CreateNotJsPage(golog syslog.Writer, c web.C, w http.ResponseWriter, r *http.Request, variant string, character domains.Character,site string) {

	golog.Info("CreateNotJsPage msnbot!!! " + character.Name)
	
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
			golog.Crit(err.Error())
		}
	}
	
	var characterSite domains.CharacterSite
	
	characterSite.Site = site
	characterSite.SiteCharacter = character
		
	lp := path.Join("templates", "layout.html")

	t, err := template.ParseFiles(lp)

	err = t.Execute(w, characterSite)
	check(err)
}

func CreateNotJsPageIndex(golog syslog.Writer, c web.C, w http.ResponseWriter, r *http.Request, variant string, characters []domains.CharacterRedis, site string) {

	golog.Info("CreateNotJsPage index.html msnbot!!! ")


	funcMap := template.FuncMap{
		"PermLink": createPermLink,
	}

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
			golog.Crit(err.Error())
		}
	}
	var charactersSite domains.CharacterRedisSite
	charactersSite.Site = site
	charactersSite.SiteCharaters = characters


	lp := path.Join("templates", "layout_index.html")

	t, err := template.New("layout_index.html").Funcs(funcMap).ParseFiles(lp)
	err = t.Execute(w, charactersSite)
	check(err)

}

func createPermLink(moto string) string {
	permlink := strings.Split(moto, " ")

	return permlink[0] + "-" + permlink[1] + ".html"

}
