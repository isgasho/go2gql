// Code generated by go-bindata.
// sources:
// generator/plugins/graphql/templates/output_fields.gohtml
// generator/plugins/graphql/templates/output_map_fields.gohtml
// generator/plugins/graphql/templates/schemas_body.gohtml
// generator/plugins/graphql/templates/schemas_head.gohtml
// generator/plugins/graphql/templates/types_body.gohtml
// generator/plugins/graphql/templates/types_head.gohtml
// generator/plugins/graphql/templates/types_service.gohtml
// DO NOT EDIT!

package graphql

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

var _templatesOutput_fieldsGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\x4d\x8f\xd3\x30\x10\xbd\xef\xaf\x18\x45\x15\x4a\xaa\x36\x91\x38\x46\xea\x01\x15\xd8\x5b\x77\x59\xca\xde\x5d\x67\xe2\x9a\xba\x76\x6a\x3b\x14\x34\xf2\x7f\x47\x76\xbf\x69\x59\x2d\x48\x3b\x87\x48\x99\x79\xf3\xf5\x9e\x87\x68\x0c\xd5\x50\x18\xff\xab\xc3\x1a\x84\xf4\xcb\x7e\x51\x72\xb3\xae\x3e\xdd\xcf\xc7\xdf\x56\x96\x49\x8d\x95\x30\xef\xc5\x46\x55\x02\x35\x5a\xe6\x8d\xad\x3a\xd5\x0b\xa9\x5d\x25\x2c\xeb\x96\x1b\x55\x3e\xa1\x6e\xd0\x7e\x96\xa8\x1a\x37\x35\xda\xe3\x4f\x3f\xac\x60\x1c\xc2\x1d\x11\x58\xa6\x05\xc2\xa0\x8d\x51\xa8\x27\x30\x28\x1f\x7a\xdf\xf5\xfe\x61\xf1\x1d\xb9\x2f\x77\x59\x09\x0c\x00\x40\xf4\x47\xfc\x99\x59\xc9\x16\x0a\x67\x6c\x8d\x21\x94\x1f\x9a\x26\x65\x4c\x8d\x6e\xa5\xc8\x33\xa2\x5d\xe5\x72\x17\xcf\x46\xf0\x8e\x48\x6c\xd4\xe3\x4a\x84\xb0\x2b\x4e\xa9\x70\xb4\x88\xa9\xe1\x3a\xe7\x08\xf8\x88\x8e\x5b\xd9\x79\x69\x74\x0d\x47\xd8\x97\xde\x78\x6c\xa6\x66\xbd\x46\xed\x43\x38\xc1\xe7\x89\x35\x22\xce\x94\xda\x6f\x58\x46\x5f\xdc\x31\x4d\xbf\x27\xe3\x3c\xe7\x09\x9d\x51\x3f\xb0\x86\xb6\xd7\x3c\xef\xe0\x6c\xda\x7d\xe8\x91\x59\xb6\x76\x05\xe4\x52\x7b\xb4\x2d\xe3\x48\x61\x04\x68\xad\xb1\x05\x9c\x96\x89\xe6\xb6\xd2\xf3\x25\x38\xcb\x23\xb3\x5d\xf9\xd5\xf4\x96\x63\x99\x47\x39\x8b\x4b\x68\x34\xce\x1c\xc2\x90\x48\x98\xc3\x94\xe7\x4c\xdf\x27\x6f\x08\xf5\x55\x5e\x34\xd9\xa6\x36\x93\x09\x68\xa9\xe0\xba\xf6\xc1\x2c\xfa\xde\xea\x08\x1a\xc5\xcf\x4d\x5c\xb8\xe9\x75\x71\x87\xa1\xb3\xfc\x66\x94\x48\xb6\x07\x8e\x67\x88\xcd\x94\x39\x7f\x7c\x35\x2f\x0c\x72\x54\x31\x26\xcc\x4d\x08\xf9\xa5\x5e\xcf\x4c\xf5\x08\x99\xcb\xae\x45\x2b\xfe\xbe\x01\x11\x2a\x87\xaf\xeb\xff\xca\x6e\x2f\x35\x03\xd4\xcd\xcd\x66\x49\xd1\xff\x13\xf4\xad\x09\xb5\x3c\x83\xc1\x9b\x92\x98\x3a\xfc\x0b\x8d\xb7\x58\xbc\xfc\x3b\x7f\xbd\x44\xe9\xe6\xdc\xee\x38\x67\xb8\xcd\x33\x97\xee\x0b\x4c\x0b\xbd\x5e\x69\xb3\xd5\x10\x2f\x2d\x2b\x8e\x35\xf6\x87\x1e\x8a\xbb\x93\x68\xbf\x03\x00\x00\xff\xff\x2f\x54\x41\x15\x64\x05\x00\x00")

func templatesOutput_fieldsGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesOutput_fieldsGohtml,
		"templates/output_fields.gohtml",
	)
}

