// Code generated by go-bindata.
// sources:
// generator/plugins/dataloader/templates/loaders_body.gohtml
// generator/plugins/dataloader/templates/loaders_head.gohtml
// generator/plugins/dataloader/templates/output_object_fields.gohtml
// DO NOT EDIT!

package dataloader

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

var _templatesLoaders_bodyGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x4f\x6f\x1a\x3f\x10\x3d\xdb\x9f\x62\x0e\xd1\x4f\x2c\x22\xbb\xd2\xef\x88\xc4\xa5\xa4\x41\x55\x9b\xaa\x6a\xd3\xe6\x10\xe5\xe0\x7a\x87\x8d\x85\xb1\x89\x77\x96\x82\x2c\x7f\xf7\xca\xfb\x27\x2c\x64\x23\x38\xf4\x84\x34\xf3\x76\xe6\xcd\x7b\xcf\x78\x7f\x0d\xd9\xb8\xb0\xb4\xdf\xe0\x14\x0a\x45\xcf\xd5\xef\x54\xda\x75\xf6\x71\x71\x7f\xfd\x73\xe5\x84\x32\x98\x15\xf6\xff\xe2\x45\x67\x05\x1a\x74\x82\xac\xcb\x36\xba\x2a\x94\x29\xb3\x5c\x90\xd0\x56\xe4\xe8\xd2\x2f\xf5\x4f\xf9\xc1\xe6\xfb\xb9\x35\x84\x3b\x1a\x67\x70\x1d\x02\xe7\x71\x34\x34\xed\xb9\x56\x68\xa8\x04\x65\x08\xdd\x52\x48\x04\xcf\x99\xf7\x4e\x98\x02\xe1\xaa\x99\x04\xd3\x19\x5c\x75\xe3\xea\x09\x8c\x2d\x90\xbc\x6f\xfb\xe9\x0f\x74\x5b\x25\x31\xfd\x2a\xd6\x18\x42\x33\x72\x94\x80\xf7\x85\xbd\x8f\xab\x4e\x71\x73\xa1\xf5\xa7\x6e\x61\x1c\xe7\x3d\x9a\xbc\x9e\xdc\xb1\xbb\x11\x24\xba\x8d\x25\xb9\x4a\xd2\x85\xc4\x0e\xac\xe6\xd6\x2c\x55\xd1\x92\x6a\x40\x70\xe8\x36\x85\x48\xaf\x0f\x18\xa2\x92\x1f\xa8\xb4\x3a\x7e\xc6\x7d\x7d\x57\x43\xcc\x07\xce\xb7\xc2\x0d\xe3\x60\xf6\xfe\xf7\xf1\xc3\x65\x65\x24\x2c\x90\xda\xce\x83\xa2\xe7\x16\x3b\x92\xb4\x03\xd9\x94\xd3\xb6\x3d\x01\xb1\x51\x9d\x65\x47\x06\x26\xa7\xd0\x28\x57\x6f\x73\x94\xea\xbf\x9e\xa8\xff\x40\xcc\x29\x48\x87\x82\xf0\x1d\x50\xe4\xdf\xe7\x9b\x5e\x10\x99\x64\xd2\x37\x80\x05\xce\x99\x43\xaa\x9c\x79\xbd\x2e\x0a\xf4\x4b\xe8\x0a\x9b\xf1\x83\xda\x1e\x95\x93\xe8\xe3\xd9\x53\x6b\x1b\xce\x9e\xf3\xd6\x0e\x59\x13\xbf\x38\xe9\xc9\xd9\x00\x46\xdb\xda\x93\xcf\x41\x3d\x67\x6c\x89\x24\x9f\xa7\x10\xe9\x8f\x56\xb8\x2f\xdf\x32\xf9\x8e\x2f\x15\x96\xb4\xa8\xab\x91\xc1\xe8\xf1\x69\x00\x54\x6e\xac\x29\xb1\x43\x4d\xe0\xf1\x09\x9d\xb3\x2e\x89\x7c\xfa\x39\xb8\x8d\x0b\xe7\x36\xaf\x9f\x2d\x0b\x13\xce\xd8\x1f\xa1\x68\x0a\xde\xe7\x95\x13\xa4\xac\x81\x13\xfd\x1e\x84\xa2\x9b\xb6\x77\x57\x86\x00\x63\xf0\x9e\xd4\x1a\xbf\xad\x8a\x10\xd2\x3b\xa5\xb5\x2a\x51\x5a\x93\x4f\xa2\xe5\x81\x1f\x12\xf0\xfa\x3c\x7a\xc9\xbd\x75\x76\xdd\xea\x3f\xe4\x49\x02\xe3\xfe\x7f\x87\xe7\x6c\x2b\x74\x34\x5c\xd2\x2e\x6d\x92\x33\x18\x9a\x84\x73\xa6\x96\x10\xc1\xb3\x19\x18\xa5\xeb\xcb\x5b\x2b\x8c\xd2\x47\x69\xdc\x0a\x9d\x8e\xfa\x7b\x62\xca\xfe\x06\x00\x00\xff\xff\xdb\x59\x69\xe3\xb9\x05\x00\x00")

