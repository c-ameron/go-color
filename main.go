package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func hexHandler(w http.ResponseWriter, r *http.Request) {

	keys := r.URL.Query()
	hexColor := keys.Get("hex")

	// anything other than example query is invalid
	if len(keys) > 1 || !isValidHex(hexColor) {
		http.Error(w, "Bad query, please use format: /convert?hex=ff0000", http.StatusBadRequest)
		return
	}

	response := hexToRGB(hexColor)

	fmt.Fprintf(w, "%s", response)
}

func isValidHex(h string) bool {
	valid, _ := regexp.MatchString("^[0-9a-fA-F]{6}$", h)
	return valid
}

func hexToRGB(h string) string {

	decoded, _ := hex.DecodeString(h)

	// convert decoded hex bytes to string
	var hexStr = make([]string, len(decoded))
	for i, e := range decoded {
		hexStr[i] = strconv.Itoa(int(e))
	}

	rgb := "RGB(" + strings.Join(hexStr, ", ") + ")"

	return rgb
}

func statusHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "OK")
}

func main() {

	listenPort, ok := os.LookupEnv("LISTEN_PORT")
	if !ok {
		listenPort = "8080"
	}

	http.HandleFunc("/convert", hexHandler)
	http.HandleFunc("/status", statusHandler)
	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}
