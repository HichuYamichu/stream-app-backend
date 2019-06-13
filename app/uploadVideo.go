package app

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/HichuYamichu/stream-app-server/db"
)

const storePath = "./store/"

// UploadVideo : uploads video
func UploadVideo(res http.ResponseWriter, req *http.Request) {
	// contentType := req.Header.Get("Content-type")
	// if contentType != "multipart/form-data" {
	// 	res.WriteHeader(400)
	// 	res.Write([]byte("Bad request"))
	// 	return
	// }

	file, _, err := req.FormFile("video")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	req.ParseForm()
	fileName := db.InsertVideo(req.Form["title"][0], req.Form["desc"][0])

	if err != nil {
		log.Fatal(err)
		return
	}

	filePath := storePath + "videos/" + fileName + ".mp4"
	newFile, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	newImg, err := os.Create(storePath + "miniatures/" + fileName + ".jpeg")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer newImg.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		log.Fatal(err)
		return
	}

	width := 640
	height := 360
	cmd := exec.Command("ffmpeg", "-i", filePath, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if cmd.Run() != nil {
		panic("could not generate frame")
	}

	fw := bufio.NewWriter(newImg)
	fw.Write(buffer.Bytes())

	res.WriteHeader(200)
}
