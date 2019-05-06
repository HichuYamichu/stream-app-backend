package videos

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/HichuYamichu/stream-app-server/storage"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func Upload(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("video")
	if err != nil {
		handleErr(err, res)
		return
	}
	defer file.Close()

	req.ParseForm()
	fmt.Println(req.Form["title"])
	_, err = storage.DB.Collection("videos").InsertOne(ctx, bson.M{
		"title": req.Form["title"][0],
		"desc":  req.Form["desc"][0],
	})
	if err != nil {
		handleErr(err, res)
		return
	}

	videoName := req.Form["title"][0]
	os.MkdirAll(appPath+videoName, os.ModePerm)
	newFile, err := os.Create(appPath + videoName + "/" + videoName + ".mp4")
	if err != nil {
		handleErr(err, res)
		return
	}
	newImg, err := os.Create(appPath + videoName + "/" + videoName + ".jpeg")
	if err != nil {
		handleErr(err, res)
		return
	}
	defer newImg.Close()

	fmt.Printf("File name %s\n", videoName)
	numBytesWritten, err := io.Copy(newFile, file)
	if err != nil {
		handleErr(err, res)
		return
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)

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
