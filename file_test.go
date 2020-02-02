package bk_converter

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_curFile(t *testing.T) {
	fp := curFile(0)
	base := filepath.Base(fp)
	assert.Equal(t, "file_test.go", base)
}
