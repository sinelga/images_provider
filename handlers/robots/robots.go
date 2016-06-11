package robots

import (
	"bytes"
	"log/syslog"
	"net/http"
	"strings"
)

func Generate(golog syslog.Writer, w http.ResponseWriter, r *http.Request, site string) {

	var buffer bytes.Buffer

	golog.Info("robots: " + site)

	splithost := strings.Split(site, ":")

	if len(splithost) == 1 {


		buffer.WriteString("User-agent: *\nAllow: /\nSitemap: http://" + site + "/sitemap.xml\n")


	} else if len(splithost) == 2 {

		if splithost[1] == "80" {

			buffer.WriteString("User-agent: *\nAllow: /\nSitemap: http://" + site + "/sitemap.xml\n")

		} else {

			buffer.WriteString("User-agent: *\nAllow: /\n")

		}

	}


	w.Header().Add("Content-type", "text/plain")
	w.Write(buffer.Bytes())

}