func templatesOutput_fieldsGohtml() (*asset, error) {
	bytes, err := templatesOutput_fieldsGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/output_fields.gohtml", size: 1380, mode: os.FileMode(436), modTime: time.Unix(1547025125, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOutput_map_fieldsGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x54\xb1\x6e\xdb\x30\x10\xdd\xf3\x15\x07\x21\x28\x64\xc3\x96\x80\x8e\x02\x3c\x14\x41\x9b\xa9\x49\x90\xa6\x59\x82\x0c\x34\x75\xa2\x59\xd3\xa4\x7c\xa4\xec\x1a\x04\xff\xbd\x20\xe5\xd4\x4e\x2d\x07\x68\x87\x4e\xb9\x41\x03\xef\xdd\xf1\xdd\x3d\xea\x79\x3f\x85\x72\x2c\x8c\xdb\xb5\x58\x81\x90\x6e\xd1\xcd\x0b\x6e\x56\xe5\xe7\xeb\x87\xe9\xf7\x25\x31\xa9\xb1\x14\xe6\xa3\x58\xab\x52\xa0\x46\x62\xce\x50\xd9\xaa\x4e\x48\x6d\x4b\x41\xac\x5d\xac\x55\x71\x8f\xba\x46\xfa\x22\x51\xd5\xf6\xca\x68\x87\x3f\xdd\xb8\x84\x69\x08\x17\xde\x03\x31\x2d\x10\x2e\x9b\x98\x85\x6a\x06\x97\xc5\x6d\xe7\xda\xce\xdd\xce\x7f\x20\x77\xc5\x57\xd6\xf6\x85\x09\x0f\x00\xe0\xfd\x1f\x90\x47\x46\x92\xcd\x15\xde\xb0\x15\x86\x50\x7c\xaa\xeb\x54\x71\x65\x74\x23\x45\x9e\x79\xdf\x37\x2f\xfa\x7c\x36\x81\x0f\xde\x8b\xb5\xba\x5b\x8a\x10\x8a\x04\xf5\xa9\x71\x8c\x88\xa9\xe0\xb4\xe6\x37\xe0\x21\x6d\xc2\x7b\xce\x94\xda\xb3\x2e\xe2\x59\xe4\x9d\xe8\xec\x07\x0c\xe1\x50\x73\x8f\xd6\xa8\x0d\x56\xd0\x74\x9a\xe7\x2d\x1c\x5d\xbf\x4f\xdd\x31\x62\x2b\x3b\x82\x5c\x6a\x87\xd4\x30\x8e\x3e\x4c\x00\x89\x0c\x8d\xe0\xc0\x2e\x86\xdd\x4a\xc7\x17\x60\x89\xc7\x6d\xb5\xc5\x37\xd3\x11\xc7\x22\x8f\x12\x8d\x5e\x43\x63\x70\x66\x11\xc6\xde\x0b\xf3\xc2\xf2\x78\x75\xd7\xe9\x34\x84\xea\xa4\x2e\x86\x6c\xd2\x35\xb3\x19\x68\xa9\xe0\xb4\xf7\x4b\x10\xba\x8e\x74\x04\x4d\xe2\x67\x10\x17\x06\x4f\x6d\x9c\x61\x6c\x89\x0f\x66\x37\x8c\x80\xd0\xc2\xd3\xf3\x8a\xb5\x4f\xd6\x91\xd4\xe2\xf9\x68\x43\x83\x45\x8d\x21\x58\xe2\x6e\x02\x1b\xa6\x3a\x8c\xfd\xfb\x17\xf6\x5a\xb2\xc7\x94\xcc\x6c\x76\xaa\xdb\x9b\x83\x5a\x98\x01\x6b\x5b\xd4\x75\x4e\x68\x27\x30\x4c\xec\x7c\x87\x18\xd9\x12\x77\x59\x05\x90\x68\xbe\x8d\x4c\x33\x64\x55\x3f\xcb\x79\x6c\x18\xfd\xc5\xce\xf7\x6a\x25\xfa\x43\x6a\xa5\x17\xf3\x6f\x0f\xe6\x3f\x08\x46\xfc\x5d\xb2\x8b\xf3\x15\xc7\xbf\xa2\xf7\xc9\x40\x6c\xef\x34\x37\xb8\xcd\x33\x9b\xcc\x02\x4c\x03\x9d\x5e\x6a\xb3\xd5\x10\x6d\x23\x3b\x70\xd9\xbb\x56\x18\x45\x6b\x46\x5d\x47\xd7\xfd\x15\x00\x00\xff\xff\x14\x07\xb2\x2e\x05\x06\x00\x00")

func templatesOutput_map_fieldsGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesOutput_map_fieldsGohtml,
		"templates/output_map_fields.gohtml",
	)
}

