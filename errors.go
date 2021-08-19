package docxer

import "fmt"

// 解压大小错误
func newUnzipSizeLimitError(unzipSizeLimit int64) error {
	return fmt.Errorf("unzip size exceeds the %d bytes limit", unzipSizeLimit)
}
