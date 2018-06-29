package writers

import (
	"os"
	"strconv"
	"testing"
)

func Test_Rotate(t *testing.T) {
	max, _ := strconv.Atoi(os.Args[2])
	RotateFile(os.Args[1], max)
	t.Log("Rotate ", os.Args[1], " to ", os.Args[2], " done")
}
