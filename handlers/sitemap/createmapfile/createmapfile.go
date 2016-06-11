package createmapfile

import (
	"io/ioutil"
	"log/syslog"
)

func Createmap(golog syslog.Writer, filestr string, bfile []byte) {

	if err := ioutil.WriteFile(filestr, bfile, 0644); err != nil {

		golog.Err(err.Error())
	}
}
