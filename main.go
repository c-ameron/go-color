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

    response := hexToResponse(hexColor)

    fmt.Fprintf(w, "%s", response)
}


func hexToResponse(h string) string {
	decoded, _ := hex.DecodeString(h)

	var str = make([]string, len(decoded))

	for i, e := range decoded {
		str[i] = strconv.Itoa(int(e))
	}

	response := "RGB(" + strings.Join(str, ", ") + ")"

	return response
}

func main() {
    listenAddr, ok := os.LookupEnv("LISTEN_ADDR")
    if !ok {
        listenAddr = "localhost:8080"
    }
	http.HandleFunc("/convert", handler)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
