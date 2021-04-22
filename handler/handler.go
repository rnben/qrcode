package handler

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

var WIFI string

func GenerateQRHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	e := `"` + t.Format("20060102") + `"`
	w.Header().Set("Etag", e)
	w.Header().Set("Cache-Control", "max-age=86400") // 1 day
	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, e) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	log.Println("缓存失效")

	dataString := r.FormValue("d")

	if dataString == "" {
		dataString = fmt.Sprintf(WIFI, t.Format("20060102"))
	}

	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)

	png.Encode(w, qrCode)
}
