package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordMD5(t *testing.T) {
	a := assert.New(t)
	s := "D26384628B40402F|jhp1!T!C!#okkzBoTK8!"
	a.Equal("32ace738492a26008358910db1aaf0e4", EncodeMD5(s), "test box md5 sum")
}
