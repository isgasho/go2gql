// Code generated by go-bindata.
// sources:
// generator/plugins/swagger2gql/templates/method_caller.gohtml
// generator/plugins/swagger2gql/templates/method_caller_null.gohtml
// generator/plugins/swagger2gql/templates/value_resolver_array.gohtml
// generator/plugins/swagger2gql/templates/value_resolver_datetime.gohtml
// generator/plugins/swagger2gql/templates/value_resolver_ptr_datetime.gohtml
// DO NOT EDIT!

package swagger2gql

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

var _templatesMethod_callerGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\xc1\xaa\x83\x30\x10\x45\xf7\x7e\xc5\x3c\x78\x8b\x04\x24\x1f\x50\x70\xd5\x7d\xe9\xa2\xb4\xcb\x12\x74\xc4\x80\x26\x7a\x8d\xa0\xc8\xfc\x7b\x49\x14\xda\xe5\x3d\x33\x9c\xb9\xd3\x2e\xbe\x56\xe0\x89\xf6\xfd\xdf\x80\xa7\xc7\x36\xb2\x88\x26\xf5\x3e\xc9\x3c\x1e\xa8\x24\x30\x40\x0c\x04\x68\xda\x0b\x22\x22\xf0\x5c\x26\x42\x97\x2a\x6f\xd7\xbd\x63\x1f\x9f\x16\x22\x26\xe5\x81\x63\x17\x9a\x9b\x1d\x58\x24\x1d\x31\x2f\x17\xbb\x6b\xf0\x91\xd7\xa8\xea\xb8\x6a\x9d\x35\xae\xcd\x92\xbf\x8a\xbc\xeb\x4f\xf5\xa1\x07\xa8\x4a\xb3\x1f\x14\x17\xf8\x1c\xa5\xf8\x82\xd4\xc4\xdc\xed\xd6\x07\xdb\x94\xc9\x52\x88\x3a\x1f\xca\x6d\xf4\x27\x00\x00\xff\xff\xe3\x27\x36\x2c\xe7\x00\x00\x00")

func templatesMethod_callerGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMethod_callerGohtml,
		"templates/method_caller.gohtml",
	)
}

func templatesMethod_callerGohtml() (*asset, error) {
	bytes, err := templatesMethod_callerGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/method_caller.gohtml", size: 231, mode: os.FileMode(420), modTime: time.Unix(1542968209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMethod_caller_nullGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcd\x31\x0a\xc3\x30\x10\x44\xd1\xde\xa7\xd8\x22\x85\x04\x41\x97\x48\x9f\x2a\x24\xb5\x50\xc6\x58\x10\xaf\xe2\x61\x0d\x0e\x62\xef\x1e\x6c\x5c\x4e\xf1\xe6\x8f\xab\x96\x40\x2c\xd2\xfb\x25\x11\xcb\xe3\xf7\x85\x7b\x94\x50\xd5\xc0\x31\x17\x74\xbf\x0a\xc8\xc6\x28\x7d\x10\x11\x21\x6c\xa5\x1e\xa0\x7c\x2a\xd4\x9e\x99\xee\x69\xdf\x33\x6c\x6a\xef\x7b\x9e\xe1\xbe\xbf\xa6\x57\xb5\xe9\xd6\xd4\xb0\x59\x28\xb6\xc5\x38\x78\x38\x4b\x87\x8a\xff\x00\x00\x00\xff\xff\xa7\xc5\xb0\x4f\x80\x00\x00\x00")

func templatesMethod_caller_nullGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMethod_caller_nullGohtml,
		"templates/method_caller_null.gohtml",
	)
}

