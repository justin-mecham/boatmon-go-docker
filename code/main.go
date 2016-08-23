package main
 
import(
    "fmt"
    "net/http"
    "log"
    "time"
    "os"
)

func logtime() {
    t := time.Now()
    formatedTime := t.Format(time.RFC1123)
    logFile, err := os.OpenFile("/logfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }
    defer logFile.Close()
    log.SetOutput(logFile)
//    log.Println("First log message!")
    log.Println(formatedTime)
}
 
func indexHandler( w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "logging time")
    logtime()
}


func main(){
    http.HandleFunc("/", indexHandler)
    http.ListenAndServe(":80",nil)
}
