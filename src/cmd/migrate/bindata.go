// Code generated by go-bindata.
// sources:
// migrations/20150715080601-create-schema.sql
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _migrations20150715080601CreateSchemaSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x58\x5b\x6f\xe2\x38\x14\x7e\xe7\x57\xf8\x6d\x8a\x76\x2b\x85\x5e\xa4\x99\xad\xfa\x40\xa9\x3b\x8b\x96\x09\xb3\x21\x3c\xcc\x93\x63\x12\x03\x56\x89\x9d\x75\x1c\x66\xf8\xf7\xeb\x18\x72\x71\xb8\x85\x64\x34\x62\xaa\xa8\x52\x52\x9f\xdb\x77\x3e\x7f\xe7\xa8\xb7\xb7\xe0\x8f\x90\x2e\x04\x96\x04\x4c\xa3\x4e\xf9\x75\x22\xd5\xef\x90\x30\xf9\x42\x16\x94\x1d\xfe\x13\x64\x41\x67\xe0\xc0\xbe\x0b\x81\xdb\x7f\x19\x41\xe0\x45\x38\x8e\xbf\x73\x11\x20\x41\x42\xca\x02\x22\x62\x0f\xdc\x74\x00\xf0\x48\x88\xe9\xca\x03\x6b\x2c\xfc\x25\x16\x37\xbd\x4f\xbd\x2e\x18\x8c\x47\xa3\xd4\x36\x91\xf3\x8f\xe1\xec\x01\x25\x8c\xfa\x3c\x20\xc8\xa7\xc0\x1e\xbb\xc0\x9e\x8e\x46\x7f\xa6\xb6\x92\xbf\x13\xd6\xd0\xd6\x17\x44\x25\x1b\x20\x2c\x3d\x20\x69\x48\x62\x89\xc3\x28\x3f\x02\x5e\xe1\x5b\x7f\x3a\x72\xc1\x07\x4b\xfd\xdc\xea\x07\x58\xd6\x5f\xfa\xf9\x90\x3a\xf8\x07\x7e\x3b\x54\x15\xd2\xf5\xa0\xf4\xe5\x87\xaa\x70\x57\x5e\xf7\x94\x85\xae\xa2\xb0\xd8\x16\xd5\xed\x74\x01\xb4\x3f\x0f\x6d\xf8\x3c\x64\x8c\xbf\xbe\xe4\x19\x0d\xfe\xee\x3b\x13\xe8\x3e\xa7\x05\x66\xd5\xea\x97\x52\xa9\x4f\x9d\x0a\xfc\x49\x5c\x20\x4e\x03\x0f\x50\x26\x6f\x7a\x56\x17\x24\x2c\xa6\x0b\x46\x82\xa2\xf0\xfe\xd4\x1d\xa3\xa1\xad\xcc\xbf\x40\xdb\xd5\x50\xa5\xc6\x0c\x87\xa4\x21\xd2\x95\x0e\xdf\x5b\xd6\x49\xdb\xac\xce\xdc\x3e\x83\xac\x70\xf1\x78\xc6\x85\x11\x3e\x54\x1f\x15\x3b\xb9\x48\xfb\xcc\x36\xba\xf2\xae\x79\x44\xf5\x82\x84\x33\x22\x50\x85\x50\xe7\xe2\xec\xa5\xda\x96\x54\x5e\x12\x05\xad\x1c\x7c\x75\x86\x5f\xfa\xce\x37\xcd\xb5\x9b\xb4\xd3\x9a\x79\x53\x7b\xf8\xef\x14\x6e\x09\xa8\x99\x80\xb2\x96\xa6\xd5\xfc\x97\x90\x94\x77\x79\x97\xf7\xa9\x67\x72\xe2\xb9\x67\xb5\x23\x23\x0d\x94\x44\xd0\x39\x6d\x41\xc9\x80\xfb\x31\x91\xc8\x64\xe5\xfd\xc3\x05\xb4\xd8\x79\x98\xd3\x15\x69\xef\x25\x5a\x61\x39\xe7\x22\x6c\xe7\x65\x96\xb0\x60\xd5\x32\x93\xb5\x02\x95\x72\xd6\xd0\x49\x84\x17\x04\x45\x58\x2e\x3d\xb0\xe2\x6c\x21\xc9\x0f\x79\x99\xad\xa4\xb2\x71\x05\x4b\x29\x05\xf6\xdf\x51\xcc\x13\xe1\x93\x06\x19\xcc\x30\x53\xbc\x41\x73\xc1\x43\x14\x25\xb3\x15\xf5\x8f\x5f\xf9\x2b\xbc\xa9\xe7\xae\xdd\xdd\x5d\xbb\x6b\xa7\x2e\x9d\xa0\xa4\xf1\x95\x6b\xd3\xda\x19\x0f\x36\x4d\x1a\xaa\xcc\xd4\xa0\x4c\xc7\x24\x09\x1a\xd8\xcb\x4d\x54\xca\xf8\xee\xf1\xb1\x7e\xc6\x85\x46\xa1\x93\x48\xe9\xc3\x98\xf9\xcb\x74\xbe\xe4\x91\xac\x4b\x06\x54\x2a\xbc\x35\x82\x9c\x63\xb4\x1a\x62\x7c\x5d\x97\xfe\xb1\xcf\x05\xd9\x45\xfc\x0d\xae\x46\xbe\x3e\xed\x38\x8c\x8c\xf6\x20\x25\xbd\x44\x21\xe6\xe9\xd3\xe5\xbe\xed\x1b\xee\xc0\x2e\x9b\x64\xf8\xeb\xc3\x83\xb1\x3d\x71\x9d\xfe\xd0\x76\xcf\x06\x7b\x1b\x3b\x70\xf8\xd9\xce\xd3\x34\x02\x03\x07\xbe\x41\x07\xda\x03\x38\xa9\x0e\x3c\xef\x54\xac\xbd\xfc\xcc\x28\x79\xae\x86\xff\x6c\xbb\xab\x25\x23\xbd\x8f\xed\x64\x44\x12\x1c\x36\x16\x91\x16\x6b\x24\xf6\x7d\x12\xc7\xe8\x9d\x6c\x1a\x2e\x82\x57\xca\xeb\xf2\x72\xa6\xb1\x45\x95\xc5\xac\xde\x52\x76\xdf\x7e\x38\x6c\x50\x1a\xbf\x69\x6b\xb7\x1e\xce\xeb\x58\x1a\xa3\xc6\x31\x43\xce\xb6\x79\xfd\x46\xb3\xdc\xd0\x9d\x2d\xae\x28\x03\xa8\xac\x3d\x39\x68\x07\x2d\x76\x50\x95\x0d\x32\xf4\x0e\x0a\xc8\xd1\x40\xa6\x88\x14\x41\x0d\x15\x29\xf6\x83\xa3\x0a\x75\x2c\x2f\xd3\x7f\x9e\xa3\xe1\x3e\xd3\x8d\x5a\x22\xf5\xa9\xbd\x46\x69\x29\x6d\xbc\xec\xd4\x23\x69\xcd\xd1\x2d\xf8\xaa\xe9\x1e\x72\xcd\xdc\xce\x51\x3e\x4b\xd4\xca\xf1\x4b\x86\xf0\x89\x20\x17\xb3\xee\xb8\xef\x5f\x30\x76\x1f\xdb\x51\x7a\xcd\x65\x8b\xdd\x5d\x2f\xc2\x99\x7e\x3e\x54\xf4\xb3\xae\x72\xd7\x64\xfb\x35\x53\x56\xa3\x58\x57\x89\xb7\x87\x2f\x21\xeb\x31\xf7\x3f\x43\x7f\x8f\x64\xf3\xf3\x89\xfa\x70\x21\x51\xcb\xff\x01\x7e\xe5\xdf\x59\xe7\xd5\x19\x7f\xdd\xd1\x56\x27\xfd\x54\xfe\x92\xdf\x39\xe3\x6b\x31\x5c\xf6\x0e\x9b\xe6\xa5\x55\x7a\xcf\x01\xad\x84\xd2\x65\x3f\x75\xfe\x0f\x00\x00\xff\xff\x49\x5c\xf0\xb3\xcc\x16\x00\x00")

func migrations20150715080601CreateSchemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20150715080601CreateSchemaSql,
		"migrations/20150715080601-create-schema.sql",
	)
}

func migrations20150715080601CreateSchemaSql() (*asset, error) {
	bytes, err := migrations20150715080601CreateSchemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20150715080601-create-schema.sql", size: 5836, mode: os.FileMode(420), modTime: time.Unix(1436977644, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"migrations/20150715080601-create-schema.sql": migrations20150715080601CreateSchemaSql,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"migrations": &bintree{nil, map[string]*bintree{
		"20150715080601-create-schema.sql": &bintree{migrations20150715080601CreateSchemaSql, map[string]*bintree{
		}},
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
