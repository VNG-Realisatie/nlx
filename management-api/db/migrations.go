// Package db Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// db/migrations/001_initial_schema.down.sql
// db/migrations/001_initial_schema.up.sql
// db/migrations/002_authorization.down.sql
// db/migrations/002_authorization.up.sql
// db/migrations/003_audit_log.down.sql
// db/migrations/003_audit_log.up.sql
package db

import (
	"bytes"
	"compress/gzip"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __001_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xdd\x4e\xc4\x20\x10\x85\xef\x79\x0a\x9e\xc0\x17\xe0\xaa\x5b\x1b\x6d\xb2\x6d\x4d\x97\x7b\x32\x21\x2c\x21\xda\x41\x19\xd6\xbf\xa7\x37\x2e\x41\xa3\xd9\x22\x4d\xbc\x24\xf9\xe6\x9c\xe1\xcc\xd9\x75\x37\xfd\xc8\x63\x00\x24\xd0\xd1\x79\x14\x8c\x35\x7b\xd9\xcd\x5c\x36\xbb\x7d\xc7\xf1\xe1\x55\x2d\x80\x60\xcd\x62\x30\x5e\x81\xd6\x86\x48\x05\xf3\x74\x32\x14\x49\x39\xd4\x7e\x71\x68\xf9\xf5\x3c\xdd\xf1\x76\x1a\x0f\x72\x6e\xfa\x51\xf2\xe3\xbd\x5a\x63\x15\x99\xf0\xec\xb4\x11\x15\x46\x36\x00\x46\x2a\xa8\x27\xe0\x5b\xfb\xa7\x69\x8d\xc5\x63\xf0\xfe\x58\xb2\x48\x80\xf2\xa7\x68\xfd\x46\x0b\x87\x2f\xf0\x46\xf9\xc3\x17\x4d\xce\xc8\x17\x91\x9e\xff\xac\x59\x93\x37\x99\x18\x1d\xda\x8b\x72\x3e\x58\x40\xf7\x0e\x9f\xf5\x50\x19\xcc\x9b\xb2\xf3\xc0\x8a\x66\x5a\x40\x14\x98\xb5\x92\x6c\x99\xc9\x97\xa9\x98\x49\x75\xa9\x00\xd3\xd1\x4b\x60\xba\xc3\xdf\x84\xaa\x89\x21\xa7\x9a\xf3\x3c\xb4\xb7\xdd\xd0\xfc\xa2\x04\x63\xed\x34\x0c\xbd\x14\xec\x23\x00\x00\xff\xff\x08\x17\x6e\x73\xb7\x03\x00\x00")

func _001_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initial_schemaDownSql,
		"001_initial_schema.down.sql",
	)
}

