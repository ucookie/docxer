package docxer

func NewFile() *File {
	f := newFile()

	return f
}

func newFile() *File {
	return &File{}
}
