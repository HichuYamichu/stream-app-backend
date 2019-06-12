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

const appPath = "G:/Videos/stream-app-vids/"

// UploadVideo : uploads video
func UploadVideo(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("video")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	req.ParseForm()
	videoName := db.InsertVideo(req.Form["title"][0], req.Form["desc"][0])

	if err != nil {
		log.Fatal(err)
		return
	}

	// videoName := req.Form["title"][0]
	os.MkdirAll(appPath+videoName, os.ModePerm)
	newFile, err := os.Create(appPath + videoName + "/" + videoName + ".mp4")
	if err != nil {
		log.Fatal(err)
		return
	}
	newImg, err := os.Create(appPath + videoName + "/" + videoName + ".jpeg")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer newImg.Close()

	// fmt.Printf("File name %s\n", videoName)
	_, err = io.Copy(newFile, file)
	if err != nil {
		log.Fatal(err)
		return
	}

	filename := appPath + videoName + "/" + videoName + ".mp4"
	width := 640
	height := 360
	cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if cmd.Run() != nil {
		panic("could not generate frame")
	}

	fw := bufio.NewWriter(newImg)
	fw.Write(buffer.Bytes())

	res.WriteHeader(200)
}
