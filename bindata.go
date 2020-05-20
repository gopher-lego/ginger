// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// setting/app.debug.json
// setting/app.mode.json.example
// setting/app.release.json
package main

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

// Mode return file modify time
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

var _settingAppDebugJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\xc1\x4a\xc4\x30\x10\x06\xe0\x7b\x9f\x62\xc8\x61\x51\x90\x36\xdd\xae\x25\x1b\x28\xea\x5d\x6f\xeb\x03\x4c\xdb\x91\x16\x93\x26\x26\xb3\x82\x48\xdf\x5d\xa6\x87\x0a\x7b\x9c\xe4\xe3\xff\xff\xdf\x02\x40\x2d\xe8\x49\x59\x50\xf9\x93\x1c\x71\x58\xe0\x25\x46\xf5\x20\x3f\x99\xd2\x37\x25\x65\x41\x1c\x80\x8a\x21\xb1\x48\x6b\xf4\x59\x6f\x04\x40\x61\x8c\xef\xc9\xc9\xf3\xc4\x1c\x6d\x55\xb9\x30\xa0\x9b\x42\x66\x6b\xb4\xd9\xd9\x18\x3c\xce\xcb\x1b\xce\x8b\xd0\x1b\x53\x00\xac\x5b\xe3\x88\x8c\x3d\x66\xfa\xef\xf4\x3f\xf9\xcb\xed\xa7\x90\xbc\x25\xa4\x10\xd8\xd6\xc7\xe6\xf4\xd8\x3e\xdf\xd5\xe7\x63\x59\xb7\xa6\xd4\x65\xad\xb5\x6d\x1a\xdd\xde\x57\x4c\x99\xc7\xfe\x69\x98\x30\x65\xe2\xee\xca\x1f\xc6\xf7\xa7\x43\x94\xf3\x32\x7b\xea\x2e\xe9\x4a\x07\x17\x86\xee\x55\xc6\xa8\x2d\x7f\x95\x25\xc5\x5a\xfc\x05\x00\x00\xff\xff\x1a\x19\x42\x6d\x19\x01\x00\x00")

func settingAppDebugJsonBytes() ([]byte, error) {
	return bindataRead(
		_settingAppDebugJson,
		"setting/app.debug.json",
	)
}

func settingAppDebugJson() (*asset, error) {
	bytes, err := settingAppDebugJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "setting/app.debug.json", size: 281, mode: os.FileMode(420), modTime: time.Unix(1589971897, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _settingAppModeJsonExample = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xce\xcf\x4e\x84\x30\x10\x06\xf0\x3b\x4f\x31\xe9\x61\xa3\x89\x29\x05\x56\xc2\x36\x21\xea\x5d\x6f\xeb\x03\x0c\x30\x66\x89\x2d\xad\xed\x20\x1a\xc3\xbb\x9b\x36\xf1\xcf\xf1\x6b\x7f\x33\xdf\x7c\x15\x00\x62\x41\x4b\x42\x83\x88\xaf\x64\x88\xdd\x02\x0f\xde\x8b\x9b\xf4\x13\x29\xbc\x53\x10\x1a\x92\x03\x10\xde\x05\x4e\x52\x77\xea\xa4\x32\x01\x10\xe8\xfd\x73\x30\xe9\xf9\xc2\xec\x75\x59\x6e\xdb\x26\xe9\x03\xad\x37\x24\x47\x67\x7f\xdc\xe4\x2c\xce\xcb\x13\xce\x4b\xb2\xff\x41\x01\xb0\xe7\xbe\x09\x19\x07\x8c\xf4\xd7\x68\x3f\xe3\x9b\xf9\x8d\x89\xc4\x3c\x1e\x9c\x63\x5d\xd5\xcd\xf1\xb6\xbd\xbf\xaa\x4e\xb5\xac\xda\x4e\x2a\x59\xa9\x5a\x37\x8d\x6a\xaf\x4b\xa6\xc8\xd3\x70\x37\x5e\x30\x44\xe2\x7e\xe5\x97\xce\x0e\xc7\x83\x4f\xf1\x3c\x5b\xea\xcf\x61\xa5\x83\x71\x63\xff\xe8\x46\x34\x22\xef\xdf\xd3\x25\xc5\x5e\x7c\x07\x00\x00\xff\xff\x32\x41\xd8\x2c\x17\x01\x00\x00")

func settingAppModeJsonExampleBytes() ([]byte, error) {
	return bindataRead(
		_settingAppModeJsonExample,
		"setting/app.mode.json.example",
	)
}

func settingAppModeJsonExample() (*asset, error) {
	bytes, err := settingAppModeJsonExampleBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "setting/app.mode.json.example", size: 279, mode: os.FileMode(420), modTime: time.Unix(1589971876, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _settingAppReleaseJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\xc1\x4a\xc4\x30\x10\x06\xe0\x7b\x9f\x62\xc8\x61\x51\x90\x36\xdd\xae\x25\x1b\x28\xea\x5d\x6f\xeb\x03\x4c\xdb\x91\x16\x93\x26\x26\xb3\x82\x48\xdf\x5d\xa6\x87\x0a\x7b\x9c\xe4\xe3\xff\xff\xdf\x02\x40\x2d\xe8\x49\x59\x50\xf9\x93\x1c\x71\x58\xe0\x25\x46\xf5\x20\x3f\x99\xd2\x37\x25\x65\x41\x1c\x80\x8a\x21\xb1\x48\x6b\xf4\x59\x6f\x04\x40\x61\x8c\xef\xc9\xc9\xf3\xc4\x1c\x6d\x55\xb9\x30\xa0\x9b\x42\x66\x6b\xb4\xd9\xd9\x18\x3c\xce\xcb\x1b\xce\x8b\xd0\x1b\x53\x00\xac\x5b\xe3\x88\x8c\x3d\x66\xfa\xef\xf4\x3f\xf9\xcb\xed\xa7\x90\xbc\x25\xa4\x10\xd8\xd6\xc7\xe6\xf4\xd8\x3e\xdf\xd5\xe7\x63\x59\xb7\xa6\xd4\x65\xad\xb5\x6d\x1a\xdd\xde\x57\x4c\x99\xc7\xfe\x69\x98\x30\x65\xe2\xee\xca\x1f\xc6\xf7\xa7\x43\x94\xf3\x32\x7b\xea\x2e\xe9\x4a\x07\x17\x86\xee\x55\xc6\xa8\x2d\x7f\x95\x25\xc5\x5a\xfc\x05\x00\x00\xff\xff\x1a\x19\x42\x6d\x19\x01\x00\x00")

func settingAppReleaseJsonBytes() ([]byte, error) {
	return bindataRead(
		_settingAppReleaseJson,
		"setting/app.release.json",
	)
}

func settingAppReleaseJson() (*asset, error) {
	bytes, err := settingAppReleaseJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "setting/app.release.json", size: 281, mode: os.FileMode(420), modTime: time.Unix(1589971904, 0)}
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
	"setting/app.debug.json":        settingAppDebugJson,
	"setting/app.mode.json.example": settingAppModeJsonExample,
	"setting/app.release.json":      settingAppReleaseJson,
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
	"setting": &bintree{nil, map[string]*bintree{
		"app.debug.json":        &bintree{settingAppDebugJson, map[string]*bintree{}},
		"app.mode.json.example": &bintree{settingAppModeJsonExample, map[string]*bintree{}},
		"app.release.json":      &bintree{settingAppReleaseJson, map[string]*bintree{}},
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
