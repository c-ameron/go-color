package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {

	keys := r.URL.Query()
	hexColor := keys.Get("hex")
	if len(keys) > 1 || hexColor == "" {
		http.Error(w, "bad params", http.StatusBadRequest)
		return
	}
	valid, _ := regexp.MatchString("^[0-9a-fA-F]{6}$", hexColor)
	if !valid {
		http.Error(w, "bad params", http.StatusBadRequest)
		return
	}

	decoded, _ := hex.DecodeString(hexColor)

	var strArray = make([]string, len(decoded))

	response := "RGB("

	for i, e := range decoded {
		strArray[i] = strconv.Itoa(int(e))
	}

	response += strings.Join(strArray, ",")

	response += ")"

	fmt.Fprintf(w, "%s", response)
}

func main() {
	http.HandleFunc("/convert", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
