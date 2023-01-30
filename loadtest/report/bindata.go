// Code generated by go-bindata. DO NOT EDIT.
// sources:
// comparison.tmpl.json (28.679kB)

package report

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _comparisonTmplJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x5b\x6f\xe3\xb8\x15\x7e\xcf\xaf\x20\x88\xb6\xc8\xb6\xc9\xd4\x76\x6e\x9b\x00\x7d\xc8\x66\x30\xd3\x01\x32\x6d\x9a\xcc\x6c\x1f\x66\x03\x83\x96\x8e\x65\x22\x14\xa9\x21\x29\x27\xde\xc0\xfd\xed\x05\x29\x59\xa6\x2e\xf1\x2d\x76\xe2\xec\xf0\x25\x11\x0f\x29\x8a\xfc\x78\x78\xf8\x7d\x94\x2c\x3d\xee\x20\x84\x09\xe7\x42\x13\x4d\x05\x57\xf8\x0c\x19\x13\x42\x98\x51\xa5\xf1\x19\xfa\x66\x53\x28\xb7\xda\x9c\x5e\x4a\x99\xfe\xc4\xf1\x19\x6a\xef\x4d\xad\x21\xd1\x44\x89\x54\x06\x80\xcf\x10\xde\xdf\x47\x1f\x25\xe9\x13\x4e\xd0\xfe\x3e\x76\x8a\x01\x27\x3d\x66\x8a\x68\x99\x82\x63\x1f\xd0\xb0\xc1\x4a\x03\xc1\x2f\x04\x13\xd2\xd4\x29\xa3\x1e\xd9\x6d\xed\xa1\x4e\xbb\xbd\x87\x3a\x47\x47\x7b\xa8\xfd\x93\x5b\x35\x27\xb1\xbd\xf6\xf9\xb4\x3b\xe8\x2f\xe8\x9c\x81\xd4\xca\x2d\xa7\x47\x89\x2d\x17\x12\x35\xe8\x09\x22\x43\x9c\xe7\x8d\xed\xff\xdb\x1d\x84\xc6\xa6\x38\x86\x90\xea\x4a\x6b\x71\xc4\x41\x7f\x0a\xf1\x19\xea\x1c\x1d\x76\x32\x8b\x24\xc9\xe0\x8b\x10\x4c\xd3\x64\x82\x09\xa6\xa6\x08\x4f\x19\xcb\x52\x1a\xa4\x6d\x90\xc9\x3f\x6e\x1d\x9c\x9e\x9c\x1c\x1f\x1d\x1e\x1f\x1d\xda\x5c\x46\xf9\x9d\x01\xfe\xdb\xad\x4d\x26\x84\x03\x53\x05\xf4\x13\xe0\x31\x61\x94\x28\x0b\x86\x1d\xa5\xf1\xa4\x47\xb8\x47\xac\xa5\x4f\x98\x2a\xb0\xb3\x9d\xbb\x04\x1e\xe9\x81\xb9\x66\xab\x64\x87\xa6\xe2\xee\xe8\x31\x41\x42\x0d\x4a\xef\xe7\xa6\xa2\x58\x03\x22\x99\x5d\x4a\x3b\x46\xe5\x4a\xfb\x94\x31\xd7\x4b\xac\xe1\xa3\x24\x21\x05\x6e\x7c\x6b\xda\xaa\x48\xd2\xf0\x4a\x4c\xbd\x2f\x73\x09\x7c\x86\x4e\x9d\x71\xbb\x37\x75\x75\x1c\xc3\x83\x5b\x07\x42\x78\x64\xd2\x93\xb1\x2c\xea\x1e\xd0\x30\x04\x7e\x03\x92\x36\xf4\xdb\x8e\xd3\x49\x91\x64\x10\x01\x0f\xcb\xcd\x20\xc3\xa8\xee\x98\x41\x2a\x65\xd6\x89\x52\x7d\x08\xe1\x98\x3c\xd4\x8b\xc7\x94\xd7\x8d\x6a\x20\xee\xeb\x56\x2d\x34\x61\x0d\xf5\x0e\x09\x4b\x6d\x07\x4c\xf9\x5a\x27\x19\xe5\x45\x66\xc9\x78\x4f\xc3\xcc\x05\x5c\xab\xe3\x6e\xd9\xcc\x49\x19\xbb\x12\x94\xeb\xcf\xc2\xce\x42\x1c\x08\xce\x21\xd0\x10\x4e\x87\x5e\x24\xe5\x00\x51\xb8\xcd\x65\x51\x5f\xad\x55\x09\xc8\x00\xb8\x26\x11\xd4\x80\x4f\xcc\xe5\x8c\x2b\xa4\xe6\xdc\xa3\xb2\xbd\x3e\x4e\x12\x78\x08\x12\x6c\x1c\xe8\x33\xa1\xa7\xed\x52\x76\x60\xff\x3d\x04\x29\x69\x08\x95\x8e\xa9\x84\x04\xd0\x34\x0d\x94\x26\xc1\x5d\xed\x2a\x4a\x43\x92\x40\x78\x49\x79\xbd\xc1\x9a\xc8\x08\xb4\x72\x42\xa2\x1b\x14\xcd\x1c\x78\x48\x6c\xf3\x54\x1a\xef\x4a\xa2\x61\x37\x26\x5a\x83\x8c\x85\xd2\xdd\xc4\xfc\xb1\x23\xfb\x48\xb9\xd2\x84\x07\xf0\x8f\xff\xfd\x86\xff\xa4\x40\x0e\x41\xfe\x86\xc7\xdf\xda\xf1\x2d\x12\xfd\xbe\x02\x8d\x1e\x1f\xdf\x65\x47\xe3\xb1\xfa\xc9\x8d\x70\x66\x06\x09\x19\x13\xe3\x76\x58\xd3\x18\xba\x59\xe7\xcb\x45\x28\xd7\x20\x87\x84\x7d\x20\x81\xb6\x93\xb2\x5d\xca\xce\x5c\xfc\x43\x51\xcf\xe3\xe3\xbb\x1e\x51\x70\x49\x7a\xc0\xc6\xe3\x72\x55\x31\x68\x49\x03\x53\xaa\xb1\x2f\xe5\xc2\x12\xfa\x36\x30\xe2\xf3\xb2\xdd\x80\x6a\x86\xb8\xb0\x8d\xf7\x36\x04\xe0\x8b\x80\xc5\xe1\x7e\x8d\x58\xfd\x32\x0f\xab\xfc\x68\xea\xd2\x7a\x20\x41\x0d\x04\x0b\x2b\xae\x6e\xba\xf8\x41\x8a\xd8\x59\x77\x0a\xfb\x35\x44\xf9\xdc\xad\x9c\x70\x33\xa0\x7d\x5d\x3f\x43\xdb\x08\x8f\xaf\x84\xd2\x0a\x25\x20\xd1\x0d\x04\x82\x3b\xc1\x40\x17\xcb\x9d\x13\x0c\x62\x75\x0d\x4a\xb0\x34\x5f\xe8\xaa\x01\x4c\x0d\x88\x84\xb0\x21\x0c\x0a\xa9\x2b\x81\xdc\xc6\xba\xee\x64\x9d\xa6\x3c\xa4\x43\x1a\xa6\x84\xe1\x5a\x84\x99\x94\xb1\x8b\xf0\xb4\x7d\x0f\xe4\x81\x56\x42\x55\x2f\x0d\xee\xb2\xf9\xeb\x76\xd6\x34\x3b\x8f\x79\x06\x8f\x06\x3a\x51\x29\xdd\x1c\xb4\x8b\xe0\xdc\x10\x04\x47\xe4\x01\x66\x84\x8d\xa9\x93\xaa\x81\x41\xa2\xec\x7f\xc6\xd3\x4c\xde\x85\x48\x79\x35\x4f\x44\xbf\x10\x05\x35\x9f\xcd\x16\xa0\x72\xb3\x8b\x25\xa8\x66\x76\xfa\x33\x77\x82\x2e\xd4\xd4\xda\x15\x36\xd8\xce\xda\xe4\x18\xd5\xc7\x9d\x30\x1a\x35\xb9\xa3\xb5\x5f\xc2\xb0\x68\x74\x89\x04\xe6\x10\x78\xf2\x35\x49\x37\x92\xaf\x92\x61\x55\xf6\xe5\xe0\xe3\xe9\x97\xa7\x5f\x9b\xa0\x5f\x03\xad\x93\xae\x84\xef\x29\x28\xad\xfe\x28\x3c\xcc\x76\xca\xce\x7d\xf5\xe2\x74\x6c\x39\x40\xb7\x8f\x97\x2d\x8a\xdd\x56\xd3\xb3\xeb\x1c\x7f\xcf\xd0\x3c\x43\x5b\xa6\xa9\x9e\xa1\xfd\x38\x0c\xad\xba\x3d\x76\xba\x02\x41\x3b\x5e\x80\x9f\x55\xc7\x6e\x1e\x41\xab\x59\x69\x73\x40\x5a\x89\xa2\xd9\x0c\xcf\xd1\xb6\x9b\xa3\x55\x57\xe3\x7b\xe8\x29\x61\x17\x80\x99\x7c\xa2\x91\x9c\xe1\x95\x99\x95\xc3\x92\x96\x22\x44\x2b\xb6\xfe\x35\x89\xd0\xdb\x20\x35\xff\x4a\xe3\x1e\x48\x24\xfa\xe8\x62\x32\xc1\xd0\x7b\x18\xd2\x00\x14\xda\xfd\x2f\xf4\x6e\x2c\xca\x93\x4c\x73\x85\x9f\x3c\xe9\xf1\xa4\xc7\x93\x1e\x4f\x7a\x10\x7a\x6a\x5b\x6a\x15\xd6\xd3\x5e\x80\xf5\xf8\x5d\x29\xcf\x78\x56\x62\x3c\x61\xaf\x1b\x13\xa5\x41\x76\x83\xe9\x42\xf6\x5c\xde\xb3\x05\x7b\x52\x33\xfa\xb5\xc1\xdd\xa9\x67\x01\xbb\x75\x7b\x53\xcb\x63\xf8\xb6\x08\x9d\x7d\x2e\x47\x0b\xf4\xd9\x76\x12\xbd\x27\x9a\x18\x27\xf3\x2c\xce\xb3\x38\xcf\xe2\x3c\x8b\x43\xe8\x89\xad\xab\xf6\xcf\xab\xb0\x38\x4f\xe3\x3c\x8d\x43\x6b\xa5\x71\xd5\x7b\x61\x24\xa1\xdd\x8c\x2b\xa4\xf1\xd3\xb7\x0f\xd1\xdf\xd1\xcc\x93\x03\x13\x83\xd7\x78\xf7\x31\x7f\x9e\xb6\xe2\x69\xcb\xb3\x17\x32\x8c\xd0\xd2\x34\xd0\x85\xe4\x05\xef\x4a\x56\x47\x62\x09\xd4\x5f\x1f\xe1\xa5\x38\xe2\x02\x00\xcf\x25\x85\x73\x00\x1e\x50\xa5\x45\x24\x49\xdc\xfd\x9e\x12\xae\x29\x83\xdd\xd6\xbb\xd3\xd3\xbd\xd9\x80\x66\x44\x68\xc6\x34\xe8\x8d\xd0\x2e\x83\x75\xc2\x5d\xc5\x33\x39\x3d\x9d\xe9\xb1\x05\x42\x17\x8b\x6f\xf9\x3e\x17\x8c\x97\xec\xf8\xdc\xfd\xdf\xf7\x78\x3b\x94\xc1\xf9\xd5\x27\x94\xdf\xc3\x46\x5f\x68\x0c\x68\xf7\xc6\x8a\x32\xbf\xa1\xbb\x61\x29\xf0\x84\x0c\x78\x4a\x01\x74\xbc\x02\x98\x42\xec\x15\x40\x3e\x58\xaf\xa0\x00\x3a\x27\x2b\x28\x80\x83\x43\xaf\x00\xbc\x02\xd8\xa8\x02\x08\x7b\x5d\xa5\x85\x84\xd5\x65\x40\xb9\x06\xaf\x05\x9e\xaf\x05\x1a\xc7\x64\x59\xfc\x5f\x1f\xeb\x37\xad\x0a\xca\xa8\x7a\x69\xf0\x34\x22\x5e\x1f\x34\xe8\x83\x1b\x83\x14\xba\x20\x8c\x79\x79\xe0\xe5\x81\x97\x07\x5e\x1e\x94\xe5\x41\xed\x31\x8f\x95\xee\x10\x1c\x74\xbc\x3e\xf0\xfa\x60\xfd\xfa\x60\x32\x91\x4a\xbf\x95\x99\xc1\x46\x67\x94\xdf\x22\x3e\xba\xbd\xdc\x7f\x3e\xde\xf3\x15\xd9\x22\x63\xb0\x9d\x9a\xec\xed\xe9\x84\x19\x58\xff\xb0\x4a\x61\x21\x4c\xbc\x56\x58\xe8\x5e\xc2\x05\x33\x54\xc3\x8b\x05\x2f\x16\xbc\x58\xc8\x32\x7e\x68\xb1\xb0\x96\x7b\x09\xed\x03\xaf\x15\xbc\x56\x58\x9b\x56\x20\xc3\xa8\xfe\xa2\x23\x29\x02\x50\xaa\x1b\x24\x69\x57\xd9\xdf\x74\xaf\xfa\xc2\x82\xbf\xa2\x76\xab\xf5\x9a\x4f\x88\xaf\x8d\xe3\xaf\x0b\xa7\x97\xc3\xe4\x6d\xff\x06\xef\xe2\xea\x2b\xfa\xaa\x29\xa3\xbf\xdb\x97\x16\xa2\x6b\xa2\x01\xed\xfe\xd9\x73\xa9\xd7\x7b\x44\x7b\xe6\x70\xa0\x0d\x73\x17\xcf\xb1\x3c\xc7\x5a\x69\x43\x76\x25\x92\x75\xe4\x39\x96\xe7\x58\x6b\xe5\x58\x91\xe8\xc6\x10\x2b\x4d\xb4\xea\xda\x4b\x75\x29\x4f\x15\x74\x7b\x23\x0d\x6a\xa3\xbf\xb7\x5b\xd7\x2e\xdf\x8d\x69\xf5\x39\x63\x22\x58\x70\x13\xd6\xed\x32\x31\xe7\x65\x9d\xdd\xe8\x4f\xf1\x56\x85\x7a\x5b\x60\x9d\xbf\x87\xba\x0a\xaa\xcf\xdd\x49\xad\xa2\x3a\x00\x92\xbc\x39\xff\xfd\x27\x90\x64\xd3\xee\x7b\xf1\x3a\x40\x6f\x09\xa8\x9b\x71\xde\xf7\xdb\x2c\x53\x3e\x43\x2c\xe4\x08\x7d\x55\x66\x45\xf3\xd2\x64\x93\xd2\xc4\x7a\xca\x92\x5b\xbd\x3f\x90\x0c\x71\x5f\x38\xe5\x75\xc8\xc4\xbe\x3d\x3a\xe4\xe0\x78\x05\x1d\xe2\x9f\x1b\xf7\x3a\x04\xad\x4d\x87\xa8\x34\x36\xec\x22\x12\x52\xa4\xda\x0c\xcb\xdb\x7b\xcf\x47\xa9\xf9\x1b\x7c\x9c\x63\x11\x98\xb6\xe3\xad\x1d\x8b\x20\xb2\xd5\x3b\xbd\xd3\x97\x73\x7c\x14\xe8\xba\xda\x11\xcf\xa5\xfc\x9b\x38\xfe\xe0\xfb\xba\x3b\x79\xb5\x66\xbe\x9a\x59\xe7\xd6\x80\x55\x30\x80\x98\xfc\x0a\x52\x65\xae\xde\xc9\x3e\xf8\xa4\xf4\x28\x9b\x3d\x21\x91\x77\x16\x1d\xac\x49\x34\x9d\x7b\x58\x43\x9c\x30\xa2\x29\x8f\x16\xf9\x9c\x16\x61\xec\x57\xe3\x57\x75\x27\x9c\xd2\x89\xd2\xf8\x68\x78\xc8\x1e\xfd\x4a\x92\xfd\xd6\xd9\xcf\xad\xe3\x13\xf4\x37\x64\x12\x6d\x9b\x28\x8f\xd7\x30\xaf\xfa\x9b\x63\x44\xee\xb9\xa5\xe2\x79\x4e\x5e\x91\x93\x71\xdb\xe4\x35\x0b\x32\x44\x53\x10\xfa\x94\xd3\x3c\x62\x64\x3e\xd4\xcd\x26\xd3\xee\x24\xbe\x97\x3e\xe5\x95\x4b\x64\x37\x58\x50\x1e\xb0\x34\x84\x73\xd6\x44\x83\x8a\x09\x94\x2d\x10\x6e\x55\x71\xca\x34\xad\x4f\xe5\xc9\xc7\xc2\xea\x27\x4c\x29\xcd\x34\x94\x22\x84\xbf\xa7\x20\x47\x8b\xb5\x7e\xea\x4b\xed\x92\x35\x82\x87\x8a\x60\xc2\xea\x8e\x26\x5f\x25\xbb\x19\xf1\xa0\x29\x96\xd6\x63\xa6\x26\x91\xf5\x16\xf5\x9f\x49\x7b\x70\x39\xb7\xd6\x70\x63\x6b\x2e\x9c\x07\xd5\xac\x67\x4e\x46\xaa\xe0\x4b\x56\x51\xe9\x4d\xba\x3b\xa8\xf4\x75\x34\x1b\x48\x0b\xff\xee\x67\x0b\x93\x59\x37\xcd\x61\xb1\x66\x62\x2d\x72\xb3\x16\xe3\x31\x2e\x9d\x9d\xd0\xe0\xce\x92\xb7\xbc\x8e\x1c\xb7\xee\x64\xa9\x76\xc3\x29\x3e\x72\xd6\xa5\x76\xcb\x49\x1c\xb8\x89\xf6\xf4\x81\x46\x7c\xe4\x1c\xb7\xdd\xc4\x41\xcb\xcd\x71\xd6\x93\x8e\x73\xdc\xce\x3f\x0b\x77\x3b\xe9\x87\x21\x16\x8e\x6f\xcc\xbd\x8a\x5b\xf1\xb1\x5b\xb1\x7b\x95\xce\xa1\x9b\x70\x88\xf5\x49\xe8\xb6\x77\xd2\x96\x12\x7c\xbf\x0b\x4b\x48\x71\x4f\x8a\x7b\x95\xfb\xf0\x74\x69\x37\x88\x9b\xe3\x7c\x24\x70\x5a\xfe\x24\xdd\xb0\x88\x6a\x87\x3b\xe3\x9d\xff\x07\x00\x00\xff\xff\x93\x3d\xab\x74\x07\x70\x00\x00")

func comparisonTmplJsonBytes() ([]byte, error) {
	return bindataRead(
		_comparisonTmplJson,
		"comparison.tmpl.json",
	)
}

func comparisonTmplJson() (*asset, error) {
	bytes, err := comparisonTmplJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "comparison.tmpl.json", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x35, 0x1c, 0x4e, 0x2a, 0x53, 0x60, 0x0, 0x2c, 0x40, 0x8b, 0x9e, 0x24, 0x18, 0xb3, 0x95, 0xd0, 0x9f, 0x6c, 0x3a, 0x5c, 0x14, 0xed, 0xeb, 0xfa, 0xc1, 0x34, 0xee, 0x9d, 0x7, 0xc0, 0x52, 0x79}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"comparison.tmpl.json": comparisonTmplJson,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"comparison.tmpl.json": {comparisonTmplJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
