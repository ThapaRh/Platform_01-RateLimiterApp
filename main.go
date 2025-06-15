package main
import
(
	"fmt"
	"log"
	 "net/http"
)

func main() {
 http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "api endpoint is created!")
 })

 fmt.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}