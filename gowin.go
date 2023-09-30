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

// GetAssembly
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

// GetTypeDefs
func GetTypeDefs() []*winmd.TypeDef {
	var td []*winmd.TypeDef
	for i := 0; i < int(metaData.Tables.TypeDef.Len); i++ {
		ts, err := metaData.Tables.TypeDef.Record(winmd.Index(i))
		if err != nil {
			logger.Errorln(`skipped typedef`, err)
			continue
		}
		logger.Infoln(`typedef :`, ts.Name.String())
		td = append(td, ts)
	}
	return td
}

// GetModules
func GetModules() []*winmd.Module {
	var md []*winmd.Module
	for i := 0; i < int(metaData.Tables.Module.Len); i++ {
		ts, err := metaData.Tables.Module.Record(winmd.Index(i))
		if err != nil {
			logger.Errorln(`skipped module`, err)
			continue
		}
		logger.Infoln(`module :`, ts.Name.String())
		md = append(md, ts)
	}
	return md
}

func GetCustomAttributes() []*winmd.CustomAttribute {
	var md []*winmd.CustomAttribute
	for i := 0; i < int(metaData.Tables.CustomAttribute.Len); i++ {
		ts, err := metaData.Tables.CustomAttribute.Record(winmd.Index(i))
		if err != nil {
			logger.Errorln(`skipped custom attribute`, err)
			continue
		}
		logger.Infoln(`custom attribute :`, ts.Parent.Index, `type: `, ts.Type, ts.Value)
		md = append(md, ts)
	}
	return md
}

// metaDataから生成
func getGUID(t *winmd.AssemblyRef) *ole.GUID {
	return ole.NewGUID(``)
}
