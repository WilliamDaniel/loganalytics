package impllogreader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdapter(t *testing.T) {
	const filepath = "../testdata/logs_test.txt"

	adapter := NewLogReaderAdapter(filepath)
	logFile, err := adapter.Read()
	if err != nil {
		panic(err)
	}
	assert.NotEmpty(t, logFile)
	assert.NoError(t, logFile.Check())

}
