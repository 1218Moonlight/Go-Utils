package Log

import (
	"io"
	"log"
	"os"
)

type logger struct {
	Info     *log.Logger
	infoFile *os.File
	Err      *log.Logger
	errFile  *os.File
}

func LoggerInit() logger {
	infoLogFile, err := os.OpenFile("InfoLog.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	errLogFile, err := os.OpenFile("ErrLog.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	infoLog := log.New(os.Stdout, "[ INFO ] ", log.LstdFlags|log.Ltime|log.Lshortfile)
	multiWriter := io.MultiWriter(infoLogFile, os.Stdout)
	infoLog.SetOutput(multiWriter)

	errLog := log.New(os.Stdout, "[ ERR ]", log.LstdFlags|log.Ltime|log.Lshortfile)
	multiWriter = io.MultiWriter(errLogFile, os.Stdout)
	errLog.SetOutput(multiWriter)

	return logger{
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
