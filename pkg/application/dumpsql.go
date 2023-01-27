package application

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func (a *applicationType) DumpSqlClear() {
	if !a.Configuration.Debug {
		return
	}
	data := []byte("")
	dumpname := filepath.Join(a.GetOutput(), "sqldump.sql")
	if err := ioutil.WriteFile(dumpname, data, 0644); err != nil {
		a.ErrorLog().AnErr("writeSql", err).Send()
	}
}

func (a *applicationType) DumpSql(s string) {
	if !a.Configuration.Debug {
		return
	}
	data := []byte(s)
	dumpname := filepath.Join(a.GetOutput(), "sqldump.sql")

	if err := ioutil.WriteFile(dumpname, data, 0644); err != nil {
		a.ErrorLog().AnErr("writeSql", err).Send()
	}
}

func (a *applicationType) DumpSqlAppend(s string) {
	if !a.Configuration.Debug {
		return
	}
	s = "\n" + s
	data := []byte(s)
	dumpname := filepath.Join(a.GetOutput(), "sqldump.sql")
	f, err := os.OpenFile(dumpname, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.Write(data); err != nil {
		panic(err)
	}

}
