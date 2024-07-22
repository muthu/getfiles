package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var path string

type fileInfo struct {
    Name string `json:"name"`
    IsDir bool `json:"isDir"`
}

type directory struct {
    Files []fileInfo
}

func viewFile (w http.ResponseWriter, r *http.Request) {

}

func viewHandler (w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
    }

    // requestFilePath = http.StripPrefix(r.URL.Path, "/view/")

    requestFilePath, _ := strings.CutPrefix(r.URL.Path, "/view")
    fmt.Println(requestFilePath)

    // if endpointSplit := strings.Split(r.URL.Path, "/"); len(endpointSplit) > 1 {
    //     requestFilePath = endpointSplit[1]
    // }
    
    // needs edge case test where number of path

    path += requestFilePath
    info, err := os.Stat(path)
    if err != nil {
        log.Fatal(err)
    }

    if !info.IsDir() {
        viewFile(w, r)
    }

    entries, err := os.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }
    var dir directory

    for _, file := range entries {
        f := fileInfo{file.Name(), file.IsDir()}
        dir.Files = append(dir.Files, f)
    }

    // fmt.Printf("dir.files: %v\n", dir.Files)

    responseJson, err := json.Marshal(dir)
    if err != nil {
        log.Fatal(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseJson)
    // fmt.Println(string(responseJson))
    // fmt.Printf("%T\n", entries)
    // for _, e := range entries {
    //     fmt.Fprintln(w, e.Name())
    // }
}

func downloadFilesHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
    }
    w.Header().Set("Content-Type", "application/json")

    // path := "/Users/muthu/Desktop/scratch/notes/"
    filename := r.PathValue("filename")
    path = filepath.Join(path, filename)

    // force a download with the content- disposition field
    w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(path))
    fmt.Fprintln(w, path)
}

func main() {

    var err error
    path, err = os.Getwd()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    pathFlag := flag.String("p", path, "absolute path to directory (or) relative path from current directory")
    flag.Parse()
    path = *pathFlag
    fmt.Println(path)

    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/download/{filename}", downloadFilesHandler)

    fmt.Println("starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
    // log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/Users/muthu/Desktop/scratch/"))))
}
