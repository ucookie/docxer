package docxer

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type File struct {
	options *Options
	Path    string
	Pkg     sync.Map
}

// 选项
type Options struct {
	UnzipSizeLimit int64
}

/*
打开文件
*/
func OpenFile(filename string, opt *Options) (*File, error) {
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	f, err := OpenReader(file, opt)
	if err != nil {
		return nil, err
	}
	f.Path = filename
	return f, nil
}

/*打开文件
 */
func OpenReader(r io.Reader, opt *Options) (*File, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	f := newFile()

	f.options = opt
	if f.options.UnzipSizeLimit == 0 {
		f.options.UnzipSizeLimit = UnzipSizeLimit
	}

	zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, err
	}
	file, err := ReadZipReader(zr, f.options)
	if err != nil {
		return nil, err
	}

	for k, v := range file {
		f.Pkg.Store(k, v)
	}

	return f, nil
}

func (f *File) debug() {
	content := f.readXML("word/document.xml")
	fmt.Println(string(content))
}
