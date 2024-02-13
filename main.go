package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	/*fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/", handlers.Login)
	http.HandleFunc("/home", handlers.Home)
	// http.HandleFunc("/error", handlers.HandlerError)
	fmt.Println("(http://localhost:8080) - port started ")
	http.ListenAndServe(":8080", nil)*/
	fmt.Println("first hash", Encryptage("Fatima"))
	fmt.Println("second hash", Encryptage("Fatima"))
	if Encryptage("Fatima") == Encryptage("Fatima") {
		fmt.Println(true)
	}
}
func Encryptage(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	Hasher := hash.Sum(nil)
	return hex.EncodeToString(Hasher)
}
