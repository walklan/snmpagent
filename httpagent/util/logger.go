package util

import (
	"config"
	"errors"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// get path
func Dir(filename string) string {
	if runtime.GOOS == "windows" {
		index := strings.LastIndex(filename, "\\")
		return string([]byte(filename)[0:index])
	} else {
		return path.Dir(filename)
	}
}

func OpenFile(filename string, flag string) (*os.File, error) {

	var f *os.File
	var err error

	//create dir
	filedir := Dir(filename)
	if !CheckFileIsExist(filedir) {
		log.Printf("mkdir:%s", filedir)
		//os.MkdirAll(filedir, 0755)

		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("md", filedir)
		} else {
			cmd = exec.Command("/bin/mkdir", "-p", filedir)
		}

		cmd.Run()
		//cmd.Wait()
	}

	//rewrite
	if flag == ">" {
		f, err = os.OpenFile(filename, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0755)
		f.Chown(os.Geteuid(), os.Getegid())
	} else if flag == ">>" {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
		f.Chown(os.Getuid(), os.Getgid())
	} else {
		err = errors.New("unknown flag:" + flag)
	}

	//return
	if err != nil {
		return nil, err
	} else {
		return f, nil
	}
}

func GetLogger(logname string) *log.Logger {

	log_file := config.Logdir + "/" + logname

	//logger
	file, err := OpenFile(log_file, ">>")
	if err != nil {
		log.Panicf("%s\n", err.Error())
	}

	// log.Println("log:" + log_file)
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func Debug(v ...interface{}) {
	log.Println("[DEBUG]", v)
}

func Info(v ...interface{}) {
	log.Println("[INFO]", v)
}

func Warn(v ...interface{}) {
	log.Println("[WARN]", v)
}

func Error(v ...interface{}) {
	log.Println("[ERROR]", v)
}
