// Code generated by go-bindata.
// sources:
// config.yml
// assets/css/styles.css
// assets/index.html
// assets/js/translate.js
// DO NOT EDIT!

package binded

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

var _configYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\xc9\x4e\x1b\x4d\x10\xbe\xfb\x29\x5a\x3e\xfc\x47\x64\xf8\xb3\x90\xb9\x21\x11\x14\x21\x39\x1c\x40\x42\x39\x59\x0d\x6e\x2f\xc1\xf4\x58\xe3\x19\x90\x6f\x59\x94\xf8\x10\x45\x39\xa1\x2c\x8a\x64\x45\x22\x3d\x04\x8c\x41\x1c\x12\x47\x8a\xe0\x55\xea\x51\xa2\xaa\xae\x9e\x8d\x31\xb9\xcc\xd4\xf2\x55\x75\x55\x75\xd7\x27\xfb\x7d\xaf\x22\x44\xdf\x0f\x42\x4f\x2c\xd7\x6a\xcb\x15\x21\x0e\x54\x30\xe8\xfa\xda\x13\xd5\xa5\x85\xc5\x85\x5a\xb5\xd2\x93\xba\x1d\xc9\xb6\x42\xa4\xd2\xed\x5e\x77\xd0\x41\x51\x08\x75\xa0\x82\xa1\x27\xaa\x8f\xf1\x2f\xaa\xa9\xad\xb1\xef\xeb\xb0\xd3\xf0\x5b\x89\xb3\x8e\x06\xb1\xd1\xca\xa1\x9a\x12\xa3\x85\x45\xac\xca\xa1\xf5\xf9\x5a\x0b\x34\x6f\x68\xe1\x0c\xd9\x7c\x68\x2f\x64\x93\x61\x83\x8f\xed\xea\x28\x54\x08\x5a\x09\x39\x6d\x9d\x4c\xa5\xb8\x86\xdf\x6a\x74\xfc\x28\x28\xc1\x63\xee\x27\x7e\x14\xf0\x01\x4d\x39\x44\x70\xd8\x51\x8d\x43\xa5\xf6\x30\x60\x55\x0e\x11\xb4\xd5\x51\x62\x5b\xa9\xbd\xdb\x40\x2a\xb9\x80\xb4\x75\x5b\xa8\x6b\x89\xc0\x43\x25\xa9\x8c\xa4\x31\x44\x3f\x53\x32\x48\x3a\x14\x5c\x25\xeb\xba\x89\xfa\x7f\xac\x86\x3e\x6a\xa1\xcf\x6a\x2b\xf0\xf7\xd1\x80\xff\xb4\x30\x57\x0b\x1f\x9f\x8c\x8a\x3b\xb6\x66\x37\x8f\x4c\xf3\x83\x48\xdb\xe8\x4d\x12\x92\xea\xad\xb1\xee\xa7\xc6\x30\x52\x03\x6b\xdd\xb2\x92\x35\x1f\xaa\xa6\x76\x8e\x6d\x27\x73\x44\x27\x0a\x5c\x08\x8b\xae\x85\xae\x35\xaf\x91\xc0\x95\xc8\x30\x0a\xb8\x16\x16\xad\xe3\xb9\xd4\x91\xa4\xa7\xb8\x6e\x25\x4e\xa2\x76\x02\xb6\xaf\xb1\xc8\xd5\xcb\x60\x17\xef\xa6\x8e\x7f\x9e\x68\x3f\xe8\xf6\x3c\x51\x5d\xc1\xbf\x43\x51\x83\xc9\x21\x91\xc6\x79\xad\x47\x5a\x39\x43\x8f\x8e\x8c\x7a\x8c\x90\x51\x3b\x1a\x84\x98\x84\x04\x2e\x5a\xf5\x43\xb5\xbf\xa3\x70\xac\x9b\x4e\xe6\x87\xbd\x1b\xfa\xd6\xb1\x61\x25\x6b\xd6\xfe\x81\x0b\x78\xca\x22\x5f\xa2\xda\x75\x8e\x55\x16\xd1\xb1\x23\x75\xbb\x27\xf3\x2b\x09\xe6\x14\xe2\xf7\x60\x2e\xc1\x1c\x83\xb9\x99\xb7\x9c\x45\x18\x98\x29\x98\x6b\x30\x33\xaf\x64\x4f\x0b\x60\x30\x86\xbe\x27\xf9\xb5\x2d\xdd\xd9\x62\xde\x92\x8d\x2d\x2d\x05\xd3\xd3\x77\x0c\xf1\xa8\x7c\xdb\xb3\x5b\xfc\x8f\x1c\x02\xcc\x27\x54\x10\x32\x06\x73\xed\xcd\x5f\x6f\x30\x33\x2a\xd8\xe6\x3b\x26\xf9\x37\x65\x39\x12\x49\xe3\x77\x2c\x7d\xd2\xf0\x9c\x98\xd2\xed\x07\x33\xa1\xa0\x4b\x3a\xf4\x57\x12\xca\xa9\xf2\x5c\x40\xb9\xa7\x60\x2e\xc0\xfc\xc9\xb3\x02\x98\x0f\x94\xe9\x55\x8e\x1d\xc0\x7c\x87\x78\x04\xe6\x28\x9d\x24\xf3\x44\x8e\x21\x92\x42\xbd\x22\x4f\x14\x86\xe9\xe5\x09\xa3\x38\xda\x02\x77\xd0\x9d\x4c\x28\x92\x7b\x2c\xf0\x08\xb6\x13\xbf\xa3\x33\x72\x80\x94\x53\xc8\xf7\x99\x4e\xf8\x08\xe6\xaa\x80\xcb\x92\x0c\xba\xe2\x97\x60\xe2\x62\xae\x94\x6d\x08\xf2\x9a\xee\x74\x46\x29\x4f\xdd\xa3\xce\x45\x24\x34\x04\xe6\x27\xa5\x3c\x72\xef\x2b\x07\xcb\x10\x13\x02\x79\x44\x39\x48\x4a\x51\x60\xbe\x92\xe3\x84\x12\xe2\xf5\xb9\x4b\x7f\x81\x35\xc4\xa3\x5b\xe4\x05\xe6\x8c\xae\x6e\xc2\x87\xdf\x19\xe7\xb8\xcd\x3d\x1b\xfb\x98\xbe\x24\x7e\x47\x74\xf4\x4a\x92\x75\xb9\xc1\x89\xa6\x29\xdc\xbc\x9d\x85\xb9\x0f\x4b\xc7\xc3\x4f\x32\x8e\xde\x30\xe3\xb8\xa2\x23\x47\x22\x5b\x50\xc2\x8a\x60\xde\xd2\xd5\xcd\xf8\x95\x38\x40\x96\x21\xc9\x3b\x72\x85\x8d\x49\x9e\x92\x3c\xa1\x56\x46\x45\xe6\x04\xf3\xc6\xdd\xca\x98\x1e\x50\x0e\x97\xa1\x52\xba\x95\xf3\x79\x09\x33\xd4\x0a\xe6\x1b\xcd\x63\x36\x0f\x5b\x43\x50\x6c\xac\xb2\x48\x4a\x6c\x95\x25\x52\x98\x10\xff\x27\xe5\x87\x55\xee\x91\x72\x6a\x95\xfb\xa4\x9c\x59\xe5\x01\x29\x13\xab\x3c\x24\xe5\xdc\x2a\xcb\xa4\x4c\xad\xf2\x88\x94\x8b\x6a\xe5\x6f\x00\x00\x00\xff\xff\x06\xc9\x00\x2f\xb3\x09\x00\x00")

