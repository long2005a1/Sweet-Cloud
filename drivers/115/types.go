package _115

import (
	"time"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/pkg/utils"
	"github.com/lm379/115driver/pkg/driver"
)

var _ model.Obj = (*FileObj)(nil)

type FileObj struct {
	driver.File
}

func (f *FileObj) CreateTime() time.Time {
	return f.File.CreateTime
}

func (f *FileObj) GetHash() utils.HashInfo {
	return utils.NewHashInfo(utils.SHA1, f.Sha1)
}
