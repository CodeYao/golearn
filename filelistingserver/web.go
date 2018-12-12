package main

import (
	"net/http"
	"os"

	"github.com/golearn/filelistingserver/filelisting"
	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			log.Error("Painc : [%v]", r)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}()

		err := handler(w, r)
		if err != nil {
			log.Warn("Error handling request :%s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}
func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
