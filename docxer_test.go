package docxer

import (
	"fmt"
	"testing"
)

func TestOpenFile(t *testing.T) {
	f, err := OpenFile("test.docx", &Options{UnzipSizeLimit: 0})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("xx>>>%+v\n", *f)

	content := f.readXML("word/document.xml")
	fmt.Println(string(content))

	r := NewRunParser(content)
	r.Execute()

}

// go test . -count=1 -v -test.run TestOpenFile
