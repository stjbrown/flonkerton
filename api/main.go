package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const MAX_UPLOAD_SIZE = 51200 * 1024 // 1MB

// Progress is used to track the progress of a file upload.
// It implements the io.Writer interface so it can be passed
// to an io.TeeReader()
type Progress struct {
	TotalSize int64
	BytesRead int64
}

// Write is used to satisfy the io.Writer interface.
// Instead of writing somewhere, it simply aggregates
// the total bytes on each read
func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.Print()
	return
}

// Print displays the current progress of the file upload
func (pr *Progress) Print() {
	if pr.BytesRead == pr.TotalSize {
		fmt.Println("DONE!")
		return
	}

	fmt.Printf("File upload in progress: %d\n", pr.BytesRead)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	//Parse Form Data
	r.ParseForm()

	filename := r.Form.Get("fname")
	//fmt.Println(filename)

	data := r.Form.Get("data")
	fmt.Println(data)

	//Create file
	f, err := os.Create(fmt.Sprintf("/uploads/%s", filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	referer := r.Header.Get("Referer")
	//redir := fmt.Sprintf("%s?recieved=true", referer)
	http.Redirect(w, r, referer, 303)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("uploadHandler - triggered")
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Loop over header names
	/*
		for name, values := range r.Header {
			// Loop over all values for the name.
			for _, value := range values {
				fmt.Println(name, value)
			}
		}
	*/

	// 32 MB is the default used by FormFile
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get a reference to the fileHeaders
	files := r.MultipartForm.File["files"]

	//fmt.Println(files)

	for _, fileHeader := range files {
		filename := fileHeader.Filename
		fmt.Println(filename)
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			http.Error(w, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.MkdirAll("/uploads", os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Create file
		f, err := os.Create(fmt.Sprintf("/uploads/%s", filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer f.Close()

		pr := &Progress{
			TotalSize: fileHeader.Size,
		}

		_, err = io.Copy(f, io.TeeReader(file, pr))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	referer := r.Header.Get("Referer")
	fmt.Println(referer)
	http.Redirect(w, r, referer, 303)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/fileupload", uploadHandler)
	mux.HandleFunc("/api/formpost", formHandler)

	if err := http.ListenAndServeTLS(":4500", "./certs/localhost.crt", "./certs/localhost.key", mux); err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

}
