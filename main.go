package main
import (
	"fmt"
	"net/http"
	"os"
	"github.com/uninoma/calculator/webServer"
)
func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func main() {
	webServer.Init()


}