func templatesMethod_caller_nullGohtml() (*asset, error) {
	bytes, err := templatesMethod_caller_nullGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/method_caller_null.gohtml", size: 128, mode: os.FileMode(420), modTime: time.Unix(1543234503, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesValue_resolver_arrayGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\xc1\x6a\xf3\x30\x0c\xc7\xcf\xf6\x53\xa8\xa1\xf0\x39\x7c\x59\x1e\x60\xd0\xd3\xd8\x75\x8c\x31\xda\x43\x29\x43\x64\x4a\x66\xe2\xda\x45\x76\xb6\x15\xe3\x77\x1f\x76\x9a\xd1\xcb\x06\xcb\x2d\xb2\xf5\xd3\x4f\x7f\xf7\x93\xed\x40\x21\x0f\xa0\x6d\x20\xee\xb1\xa3\x98\x6a\x50\x31\x76\x68\x0c\xac\x5b\x26\x3f\x99\xf0\x7c\x3e\x51\x4a\x0d\x10\xb3\xe3\x1a\xa2\x14\x64\xe8\x48\x36\xf8\x06\xdc\x08\xb7\x1b\x40\x1e\x5a\xb5\x3f\x5c\x53\xa4\xd0\x3d\xac\xdc\x98\xaf\x0b\xa6\x30\xb1\x05\xab\x4d\x03\x31\x16\x8e\x7f\x1c\x87\x94\xda\x07\xfa\x50\x15\xf2\x30\x65\x1e\x68\x0f\xd6\x05\x40\x66\x3c\x57\xb5\x14\x49\x0a\x26\x9f\x27\x1c\x71\xa4\x9f\xbc\x0c\x59\xb5\x18\xd5\xb5\x14\xbd\x63\xd0\x0d\x5c\x4a\xb9\x9b\xd1\x0e\xb4\x14\x7c\x36\x7a\x81\xcd\xf2\x2f\x85\x88\x11\x74\x0f\xeb\x36\x57\x9e\xc8\x3b\xf3\x4e\xbc\xd3\xe1\xed\x9e\x19\x6e\x52\x92\x42\x08\x32\x5b\x34\x25\x83\x0c\xfc\x36\xb9\xee\x80\xea\x42\xac\xb2\xa2\x73\xe1\x2e\x7c\xce\xcd\xba\x2f\x8d\xab\x4d\x8e\x20\x4a\xb8\xfa\x7e\xc9\x66\xc7\x78\x52\xc4\xdc\x40\xd5\xa1\xfd\x17\x80\xe7\x41\x73\x3e\x8b\x7e\xce\x49\x88\x32\x86\xc9\xef\xf5\xa1\x6c\xb6\x45\x33\xef\x45\xc6\xd3\xb2\xc3\xe5\xfc\xff\x5f\xfd\x33\xc6\xbe\xce\x94\xf2\x24\xc5\x98\xc9\x37\x59\x5b\x26\x15\xe3\xba\x45\x1e\x52\xaa\xbf\x02\x00\x00\xff\xff\xb9\x59\x43\x9b\x53\x02\x00\x00")

func templatesValue_resolver_arrayGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesValue_resolver_arrayGohtml,
		"templates/value_resolver_array.gohtml",
	)
}

func templatesValue_resolver_arrayGohtml() (*asset, error) {
	bytes, err := templatesValue_resolver_arrayGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/value_resolver_array.gohtml", size: 595, mode: os.FileMode(436), modTime: time.Unix(1539090323, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesValue_resolver_datetimeGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xc1\x4a\xc4\x30\x10\x86\xef\xfb\x14\x43\x11\x49\x60\xe9\x41\xc5\x83\xd0\x9b\x67\xf1\xa0\xa7\x65\x91\xa1\x9d\x96\x60\x3b\x2d\x93\x59\x14\x66\xf3\xee\x92\x54\xdd\x5e\xb6\xd0\x84\xfc\xc9\x7c\xf3\x25\xfd\x89\x5b\x87\x32\x40\x60\x25\xe9\xb1\x25\x4b\x1e\xdc\x07\x98\x45\x95\x7e\xd2\xd7\xcf\x21\xa5\xfa\x19\x95\xde\xc2\x44\x7b\x20\x91\xfc\xcf\xe2\xc1\x76\x00\x00\xa1\x87\x5c\xdf\x34\xc0\x61\xfc\xcd\xf2\x27\xa4\x27\xe1\xb2\x4c\x65\x44\x78\x6a\xf2\xd1\xda\x4d\xb8\x1c\xa2\x4a\xe0\xe1\xb8\x6d\xfb\x8f\x3b\x54\x91\xda\x99\xbb\x58\x1d\xff\xb8\xe7\x73\x8e\x19\x79\xde\x84\x97\x66\xd9\xaa\x01\xb3\x62\x16\x57\xe7\x17\xfa\x72\x15\xcf\x0a\x38\x8e\xd0\xa1\x92\x86\x89\x60\x41\xc1\x89\x94\x24\xc2\x82\x31\x52\x57\xf9\xeb\xca\x9a\x95\xcd\x72\xe1\xca\x7c\xe7\xf0\xed\xb6\x7e\xb5\x0b\xac\x8f\x0f\x7e\x0f\x65\x76\x17\xc9\xb2\x73\x7f\xe7\xfd\xca\x5f\xd9\x57\x9e\xd5\xdd\xaa\xdf\xe7\x2b\xed\x92\x33\xbb\xa9\x51\x86\x94\xfc\x4f\x00\x00\x00\xff\xff\x9e\xeb\xe6\xc5\x9c\x01\x00\x00")

