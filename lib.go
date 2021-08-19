package docxer

import (
	"archive/zip"
	"bytes"
	"io"
	"strings"
)

// 读取内存中的zip文件
func ReadZipReader(r *zip.Reader, o *Options) (map[string][]byte, error) {
	var (
		err     error
		docPart = map[string]string{
			"[content_types].xml":  "[Content_Types].xml",
			"xl/sharedstrings.xml": "xl/sharedStrings.xml",
		}
		fileList  = make(map[string][]byte, len(r.File))
		unzipSize int64
	)
	for _, v := range r.File {
		unzipSize += v.FileInfo().Size()
		if unzipSize > o.UnzipSizeLimit {
			return fileList, newUnzipSizeLimitError(o.UnzipSizeLimit)
		}
		fileName := strings.Replace(v.Name, "\\", "/", -1)
		if partName, ok := docPart[strings.ToLower(fileName)]; ok {
			fileName = partName
		}
		if fileList[fileName], err = readFile(v); err != nil {
			return nil, err
		}
	}
	return fileList, nil
}

// 以字符串形式读取存档文件中的文件内容
func readFile(file *zip.File) ([]byte, error) {
	rc, err := file.Open()
	if err != nil {
		return nil, err
	}
	dat := make([]byte, 0, file.FileInfo().Size())
	buff := bytes.NewBuffer(dat)
	_, _ = io.Copy(buff, rc)
	return buff.Bytes(), rc.Close()
}

// readXML提供了一个将XML内容读取为字符串的函数
func (f *File) readXML(name string) []byte {
	if content, _ := f.Pkg.Load(name); content != nil {
		return content.([]byte)
	}
	return []byte{}
}
