package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var path string

type directory struct {
    Files []string
}

func viewFile (w http.ResponseWriter, r *http.Request) {

    requestFilePath, _ := strings.CutPrefix(r.URL.Path, "/view")

    requestFilePath = path + requestFilePath

    mimeType := mime.TypeByExtension(filepath.Ext(requestFilePath))
    w.Header().Set("Content-Type", mimeType)

    fileReader, err := os.Open(requestFilePath)
    if err != nil {
        log.Fatal(err)
    }
    io.Copy(w, fileReader)
}

func viewHandler (w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
    }

    requestFilePath, _ := strings.CutPrefix(r.URL.Path, "/view")

    // requestFilePath = path + requestFilePath
    requestFilePath = filepath.Join(path, requestFilePath)
    info, err := os.Stat(requestFilePath)
    if err != nil {
        log.Fatal(err)
    }

    if !info.IsDir() {
        viewFile(w, r)
    } else {
        entries, err := os.ReadDir(requestFilePath)
        if err != nil {
            log.Fatal(err)
        }
        var dir directory

        for _, file := range entries {
            // f := fileInfo{file.Name(), file.IsDir()}
            dir.Files = append(dir.Files, file.Name())
        }

        responseJson, err := json.Marshal(dir)
        if err != nil {
            log.Fatal(err)
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(responseJson)
    }
}

func downloadFilesHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
    }
    requestFilePath := r.PathValue("filename")
    fmt.Println(requestFilePath)
    requestFilePath = filepath.Join(path, requestFilePath)

    // force a download with the content- disposition field
    w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(requestFilePath))
    // fmt.Fprintln(w, requestFilePath)

    mimeType := mime.TypeByExtension(filepath.Ext(requestFilePath))
    w.Header().Set("Content-Type", mimeType)

    fileReader, err := os.Open(requestFilePath)
    if err != nil {
        log.Fatal(err)
    }
    io.Copy(w, fileReader)
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