func templatesOutput_map_fieldsGohtml() (*asset, error) {
	bytes, err := templatesOutput_map_fieldsGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/output_map_fields.gohtml", size: 1541, mode: os.FileMode(436), modTime: time.Unix(1547025125, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSchemas_bodyGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\x4b\x6f\x1a\x31\x10\x3e\xef\xfe\x8a\xd1\x0a\xb5\x4b\x44\x16\xa9\x47\x24\x2e\xa5\x69\xda\x43\xf3\x68\xe8\xa9\xaa\x2a\x63\x86\xc5\x8d\xb1\x89\xd7\x9b\x34\xb2\xfc\xdf\x2b\x3f\xf6\x01\x04\x92\x5b\x39\x99\x99\xf1\x7c\xdf\xcc\x37\x9e\x35\xe6\x1c\xc6\x67\xa5\xd4\xcf\x5b\x9c\x40\xc9\xf4\xba\x5e\x14\x54\x6e\xc6\x17\x97\xf3\xf3\x1f\xf7\x8a\x30\x81\xe3\x52\x7e\x28\x1f\xf8\xb8\x44\x81\x8a\x68\xa9\xc6\x5b\x5e\x97\x4c\x54\xe3\x52\x91\xed\xfa\x81\x17\x77\x74\x8d\x1b\xf2\x51\x2e\x9f\x67\x52\x68\xfc\xab\xcf\xc6\x70\x6e\x6d\xea\xb2\x82\x31\x83\x18\x70\x45\x36\x68\x6d\x38\xcf\x38\x43\xa1\x2b\xa8\xb4\xaa\xa9\x06\x93\x26\xc6\x80\x22\xa2\x44\x18\x54\xa8\x1e\x19\x45\x98\x4c\x61\x50\xdc\x85\x3f\x95\x4f\x98\x24\xc6\x34\xee\x22\xa4\x0b\x89\xc0\x98\x52\xce\x1d\x5c\xeb\x0e\x8e\x4b\x6f\x75\x57\x8d\x01\x14\x4b\x9f\xc6\xa6\xe9\x5b\xd0\x56\xb5\xa0\x90\xd3\xd3\x15\x0c\xe1\x12\xf5\x11\x56\xf9\xf0\x55\x5e\xae\x72\x85\xba\x56\x02\x68\x71\x24\x4d\x6a\xd3\x1e\xf9\x40\x2b\x80\xbe\x40\x2b\xa7\xbc\x3a\xcd\x78\x04\x6c\x0d\x67\xc6\x30\xa1\x51\x51\xdc\x6a\xa9\xaa\x9b\xfb\xd2\xda\xe2\x6b\x67\xf9\x42\xc4\x92\xa3\x72\xf3\xc1\x56\x30\x28\xe6\x8a\x50\x54\x17\x82\x2c\x38\x7a\x1e\x23\xd0\x0a\x8c\x91\x5b\x14\x5a\x11\xca\x44\x19\x72\x84\x40\x77\x2f\x12\x1e\x42\x6e\x4c\xf9\xc0\x83\x3b\x30\x19\x01\x2a\x25\xd5\xf0\xed\xba\xb3\x15\x50\x5e\x1d\xeb\x10\x4c\xa7\x20\x18\x77\xe9\x92\xa6\x9d\x07\xa0\xc6\x8e\xc0\x18\x0f\x1c\xeb\xbd\x70\xe7\x55\x9e\x45\x2c\xa0\xcd\x2c\xed\x81\x00\x25\xe2\xbd\x86\x05\x7a\x10\xc1\x78\x36\x4c\x93\x64\x77\xa6\xde\x58\xc7\x23\x51\x87\xf9\x6f\x6b\x54\xcf\x9f\x19\xf2\x65\x05\x53\x30\x26\x7a\x67\x52\x84\xe7\x21\x15\x64\x3e\x26\xeb\x92\x0f\xac\xcd\x4f\x74\xc4\x89\x7c\x5a\xbc\xbe\x44\x91\xd8\x6f\x8f\x7e\x9c\xdc\x31\xfe\xdf\x6a\x4d\x34\x93\xe2\x95\x12\x9a\xb0\xff\x51\xc5\x2e\xc5\x63\xca\xc9\xc5\x1f\xa4\x3a\x08\x77\xed\xcf\x7b\xba\x85\x80\x66\x2c\xa6\xfd\x21\xbb\xc2\xa7\x70\xa5\x3f\xee\xc1\x32\x93\x62\xc5\x4a\x3f\x9c\xee\xe6\x04\xb2\xfd\x54\xd9\xc8\x39\x8d\x71\x95\x8a\x96\x48\x71\x5b\x4b\x8d\xcb\x99\xdc\x6c\xdc\x5c\x66\x59\x24\x93\x24\x9f\xb0\xa2\x8a\x6d\x5d\x45\x93\x1e\xaf\x9d\x78\x6b\x9b\xa4\x6d\xa5\x49\x12\xea\x9f\xf4\x89\x07\x93\x67\x17\x19\x70\x14\x2d\x85\xa8\x69\x03\xdc\x6b\xd6\x8a\x2f\x7d\xa7\x8e\x04\xc6\x5c\x2e\xac\x79\x04\x3d\x67\xe2\x5a\xe0\x5c\xb1\x7e\x5f\x45\x2f\x34\xda\xbb\xd2\xdc\xba\x0c\xb6\x00\xf4\x73\x2f\xc1\xaf\x51\x07\x8b\xbc\x3a\x8d\xf5\x6e\xbf\x7c\xd3\x84\xf6\x04\xea\xdd\x18\xb5\xee\xb9\xff\x50\x46\xef\x75\x5f\xc2\x2e\x66\x5f\x1c\x17\xfa\x92\x32\xfe\xf7\x1d\x2b\xc9\x1f\x71\x02\x6e\xad\xe7\xdb\xbe\x30\xd1\x75\x43\x14\xd9\x54\x43\xc8\xfd\xc2\x5e\x11\x8a\x6e\x93\x75\x0b\xb4\xf9\xc5\xbd\x17\xde\x9b\xb1\x2e\x48\x30\xde\xfa\x3b\x4c\xdb\x6f\x55\x37\x1a\x07\xff\x0f\x3a\x99\x09\x19\x9a\x7f\xaa\x85\xb1\x81\x6d\x68\x03\xd6\x74\xae\xdb\xca\x5a\x31\x51\x46\x77\xe4\xb4\xcb\xc0\x1b\xdd\xbb\xee\x9b\x0f\xd7\xfb\x15\x3e\xc5\x6f\xdf\xc1\xce\xef\x5e\x9e\x5f\x63\x5e\x8e\xc2\x1f\x83\x76\x41\x89\x38\xa9\x45\xb3\x24\x82\xaf\x61\xd1\x58\xc3\xe5\xdd\x98\xf6\x7e\x4b\xcf\x0e\x53\xfb\x2f\x00\x00\xff\xff\xb8\x44\x15\x75\x56\x09\x00\x00")

func templatesSchemas_bodyGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesSchemas_bodyGohtml,
		"templates/schemas_body.gohtml",
	)
}

func templatesSchemas_bodyGohtml() (*asset, error) {
	bytes, err := templatesSchemas_bodyGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/schemas_body.gohtml", size: 2390, mode: os.FileMode(436), modTime: time.Unix(1547025125, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSchemas_headGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xc1\x4a\xc4\x30\x14\x85\xe1\xb5\x79\x8a\xc3\xd0\x85\x2e\x26\x01\x97\x82\x0b\x61\x06\xe9\xc6\xba\x88\x0f\x70\xdb\x5e\xd3\xd0\x36\xad\x49\x44\xe4\x72\xdf\x5d\xd0\x82\xbb\x9f\xef\x1c\xe7\xe0\xa7\x58\xf0\x1e\x17\xc6\x17\x15\x04\x4e\x9c\xa9\xf2\x88\xfe\x1b\x21\xd6\xe9\xb3\xb7\xc3\xb6\xba\xeb\xb3\x3f\xbf\xcd\x99\x62\x62\x17\xb6\xfb\xf0\xb1\x58\x5c\x3a\xbc\x74\x1e\xd7\x4b\xeb\xd1\x7a\xb3\xd3\x30\x53\x60\x88\x34\xf6\x68\x55\x63\xe2\xba\x6f\xb9\xe2\xd6\x88\x64\x4a\x81\xd1\x1c\xf2\xf0\x88\xc6\xfe\x75\xc1\x59\xd5\xdc\x88\x1c\x9b\x7d\x5a\x22\x15\x55\x9c\xfe\xe9\x95\xea\xa4\x7a\x32\x22\x9c\xc6\xdf\xff\x9d\xf9\x09\x00\x00\xff\xff\xc2\xe3\x00\xe2\xbf\x00\x00\x00")