func templatesValue_resolver_datetimeGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesValue_resolver_datetimeGohtml,
		"templates/value_resolver_datetime.gohtml",
	)
}

func templatesValue_resolver_datetimeGohtml() (*asset, error) {
	bytes, err := templatesValue_resolver_datetimeGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/value_resolver_datetime.gohtml", size: 412, mode: os.FileMode(436), modTime: time.Unix(1539090323, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesValue_resolver_ptr_datetimeGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x31\x4b\xc5\x40\x0c\xc7\xf7\x7e\x8a\x50\x44\xee\xa4\x74\x50\x71\x10\xba\x39\x8b\x83\x4e\x8f\x37\x84\x36\x2d\x87\x6d\x5a\x72\x79\x28\xe4\xdd\x77\x97\x6b\x9f\xda\x41\x6f\xc8\xc1\xef\x2e\xf9\xff\x48\x7f\xe2\xd6\xa1\x0c\x10\x58\x49\x7a\x6c\xc9\x92\x07\x77\x63\x16\x55\xfa\x49\x5f\xde\x87\x94\xea\x27\x54\x7a\x0d\x13\x55\x40\x22\xb3\x78\xb0\x02\x00\x20\xf4\x90\x5b\x9b\x06\x38\x8c\x17\x96\x8f\x90\x9e\x84\x33\xac\x72\x59\x79\x5a\x2b\xc2\x63\x93\x7b\x6a\x37\xe1\x72\x88\x2a\x81\x87\xe3\x3e\xfa\x67\xee\xa1\x8c\xd4\xce\xdc\xc5\xf2\xf8\x1d\x70\x3e\x67\xcc\xc8\xf3\x0e\xfe\x9d\x6a\xb6\x8a\xc6\x4d\xff\x99\x3e\x5c\xc9\xb3\x02\x8e\x23\x74\xa8\xa4\x61\x22\x58\x50\x70\x22\x25\x89\xb0\x60\x8c\xd4\x95\x7e\x67\xaa\xd9\xd4\x2c\xff\xdc\x86\xbc\x71\xf8\x74\x7b\xad\xda\x05\xd6\x87\x7b\x5f\xc1\x7a\xbb\x5f\xb7\xf5\xe5\xee\xd6\xfb\x6d\xe0\x45\xec\xbf\xa5\x7a\x77\xad\x7e\x5b\x54\x72\x66\x57\x35\xca\x90\x92\x2f\x8a\xa2\xf8\x0a\x00\x00\xff\xff\xf1\x9b\x6c\x88\x9e\x01\x00\x00")

func templatesValue_resolver_ptr_datetimeGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesValue_resolver_ptr_datetimeGohtml,
		"templates/value_resolver_ptr_datetime.gohtml",
	)
}

func templatesValue_resolver_ptr_datetimeGohtml() (*asset, error) {
	bytes, err := templatesValue_resolver_ptr_datetimeGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/value_resolver_ptr_datetime.gohtml", size: 414, mode: os.FileMode(436), modTime: time.Unix(1539090323, 0)}
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
	"templates/method_caller.gohtml": templatesMethod_callerGohtml,
	"templates/method_caller_null.gohtml": templatesMethod_caller_nullGohtml,
	"templates/value_resolver_array.gohtml": templatesValue_resolver_arrayGohtml,
	"templates/value_resolver_datetime.gohtml": templatesValue_resolver_datetimeGohtml,
	"templates/value_resolver_ptr_datetime.gohtml": templatesValue_resolver_ptr_datetimeGohtml,
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
		"method_caller.gohtml": &bintree{templatesMethod_callerGohtml, map[string]*bintree{}},
		"method_caller_null.gohtml": &bintree{templatesMethod_caller_nullGohtml, map[string]*bintree{}},
		"value_resolver_array.gohtml": &bintree{templatesValue_resolver_arrayGohtml, map[string]*bintree{}},
		"value_resolver_datetime.gohtml": &bintree{templatesValue_resolver_datetimeGohtml, map[string]*bintree{}},
		"value_resolver_ptr_datetime.gohtml": &bintree{templatesValue_resolver_ptr_datetimeGohtml, map[string]*bintree{}},
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

