package util

import (
	"fmt"
	"testing"

	"github.com/bpfs/defs/paths"
)

func TestGetDefaultDownloadPath(t *testing.T) {
	path := paths.DefaultDownloadPath()
	fmt.Printf("默认路径：\t%s", path)
}
