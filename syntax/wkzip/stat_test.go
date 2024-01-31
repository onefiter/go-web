package wkzip_test

import (
	"testing"

	"github.com/go-web/syntax/wkzip"
)

func TestStat(t *testing.T) {
	wkzip.Stat("./txt/9100227.txt")
}
