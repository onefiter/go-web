package wkzip_test

import (
	"fmt"
	"testing"

	"github.com/go-web/syntax/wkzip"
)

func TestZip(t *testing.T) {
	err := wkzip.DeCompressZip("txt.zip", "./tmp", "123456", nil, 0)
	if err != nil {
		fmt.Println(err)
	}
}
