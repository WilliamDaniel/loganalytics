package shared

type LogFile struct {
	Content []string
}

func (l *LogFile) Check() error {
	if len(l.Content) == 0 {
		return ErrEmptyLogFile
	}
	return nil
}