func _001_initial_schemaDownSql() (*asset, error) {
	bytes, err := _001_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_initial_schema.down.sql", size: 951, mode: os.FileMode(420), modTime: time.Unix(1610437247, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __001_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x5d\x73\xab\x36\x10\x7d\xe7\x57\xec\x5b\xec\x19\x4f\x27\xed\xe4\x3e\xe5\x89\x38\xca\x0d\x73\x1d\xdc\x62\xd2\x69\x9e\x34\xba\x62\xed\xa8\xc6\x12\x95\x44\x3e\xee\xaf\xef\x00\xc6\x80\x6d\x61\x3c\xb7\xe9\x84\x27\xe3\x3d\x2c\xab\x3d\x67\x8f\xd0\x0d\xf9\x1a\x84\x60\x35\x93\x86\x71\x2b\x94\xbc\x06\xcf\x9b\x46\xc4\x8f\x09\x2c\xa6\xf7\xe4\xc1\x07\x99\xbe\xd1\x0d\x93\x6c\x85\x1b\x94\xf6\x7a\x17\x8e\xfd\x9b\x19\xd9\x8b\xfe\x62\x50\xbf\x08\x8e\x06\x46\x1e\xc0\xf6\x86\x8a\x04\x16\x24\x0a\xfc\x19\xfc\x1e\x05\x0f\x7e\xf4\x04\xdf\xc8\xd3\xc4\x03\x90\x6c\x83\xf0\xa7\x1f\x4d\xef\xfd\x68\xf4\xdb\x97\xcb\x31\x84\xf3\x18\xc2\xc7\xd9\x0c\x1e\xc3\xe0\x8f\x47\x52\x80\x50\x26\x99\x12\xd2\xd2\x5c\xa7\xc7\xc1\x05\x2a\x51\x3c\x2f\x2a\x60\xc5\x22\xfa\xa1\x2c\x13\xd4\x64\xc8\xc5\x52\xf0\x01\x70\x21\x2d\x6a\xc9\x52\xf8\xae\x54\x8a\x4c\x36\x35\xde\x92\x3b\xff\x71\x16\xc3\x92\xa5\x06\x0b\xa8\x45\xfe\x4c\x4d\x9e\x65\x4a\x5b\xca\x95\xb4\x8c\x5b\x77\xe2\x2c\xff\x9e\x0a\x3e\x1c\xcf\x35\x32\x8b\x09\x65\x16\xac\xd8\xa0\xb1\x6c\x93\xc1\xab\xb0\xcf\xe5\x2d\xfc\x50\x12\x3b\xf8\x3c\x4b\x06\xe3\xc7\xd7\x5e\x3f\xb1\x8c\x73\x34\x86\x6a\xfc\x27\x47\x63\x0d\x15\x92\xab\x8d\x90\xab\x92\xe8\x6e\x70\x17\x73\x13\xdf\x52\x46\x10\xc6\x9d\xaa\x95\x5e\x31\x29\x7e\x54\xc4\xb8\x05\xd2\x6a\xe0\x1a\xdf\xe9\x52\xc8\x15\xea\x4c\x0b\xd9\x34\xf0\xea\xaa\x0b\x37\x96\xd9\x26\x5b\x91\x6c\x7a\x4f\xa6\xdf\x60\x54\x05\x82\x10\x46\x17\x1a\x39\x8a\x17\x4c\x2e\x26\x70\xc1\xb2\x4c\xab\xed\x6f\x8d\x7f\x23\xb7\xf5\xef\x17\xb5\xc6\xe4\x62\x3c\x3e\x94\x42\x2b\xc1\x07\x33\x56\xe0\xa7\xf3\x70\x11\x47\x7e\xd1\xc2\xe5\x9a\xba\x28\xa2\xdb\x6e\x7b\x00\x00\x77\xf3\x88\x04\x5f\xc3\x82\x08\x18\x35\x34\x8c\xcb\x60\x44\xee\x48\x44\xc2\x29\x59\xb8\xe7\xba\xf5\x0c\x94\x0f\xcd\x43\xb8\x25\x33\x12\x13\x88\xc8\x22\x8e\x82\x69\xec\x8d\x1b\x97\x08\xc2\x5b\xf2\xd7\x90\xea\x8a\x3c\x83\x35\xd7\x2a\x62\xa0\x70\x57\x9a\x49\x6b\xda\x6a\x2d\xff\x71\x4b\xb4\x47\xd3\xfb\x92\x3d\x97\xe6\xad\x80\xfa\xf1\x6d\x55\xf5\xf2\x5d\xad\xac\x29\xb0\x5b\xf8\x21\xe9\xee\x85\x9d\x12\x81\x9b\x8f\x53\x39\xcf\xd3\x48\xff\x8a\xdc\x42\xa9\x39\xee\xa9\xe6\x5c\x9b\x53\xb9\x5d\x29\x87\xcd\xd5\x31\xb7\x86\xdc\x5e\xf6\xeb\xe5\x9e\x97\xd5\x8a\xee\x77\xbc\x21\x16\xb6\x55\x63\xe1\x54\x4b\x26\xd2\xda\xb3\x7e\xd2\xd7\x76\x59\xcf\xf7\x5d\x8d\x4b\xd4\x28\x2b\xbb\x6f\xae\xf6\x14\xed\xde\x73\x59\xee\xf8\x5a\x2b\x4d\xb9\x4a\x10\xe0\x9c\x07\x58\x6e\x3a\x4f\x34\x7d\xba\x1c\x37\x30\x63\x19\x5f\x53\xab\x19\xdf\x82\x6f\x9e\x62\xe2\x17\xf1\x54\xf1\x75\xb7\xc6\xf2\xca\x73\x91\xb4\xbe\x47\x4a\x14\xbe\x65\x42\xa3\x29\x46\x78\x7b\xb9\x26\x79\x02\xf0\xc1\x3b\x81\x07\x70\x30\x4c\x22\x79\xa3\x1d\xf5\x0d\x77\xd9\x46\xf2\x07\xf2\x9d\x74\x64\x3a\x74\x98\x32\xad\xd4\xb2\x63\xbd\xe5\x3f\x83\xad\xb7\x3d\x67\x85\x06\x3e\x91\xf7\x56\x4b\x6b\x2a\x3c\xd3\x7b\x5b\x2b\x3b\xd7\x7b\x1b\x96\x4e\xe5\x3c\xcf\x7b\xfb\x57\xe4\x96\x4f\x4d\x72\x4f\x35\xa7\xe4\x22\xe4\x2b\x7b\x37\x30\x82\xf2\xa3\xfb\x95\xbd\xff\xe4\xc1\xc1\x60\xba\xa4\x2c\x49\x34\x1a\xe3\x36\xde\x17\xd4\x46\x28\xe9\x06\x3c\x2b\x63\xfb\x5d\x59\x64\x07\xaf\xb9\xfa\xf2\x89\xbe\xdd\xab\xc6\xd2\xce\xd9\x6c\xbf\xc1\xc7\xf6\xa2\xe3\xd1\x16\x15\x30\xaa\xd3\x4c\x5a\x0f\x8d\x0f\xc7\xa5\x82\xd5\x05\x54\xb7\x87\xd3\x51\x27\x3b\x35\x0b\xb5\x50\xba\xf8\x43\x9d\x9f\xac\xe3\xa3\x3e\x8c\xfb\xe6\xee\xc4\xf1\xd9\x5a\x21\x57\xf5\xf1\xb9\xba\x71\x8f\xc1\x8e\xc4\x72\x6b\x3c\xe6\x5a\x42\x1a\xb1\x7a\xb6\xb4\x38\xf3\xf6\x9f\x72\xf5\x86\x95\xfd\x40\xdd\x0f\xfc\x7f\xcf\x34\x7b\x7b\x58\xdd\x90\xff\x5a\x3f\xc7\x0c\xb1\xe7\xd5\x47\x6c\xb0\x61\x6e\x97\xb5\x24\x7b\xfe\xf0\x10\xc4\xd7\xde\xbf\x01\x00\x00\xff\xff\x28\x2e\x3c\x2a\x68\x11\x00\x00")

