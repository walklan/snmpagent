package main

import (
	"config"
	"github.com/go-martini/martini"
	"httpagent/route"
	"httpagent/util"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var logger *log.Logger

func main() {

	port := config.Port
	os.Setenv("PORT", port)
	os.Setenv("MARTINI_ENV", martini.Prod)
	mux := martini.Classic()
	filename := strings.TrimSuffix(filepath.Base(os.Args[0]), path.Ext(os.Args[0]))
	logger = util.GetLogger(filename + ".log")
	mux.Map(logger)

	//global inject
	// sysvar := &route.Syslvlvar{timeout, retry, Debug}
	// mux.Map(sysvar)

	// support get and post method
	mux.Get("/snmpagent", route.SnmpAgent)
	mux.Post("/snmpagent", route.SnmpAgent)
	mux.Run()
	cpunum := runtime.NumCPU()
	//ret := runtime.GOMAXPROCS(cpunum)

	util.Info("listen port:", port, "cpunum:", cpunum)
	errs := http.ListenAndServe(":"+port, nil)
	if nil != errs {
		log.Fatalf("listen port %s error:%s", port, errs)
	}
}
