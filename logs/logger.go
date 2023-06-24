package logs

import (
	logg "log"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	InfoL *logg.Logger
	ErrL  *log.Logger
	ReqL  *log.Logger

	info *os.File
	errl *os.File
	req  *os.File
)

func createFile(name string) *os.File {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
func init() {
	info = createFile("logs/info.log")
	errl = createFile("logs/error.log")
	req = createFile("logs/requests.log")

	ErrL, ReqL = log.New(), log.New()
	InfoL = logg.New(info, "INFO: ", logg.Ldate|logg.Ltime|logg.Lshortfile)
	//InfoL.SetOutput(info)
	ErrL.SetOutput(errl)
	ReqL.SetOutput(req)
}

func GetLogger() *os.File {
	return req
}

func CloseLogger() {
	err := info.Close()
	if err != nil {
		log.Fatalf("can't close file %v, %v", info, err)
	}
	err = errl.Close()
	if err != nil {
		log.Fatalf("can't close file %v, %v", errl, err)
	}
	err = req.Close()
	if err != nil {
		log.Fatalf("can't close file %v, %v", req, err)
	}
}
