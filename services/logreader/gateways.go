package logreader

import "github.com/WilliamDaniel/loganalytics/shared"

type LogReaderGateway interface {
	Read() (*shared.LogFile, error)
}