func _001_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initial_schemaUpSql,
		"001_initial_schema.up.sql",
	)
}

func _001_initial_schemaUpSql() (*asset, error) {
	bytes, err := _001_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_initial_schema.up.sql", size: 4456, mode: os.FileMode(420), modTime: time.Unix(1611732984, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __002_authorizationDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x72\x75\xf7\xf4\x53\x28\x29\x4a\xcc\x2b\x4e\x4c\x2e\xc9\xcc\xcf\xb3\x56\xe0\xe2\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\xcb\xa9\x88\xcf\x4d\xcc\x4b\x4c\x4f\xcd\x4d\xcd\x2b\xd1\x2b\x48\x2d\xca\xcd\x2c\x2e\xce\xcc\xcf\x2b\x8e\x2f\xca\xcf\x49\x2d\xb6\xc6\xa3\xb8\xb4\x38\xb5\x88\x08\x65\x48\x66\x12\x34\x0d\x9f\x02\xa8\x45\x5c\xce\xfe\xbe\xbe\x9e\x21\xd6\x5c\x80\x00\x00\x00\xff\xff\x08\xc2\x5e\x93\xdb\x00\x00\x00")

func _002_authorizationDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__002_authorizationDownSql,
		"002_authorization.down.sql",
	)
}

func _002_authorizationDownSql() (*asset, error) {
	bytes, err := _002_authorizationDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "002_authorization.down.sql", size: 219, mode: os.FileMode(420), modTime: time.Unix(1612279495, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __002_authorizationUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x54\xcf\x6f\x9b\x30\x14\xbe\xf3\x57\x7c\xb7\x80\x64\x4d\xd3\xa4\x9d\x72\x72\xe9\x4b\x67\x0d\x4c\x67\x4c\xa7\x9e\x90\x55\xbc\x0d\x2d\x40\x04\xae\x36\xed\xaf\x9f\xa0\x99\x6a\x46\x92\xa6\x87\x48\xd3\x4e\xd8\xbc\x9f\xdf\xfb\x3e\xbf\x2b\xba\x11\x12\xae\x37\xed\x60\x1e\x5c\xdd\xb5\x6b\x04\x41\xac\x88\x6b\x82\xe6\x57\x09\xa1\xdd\xfe\x2c\x1b\xd3\x9a\xaf\xb6\xb1\xad\x7b\xb3\xb3\x7d\x53\x0f\x43\xdd\xb5\x03\xc2\x00\x78\xbe\x97\x75\x85\x9c\x94\xe0\x09\x6e\x95\x48\xb9\xba\xc7\x47\xba\x67\x01\xd0\x9a\xc6\xe2\x8e\xab\xf8\x03\x57\xe1\xbb\xf7\x6f\x23\xc8\x4c\x43\x16\x49\x32\x5a\x1f\xba\xea\x88\x15\x85\x14\x9f\x0a\x9a\x9c\x7a\x6b\x9c\xad\x4a\xe3\xe0\xea\xc6\x0e\xce\x34\x3b\xfc\xa8\xdd\xb7\xe9\x8a\x5f\x5d\x6b\x67\x49\x1f\x77\xd5\xd9\xfe\x41\xb4\x3e\x8d\xf9\x71\xb0\xfd\x13\xda\xf1\x74\x1c\xa7\x6d\x4c\xbd\xfd\xb7\xa1\xf4\xdd\xd6\x3e\x41\x19\x4f\xff\x2f\x65\x9e\x4c\xcb\xa3\x98\xfd\xea\x07\x75\xec\x3b\x78\x03\x42\xb8\x4f\xc4\xe6\x61\xd1\xa5\x61\xb3\x20\x00\xe2\x4c\xe6\x5a\x71\x21\x35\xbe\x7c\x9f\xc0\x05\x00\xb0\xc9\x14\x89\x1b\x39\xeb\x2f\x9a\x2c\x8a\x36\xa4\x48\xc6\x94\x1f\x54\xc3\xdc\x39\x93\xb8\xa6\x84\x34\x41\x51\xae\x95\x88\x35\x5b\xd4\x7c\x06\xbd\xac\x3c\x1f\xc8\x0b\xf5\x3d\x96\x0e\x05\x2e\x7b\x39\xef\xa9\x7a\x8c\xff\xf5\x60\xfd\xc1\x9f\x10\xc3\x8c\xeb\x7d\x0a\xf6\x27\xe0\xd5\x2c\xe3\x9a\x36\xbc\x48\x34\xe2\x42\x29\x92\xba\xd4\x22\xa5\x5c\xf3\xf4\xf6\xb5\xfc\x9f\xca\xb4\x60\x69\xec\x7b\xc9\xcf\x1e\xcd\x4b\xcc\x4c\x73\x9c\x3b\x9f\xa3\x8c\xcb\xab\x71\x52\x80\x90\x39\x29\x0d\x21\x75\x76\x64\xc1\x8d\xfb\x8b\x4d\x7b\x8a\x79\x5c\x31\x6f\xda\x11\xee\x78\x52\x50\x8e\x70\xc5\xab\xa6\x6e\xeb\xc1\xf5\xc6\x75\xfd\x8a\x61\x65\xc6\x1f\x2b\x06\x99\x7d\x0e\xa3\xfd\x67\x52\x5e\x96\xa6\x42\xaf\x83\xdf\x01\x00\x00\xff\xff\x7d\x69\x12\xb3\x36\x07\x00\x00")

func _002_authorizationUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__002_authorizationUpSql,
		"002_authorization.up.sql",
	)
}

