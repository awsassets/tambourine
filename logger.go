package tambourine

import (
	"fmt"
	"io"
)

// Nóva says libraries should not log,
// but here you can define a logger.
type Logger interface {
	Success(format string, a ...interface{})
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Critical(format string, a ...interface{})
	Deprecated(format string, a ...interface{})
	Always(format string, a ...interface{})
	SetWriter(w io.Writer)
}

type StdOutLogger struct {
	writer io.Writer
}

func (l *StdOutLogger) Success(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
func (l *StdOutLogger) Debug(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
func (l *StdOutLogger) Info(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
func (l *StdOutLogger) Warning(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
func (l *StdOutLogger) Critical(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
func (l *StdOutLogger) Deprecated(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
func (l *StdOutLogger) Always(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
func (l *StdOutLogger) SetWriter(w io.Writer) {
	l.writer = w
}
