package processor

import (
	"ledgerCo-loans/entity"
)

// IProcess interface
type IProcess interface {
	Process() bool
}

// New returns IProcess interface
func New(t entity.ProcessorType, args []string) IProcess {
	switch t {
	case entity.FileProcessor:
		filePath := ""
		if len(args) >= 1 {
			filePath = args[1]
		}
		return FileProcessor{
			FilePath: filePath,
		}
	}
	return nil
}
