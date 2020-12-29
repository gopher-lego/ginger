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

var _settingAppDebugJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xcd\x6e\xe3\x20\x14\x85\xf7\x7e\x0a\xc4\x22\x9a\x19\x65\x6c\xec\xfc\x88\x20\x45\x33\xdd\xb7\x6a\x17\xe9\x03\x10\x7c\x63\xa3\x82\x2f\xbd\x90\xa6\x3f\xca\xbb\x57\x58\x6a\xea\x66\xc9\xd1\x77\xcf\x39\x9c\x8f\x82\x31\x3e\x68\x0f\x5c\x31\xde\x61\xe8\x81\xfe\x3a\xe8\xb0\x8a\x4f\xe0\x20\xe1\xc0\x6e\x42\xe0\xf3\x4c\x45\xa0\x17\x20\xae\x58\xbe\x61\x8c\x07\xa4\x94\xaf\x94\x14\x1b\x31\x22\x8c\x71\x1d\xc2\x23\xb9\x2c\xf7\x29\x05\x55\x55\x0e\x8d\x76\x3d\xc6\xa4\xa4\x90\x17\xac\x45\xaf\xed\x70\xa7\xed\x90\xd1\x2b\xa6\x60\xec\x3c\x26\xb6\x3a\xe9\xbd\x8e\xf0\x9d\xe9\xdf\xe2\xb3\xbb\x3c\x33\x12\x47\x07\x42\x4c\xaa\x6e\x16\xcb\xd5\xfa\xff\xaf\x7a\xd3\x94\xf5\x5a\x96\xa2\xac\x85\x50\x8b\x85\x58\xff\xae\x12\xc4\xd4\xee\xff\x99\x5e\x53\x84\xb4\x3d\xa6\x83\xf4\xfb\xe5\x2c\xe4\xe7\xce\x7a\xd8\xee\xe8\x08\x33\x87\x66\x7b\x9b\xcb\xf0\xd1\xff\x7c\x69\x62\xb4\xe9\x27\x35\x0e\x04\xf0\x53\xca\xfb\xd8\xf7\x71\xc5\x55\x23\x1b\x29\xc5\xd7\x5f\x19\xe3\x9d\x79\x00\x32\x30\x8c\x73\x35\xe2\xda\xdc\xdb\xb6\x75\x70\xd2\x34\x49\x30\x48\x71\x6a\xae\x9d\xc3\xd3\x3d\xd9\xce\x0e\x59\xe7\x7f\xe6\x79\xdf\xa8\xaa\x0a\x5e\xb5\x0f\x0e\x4a\x83\x7e\x62\x5c\x9c\x8b\xcf\x00\x00\x00\xff\xff\xb1\xeb\xbe\xfa\xdb\x01\x00\x00")

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

	info := bindataFileInfo{name: "setting/app.debug.json", size: 475, mode: os.FileMode(420), modTime: time.Unix(1609176031, 0)}
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

var _settingAppReleaseJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xcd\x6e\xe3\x20\x14\x85\xf7\x7e\x0a\xc4\x22\x9a\x19\x65\x6c\xec\xfc\x88\x20\x45\x33\xdd\xb7\x6a\x17\xe9\x03\x10\x7c\x63\xa3\x82\x2f\xbd\x90\xa6\x3f\xca\xbb\x57\x58\x6a\xea\x66\xc9\xd1\x77\xcf\x39\x9c\x8f\x82\x31\x3e\x68\x0f\x5c\x31\xde\x61\xe8\x81\xfe\x3a\xe8\xb0\x8a\x4f\xe0\x20\xe1\xc0\x6e\x42\xe0\xf3\x4c\x45\xa0\x17\x20\xae\x58\xbe\x61\x8c\x07\xa4\x94\xaf\x94\x14\x1b\x31\x22\x8c\x71\x1d\xc2\x23\xb9\x2c\xf7\x29\x05\x55\x55\x0e\x8d\x76\x3d\xc6\xa4\xa4\x90\x17\xac\x45\xaf\xed\x70\xa7\xed\x90\xd1\x2b\xa6\x60\xec\x3c\x26\xb6\x3a\xe9\xbd\x8e\xf0\x9d\xe9\xdf\xe2\xb3\xbb\x3c\x33\x12\x47\x07\x42\x4c\xaa\x6e\x16\xcb\xd5\xfa\xff\xaf\x7a\xd3\x94\xf5\x5a\x96\xa2\xac\x85\x50\x8b\x85\x58\xff\xae\x12\xc4\xd4\xee\xff\x99\x5e\x53\x84\xb4\x3d\xa6\x83\xf4\xfb\xe5\x2c\xe4\xe7\xce\x7a\xd8\xee\xe8\x08\x33\x87\x66\x7b\x9b\xcb\xf0\xd1\xff\x7c\x69\x62\xb4\xe9\x27\x35\x0e\x04\xf0\x53\xca\xfb\xd8\xf7\x71\xc5\x55\x23\x1b\x29\xc5\xd7\x5f\x19\xe3\x9d\x79\x00\x32\x30\x8c\x73\x35\xe2\xda\xdc\xdb\xb6\x75\x70\xd2\x34\x49\x30\x48\x71\x6a\xae\x9d\xc3\xd3\x3d\xd9\xce\x0e\x59\xe7\x7f\xe6\x79\xdf\xa8\xaa\x0a\x5e\xb5\x0f\x0e\x4a\x83\x7e\x62\x5c\x9c\x8b\xcf\x00\x00\x00\xff\xff\xb1\xeb\xbe\xfa\xdb\x01\x00\x00")

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

	info := bindataFileInfo{name: "setting/app.release.json", size: 475, mode: os.FileMode(420), modTime: time.Unix(1609176055, 0)}
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
