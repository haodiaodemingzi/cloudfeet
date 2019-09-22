package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderTemplate(t *testing.T) {
	a := assert.New(t)
	ss, err := RenderBoxScript("ss.csdc.io", "dacha", "Diveinedu",
		7006, "chacha20", "https://cloudfeet.com/files/gfw.list")
	a.NoError(err, "生成模板成功")
	fmt.Println(ss)
}