func configYmlBytes() ([]byte, error) {
	return bindataRead(
		_configYml,
		"config.yml",
	)
}

func configYml() (*asset, error) {
	bytes, err := configYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.yml", size: 2483, mode: os.FileMode(420), modTime: time.Unix(1559226509, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCssStylesCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x52\x4d\xae\xd3\x30\x10\xde\xe7\x14\xa3\x87\xd8\x51\x94\x34\x94\x85\xbb\xe3\x10\x6c\xd1\x38\x99\x24\x56\xdd\x99\xc8\x9e\x40\x1f\x88\xbb\x23\xc7\xa1\xb8\xa8\x4f\x02\x41\x97\x1e\xcf\xcf\xf7\x67\xa5\x7f\x86\x6f\x15\xc0\x19\xc3\xe8\xd8\xc0\xa1\x9e\x2f\xc7\x0a\xa0\x13\x2f\xc1\xc0\xab\xb6\x6d\xd3\x73\x10\xd6\xdd\x80\x67\xe7\x9f\x0d\x7c\xa4\xd0\x23\xe3\x1b\x88\xc8\x71\x17\x29\xb8\xe1\xda\x13\xdd\x57\x32\xd0\xbc\xcf\x5b\x2c\x76\xa7\x31\xc8\xc2\xfd\x6e\x5b\xf8\xc1\x63\x77\x3a\x56\xdf\xab\xa9\x49\x67\x95\x2e\xba\x43\xef\x46\x36\xd0\x11\x2b\x85\xf4\x57\xbd\xa5\xcb\x1c\x28\x46\x27\xfc\xc9\x2e\xd6\x7a\x5a\x41\xce\x12\x9d\x3a\x61\x03\x81\x3c\xaa\xfb\x4c\xe9\xc8\x17\xd7\xeb\x64\xa0\xdd\x6f\xd0\x27\x72\xe3\xa4\x06\x9a\xfd\x21\x15\xd2\x1c\xf6\xbd\xe3\xd1\x40\xb3\xb5\x58\x09\x3d\x85\x5d\xc0\xde\x2d\xd1\xc0\xa1\xac\x1a\x88\xe2\x5d\x0f\x3e\x6d\xb1\x7e\xa1\x5b\x22\xa6\xfc\xb8\x0b\xd5\x58\x1a\x24\x64\xc4\x9d\xb0\x12\xab\x79\x7a\x3a\xde\xe0\x47\x1b\xc5\x2f\xba\xee\x0e\x1b\xdc\xba\x7e\x9d\x9e\x2a\xb3\x81\xfd\x26\xe0\xc6\xad\x2e\x79\xd5\x2b\xa7\x8d\xc2\xda\xdd\xb4\xf3\x65\x43\xad\x01\x39\xce\x18\x88\xb5\x24\x9a\x27\xd3\xd6\xbb\xec\x72\x97\x15\x55\x39\xbf\xbc\xee\xbe\x31\xc9\x42\x0c\x84\xbf\x39\xe4\x78\xa2\xe0\xb4\x20\x51\x54\x7e\x52\x29\x4a\xa5\xc2\xd7\xf2\x6d\xaa\xde\xad\xee\x55\x45\x5a\x57\x5d\xae\xee\x96\xc2\x18\x60\xe1\x2c\x2f\xe5\xe9\xfc\xbe\xcf\x21\xce\xc8\x8f\xc0\x9f\xae\x29\xfa\xd3\x1f\x67\xb8\x7e\x31\xc3\x83\x17\x54\x93\xc3\x92\x45\x78\x68\xaa\x7f\xa1\xfe\xdb\x38\x7b\x1a\x1e\x9c\xe6\x7c\xe1\x3f\x84\xb9\xb0\xe6\xc1\x29\xfe\x97\xc0\xfe\x08\x00\x00\xff\xff\xef\x51\x4b\x97\xa4\x05\x00\x00")

func assetsCssStylesCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsCssStylesCss,
		"assets/css/styles.css",
	)
}

func assetsCssStylesCss() (*asset, error) {
	bytes, err := assetsCssStylesCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/css/styles.css", size: 1444, mode: os.FileMode(420), modTime: time.Unix(1558602511, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\x4d\x8f\xd3\x30\x10\xbd\xf7\x57\x0c\x3e\x22\x1a\xab\xda\x0b\x5a\x39\xd1\xa2\xb2\x17\x2e\x20\xb5\x17\x4e\x68\xe2\x0c\x8d\x83\x63\x07\xcf\x34\xdb\x68\xd9\xff\x8e\x12\xa7\x68\x97\x0f\xa1\x1c\x92\x79\x6f\xe6\xbd\xc9\xb3\xcd\xab\xf7\x1f\xf7\xc7\xcf\x9f\xee\xa1\x95\xde\x57\x9b\x8d\x99\xdf\xe0\x31\x9c\x4a\x45\x41\x55\x9b\x0d\x80\x69\x09\x9b\x6a\x03\x00\x60\x7a\x12\x04\xdb\x62\x62\x92\x52\x9d\xe5\xeb\xf6\xad\xd2\xcf\xb9\x80\x3d\x95\x6a\x74\xf4\x30\xc4\x24\x0a\x6c\x0c\x42\x41\x4a\xf5\xe0\x1a\x69\xcb\x86\x46\x67\x69\xbb\x14\x6f\xc0\x05\x27\x0e\xfd\x96\x2d\x7a\x2a\x77\xd9\x0e\xc0\x88\x13\x4f\xd5\x3e\xc5\x20\xe8\xbf\x19\x9d\xeb\xcc\xb1\x4d\x6e\x10\xe0\x64\x4b\xd5\x8a\x0c\x7c\xab\x35\x76\x78\x29\x4e\x31\x9e\x3c\xe1\xe0\xb8\xb0\xb1\x5f\x30\xed\x5d\xcd\xba\xfb\x7e\xa6\x34\xe9\x9b\xe2\xa6\xd8\xad\x45\xd1\xbb\x50\x74\xac\x2a\xa3\xb3\xde\x4b\x71\x99\x06\x2a\x95\xd0\x45\x74\x87\x23\x66\x54\x55\x77\x8f\x8f\x50\xbc\x63\x26\xe1\xe2\x98\x30\xb0\x47\xa1\x0f\x07\xf8\x01\x1d\xc3\xd3\xd3\xdd\xef\x62\x32\x79\x7a\x31\x74\x98\x11\xde\x1f\xe6\x11\xcb\xd7\x99\xa5\x6d\xce\x59\xe7\xa0\xe7\xcf\x3a\x36\xd3\x2a\xd3\xee\x96\x24\x8e\x4b\x12\xed\x6e\x45\x1b\x37\x82\xf5\xc8\x5c\x2a\xba\x0c\x89\x98\x5d\x0c\x5f\xea\x73\x5d\x7b\x52\xb9\x67\x5e\x61\xc0\x50\xdd\xff\xe2\x6f\x8d\x5e\x90\x2b\x3d\xff\x21\x26\x42\x70\xcd\x73\x19\x55\xbd\x86\xf5\x31\xfa\xda\xb3\xfa\xea\xc6\x8d\x7f\xae\x30\x1f\xd3\x5f\xbd\xaf\x31\xfd\xcf\x3c\x11\x9f\xbd\x28\x48\x84\x4d\x0c\x7e\xaa\xfe\x61\xbc\x84\xa3\x73\x3a\x46\xe7\x4b\xfb\x33\x00\x00\xff\xff\x3f\x0d\x7e\x7d\xc5\x02\x00\x00")

func assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsIndexHtml,
		"assets/index.html",
	)
}

func assetsIndexHtml() (*asset, error) {
	bytes, err := assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/index.html", size: 709, mode: os.FileMode(420), modTime: time.Unix(1558602511, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsTranslateJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\xbd\x6e\x2a\x31\x10\x85\xfb\x7d\x8a\xd1\x5e\x24\x6c\xc9\xf2\x52\x22\x10\xfd\x55\x8a\x44\x0a\x14\x69\x9d\x65\x08\x0e\xce\x78\x33\x1e\x13\x10\xe2\xdd\xa3\xfd\x01\x14\x85\x2a\x8d\x5d\x9c\x99\xef\x3b\xf6\xde\x31\x64\x0e\xb0\x80\x72\x2b\xd2\xcc\xaa\x2a\xc4\xda\x85\x6d\x4c\x32\x9b\x4e\x26\xd3\xaa\xe6\x48\xe2\xc2\xae\x12\x76\x94\x82\x13\x2c\xe7\x45\x31\x52\xeb\x58\xe7\x0f\x24\xd1\x96\xd1\xad\x8f\x6a\x93\xa9\x16\x1f\x49\xe9\x53\x01\x30\x52\xe5\x3f\x3c\x34\x8c\x29\xf9\x48\xa5\xb6\x99\x5e\x3d\xad\x95\x9e\xdf\x0b\xbb\xa8\xdc\xe1\x31\x37\xa5\xf9\x09\x02\x68\x1b\xb6\xd3\xb0\xf8\xbd\xb8\x77\xa1\x47\x02\x8c\x6c\x13\x93\xa8\xf6\x2d\x06\x1e\x96\x4f\x8f\x36\x09\x7b\x7a\xf3\x9b\xa3\x3a\xdd\x96\x66\x1d\xeb\xac\xc1\xc0\x55\xc4\x98\x9a\x48\x09\xc1\x40\x12\x27\x39\x0d\xe6\xbe\x29\x63\xca\x41\x06\xd9\x78\x3c\xe8\xee\x64\x17\xcc\x30\x71\xd6\x76\xe3\x7c\xb8\x7d\xcc\xfb\xe7\xcb\xff\x67\x30\x20\x78\x90\x65\xe7\x01\x03\xc8\x1c\x79\xb5\xe5\xf8\x45\x7f\xb3\x76\x54\x7b\x71\xaf\xf0\x20\x57\x7f\x7b\xf7\x27\xa3\x64\xa6\x79\x71\xd6\xc5\x77\x00\x00\x00\xff\xff\x39\xd9\xa8\x21\xf1\x01\x00\x00")

func assetsJsTranslateJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsJsTranslateJs,
		"assets/js/translate.js",
	)
}

func assetsJsTranslateJs() (*asset, error) {
	bytes, err := assetsJsTranslateJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/js/translate.js", size: 497, mode: os.FileMode(420), modTime: time.Unix(1558602511, 0)}
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
	"config.yml": configYml,
	"assets/css/styles.css": assetsCssStylesCss,
	"assets/index.html": assetsIndexHtml,
	"assets/js/translate.js": assetsJsTranslateJs,
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
	"assets": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"styles.css": &bintree{assetsCssStylesCss, map[string]*bintree{}},
		}},
		"index.html": &bintree{assetsIndexHtml, map[string]*bintree{}},
		"js": &bintree{nil, map[string]*bintree{
			"translate.js": &bintree{assetsJsTranslateJs, map[string]*bintree{}},
		}},
	}},
	"config.yml": &bintree{configYml, map[string]*bintree{}},
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

