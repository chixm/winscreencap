package gowin_test

import (
	"log"
	"testing"

	"github.com/chixm/gowin"
)

func TestLoad(t *testing.T) {
	gowin.LoadWinMetaData()
}

func TestAssembly(t *testing.T) {
	gowin.LoadWinMetaData()

	assembly := gowin.GetAssembly()

	log.Println(len(assembly), `found`)

	gowin.GetTypes()

	//	gowin.PrintMetaFile()
}
