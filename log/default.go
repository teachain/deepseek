package log

import (
	"github.com/mattn/go-colorable"
	"io"
	"os"
)

// DefaultHandler
/**
* 1 error
* 2 warning
* 3 info
* 4 debug
* 5 trace
 */
func DefaultHandler(verbosity int) {
	var ioStream Handler
	output := io.Writer(os.Stderr)
	logger := NewGlogHandler(StreamHandler(os.Stderr, TerminalFormat(false)))
	output = colorable.NewColorableStderr()
	ioStream = StreamHandler(output, TerminalFormat(true))
	logger.SetHandler(ioStream)
	logger.Verbosity(Lvl(verbosity))
	Root().SetHandler(logger)
}
