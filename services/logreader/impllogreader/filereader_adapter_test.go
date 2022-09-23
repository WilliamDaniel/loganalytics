package impllogreader

import "testing"

func TestAdapter(t *testing.T) {
	const filepath = "../testdata/logs_test.txt"

	adapter := NewLogReaderAdapter(filepath)
	adapter.Read()

}