func templatesSchemas_headGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesSchemas_headGohtml,
		"templates/schemas_head.gohtml",
	)
}

func templatesSchemas_headGohtml() (*asset, error) {
	bytes, err := templatesSchemas_headGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/schemas_head.gohtml", size: 191, mode: os.FileMode(436), modTime: time.Unix(1539090323, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTypes_bodyGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\xdb\x6f\xdb\x36\x17\x7f\x96\xfe\x8a\xf3\xe9\xcb\x17\xc8\x85\x2a\x7f\xd8\xa3\x07\x3f\x64\x5d\x53\x14\x5d\x9b\xb6\xeb\xda\x87\xd4\x48\x58\x9b\x56\x38\xc9\x94\x42\x5d\x9a\x80\xd0\xff\x3e\xf0\x26\x51\x17\xdb\x4a\xd7\x0e\xdb\x50\x3d\x04\x16\x75\x6e\xfc\x9d\x1f\x0f\xc9\x13\xce\x1f\xc3\xfc\x51\x94\x16\xf7\x19\x5e\x40\x44\x8a\x9b\xf2\x53\xb8\x4e\x77\xf3\xa7\xcf\xde\x3d\xfe\x2d\x66\x88\x50\x3c\x8f\xd2\x1f\xa2\xdb\x64\x1e\x61\x8a\x19\x2a\x52\x36\xcf\x92\x32\x22\x34\x9f\x47\x0c\x65\x37\xb7\x49\xf8\x53\xba\xb9\x7f\x92\xd2\x02\xdf\x15\x8f\xe6\xf0\xb8\xae\xdd\xf9\x1c\x9e\xd2\x72\x97\xbb\x9c\x33\x44\x23\x0c\x27\x98\x96\x3b\x58\x2c\x21\x3c\x27\x09\x0e\xe5\x47\x29\xe9\x54\x88\x01\xe7\xf2\x7b\xf8\x1e\x31\x82\x3e\x25\xf8\x15\xda\xe1\xba\x86\x25\x70\x1e\xdd\x26\xaf\xe3\xa8\xae\xc3\x57\xf8\xb3\xd0\xf2\xad\x21\xf1\xfe\x24\xa5\x5b\x12\x71\xd7\x71\x84\xd2\x02\xf4\xe3\x19\x93\xcf\x44\x8c\x6f\x7e\x51\x16\xbd\xc0\x75\x1c\xce\x81\x6c\x55\x40\xe1\x93\x74\xb7\xc3\xb4\x50\x91\x38\xce\xcf\x38\x5f\x33\x92\x15\x24\xa5\x8b\x26\x28\x2d\x53\xd7\x5a\x19\xd3\x8d\x96\x7f\x8f\x92\x12\xe7\x0b\xe8\x85\x24\x87\x55\x5c\x2f\x51\xc6\x5d\xb0\x9e\x06\x8f\x4a\x08\x09\x40\xcc\xcc\x85\x29\x13\x87\x23\xc2\x97\x12\xa1\x0e\x7c\x01\xa7\xfb\xbd\x70\xa9\xa3\xc2\x91\x71\x2b\x55\xf9\xae\xc2\x76\xcc\xb4\xa9\x71\xdd\x4c\xfd\xda\xf3\xae\x1b\xbf\x43\x08\x3a\xc2\xb6\xb1\x16\x06\xc7\x51\xc3\x9c\xb7\x63\x62\xa4\x9e\xb9\xed\x90\x2b\x38\xf1\x9c\x66\x65\x01\xe9\xa7\xdf\xf1\xba\x70\x39\x07\x8d\x85\x1a\x68\xd9\x21\xc5\x2e\xe4\x60\x97\x24\x4a\xf0\x18\x4d\x2c\x75\x9b\x2d\xd6\xb0\x81\xcd\x64\x45\x51\xc7\x6b\x3d\xf4\x59\x63\x04\xcf\x09\x4e\x36\xdd\x8c\x0f\xcc\x4a\x19\x91\xf8\x5a\xe9\xd5\x33\xd7\xd9\x96\x74\x0d\x84\x92\xc2\x9f\x71\xc9\x22\x3d\xf3\xad\x90\x95\x2c\xd0\x8e\x95\x03\x03\xec\xbe\x29\x87\x67\x9b\x8d\x94\x54\x1e\x7d\x11\xb9\x34\x65\xe8\x12\x74\xe8\x62\x85\x68\x69\xf1\x77\x72\xd1\x73\xbe\x46\x49\xa2\x43\x09\xc5\x18\x9c\xd4\x75\x00\x7d\x22\xa8\xef\x6f\xca\xb4\xc0\x9b\x86\x0e\x62\x6e\x76\xde\x6b\xd7\x62\x46\x3f\xe3\x39\x30\x9c\xa7\x49\x85\x59\x6e\x25\xdf\x8c\x8d\xa6\xff\xad\x51\x90\x06\x25\x8a\x9c\x37\x2a\xe1\x79\x49\xd7\x22\x40\x35\x6b\x7f\x5d\xdc\x89\xe9\x14\x77\x6a\xda\xba\x24\x05\x40\x80\xd0\x02\xb3\x2d\x5a\x63\x5e\xcf\xc0\xbf\x12\xf9\x4b\xd5\x54\x1b\x5b\x17\x65\x91\x95\xc5\x33\x39\x2c\xe6\xcf\x30\x63\x80\x19\x4b\x99\x48\x19\xd9\x02\x81\xe5\x12\x28\x49\x40\x64\x90\xe1\xa2\x64\x54\xbc\x06\xe2\x8f\x98\xb9\x83\x58\x94\x07\x70\x25\x26\x42\x42\x7f\x87\xb2\xcb\xbc\x60\x84\x46\x2b\xdb\xbb\xeb\x5c\xc1\x12\x84\xa8\x62\x35\xc3\x79\x99\x14\xb0\x04\x8a\x3f\xfb\x26\xac\xf3\x94\xbd\xc2\x9f\xf7\x06\x37\x73\x9d\x16\xbf\x96\x42\x2d\x2c\x36\x89\x74\xb9\x53\xe9\x7b\x9e\x9f\xb3\x74\x77\xc6\xa2\x86\x62\x64\x2b\x83\xb9\x6c\x29\xa4\xb9\x2f\xb3\x20\x2d\x69\x4e\xad\xe0\x3f\xed\xfc\xbb\x05\x40\xf9\x20\xf9\x19\x63\xe8\xde\xf8\x52\xd1\x36\x35\x82\x50\x11\xe4\x54\x5f\xa1\x7f\xd9\x83\xcd\x11\xa0\x0b\xac\xc2\x46\x5b\xa1\x62\xe9\xc1\x12\x76\x28\xc6\x7e\x9b\x5e\x3b\x14\x91\xd5\x04\x53\x9f\xd0\x99\xb2\xb7\x4d\x19\x90\x00\x2a\x94\x88\xd0\x14\x9e\x84\x02\xb7\x0b\xa6\xb6\x60\x88\xf8\x81\x14\x37\x4f\x05\x29\xa0\xa9\x98\x5a\xdc\x71\xaa\x40\xf0\x45\x98\xea\xae\x29\x59\x87\x8d\x01\xf0\x2a\x94\x78\x62\x89\x19\x35\xb2\x95\x6a\x16\xb8\xea\xb1\x29\xc6\xb9\x64\x62\xae\x88\xfd\x81\xa1\xcc\xc7\x8c\x05\xe0\x6d\x11\x49\xf0\x06\x8a\xd4\xac\x2d\xb8\xde\xbf\x40\xae\x81\x58\xcb\x11\x64\x74\xde\xcc\x38\x6c\x02\x3a\x0a\xf3\x25\x59\xc1\x12\x2a\xb7\xa3\x27\x18\x91\xe4\xd8\xda\x4a\x26\xda\x79\x00\x58\xfd\x6d\xc7\x1d\xf1\x7b\x24\x6d\x4d\x78\x06\xe8\x69\x59\xf3\xd7\x29\x5d\xa3\x02\x3c\x49\xe0\x8f\x9e\x07\x87\x18\x0c\xde\x47\x6f\xe5\xcd\xac\x2c\xef\x49\xf2\x5f\x9c\xe3\xc9\x99\x69\xd3\xdb\x80\xda\x4d\xee\x14\x0b\xdf\x0e\xcf\x41\xf5\x69\xdf\x8e\x94\xbb\xba\x57\xbb\xba\xbf\xcd\x79\x84\xe2\x74\xdb\x2d\xa9\x17\x14\x5f\x6c\x7b\x75\x55\x4b\x13\xba\xc1\x77\x41\x67\x2f\x17\xfa\x83\xad\x5c\x50\xe0\x56\x8b\xc3\xff\xad\xfa\x7b\xac\x1a\x5e\x05\x90\xc6\x0f\x29\x9e\x3f\x0a\xf9\xd3\xd3\xe3\x86\x5b\x3e\x42\xef\x99\xb0\x8a\xfa\x2a\xe2\x79\xd8\x62\x3a\x9c\xf2\x2b\x9d\xf0\x31\x3f\xfd\xe5\x34\x26\x23\x9e\x2f\x5c\x5f\xa9\xc8\x76\x67\x09\x1d\xc5\xd2\x9b\x8d\x06\x31\x0c\xdf\x5e\x4a\xa3\x10\x7e\x5b\xf4\xda\xc3\xda\xc1\x94\x9f\xe5\x39\x89\x28\xa1\x91\xc0\x29\xc3\xfb\x33\xde\x16\x02\xc5\xfa\xe3\x85\x60\x60\xda\xab\xbc\x3d\xa1\x1e\x46\x6a\x8a\xeb\x6a\xd4\x6a\x5b\x3b\x6a\xce\x8d\x0f\xe5\xec\xfb\x7a\xfc\xbe\x1e\xbb\x10\x7e\x5f\x8f\x13\x91\xfa\x5a\xeb\xd1\x5c\x1f\xf5\xcd\x52\x6d\xcd\xea\x87\x6b\x4e\x4c\xca\x97\xba\x7b\x75\x2e\x9d\xf3\x39\x28\xbf\xe6\xd2\x39\xda\x67\x38\x51\x37\x4d\x25\xf9\xe5\x9d\x86\x61\x93\xa1\xdb\x5f\x30\x4d\xa9\x03\x9d\x05\xc7\x19\xe9\x29\xa8\x21\xae\x9b\x28\xc3\xee\x41\x33\xa1\x16\xe1\xb7\x98\x6e\x30\x53\xd7\xe8\x13\x1b\x7a\xf3\xa1\xd3\x54\x18\x7e\x0e\xd5\x0f\x7d\x68\x31\x50\x9d\xd4\xfd\x03\x53\x1f\xec\x97\x28\xcb\x3b\x47\xce\x71\xc0\x15\xde\x2f\x51\xf6\xef\xe8\xed\xbc\xbb\x29\x69\xec\x8b\xbc\xf8\xb3\x49\x0a\xbd\x9a\xa8\x59\x3c\xa9\x8d\xa4\x1a\x82\x31\xbe\xef\x35\x01\x27\x75\x75\xf4\x9c\x5f\xe0\x7b\x25\x69\xfa\x3b\xba\x37\xe5\x98\x80\x3c\xd9\xe7\xfb\x13\x1e\x64\x4d\x1c\xf5\x61\x9e\xb6\x68\xd4\xb3\xa6\x33\x76\x8c\x4c\x13\x5b\x46\x5d\x62\x75\xbb\x46\xce\xd7\x6b\x1b\xed\x50\x76\x39\xd2\x3a\x7a\x81\xef\x4d\x87\x61\x35\xf2\x59\x42\xb3\xaf\xb1\x34\xe8\x2c\x0d\x5b\x4b\xf2\xd2\xa2\xda\x27\x64\xa4\x2f\xa2\x3b\x48\x0b\xdd\xfb\xf8\x0a\x31\x0a\xab\xba\x37\x42\x46\x9b\x23\x57\xb0\x04\x22\x7e\xe8\xaf\x42\xea\x40\xc7\xcb\x71\xe2\x40\xed\xa3\x15\x4a\x2e\x25\x95\x57\x81\xfa\xad\x68\xb7\x92\x46\x03\x10\x76\xe3\xa0\x6a\xee\x4b\x9d\xf0\xf7\xdf\xe5\xe3\x78\x70\xcc\x19\x53\x04\x2f\x6e\x7b\x09\xa3\x17\xf2\xc3\xc7\x93\xed\xde\xf3\xc9\x7f\xff\xb7\x11\xdc\x00\x9c\x60\xd9\x53\x17\x33\x0c\x80\xcc\xec\x2e\x45\xe7\xfe\x1c\xc7\x0f\x09\xd6\x6e\xac\xf7\x81\xe9\x9c\x45\x86\xd0\x54\xc3\x13\xe0\xb8\x6a\xb3\xeb\x7f\x73\x70\x54\xca\x0f\xc2\x53\x55\x0f\x0b\xb8\x03\x90\x5a\x0f\x97\x71\xbc\x5a\x56\x95\x5e\x3d\x23\xa7\x86\x61\xaf\x5a\x56\x9f\xf4\xf8\xe1\xa1\x29\x39\x7f\x83\xe3\x83\xae\xa8\x07\x0e\x11\xed\x3f\x63\x86\xe7\x88\x89\xff\x5d\x50\x74\x3e\xed\xdb\x96\xbc\xd0\x11\x4a\x11\xf1\x3e\x6d\xfb\x91\xa2\x3a\x93\x0b\x90\xfb\x68\x66\xc7\xae\x3f\xbd\x46\x0c\xed\xf2\x19\xf8\x56\x3d\x09\x74\xed\xd4\xb4\xcc\xd9\x5a\xe4\x24\x0b\x7f\x4d\x4b\xb6\xc6\x07\xab\x90\xe0\xb5\x90\x5f\xee\xe5\xb5\xaa\xb7\xa6\xeb\xa5\xbf\xe4\x6c\xdd\x14\x2d\x2d\x20\x27\xa0\xff\xf9\x31\x11\x44\x43\xfb\x83\x30\x6a\xa1\xbd\x40\x8e\xec\xb2\xff\x3c\x28\x75\xcd\x1f\x82\x59\x5b\xff\x2f\xfc\x23\x00\x00\xff\xff\x2c\x78\xbe\xc9\x9a\x1e\x00\x00")

