package startones

import (
//	"code.google.com/p/gcfg"
	"gopkg.in/gcfg.v1"	
	"log"
	"log/syslog"
	"github.com/sinelga/images_provider/domains"

)



var config domains.Config

//func Start(golog syslog.Writer) ([]string,map[string]struct{}) {
func Start() (syslog.Writer, domains.Config) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")	

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	
	err = gcfg.ReadFileInto(&config, "/home/juno/git/imagehoster_redis/config.ini")
	if err != nil {
		
		golog.Crit("cannot read configuration file config.ini" + err.Error())

	}
	
//	golog.Info(config.Database.ConStr)

	return *golog, config

}
