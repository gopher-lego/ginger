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

var _settingAppDebugJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xcd\x4e\xc3\x30\x10\x84\xef\x79\x8a\x95\x0f\x15\xa0\x2a\x71\xd2\x1f\xb9\x96\x2a\xe0\x0e\x82\x43\x79\x80\xad\xb3\x6d\x2c\xec\xd8\xd8\x2e\xe5\x47\x7d\x77\xe4\x48\x94\xd0\xe3\x8e\x3f\xcd\x8c\xe7\xbb\x00\x60\x3d\x5a\x62\x12\x58\x7c\x25\x43\xc9\xf5\x70\xef\x3d\x9b\xe6\x97\x48\xe1\x9d\x02\x93\x90\x39\x00\xe6\x5d\x48\x99\x94\x82\xaf\xf8\x80\x00\x30\xf4\xfe\x25\x98\x2c\x77\x29\x79\x59\x55\xc6\x29\x34\x9d\x8b\x49\x0a\x2e\xce\x58\xeb\x2c\xea\xfe\x11\x75\x9f\xd1\x0b\xa6\x00\x38\x0d\x89\x2d\x26\xdc\x62\xa4\xbf\x4c\xfb\x19\xdf\xcc\xf9\xcc\x48\x1c\x1c\x82\x73\x49\xd6\xcd\x6c\xbe\x58\xde\x5d\xd5\xab\xa6\xac\x97\xa2\xe4\x65\xcd\xb9\x9c\xcd\xf8\xf2\xba\x4a\x14\x53\xbb\xbd\x55\x1d\x86\x48\x69\x7d\x48\x3b\x61\xb7\xf3\x89\xcf\xe7\x46\x5b\x5a\x6f\xc2\x81\x26\xc6\xa9\xf5\x43\x2e\xc3\x06\xff\xd3\xb9\x89\x42\xd5\x8d\x6a\xec\x02\xd1\x7f\x29\xef\xa3\xbf\x86\xe5\x16\x8d\x68\x84\xe0\xbf\x7f\x05\x60\x7b\xf5\x4c\x41\x51\x3f\xcc\xd5\xf0\x4b\x73\xab\xdb\xd6\xd0\x11\xc3\x28\x41\xb9\x10\xc7\xe6\x68\x8c\x3b\x3e\x05\xbd\xd7\x7d\xd6\xd9\xcd\x34\xef\x1b\x65\x55\xd1\x07\x5a\x6f\xa8\x54\xce\x8e\x8c\x8b\x53\xf1\x13\x00\x00\xff\xff\x23\x12\x23\x70\xcf\x01\x00\x00")

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

	info := bindataFileInfo{name: "setting/app.debug.json", size: 463, mode: os.FileMode(420), modTime: time.Unix(1591189547, 0)}
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

var _settingAppReleaseJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xcd\x4e\xc3\x30\x10\x84\xef\x79\x8a\x95\x0f\x15\xa0\x2a\x71\xd2\x1f\xb9\x96\x2a\xe0\x0e\x82\x43\x79\x80\xad\xb3\x6d\x2c\xec\xd8\xd8\x2e\xe5\x47\x7d\x77\xe4\x48\x94\xd0\xe3\x8e\x3f\xcd\x8c\xe7\xbb\x00\x60\x3d\x5a\x62\x12\x58\x7c\x25\x43\xc9\xf5\x70\xef\x3d\x9b\xe6\x97\x48\xe1\x9d\x02\x93\x90\x39\x00\xe6\x5d\x48\x99\x94\x82\xaf\xf8\x80\x00\x30\xf4\xfe\x25\x98\x2c\x77\x29\x79\x59\x55\xc6\x29\x34\x9d\x8b\x49\x0a\x2e\xce\x58\xeb\x2c\xea\xfe\x11\x75\x9f\xd1\x0b\xa6\x00\x38\x0d\x89\x2d\x26\xdc\x62\xa4\xbf\x4c\xfb\x19\xdf\xcc\xf9\xcc\x48\x1c\x1c\x82\x73\x49\xd6\xcd\x6c\xbe\x58\xde\x5d\xd5\xab\xa6\xac\x97\xa2\xe4\x65\xcd\xb9\x9c\xcd\xf8\xf2\xba\x4a\x14\x53\xbb\xbd\x55\x1d\x86\x48\x69\x7d\x48\x3b\x61\xb7\xf3\x89\xcf\xe7\x46\x5b\x5a\x6f\xc2\x81\x26\xc6\xa9\xf5\x43\x2e\xc3\x06\xff\xd3\xb9\x89\x42\xd5\x8d\x6a\xec\x02\xd1\x7f\x29\xef\xa3\xbf\x86\xe5\x16\x8d\x68\x84\xe0\xbf\x7f\x05\x60\x7b\xf5\x4c\x41\x51\x3f\xcc\xd5\xf0\x4b\x73\xab\xdb\xd6\xd0\x11\xc3\x28\x41\xb9\x10\xc7\xe6\x68\x8c\x3b\x3e\x05\xbd\xd7\x7d\xd6\xd9\xcd\x34\xef\x1b\x65\x55\xd1\x07\x5a\x6f\xa8\x54\xce\x8e\x8c\x8b\x53\xf1\x13\x00\x00\xff\xff\x23\x12\x23\x70\xcf\x01\x00\x00")

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

	info := bindataFileInfo{name: "setting/app.release.json", size: 463, mode: os.FileMode(420), modTime: time.Unix(1591189556, 0)}
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