func templatesTypes_bodyGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTypes_bodyGohtml,
		"templates/types_body.gohtml",
	)
}

func templatesTypes_bodyGohtml() (*asset, error) {
	bytes, err := templatesTypes_bodyGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/types_body.gohtml", size: 7834, mode: os.FileMode(436), modTime: time.Unix(1548415500, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTypes_headGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xc1\x4a\xc4\x30\x14\x85\xe1\xb5\x79\x8a\xc3\xd0\x85\x2e\x26\x01\x97\x82\x0b\x61\x06\xe9\xc6\xba\x88\x0f\x70\xdb\x5e\xd3\xd0\x36\xad\x49\x44\xe4\x72\xdf\x5d\xd0\x82\xbb\x9f\xef\x1c\xe7\xe0\xa7\x58\xf0\x1e\x17\xc6\x17\x15\x04\x4e\x9c\xa9\xf2\x88\xfe\x1b\x21\xd6\xe9\xb3\xb7\xc3\xb6\xba\xeb\xb3\x3f\xbf\xcd\x99\x62\x62\x17\xb6\xfb\xf0\xb1\x58\x5c\x3a\xbc\x74\x1e\xd7\x4b\xeb\xd1\x7a\xb3\xd3\x30\x53\x60\x88\x34\xf6\x68\x55\x63\xe2\xba\x6f\xb9\xe2\xd6\x88\x64\x4a\x81\xd1\x1c\xf2\xf0\x88\xc6\xfe\x75\xc1\x59\xd5\xdc\x88\x1c\x9b\x7d\x5a\x22\x15\x55\x9c\xfe\xe9\x95\xea\xa4\x7a\x32\x22\x9c\xc6\xdf\xff\x9d\xf9\x09\x00\x00\xff\xff\xc2\xe3\x00\xe2\xbf\x00\x00\x00")

