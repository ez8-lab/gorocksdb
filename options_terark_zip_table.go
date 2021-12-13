package gorocksdb

// #include "rocksdb/c.h"
// #include "gorocksdb.h"
import "C"
import (
	"unsafe"
)

// TerarkZipTableOptions represents terark-zip table options.
type TerarkZipTableOptions struct {
	c *C.rocksdb_terark_zip_table_options_t
}

// NewDefaultTerarkZipTableOptions creates a default TerarkZipTableOptions object.
func NewDefaultTerarkZipTableOptions() *TerarkZipTableOptions {
	return NewNativeTerarkZipTableOptions(C.rocksdb_terark_zip_options_create())
}

// NewNativeTerarkZipTableOptions creates a TerarkZipTableOptions object.
func NewNativeTerarkZipTableOptions(c *C.rocksdb_terark_zip_table_options_t) *TerarkZipTableOptions {
	return &TerarkZipTableOptions{c: c}
}

// Destroy deallocates the TerarkZipTableOptions object.
func (opts *TerarkZipTableOptions) Destroy() {
	C.rocksdb_terark_zip_options_destroy(opts.c)
	opts.c = nil
}

func (opts *TerarkZipTableOptions) SetLocalTempDir(dir string) {
	var (
		cDir = C.CString(dir)
	)
	defer C.free(unsafe.Pointer(cDir))
	C.rocksdb_terark_zip_options_set_local_temp_dir(opts.c, cDir)
}

func (opts *TerarkZipTableOptions) SetIndexNestLevel(indexNestLevel int32) {
	C.rocksdb_terark_zip_options_set_index_nest_level(opts.c, C.int32_t(indexNestLevel))
}

func (opts *TerarkZipTableOptions) SetSampleRatio(sampleRatio float64) {
	C.rocksdb_terark_zip_options_set_sample_ratio(opts.c, C.double(sampleRatio))
}

func (opts *TerarkZipTableOptions) SetTerarkZipMinLevel(terarkZipMinLevel int) {
	C.rocksdb_terark_zip_options_set_terark_zip_min_level(opts.c, C.int(terarkZipMinLevel))
}
