package server

import (
	"crypto/sha1"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"sort"
	"strings"
)

var (
	token string
)

func InitWxToken(conf *viper.Viper) {
	token = conf.GetString("wxtoken")
}

func makeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func validateUrl(w http.ResponseWriter, r *http.Request) bool {
	timestamp := strings.Join(r.Form["timestamp"], "")
	nonce := strings.Join(r.Form["nonce"], "")
	signatureGen := makeSignature(timestamp, nonce)

	signatureIn := strings.Join(r.Form["signature"], "")
	if signatureGen != signatureIn {
		return false
	}
	echostr := strings.Join(r.Form["echostr"], "")
	fmt.Println(echostr)
	fmt.Fprintf(w, echostr)
	return true
}