func templatesTypes_headGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTypes_headGohtml,
		"templates/types_head.gohtml",
	)
}

func templatesTypes_headGohtml() (*asset, error) {
	bytes, err := templatesTypes_headGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/types_head.gohtml", size: 191, mode: os.FileMode(436), modTime: time.Unix(1539090323, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTypes_serviceGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x4d\x8f\xdb\x36\x10\xbd\xef\xaf\x98\x08\x69\x20\x07\x8a\x0c\xf4\xe8\xc2\x87\xd4\xdd\x4d\x03\xa4\xf9\x5a\x17\x39\x06\x5c\x69\x24\x13\x96\x49\x79\x44\x6d\x76\x41\xe8\xbf\x17\xa4\x28\xaf\x6c\xeb\xcb\x4d\x5a\x84\xb7\x35\xc9\xe1\xcc\xbc\x37\x6f\x46\xab\xf5\x2b\x98\xbf\x4c\xa5\x7a\xcc\x71\x01\x29\x57\x9b\xf2\x2e\x8c\xe4\x6e\x7e\xfd\x66\xfd\xea\xef\x2d\x31\x2e\x70\x9e\xca\x5f\xd3\x7d\x36\x4f\x51\x20\x31\x25\x69\x9e\x67\x65\xca\x45\x31\x4f\x89\xe5\x9b\x7d\x16\xde\x22\xdd\xf3\x08\x57\x52\x28\x7c\x50\x2f\xe7\xf0\xaa\xaa\xae\x92\x52\x44\xf0\x06\x95\xd6\xcd\x7e\xf8\x9e\xed\xb0\xaa\xdc\x5f\x5a\x87\x37\x1c\xb3\x78\xfd\x98\x63\x55\xfd\x85\x6a\x23\xe3\xc2\x8f\x40\xeb\x54\x9a\xdf\xe0\x70\x6d\xc5\xb2\xec\xad\x50\x48\x09\x8b\xb0\xaa\x02\xe0\x1b\x78\xa9\x35\x37\x3f\x45\x98\x2b\x49\xc5\xc7\x6d\x5a\x55\xe1\xdb\xa7\x5f\xfe\x64\x22\xce\x90\x40\x6b\xe0\x09\x3c\x0f\xd7\xc4\x22\xa4\x6b\xc1\xee\x32\x8c\xa1\xaa\x20\x50\x66\x53\xe6\x28\x14\xb1\x88\x8b\xb4\x36\x51\x9f\x03\xad\x51\xc4\x55\x35\x33\xde\xec\xb3\x7a\xcb\x7a\x5b\x80\xbe\x02\x00\x67\xb7\x71\xd1\x79\x6f\xe3\x06\xb7\x08\x55\x49\xe2\xdc\x80\x3e\x9c\xa8\xed\x10\x13\x29\xc2\xf3\x9d\x35\x01\x8b\xe5\xa0\xd1\x66\x79\x5a\xbb\x1b\x2e\xa9\xde\x02\x5e\x9c\x3e\xa5\xcf\xae\x99\x65\xce\x2f\x3a\x0c\x04\x9d\xa7\xff\xc0\x22\x22\x9e\x2b\x2e\xc5\x02\x9e\xee\x7c\x2a\xa5\xc2\x78\x25\x77\x3b\x14\xaa\xaa\xba\xef\xae\x2d\xa9\xb4\x8e\x58\x96\x35\x01\x86\x6f\x0c\x67\x3e\xbd\xfb\x50\xaa\xbc\x54\x16\xe7\xe7\xe1\xef\x32\x7e\x74\xe4\xe9\xb3\xe5\x70\x74\x46\x5e\x53\x5a\x9a\x97\xbb\x93\xd3\xac\xd7\x94\x16\x8b\x33\x00\x56\x52\x24\x3c\x6d\x2c\x74\xe7\xa8\xf5\xaa\x83\x87\x51\x6a\xb0\xb9\xcc\x81\x66\x99\x64\x33\x4a\xbb\xa1\x6a\x4c\xd5\x7e\xe9\xe3\xa4\x99\x5b\x9d\x49\x3a\x03\xc6\x9c\x3c\x41\xa5\x27\x95\xad\xe0\x50\xc4\x83\x01\xf4\x83\x31\x74\xf3\x33\x16\x32\xbb\xc7\x05\x18\x09\xf0\xf3\x36\x02\x6e\xeb\x23\x23\xb6\x2b\x66\xe0\x7f\x05\xde\x14\xb6\xae\x02\x20\x24\x02\x24\x92\x34\x83\x7e\x64\x22\xf5\x60\xb0\xc8\x43\x97\x8e\xde\x83\x5f\x61\x69\x0e\xf7\xee\x77\x8b\xc3\x18\xa2\xf7\x8c\x20\x67\x84\x42\xdd\xe6\x4c\xac\xd4\x43\x97\x8e\xd8\xad\x11\xf7\xcc\xe2\x89\xb3\x65\x22\xea\xb1\x73\x43\x72\xe7\x6c\xf9\x87\xa0\x67\xbf\x35\x17\x9f\x2d\x41\xf0\x6c\x20\x5f\x6e\x1d\xfb\xbc\x74\x7f\x37\xf6\xfc\xd9\xe0\xfd\xe1\x94\x14\x39\x13\x26\x00\x45\xe1\xad\x62\x64\x1f\xf1\x0d\xeb\x4f\xa4\x3f\x3c\x55\x9d\x86\x2b\xe4\x05\x5d\xd1\xaf\x36\x3c\x8b\x3f\x24\xfe\x91\xeb\xb3\x61\x4f\x8b\xa7\xcc\xf7\xe4\xd4\xed\x7e\xe1\x6a\x63\x1d\x3d\xe4\x34\xb0\x97\x87\xcd\xc7\x98\x20\xd9\x73\xe1\x0d\x17\xbc\xd8\x8c\x24\xee\x60\x1c\x96\x6d\xd7\x06\xef\x44\x16\x9f\xa9\xa7\x6b\x8f\x6c\xad\xcd\x46\x49\x60\x08\x67\xab\x6c\x2a\x6b\xec\xb2\xe1\xde\xa2\x5a\xb3\xd4\xf7\x6c\x7d\x7a\x01\x28\x2a\x71\x16\xbe\x93\x69\xdd\xd7\x7c\xad\x33\xe9\x32\x7c\x6d\x8e\xf8\xe6\x9d\x11\xb0\x60\x94\x5a\xd5\x40\x7e\x6d\x97\x1e\xac\x57\x9e\x98\x61\x61\x39\x25\xd6\xe3\x0e\xf3\x19\xf7\x25\x16\xaa\xa1\xe7\x24\x99\x27\xdc\x07\x46\xbd\x6a\xda\x1d\x75\xbd\x53\x73\x5e\x6e\x74\xbf\xf0\x4e\x75\xdd\xcc\x63\x3c\x01\x21\x55\xdf\x55\xc3\xda\x6b\xb2\x0e\x05\xb5\x54\xfa\x82\x67\x33\x37\xb0\x4c\x81\xff\x62\xf4\xdd\x24\x23\x78\x66\x5f\xfc\x4e\x40\xa1\x3d\x1b\x1d\x25\x69\x95\x71\x14\xaa\x1e\x7b\xcc\xcc\x67\xf2\x14\x79\xe0\x11\xee\xcf\x12\x35\xda\xd9\xb2\x02\xa7\x82\x66\xf0\x12\xf8\xcd\x6f\x26\xcf\x1b\x49\xef\xf1\xdb\x29\x00\xf5\x9c\x3a\x4e\xe7\xff\x21\xb6\xb1\xae\xdd\x5f\x0e\xae\x75\xbe\xe8\x9a\x9e\xdd\xfb\xc3\x9c\x70\x7a\x5e\x0f\x90\x27\xea\xde\x33\x42\x36\xab\x0e\x7e\xfa\xe8\xd9\xac\x7a\x58\x58\x40\xde\x7f\xac\x3f\xe2\x76\x4d\xf2\x4d\x33\x7f\x98\xda\xf3\x4d\x32\x82\x5a\x37\x6d\x5e\x3a\x3f\x29\x0e\x9d\x41\x18\x09\xef\x3a\xd1\x32\xf9\x56\xdc\xcb\x2d\xd2\x0c\x7c\xc2\xa2\xcc\xd4\xf1\x70\x33\x69\xb6\x81\x1f\xa4\x44\x5d\x1c\xfc\x29\x44\xe8\xb2\xda\xac\x85\x67\x6a\x6d\x06\x46\xa3\xbe\xbb\x7a\xfa\x4b\xfc\x12\xf1\x6c\x8b\xa6\xd6\x36\x47\x8e\x30\x5f\x88\xe5\x3e\x12\x05\xe0\x25\x8c\x9b\x89\x53\x49\xa0\x3a\xb3\xc0\x0c\x20\xfd\x1e\x0c\x11\xdd\xbe\xc7\x37\xf6\x73\xd9\x91\xdb\xb2\x7f\x2a\xc3\x8d\x12\x1e\x11\xb6\x97\xf2\xe6\x85\xba\x9a\x7f\x10\xe3\x29\x00\xb9\x35\x25\x4a\xb8\x0f\x0f\x58\xff\x1b\x05\xe6\x09\x3c\x93\xdb\x09\x8d\xad\x1f\x9f\xf7\x96\x6e\xc9\x4e\x35\x13\x38\x71\xa1\x12\xdf\xfb\xdc\x82\x08\x5a\x49\x71\xa6\x0a\xb8\x63\xb1\x09\xc0\xb8\x0a\xea\x31\x47\xff\x97\xf5\x2c\x84\xdb\x8d\x2c\xb3\x18\xee\xec\x57\xdd\x50\x60\x9e\x85\x60\x64\x60\x1a\x2e\x9a\x7a\x70\x34\x58\x87\xb5\x6a\x8e\x7e\x22\x81\x2d\x8a\x96\xdc\x7c\x64\x8f\x99\x64\xb1\x9d\xe0\x56\x1b\x8c\xb6\x93\x25\xa7\xe8\x1b\x7e\x7a\x7b\xdf\x65\x9d\x0f\x7e\x9a\xf1\x85\x27\xa7\x21\x76\x65\xcd\x23\x2c\xbc\xaa\x9a\xe8\xe5\xb1\xe8\xb7\xed\xbd\x8e\x22\x2c\x0a\x39\x0d\x86\x83\x8b\x35\x07\x9e\xac\xc0\x19\x2c\x9d\x6f\x38\xa7\xa7\xba\x3c\x59\xca\x87\xfc\x22\x2c\x26\xbf\x37\x22\xde\xcd\x9a\xdc\x5b\x1a\xda\x8e\x14\xc8\xa5\x2d\xeb\xbf\xa9\x80\xf1\xef\x9d\x1e\x75\xec\xf8\x3f\xce\xc9\x4f\x1d\xb9\xad\xae\x7a\x62\x7f\x2a\xa8\xab\x63\xaf\xaa\xab\x7f\x02\x00\x00\xff\xff\x58\x91\x10\x7d\x50\x16\x00\x00")

