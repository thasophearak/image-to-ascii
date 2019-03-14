package handler

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/mattes/go-asciibot"
	ascii "github.com/sophearak/goasciiart"
	ui "github.com/sophearak/image-to-ascii/ui"
)

type asciiText struct {
	ASCII string `json:"ascii"`
}

func asciirize(url string) []byte {
	reqImg, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(reqImg.Body)
	if err != nil {
		panic(err)
	}

	img, _, _ := image.Decode(bytes.NewReader(body))
	asciiBytes := ascii.Convert2Ascii(ascii.ScaleImage(img, 120))

	reqImg.Body.Close()

	return asciiBytes
}

func H(w http.ResponseWriter, r *http.Request) {
	imgURL := r.URL.Query().Get("img")
	o := r.URL.Query().Get("output")
	ua := r.Header.Get("User-Agent")

	if strings.Contains(ua, "curl") || strings.Contains(ua, "HTTPie") {
		if len(imgURL) > 0 {
			asciiBytes := asciirize(imgURL)
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", strconv.Itoa(len(string(asciiBytes))))

			w.Write(asciiBytes)
			return
		}

		bot := asciibot.Random()
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", strconv.Itoa(len(bot)))

		w.Write([]byte(bot))
	} else {
		if len(imgURL) > 0 {
			asciiBytes := asciirize(imgURL)

			rgba, err := ascii.TextToImage(string(asciiBytes))
			if err != nil {
				panic(err)
			}

			buffer := new(bytes.Buffer)

			switch o {
			case "jpg":
				jpeg.Encode(buffer, rgba, nil)

				w.Header().Set("Content-Type", "image/jpeg")
				w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

				w.Write(buffer.Bytes())
			case "png":
				png.Encode(buffer, rgba)

				w.Header().Set("Content-Type", "image/png")
				w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

				w.Write(buffer.Bytes())
			default:
				r := asciiText{ASCII: string(asciiBytes)}
				jm, _ := json.Marshal(r)
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Content-Length", strconv.Itoa(len(string(jm))))

				w.Write(jm)
			}
		} else {
			ui.Index(w)
		}
	}
}
