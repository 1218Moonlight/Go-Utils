package Log

import (
	"io"
	"log"
	"os"
	"strings"
)

type logger struct {
	config   Config
	Info     *log.Logger
	infoFile *os.File
	Err      *log.Logger
	errFile  *os.File
}

// infoFileName = "InfoLog.txt"
// errFileName  = "ErrLog.txt"
// infoPrefix   = "[ INFO ] "
// errPrefix    = "[ ERR ] "
func Init(c Config) logger {
	infoLogFile, err := os.OpenFile(c.InfoFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	errLogFile, err := os.OpenFile(c.ErrFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	infoLog := log.New(os.Stdout, c.InfoPrefix, log.LstdFlags|log.Ltime|log.Lshortfile)
	multiWriter := io.MultiWriter(infoLogFile, os.Stdout)
	infoLog.SetOutput(multiWriter)

	errLog := log.New(os.Stdout, c.ErrPrefix, log.LstdFlags|log.Ltime|log.Lshortfile)
	multiWriter = io.MultiWriter(errLogFile, os.Stdout)
	errLog.SetOutput(multiWriter)

	return logger{
		config:   c,
		Info:     infoLog,
		infoFile: infoLogFile,
		Err:      errLog,
		errFile:  errLogFile,
	}
}

func (l logger) Close() {
	l.infoFile.Close()
	l.errFile.Close()
}

// return (origin Line, revert Line)
func (l logger) ReadTXT(prefix string) ([]string, []string) {
	if strings.Contains(l.config.InfoPrefix, prefix) {
		return readFile(l.config.InfoFileName)
	} else if strings.Contains(l.config.ErrPrefix, prefix) {
		return readFile(l.config.ErrFileName)
	} else {
		panic("Prefix Empty")
	}
}