func templatesTypes_serviceGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTypes_serviceGohtml,
		"templates/types_service.gohtml",
	)
}

func templatesTypes_serviceGohtml() (*asset, error) {
	bytes, err := templatesTypes_serviceGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/types_service.gohtml", size: 5712, mode: os.FileMode(436), modTime: time.Unix(1543843328, 0)}
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
	"templates/output_fields.gohtml": templatesOutput_fieldsGohtml,
	"templates/output_map_fields.gohtml": templatesOutput_map_fieldsGohtml,
	"templates/schemas_body.gohtml": templatesSchemas_bodyGohtml,
	"templates/schemas_head.gohtml": templatesSchemas_headGohtml,
	"templates/types_body.gohtml": templatesTypes_bodyGohtml,
	"templates/types_head.gohtml": templatesTypes_headGohtml,
	"templates/types_service.gohtml": templatesTypes_serviceGohtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"output_fields.gohtml": &bintree{templatesOutput_fieldsGohtml, map[string]*bintree{}},
		"output_map_fields.gohtml": &bintree{templatesOutput_map_fieldsGohtml, map[string]*bintree{}},
		"schemas_body.gohtml": &bintree{templatesSchemas_bodyGohtml, map[string]*bintree{}},
		"schemas_head.gohtml": &bintree{templatesSchemas_headGohtml, map[string]*bintree{}},
		"types_body.gohtml": &bintree{templatesTypes_bodyGohtml, map[string]*bintree{}},
		"types_head.gohtml": &bintree{templatesTypes_headGohtml, map[string]*bintree{}},
		"types_service.gohtml": &bintree{templatesTypes_serviceGohtml, map[string]*bintree{}},
	}},
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

