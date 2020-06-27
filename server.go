// static.v2
// Copyright (c) 2020, HuguesGuilleus
// BSD 3-Clause License

// Create easily a http Handler for static file.
package static

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	// Dev disable the minifing and read all SleepDev the files.
	Dev bool = false
	// The duration of sleeping in development mode
	SleepDev time.Duration = 100 * time.Millisecond
)

// Server a static content with a min Content-Type header.
//
// The content is by default d. If f is non empty, the function read recurrent
// from f serve it. The reading error are silent.
//
// The served content are minify (expect if Dev is enable) with min. If min
// is nil, the content are not minify.
func File(b []byte, f string, mime string, min Minifier) http.HandlerFunc {
	go func() {
		b = min.min(b)
		for range ticker() {
			if content := readFileOnce(f, min); content != nil {
				b = content
			}
		}
	}()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mime)
		w.Write(b)
	}
}

// Send at begin a signal. If Dev is enalble, return a timer,
// else close the channel.
func ticker() <-chan struct{} {
	ch := make(chan struct{}, 0)

	go func() {
		ch <- struct{}{}
		if Dev {
			for range time.NewTicker(SleepDev).C {
				ch <- struct{}{}
			}
		} else {
			close(ch)
		}
	}()

	return ch
}

func readFileOnce(f string, m Minifier) []byte {
	if f == "" {
		return nil
	}

	data := make([]byte, 0)
	filepath.Walk(f, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		d, err := ioutil.ReadFile(p)
		if err != nil {
			return nil
		}
		data = append(data, m.min(d)...)
		return nil
	})
	return data
}

type Minifier func([]byte) []byte

// Minify even m if nil
func (m Minifier) min(in []byte) []byte {
	if Dev || m == nil || len(in) == 0 {
		return in
	}
	return m(in)
}

func (m Minifier) minS(in []byte) string {
	return string(m.min(in))
}