func templatesLoaders_bodyGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesLoaders_bodyGohtml,
		"templates/loaders_body.gohtml",
	)
}

func templatesLoaders_bodyGohtml() (*asset, error) {
	bytes, err := templatesLoaders_bodyGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/loaders_body.gohtml", size: 1465, mode: os.FileMode(436), modTime: time.Unix(1548156719, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLoaders_headGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8c\xc1\x4a\x04\x31\x0c\x86\xef\x79\x8a\x30\xec\x41\x17\x66\x0a\x1e\x17\x3c\x88\x88\x0a\x1e\x3c\xe8\x03\xc4\x69\xc8\x96\xed\xb6\xb5\xcd\x82\x12\xfa\xee\x42\x77\xc0\x53\xc2\xff\x7d\xff\x6f\x36\xa3\xdb\x4b\xd6\xdf\xc2\x07\x94\xa0\xc7\xcb\xd7\xb2\xe6\xb3\x7b\x7a\xfe\x98\x3f\x4f\x95\x42\x62\x27\xf9\x4e\xbe\xa3\x13\x4e\x5c\x49\x73\x75\x25\x5e\x24\xa4\xe6\x3c\x29\xc5\x4c\x9e\xeb\xf2\x36\x4e\x7b\x61\xf2\x8f\x39\x29\xff\xe8\xde\xe1\xdc\x3b\x40\xa1\xf5\x44\xc2\x78\x15\x1b\x40\x38\x97\x5c\x15\x6f\x60\x5a\xaf\xe6\x04\x60\x56\x29\x09\xe3\x6e\x83\x87\x7b\xdc\x2d\xaf\xe3\x6f\x63\x06\x11\xd1\x6c\xc3\xcb\x43\x0c\xd4\x7a\xc7\xe9\x3f\x7a\x27\x3d\xf6\x3e\x81\x19\x27\x3f\x2a\xb7\xf0\x17\x00\x00\xff\xff\x95\x1a\xd0\x59\xde\x00\x00\x00")

func templatesLoaders_headGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesLoaders_headGohtml,
		"templates/loaders_head.gohtml",
	)
}