func _002_authorizationUpSql() (*asset, error) {
	bytes, err := _002_authorizationUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "002_authorization.up.sql", size: 1846, mode: os.FileMode(420), modTime: time.Unix(1612279495, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __003_audit_logDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\xcb\xa9\x88\xcf\x4d\xcc\x4b\x4c\x4f\xcd\x4d\xcd\x2b\xd1\x4b\x2c\x4d\xc9\x2c\x89\xcf\xc9\x4f\x57\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x0b\x0e\x09\x72\xf4\xf4\x0b\x51\x48\xcb\x8e\x2f\x2d\x4e\x2d\xb2\xe6\xe2\x02\xcb\xe0\xd7\x6d\xcd\xc5\x05\x08\x00\x00\xff\xff\xa3\x24\xde\x0e\x65\x00\x00\x00")

func _003_audit_logDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__003_audit_logDownSql,
		"003_audit_log.down.sql",
	)
}

func _003_audit_logDownSql() (*asset, error) {
	bytes, err := _003_audit_logDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "003_audit_log.down.sql", size: 101, mode: os.FileMode(420), modTime: time.Unix(1613407368, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __003_audit_logUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x41\x6f\xd4\x3e\x10\xc5\xcf\xf5\xa7\x98\x5b\xb3\xd2\x5f\x7f\x21\x24\x4e\x15\x87\x34\xeb\x6e\x43\x83\x83\x1c\x83\xe8\xc9\xb2\x92\x59\xd7\x6c\x62\x2f\xb1\x53\x4a\x3f\x3d\x72\x9a\xc2\x06\xb2\x5b\x38\x8e\xe7\xf7\x9e\xed\x99\x97\x71\x9a\x0a\x0a\x22\xbd\x2c\x28\xd8\xf6\x41\x76\xca\x2a\x8d\x1d\xda\xf0\xbf\x1a\x1a\x13\x64\xeb\xb4\x4f\x08\xc0\xcf\x4a\x9a\x06\x2e\xf3\x4d\x45\x79\x9e\x16\xf0\x81\xe7\xef\x53\x7e\x0b\x37\xf4\xf6\x3f\x02\x30\x78\xec\x23\x90\x33\x41\x37\x94\x03\x2b\x05\xb0\x8f\x45\x11\x7b\xaa\x0e\xc6\x59\x19\xbe\xef\x11\x3e\xa5\x3c\xbb\x4e\x79\xf2\xfa\xcd\xab\xd5\x0c\x1a\x0d\x94\x46\x1b\x8e\x33\xae\xd7\xca\x9a\x47\x15\xed\x66\x54\x6c\x7a\xec\xef\x4d\x8d\x7f\x9c\x37\x2a\x28\x78\x57\x95\x2c\x16\x75\x8f\x2a\x60\x23\x55\x80\x60\x3a\xf4\x41\x75\x7b\xf8\x66\xc2\xdd\x58\xc2\xa3\xb3\x38\xbb\x91\x00\x64\x25\xab\x04\x4f\x73\x26\x60\xbb\x93\xf1\x99\x04\x00\xe0\xaa\xe4\x34\xdf\xb0\xf8\x7f\x48\xa6\xdf\xaf\xc6\x0e\xa7\x57\x94\x53\x96\xd1\xea\xf7\xc1\x46\xcc\xcf\xe1\x92\xc1\x9a\x16\x54\x50\xe0\xb4\x12\x3c\xcf\x04\x59\x5d\x10\x32\xad\x27\x67\x6b\xfa\x79\xb6\x81\x07\x39\x1b\x42\xc9\x8e\x2f\x0f\x92\x43\x74\x75\x71\xca\xf4\x79\x78\xa7\xfd\x26\x2a\x3e\x30\x2d\x04\xe5\x2f\xc5\x07\x48\xba\x5e\x1f\x0e\xf0\xd7\xa5\xf5\x4e\x1e\xc4\x82\x64\xd7\x34\xbb\x81\x84\x9c\xcd\xc2\xf2\x16\xce\x5b\xa7\x8d\x95\x7e\xa8\x6b\xf4\xfe\x9c\x9c\x95\x7c\x99\xd8\x2a\xd3\x1e\x6b\xbb\x21\x9c\x74\x30\xb6\x76\x9d\xb1\x5a\xaa\x91\x91\x3d\x7e\x1d\xd0\x87\xb1\xdc\x87\x7f\x92\xf4\xf8\x05\xeb\x65\xc9\x44\xea\x5e\xd9\xc8\xdd\xbb\x1d\x2e\x72\x6e\x08\xda\x2d\x58\x3f\x45\x77\x51\x32\xed\xe5\x6f\x90\x61\xdf\xbc\x84\x34\xd8\xe2\x11\xe4\x30\x50\xd2\x63\x08\xc6\x6a\x7f\xca\x73\x26\x30\xd6\x1b\x7d\x17\x64\xed\xec\xd6\xe8\xa1\x7f\x3a\x7d\x56\xc7\xdc\xff\x08\x00\x00\xff\xff\x5a\xd6\x7e\x3c\x96\x04\x00\x00")

func _003_audit_logUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__003_audit_logUpSql,
		"003_audit_log.up.sql",
	)
}

func _003_audit_logUpSql() (*asset, error) {
	bytes, err := _003_audit_logUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "003_audit_log.up.sql", size: 1174, mode: os.FileMode(420), modTime: time.Unix(1613407368, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"001_initial_schema.down.sql": _001_initial_schemaDownSql,
	"001_initial_schema.up.sql":   _001_initial_schemaUpSql,
	"002_authorization.down.sql":  _002_authorizationDownSql,
	"002_authorization.up.sql":    _002_authorizationUpSql,
	"003_audit_log.down.sql":      _003_audit_logDownSql,
	"003_audit_log.up.sql":        _003_audit_logUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"001_initial_schema.down.sql": &bintree{_001_initial_schemaDownSql, map[string]*bintree{}},
	"001_initial_schema.up.sql":   &bintree{_001_initial_schemaUpSql, map[string]*bintree{}},
	"002_authorization.down.sql":  &bintree{_002_authorizationDownSql, map[string]*bintree{}},
	"002_authorization.up.sql":    &bintree{_002_authorizationUpSql, map[string]*bintree{}},
	"003_audit_log.down.sql":      &bintree{_003_audit_logDownSql, map[string]*bintree{}},
	"003_audit_log.up.sql":        &bintree{_003_audit_logUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
