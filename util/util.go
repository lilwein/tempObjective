package util

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

// Returns true if err is not nil
func HTTPErrors(err error, message string, w http.ResponseWriter, httpError int) bool {
	if err != nil {
		msg := fmt.Sprintf("%s%s", message, err.Error())
		fmt.Printf("\t\t%s\n\n", msg)
		http.Error(w, msg, httpError)
		return true
	}
	return false
}

// log server
func ServerDebug(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, Request, r)

	log.Printf("%s request: %s\n", ctx.Value(Request).(*http.Request).Method,
		ctx.Value(Request).(*http.Request).URL.Path)
}

// initialize viper
func InitViper() {
	viper.SetConfigName("./config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("demo")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to initialize viper: %w", err))
	}
	log.Println("viper config initialized")
}

// Get the i-th token of url
func URL_Path(url string, i int) string {
	sub := strings.Split(url, "/")
	return sub[i]
}

// Returns the s string without the last token after the last underscore, if this exists
func ExcludeLastUnderscore(s string) string {
	tokens := strings.Split(s, "_")
	if len(tokens) <= 1 {
		return s
	}
	return strings.Join(tokens[:len(tokens)-1], "_")
}
