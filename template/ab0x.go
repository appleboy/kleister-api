// generaTed by fileb0x at "2017-05-07 02:20:11.366339518 +0200 CEST" from config file "ab0x.yaml"

package template

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileIndexHTML is "./index.html"
var FileIndexHTML = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x90\xc1\x4e\xc4\x20\x10\x40\xef\xfd\x8a\x91\x0f\x28\xd9\xfb\xb4\x89\x51\x0f\xc6\x83\x1b\xe3\xc5\x23\x0b\xb3\x32\xca\x96\x86\x99\x34\xd9\x6c\xf6\xdf\x4d\x05\xa3\xf6\xc4\x84\xc7\x0b\x0f\xf0\xe6\xfe\xf9\xee\xf5\x6d\xff\x00\x51\x4f\x69\xec\x3a\xac\x2b\x00\x46\x72\x61\x1d\x00\xf0\xe0\x84\x20\x16\x3a\x0e\xe6\x72\x81\xfe\x25\x67\x85\xeb\xd5\x80\x6d\x3c\xf1\xf4\x09\x85\xd2\x60\xd8\xe7\xc9\xb4\xa3\x4e\x84\x54\x2c\x9f\xdc\x3b\x89\x3d\xba\x65\x85\x3d\xfb\x6c\xc6\x6e\xeb\x89\x9e\x13\x49\x24\xd2\x8d\x5d\x81\x75\xf3\x9c\xd8\x3b\xe5\x3c\xf5\x5e\xa4\x5d\x8d\xb6\x46\xae\xe3\x21\x87\x73\xcb\x09\xbc\x00\x87\xc1\xfc\x91\x4c\x45\xeb\xb3\x76\xe3\x53\x22\x16\xa5\x02\xb7\xfb\x47\xb4\x71\xd7\x34\x1b\x78\xf9\x29\x13\x5f\x78\x56\x90\xe2\x7f\x4b\xbe\xb7\xfe\xa7\x7c\x88\x19\xb1\x91\x1a\x54\x33\xd0\xd6\x6f\xfc\x0a\x00\x00\xff\xff\x52\x28\xac\x01\x5f\x01\x00\x00")

func init() {
	if CTX.Err() != nil {
		log.Fatal(CTX.Err())
	}

	var err error

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileIndexHTML)
	r, err = gzip.NewReader(rb)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err = FS.OpenFile(CTX, "./index.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}
}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// FileNames is a list of files included in this filebox
var FileNames = []string{
	"./index.html",
}
