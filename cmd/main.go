package main

import (
        "io/ioutil"
        "os"
        "net/http"
        "fmt"
        "log"
        "io"
        "strconv"
         "os/exec"

        "github.com/liut/jpegquality"
)


func main() {
    //isAlive function
  http.HandleFunc("/isAlive", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "isAlive!")
  })
    //check quality of images in the directory
  http.HandleFunc("/quality", func(w http.ResponseWriter, r *http.Request){
    url := r.FormValue("url")
    files, err := ioutil.ReadDir(url)
        if err != nil {
           fmt.Println(err)
        }
   // The location of the project
    pwd, err := os.Getwd()
        if err != nil {
            fmt.Println(err)
        }

     m := make(map[string]int)
    //copy all the image from the directory i get ro my images directory
    for _, f := range files {
       sourceFile, err := os.Open(url+f.Name())
        if err != nil {
            fmt.Println(err)
        }

        // Create new file
        newFile, err := os.Create(pwd+"/cmd/images/"+f.Name())
        if err != nil {
           fmt.Println(err)
        }

       bytesCopied, err := io.Copy(newFile, sourceFile)
        if err != nil {
            fmt.Println(err)
        }
        log.Printf("Copied %d bytes.", bytesCopied)

    }
   // return the quality for each image in the directory
     for _, f := range files {
       filename := string(f.Name())
       file, err := os.Open(url+f.Name())
        if err != nil {
            fmt.Println(err)
        }
        defer file.Close()
        j, err := jpegquality.New(file)// or NewWithBytes([]byte)
        if err != nil {
            fmt.Println(err)
        }
        a := j.Quality()
        if a<85 {
               m[filename] = a
             fmt.Fprintf(w, (filename+": "+strconv.Itoa(a)+", ") )
        }
      }

        //delete all the files in images directory
        out, err := exec.Command("go", "run", "/app/src/cmd/delete.go").Output()

        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(string(out))
  })

   // up server
    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }


}
