package slicingfile

import (
	"math"

	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/domain/entity/file"
)

const SizePerChunk float64 = 5 * 1024 * 1024

type SlicingFile struct {
	File       *file.File
	TotalParts int64
	UploadId   string
	UploadUrl  []string
}

func New(f *file.File) *SlicingFile {
	return &SlicingFile{
		File: f,
	}
}

func (f *SlicingFile) SetUploadId(uploadId string) *SlicingFile {
	f.UploadId = uploadId
	return f
}

func (f *SlicingFile) SetTotalParts() *SlicingFile {
	f.TotalParts = int64(math.Ceil(float64(f.File.FileSize) / SizePerChunk))
	return f
}
