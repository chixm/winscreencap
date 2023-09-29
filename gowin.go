package gowin

import (
	"debug/pe"

	"github.com/go-ole/go-ole"
	"github.com/microsoft/go-winmd"
)

const (
	RO_INIT_SINGLETHREADED = 0
	RO_INIT_MULTITHREADED  = 1
)

func init() {
	ole.RoInitialize(RO_INIT_MULTITHREADED)
	LoadWinMetaData()
}

// Windows Meta Data is placed at
// C:\Windows\System32\WinMetadata
const pathGraphics = `C:\Windows\System32\WinMetadata\Windows.Graphics.winmd`

var metaData *winmd.Metadata

func LoadWinMetaData() error {
	pefile, err := pe.Open(pathGraphics)
	if err != nil {
		return err
	}
	f, err := winmd.New(pefile)
	if err != nil {
		return err
	}
	logger.Infoln(`loaded version`, f.Version)
	metaData = f
	return nil
}

func PrintMetaFile() {
	logger.Infoln(string(metaData.Strings))
}

func GetAssembly() []*winmd.Assembly {
	var assem []*winmd.Assembly
	for i := 0; i < int(metaData.Tables.Assembly.Len); i++ {
		az, err := metaData.Tables.Assembly.Record(winmd.Index(i))
		if err != nil {
			logger.Infoln(`skipped assembly`, i)
			continue
		}
		logger.Infoln(`load :`, az.Name.String())
		assem = append(assem, az)
	}
	return assem
}

// metaDataから生成
func getGUID(t *winmd.AssemblyRef) *ole.GUID {

	return ole.NewGUID(``)
}
