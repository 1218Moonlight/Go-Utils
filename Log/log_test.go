package Log

import (
	"github.com/Kioryu/Go-Utils/Log"
	"testing"
)

var (
	c = Log.Config{
		InfoFileName: "infoLogs.txt",
		ErrFileName:  "errLogs.txt",
		InfoPrefix:   "[ INFO ] ",
		ErrPrefix:    "[ ERR ] ",
	}
)

func TestLogger(t *testing.T) {
	log := Log.Init(c)
	defer log.Close()

	for i := 0; i < 10; i++ {
		log.Info.Println("Test info", i)
		log.Err.Println("Test Err", i)
	}

}

func TestReadTXT(t *testing.T) {

	log := Log.Init(c)
	defer log.Close()

	origin, revert := log.ReadTXT(c.InfoPrefix)
	t.Log(origin)
	t.Log(revert)

	origin, revert = log.ReadTXT(c.ErrPrefix)
	t.Log(origin)
	t.Log(revert)
}
