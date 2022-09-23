package logreader

type LogReaderGateway interface {
	Read() (*LogFile, error)
}