func templatesLoaders_headGohtml() (*asset, error) {
	bytes, err := templatesLoaders_headGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/loaders_head.gohtml", size: 222, mode: os.FileMode(436), modTime: time.Unix(1548156719, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOutput_object_fieldsGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\xcb\x6e\xdb\x30\x10\x3c\xd3\x5f\xc1\x1a\x41\x21\x05\x09\x05\xf4\x68\x20\x87\x20\xaf\x43\x8b\xc4\x48\xdc\xf6\x58\xd0\xd2\x8a\x66\x4d\x91\x32\x45\xa5\x4d\x09\xfe\x7b\xb1\x24\x65\x1b\x2e\xd0\x4b\x81\x9e\x2c\x73\x77\x76\x66\x67\x24\x7a\x7f\x49\xab\x73\x61\xdc\x5b\x0f\x0b\x2a\xa4\xdb\x8c\x6b\x56\x9b\xae\xba\x7b\x58\x5d\x7e\xde\x5a\x2e\x35\x54\xc2\x7c\x10\x3b\x55\x09\xd0\x60\xb9\x33\xb6\xea\xd5\x28\xa4\x1e\x2a\x61\x79\xbf\xd9\x29\xf6\x0c\xba\x01\x7b\x2f\x41\x35\xc3\x8d\xd1\x0e\x7e\xba\xf3\x8a\x5e\x86\x30\xf3\x9e\x5a\xae\x05\xd0\xb3\x16\xab\x74\x71\x45\xcf\xd8\xd3\xe8\xfa\xd1\x3d\xad\xbf\x43\xed\xd8\x2d\x77\xfc\x93\xe1\x7b\x7c\x84\x11\xef\x4f\xda\xbe\x70\x2b\xf9\x5a\xc1\x23\xef\x20\x04\x76\xdd\x34\xb1\xfd\xc6\xe8\x56\x8a\x62\xee\x7d\x22\x60\xa9\x3e\xbf\xa0\xef\xbd\x17\x3b\xb5\xdc\x8a\x10\x58\x6c\xf5\x33\x42\xb0\xba\x20\xe4\xcf\xf6\x19\x21\xb7\x30\xd4\x56\xf6\x4e\x1a\xbd\xa0\xf3\x78\xb4\x42\x57\x08\xf1\x3e\x2f\x9a\x14\x25\xb5\x58\x43\x34\xee\x13\x25\xe6\xc5\xf3\xa6\x21\xe0\x80\x67\x18\x8c\x7a\x85\x05\x6d\x47\x5d\x17\x3d\x3d\x18\x16\xcf\x97\xdc\xf2\x6e\x28\x69\x21\xb5\x03\xdb\xf2\x1a\x7c\xb8\xa0\x60\xad\xb1\x25\x45\xbd\xa4\xe7\x16\xb4\x43\xdb\x7a\xf6\x62\x46\x5b\x03\x2b\xce\xbd\x17\x06\xe9\x4f\xad\x7c\x88\xa7\x21\x94\x33\x84\xaa\x28\x73\x40\xac\xf7\xf9\x4f\xb2\xe3\x01\xdc\xc1\xf5\xe1\xde\x9a\x2e\x6b\x2f\x7a\x96\x9f\xd2\x08\xd9\xd2\x69\xca\xd5\x15\xd5\x52\x25\x51\xc4\x82\x1b\xad\xc6\x83\xac\x76\x60\x8f\xf0\xa3\x98\xe3\xd4\x3d\x42\x1b\x47\x5b\x33\xea\x86\x4a\x4d\xeb\x34\x96\xd1\x1b\xae\xd4\xd4\x82\x42\x32\xdf\x57\xe9\x36\x59\xcf\xbc\x44\x8e\x10\x05\x78\x2f\xdb\x6c\x28\xfb\x08\x6f\x31\xc7\x17\x25\x6b\xc0\x77\x84\x10\xb7\x19\xf5\x16\x17\x9c\x06\xee\x73\x3d\xec\x97\x12\x4e\xcf\x0c\x7f\xae\x95\x5a\x21\xae\x48\xde\x1e\x30\x8f\xc6\x76\x5c\xc9\x5f\xd0\x2c\x63\x65\x22\x4c\x13\x92\x23\x79\xf3\x18\xe7\xdf\x72\x23\xaf\xdc\x66\x55\x77\xd1\xa0\x54\x9d\x65\xfb\x86\x51\xb9\x08\x88\xf1\xc4\x35\x8a\x44\x40\x5a\x63\xe9\xb7\x58\xc3\x52\xfa\x78\x62\x63\x9a\x8b\x99\x60\xed\xdd\x71\x1e\x53\xd8\x99\x0a\x03\xef\x46\xe5\x64\xe4\x4c\x99\x5f\xf7\x3d\xe8\xa6\x38\xee\x8b\x24\x65\x1a\x10\xed\xcc\xa6\x4f\x4b\x4e\x32\x8f\x31\x31\x9a\x0b\xa4\x4e\xf1\x80\x1a\xfe\x25\x8b\xff\x10\x44\x6e\x9c\x3c\x3e\xd5\xaf\x9b\x28\x1f\xbf\xd6\x50\xe2\x75\x05\xba\xc1\x2b\xe8\x77\x00\x00\x00\xff\xff\x45\xf6\x7b\xaf\x19\x05\x00\x00")

func templatesOutput_object_fieldsGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesOutput_object_fieldsGohtml,
		"templates/output_object_fields.gohtml",
	)
}

func templatesOutput_object_fieldsGohtml() (*asset, error) {
	bytes, err := templatesOutput_object_fieldsGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/output_object_fields.gohtml", size: 1305, mode: os.FileMode(436), modTime: time.Unix(1548323628, 0)}
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
	"templates/loaders_body.gohtml": templatesLoaders_bodyGohtml,
	"templates/loaders_head.gohtml": templatesLoaders_headGohtml,
	"templates/output_object_fields.gohtml": templatesOutput_object_fieldsGohtml,
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
		"loaders_body.gohtml": &bintree{templatesLoaders_bodyGohtml, map[string]*bintree{}},
		"loaders_head.gohtml": &bintree{templatesLoaders_headGohtml, map[string]*bintree{}},
		"output_object_fields.gohtml": &bintree{templatesOutput_object_fieldsGohtml, map[string]*bintree{}},
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

