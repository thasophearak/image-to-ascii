package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	ascii "github.com/sophearak/goasciiart"
)

type asciiText struct {
	ASCII string `json:"ascii"`
}

func H(w http.ResponseWriter, r *http.Request) {
	imgURL := r.URL.Query().Get("img")
	o := r.URL.Query().Get("output")

	if len(imgURL) > 0 {
		reqImg, err := http.Get(imgURL)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(reqImg.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		startDecode := time.Now()
		img, _, _ := image.Decode(bytes.NewReader(body))
		endDecode := time.Since(startDecode)

		startConvert := time.Now()
		asciiBytes := ascii.Convert2Ascii(ascii.ScaleImage(img, 120))
		rgba, err := ascii.TextToImage(string(asciiBytes))
		endConvert := time.Since(startConvert)

		if err != nil {
			log.Fatal(err)
		}
		reqImg.Body.Close()
		buffer := new(bytes.Buffer)

		switch o {
		case "ascii":
			r := asciiText{ASCII: string(asciiBytes)}
			jm, _ := json.Marshal(r)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", strconv.Itoa(len(string(jm))))

			w.Write(jm)
		case "jpg":
			startJpgEndcode := time.Now()
			jpeg.Encode(buffer, rgba, nil)
			endJpgEndcode := time.Since(startJpgEndcode)

			w.Header().Set("Content-Type", "image/jpeg")
			w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
			w.Header().Set("Time-Decoding", fmt.Sprintf("%v", endDecode))
			w.Header().Set("Time-Converting", fmt.Sprintf("%v", endConvert))
			w.Header().Set("Time-Encoding", fmt.Sprintf("%v", endJpgEndcode))

			w.Write(buffer.Bytes())
		case "png":
			startPngEndcode := time.Now()
			png.Encode(buffer, rgba)
			endPngEndcode := time.Since(startPngEndcode)

			w.Header().Set("Content-Type", "image/png")
			w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
			w.Header().Set("Time-Decoding", fmt.Sprintf("%v", endDecode))
			w.Header().Set("Time-Converting", fmt.Sprintf("%v", endConvert))
			w.Header().Set("Time-Encoding", fmt.Sprintf("%v", endPngEndcode))

			w.Write(buffer.Bytes())
		}
	}
}
