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

var _settingAppDebugJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\xcd\x6e\xe3\x20\x14\x85\xf7\x7e\x0a\xc4\x22\x9a\x19\x65\x6c\xec\xfc\x88\x20\x45\x33\xfb\x99\xaa\x59\xa4\x0f\x80\xf1\x8d\x8d\x04\x86\x5e\x70\x93\xb4\xca\xbb\x57\x58\x69\xea\xa6\xcb\x7b\xf8\x38\xf7\x70\x78\xcb\x08\xa1\xbd\xb4\x40\x05\xa1\xad\xf3\x1d\xe0\x6f\x03\xad\x23\xd2\x7b\x3a\x4f\x87\x01\xf0\x05\x90\x0a\x92\x50\x42\xa8\x77\x18\x13\x2c\x38\xdb\xb0\x11\x21\x84\x4a\xef\x9f\xd0\x24\xb9\x8b\xd1\x8b\xa2\x30\x4e\x49\xd3\xb9\x10\x05\x67\xfc\x86\x35\xce\x4a\xdd\x3f\x48\xdd\x27\xf4\x8e\xc9\x08\xb9\x8c\x1b\x1b\x19\x65\x2d\x03\x7c\xee\xb4\xe7\xf0\x6c\x6e\x63\x42\xc2\xe8\x80\xce\x45\x51\x56\x8b\xe5\x6a\xfd\xf7\x47\xb9\xa9\xf2\x72\xcd\x73\x96\x97\x8c\x89\xc5\x82\xad\x7f\x16\x11\x42\x6c\xea\x3f\xaa\x93\x18\x20\x6e\x87\x78\xe0\xb6\x5e\xce\x7c\x1a\xf7\xda\xc2\x76\x8f\x03\xcc\x8c\x53\xdb\xff\x29\x0c\x1d\xfd\x2f\xb7\x24\x4a\xaa\x6e\x12\xe3\x80\x00\x5f\xa5\xd4\x8f\x7e\x1d\xcb\x5b\x55\xbc\xe2\x9c\x7d\xbc\x95\x10\xda\xaa\x1d\xa0\x82\x7e\xac\xab\x62\x57\xf3\x6b\x15\xad\xfb\xe6\xa4\x86\x10\x9d\xfd\x07\xe7\x1d\xc2\x41\x9f\xd2\xad\x06\xea\xa1\xbd\x4f\x65\x75\xd3\x18\x38\x4a\x9c\x44\x53\x0e\xc3\xd4\x4b\x1a\xe3\x8e\x8f\xa8\x5b\xdd\x27\x9d\xfe\x9a\xa7\x8f\x09\xa2\x28\xe0\x24\xad\x37\x90\x2b\x67\x27\xc6\xd9\x25\x7b\x0f\x00\x00\xff\xff\xb4\x5d\xf9\x66\x0b\x02\x00\x00")

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

	info := bindataFileInfo{name: "setting/app.debug.json", size: 523, mode: os.FileMode(420), modTime: time.Unix(1612185124, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _settingAppModeJsonExample = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xce\x51\x4e\x84\x30\x10\x06\xe0\x77\x4e\x31\xe9\xc3\x46\x13\x2d\x05\x56\xc2\x36\x21\x7a\x00\x7d\x5b\x0f\x30\xc0\xb8\x90\xb4\xb4\xb6\x83\x68\x0c\x77\x37\x25\x51\xf7\xf1\x6f\xbf\x99\x7f\xbe\x33\x00\x31\xa3\x25\xa1\x41\x5c\x9c\x1f\x29\xdc\x1b\xba\x38\x40\xef\xc5\x5d\xfa\x8c\x14\x3e\x28\x08\x0d\x89\x02\x08\xef\x02\x27\xac\x1b\x75\x52\x3b\x01\x10\xe8\xfd\x6b\x30\xe9\x79\x64\xf6\x3a\xcf\xd7\x75\x95\xf4\x89\xd6\x1b\x92\xbd\xb3\xbf\x6e\x70\x16\xa7\xf9\x05\xa7\x39\xd9\x6b\x90\x01\x6c\x7b\xdf\x80\x8c\x1d\x46\xfa\x6f\xb4\x5f\xf1\xdd\xfc\xc5\x44\xe2\x3e\x1e\x9c\x63\x5d\x94\xd5\xf1\xa1\x7e\xba\x29\x4e\xa5\x2c\xea\x46\x2a\x59\xa8\x52\x57\x95\xaa\x6f\x73\xa6\xc8\x43\xf7\xd8\x8f\x18\x22\x71\xbb\xf0\x5b\x63\xbb\xe3\xc1\xa7\x78\x9e\x2c\xb5\xe7\xb0\xd0\xc1\xb8\xbe\x7d\x76\x3d\x1a\xb1\xef\xdf\xd2\x25\xd9\x96\xfd\x04\x00\x00\xff\xff\xf0\x39\x29\xc8\x1a\x01\x00\x00")

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

	info := bindataFileInfo{name: "setting/app.mode.json.example", size: 282, mode: os.FileMode(420), modTime: time.Unix(1609257005, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _settingAppReleaseJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\xcd\x4e\xc3\x30\x10\x84\xef\x79\x0a\xcb\x87\x0a\x50\x49\x9c\xf4\x47\xae\xa5\x0a\xee\x80\xe8\xa1\x3c\x80\xeb\x6c\x13\x4b\x76\x6c\xd6\x2e\x6d\x41\x7d\x77\xe4\xa8\x94\x50\x8e\x3b\xfe\x34\x3b\x3b\xfe\xca\x08\xa1\x9d\xb4\x40\x05\xa1\x8d\xf3\x2d\xe0\xbd\x81\xc6\x11\xe9\x3d\x1d\xa7\xc7\x00\xf8\x01\x48\x05\x49\x28\x21\xd4\x3b\x8c\x09\x16\x9c\x2d\x58\x8f\x10\x42\xa5\xf7\x6f\x68\x92\xdc\xc6\xe8\x45\x51\x18\xa7\xa4\x69\x5d\x88\x82\x33\x7e\xc1\x6a\x67\xa5\xee\x5e\xa4\xee\x12\x7a\xc5\x64\x84\x9c\xfa\x8d\xb5\x8c\x72\x23\x03\xfc\xee\xb4\xc7\xf0\x6e\x2e\x63\x42\x42\xef\x80\xce\x45\x51\x56\x93\xe9\x6c\xfe\x78\x53\x2e\xaa\xbc\x9c\xf3\x9c\xe5\x25\x63\x62\x32\x61\xf3\xdb\x22\x42\x88\xf5\xe6\x41\xb5\x12\x03\xc4\xe5\x2e\x6e\xb9\xdd\x4c\x47\x3e\x8d\x6b\x6d\x61\xb9\xc6\x1d\x8c\x8c\x53\xcb\xe7\x14\x86\xf6\xfe\xa7\x4b\x12\x25\x55\x3b\x88\xb1\x45\x80\xbf\x52\xea\x47\x7f\xf6\xe5\xcd\x2a\x5e\x71\xce\x7e\x6e\x25\x84\x36\x6a\x05\xa8\xa0\xeb\xeb\xaa\xd8\xd9\xfc\x5c\x45\xe3\xfe\x39\xa9\x5d\x88\xce\x3e\xc1\x71\x85\xb0\xd5\x87\xfe\x40\x30\x90\x9a\xb8\xca\x65\x75\x5d\x1b\xd8\x4b\x1c\x84\x53\x0e\xc3\xd0\x4d\x1a\xe3\xf6\xaf\xa8\x1b\xdd\x25\x9d\xde\x8d\xd3\xd7\x04\x51\x14\x70\x90\xd6\x1b\xc8\x95\xb3\x03\xe3\xec\x94\x7d\x07\x00\x00\xff\xff\xc1\xfa\x76\x70\x0d\x02\x00\x00")

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

	info := bindataFileInfo{name: "setting/app.release.json", size: 525, mode: os.FileMode(420), modTime: time.Unix(1612185114, 0)}
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
