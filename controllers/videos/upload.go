package videos

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Upload(res http.ResponseWriter, req *http.Request) {
	file, header, err := req.FormFile("files[]")
	if err != nil {
		handleErr(err, res)
		return
	}
	defer file.Close()
	videoName := strings.Split(header.Filename, ".")
	os.MkdirAll(appPath+videoName[0], os.ModePerm)
	newFile, err := os.Create(appPath + videoName[0] + "/" + videoName[0] + ".mp4")
	if err != nil {
		handleErr(err, res)
		return
	}
	defer newFile.Close()
	fmt.Printf("File name %s\n", videoName[0])
	numBytesWritten, err := io.Copy(newFile, file)
	if err != nil {
		handleErr(err, res)
		return
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
	res.WriteHeader(200)
}
