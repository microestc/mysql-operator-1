// Code generated by go-bindata.
// sources:
// artifacts/backup-crd.yaml
// artifacts/backup-cronjob.yaml
// artifacts/backup-pvc.yaml
// artifacts/mysql-configmap.yaml
// artifacts/mysql-crd.yaml
// artifacts/mysql-service-read.yaml
// artifacts/mysql-service.yaml
// artifacts/mysql-statefulset.yaml
// DO NOT EDIT!

package artifacts

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

var _artifactsBackupCrdYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xb1\x4e\x04\x31\x0c\x44\xfb\x7c\x85\xbf\x20\x68\x3b\x94\x12\xe8\x38\x90\x00\x89\xde\x9b\x35\x8b\xb5\x49\x1c\x62\x67\xc5\xfd\x3d\xda\x9c\x28\x8e\x86\x72\x3c\x1e\xbd\xb1\xb1\xf2\x3b\x35\x65\x29\x01\xb0\x32\x7d\x1b\x95\x43\xa9\xdf\x6e\xd5\xb3\xdc\xec\xd3\x4c\x86\x93\xdb\xb8\x2c\x01\xee\xbb\x9a\xe4\x57\x52\xe9\x2d\xd2\x03\x7d\x70\x61\x63\x29\x2e\x93\xe1\x82\x86\xc1\x01\x14\xcc\x14\x20\x9f\xf5\x2b\xcd\x18\xb7\x5e\xd5\xc7\xe6\x87\x96\x4a\x0d\x4d\x9a\x5f\x9b\x25\xbf\xb2\x7d\xf6\xd9\x47\xc9\x4e\x2b\xc5\x23\xbb\x36\xe9\x35\xc0\xbf\xfb\x17\x8a\x1e\x11\x80\x4b\xb7\xa7\xf3\xdb\xcb\xe9\x6e\x00\xc7\x34\xb1\xda\xe3\x5f\xe7\xc4\x6a\xc3\x55\x2e\x6b\x4f\xd8\xae\x8a\x0e\xa7\xa6\xde\x30\x5d\x1f\xe0\x00\x34\x4a\xa5\x00\xcf\x07\xb6\x62\xa4\xc5\x01\xec\xbf\xaf\xdb\x27\xf7\x13\x00\x00\xff\xff\xe5\xf5\xf0\x5c\x4a\x01\x00\x00")

func artifactsBackupCrdYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsBackupCrdYaml,
		"artifacts/backup-crd.yaml",
	)
}

func artifactsBackupCrdYaml() (*asset, error) {
	bytes, err := artifactsBackupCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/backup-crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsBackupCronjobYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xc1\x6e\xd3\x40\x10\xbd\xfb\x2b\x46\x81\xa8\x07\xb4\x76\xa2\x1e\x90\x2c\xf5\x00\x29\x08\x10\x29\x81\x96\x4a\xdc\x3a\x5e\x4f\x93\xa5\xde\x5d\x77\x77\x1c\x12\xa5\xf9\x77\xb4\x76\x52\xe2\x6d\x23\x95\x39\xf8\xf0\xe6\xcd\xf8\xcd\x9b\x59\xac\xd5\x35\x39\xaf\xac\xc9\xa1\x40\x96\x8b\x6c\x39\x2e\x88\x71\x9c\xdc\x29\x53\xe6\x30\x71\xd6\x7c\xb1\x45\xa2\x89\xb1\x44\xc6\x3c\x01\x30\xa8\x29\x87\xcd\x66\x9f\xbc\x40\x4d\x90\xb6\xdf\xed\x36\x01\xb0\x7f\x0c\xb9\x1f\x74\x4b\x8e\x8c\x24\x1f\x4a\x00\x04\x74\x0d\xa7\xeb\xcb\xef\x5f\xdf\xa3\xbc\x6b\xea\x16\x07\x38\x94\x20\x5d\xaa\xd7\xfe\xbe\xb2\x35\x39\x64\xeb\xd2\xb9\xe3\x2a\x9d\x2b\x5e\x34\x45\x2a\xad\xce\x96\xe3\x5d\xd5\xa3\x88\x83\x1f\x87\x68\x54\xd9\xc1\x3f\x3f\x9f\x07\xd4\xd7\x24\x83\x02\x2f\x17\x54\x36\x15\xe5\x30\x08\xd9\xcb\x9a\x64\x7a\xa5\xda\xca\x41\x02\xf0\xdb\x16\x57\xa4\xeb\x0a\x99\x3a\xbd\xfb\xba\x10\xdc\xcb\xc4\xd9\x10\x8e\x3c\xa3\xe3\x99\xad\x94\x5c\xe7\xf0\xcd\x7c\x44\x55\x35\x8e\x0e\x28\x4b\x5b\x35\x7a\xef\x46\x17\x62\x37\x45\x71\x68\x47\x17\x75\x70\xc4\x33\x19\xbe\x6e\xeb\x26\x15\x2a\x9d\xf7\x28\x00\x32\x80\x17\x7b\x23\x66\xd7\x93\x78\x13\xf1\x8f\xa4\x35\xb7\xbd\x1e\xa4\x6b\x5e\x9f\x2b\x97\xc3\xe6\x18\x5f\xcd\x85\xc6\xbe\xb8\x0e\x9e\x62\x1d\x0b\xea\xaa\xda\x0d\x26\x3d\x3a\xa3\x32\xe4\x5e\x3a\xbc\xd2\x38\xa7\x1c\xc2\xee\xb3\x15\x3b\xec\x28\x79\x58\x81\xe7\x48\x89\xd6\x68\xca\xbe\x0e\x01\x05\xfa\x45\x04\x0d\x84\x1c\x44\xd0\x43\xa4\xde\x13\x83\xa0\x55\x12\xc1\xaf\x60\x52\x59\x43\x10\xce\x1f\x6e\x9d\xd5\xa0\xd1\x33\xb9\x34\xe2\x9d\xbf\xbb\xfa\x70\x76\x53\x22\x13\x9c\xbc\x19\xfe\x12\x43\x2d\x86\xa5\x18\x7e\x12\xc3\xe9\xc9\x4d\xc4\xd5\x77\xa5\x72\x90\xb5\x4e\x65\xdd\x78\xd9\xc1\x2d\x67\xaf\x37\xa1\xdb\x36\xb6\x57\x22\x83\x10\x8e\xe4\x52\x58\x53\xad\xe1\xf1\x92\x27\x55\x13\x24\xc1\x76\x2b\x46\xe9\x33\x28\x9c\x9e\x8e\xde\xc2\x03\xac\x0a\xcf\x8e\x50\x83\x58\x81\x98\xbc\x40\xc0\x13\x2f\x66\x8e\x6a\x74\x04\xbc\xa0\xdd\xea\x62\x1f\xfe\x6d\x0c\x84\xa8\x77\x74\x21\x18\xdd\x9c\x58\x94\xca\x9d\xfd\xd7\xd8\xdd\xbb\x99\xda\xc6\xb0\x8f\xf7\x7c\xf4\x82\x00\x74\x28\x98\x21\x2f\xf2\xe3\x53\x3e\xdb\xed\xc9\x1b\xe9\xf7\x22\x96\xbb\x7e\x81\x98\x96\x11\xd5\x91\xb7\x8d\x93\xe4\xe3\xa7\x11\x52\xf7\x0d\xf9\x78\x88\x2e\x64\xdd\xe4\x30\x1e\x8d\xf4\x33\x39\x4d\xda\xba\x75\x9b\x9e\xaa\xbf\x01\x00\x00\xff\xff\x9a\x10\x14\x10\xb0\x05\x00\x00")

func artifactsBackupCronjobYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsBackupCronjobYaml,
		"artifacts/backup-cronjob.yaml",
	)
}

func artifactsBackupCronjobYaml() (*asset, error) {
	bytes, err := artifactsBackupCronjobYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/backup-cronjob.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsBackupPvcYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8d\x31\x8a\xc3\x30\x10\x45\x7b\x9d\xe2\x5f\x60\x0d\xdb\xaa\x75\xed\xc5\xac\xc1\xa9\x07\xe9\x13\x44\x2c\xc9\xd1\xc8\x81\x60\x7c\xf7\xa0\x38\x69\x06\xde\xfc\x07\xef\x16\x92\xb7\x18\x59\x34\x68\x65\xaa\x73\x5e\xb6\xc8\x7e\x91\x10\x8d\xac\x61\x6e\x43\x4e\x16\x8f\x5f\x13\x59\xc5\x4b\x15\x6b\x80\x24\x91\x16\xfb\x8e\x71\xee\xff\x24\x12\xdd\xfb\x1e\x87\xd1\x95\xae\x19\xe2\x1c\x55\x87\xec\xa9\x0d\x81\x1f\xfc\x53\xfc\xa5\x84\xca\x41\xd2\xd3\x00\x85\x9a\xb7\xe2\xbe\x42\xe1\x7d\xa3\xd6\x0f\x01\x5a\x73\x91\xeb\x99\xe9\xa6\x95\xae\x9b\xce\x4f\xcb\xbc\x02\x00\x00\xff\xff\x31\xd3\xda\x23\xb8\x00\x00\x00")

func artifactsBackupPvcYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsBackupPvcYaml,
		"artifacts/backup-pvc.yaml",
	)
}

func artifactsBackupPvcYaml() (*asset, error) {
	bytes, err := artifactsBackupPvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/backup-pvc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsMysqlConfigmapYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\xb1\x0a\xc2\x30\x10\xc6\xf1\x3d\x4f\xf1\x81\x73\x05\xd7\x6c\xe2\xec\xea\x22\x0e\xd7\xe6\xda\x06\x93\x4b\xcc\xc5\x42\xc1\x87\x97\x06\xba\x39\x38\x26\xf7\xbb\xff\x51\xf6\x37\x2e\xea\x93\x58\x2c\x27\xf3\xf4\xe2\x2c\x2e\x49\x46\x3f\x5d\x29\x9b\xc8\x95\x1c\x55\xb2\x06\x10\x8a\x6c\x11\x57\x7d\x05\x03\x04\xea\x39\xe8\xf6\x0f\x50\xce\xfb\x60\xc7\x91\xb4\x72\x39\x0e\x32\x5a\x7c\x1a\x3a\xe0\x9c\x73\x58\x51\x67\xaf\x18\xda\x05\x24\x09\x2b\x92\xa0\xce\xbc\x6f\x34\x7b\x6f\x31\xf7\x68\x8f\x90\xa6\xae\xf7\x62\x00\x0d\xb4\xf0\xbf\xcd\x86\xf5\x47\x4f\xdf\x99\x4b\x57\x98\x5c\xb7\x59\xf3\x0d\x00\x00\xff\xff\x1c\x35\x5e\xba\x03\x01\x00\x00")

func artifactsMysqlConfigmapYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsMysqlConfigmapYaml,
		"artifacts/mysql-configmap.yaml",
	)
}

func artifactsMysqlConfigmapYaml() (*asset, error) {
	bytes, err := artifactsMysqlConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/mysql-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsMysqlCrdYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xb1\x4e\x04\x31\x0c\x44\xfb\x7c\x85\xbf\x20\xe8\x3a\x94\xf6\xe8\x38\x90\x00\x89\xde\x97\x35\x8b\x75\x49\x1c\x6c\x67\xc5\xfd\x3d\xda\x5d\x51\x00\x05\xa5\x67\x3c\x7a\x63\x63\xe7\x57\x52\x63\x69\x09\xb0\x33\x7d\x3a\xb5\x75\xb2\x78\xb9\xb5\xc8\x72\xb3\x1c\xce\xe4\x78\x08\x17\x6e\x53\x82\xe3\x30\x97\xfa\x4c\x26\x43\x33\xdd\xd1\x1b\x37\x76\x96\x16\x2a\x39\x4e\xe8\x98\x02\x40\xc3\x4a\x09\xea\xd5\x3e\x4a\x2e\xc3\x9c\xd4\x62\xd6\xb8\x09\xd2\x49\xd1\x45\xe3\xac\x5e\xe2\xcc\xfe\x3e\xce\x31\x4b\x0d\xd6\x29\xaf\xe1\x59\x65\xf4\x04\xff\xee\xef\x18\x5b\x23\x00\x7b\xb9\x87\xeb\xcb\xd3\xe9\xb8\x13\x37\xb9\xb0\xf9\xfd\x1f\xeb\xc4\xe6\x9b\x6d\xdc\xe6\x51\x50\x7f\x76\xdd\xac\x5e\x86\x62\xf9\x75\x44\x00\xb0\x2c\x9d\x12\x3c\xae\xe8\x8e\x99\xa6\x00\xb0\x7c\xff\x6f\x39\x84\xaf\x00\x00\x00\xff\xff\x68\x53\xd4\x69\x4f\x01\x00\x00")

func artifactsMysqlCrdYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsMysqlCrdYaml,
		"artifacts/mysql-crd.yaml",
	)
}

func artifactsMysqlCrdYaml() (*asset, error) {
	bytes, err := artifactsMysqlCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/mysql-crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsMysqlServiceReadYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xcd\x4e\x03\x21\x14\x85\xf7\x3c\xc5\x7d\x81\xa2\x4d\x13\x17\x6c\x75\x63\xa2\x26\xb6\xd1\xfd\x2d\x73\xac\x44\xfe\xbc\x30\x63\x9a\xa6\xef\x6e\xa0\x13\x9d\x4d\x59\xb0\x38\x3f\x9c\x0f\xce\xee\x1d\x52\x5c\x8a\x86\xa6\xb5\xfa\x72\x71\x30\xb4\x83\x4c\xce\x42\x05\x54\x1e\xb8\xb2\x51\x44\x91\x03\x0c\x9d\x4e\xb4\x05\x0f\x73\xe0\x85\x03\x48\xf7\xfb\x7c\x56\x44\x9e\xf7\xf0\xa5\xa5\x89\x38\xe7\x1e\x5f\xd8\xe9\x27\x42\xb6\xf8\x80\x20\x5a\xcc\xb9\x15\x5d\x36\x9f\x8f\xbb\xd7\xa7\x7b\x3f\x96\x0a\xe9\x46\x7b\xe2\x9f\xcd\x8a\x0e\xc7\xf2\xed\x53\x86\x70\x4d\xa2\x0f\x52\xbd\x3e\xb8\xfa\x39\xee\xb5\x4d\xe1\x66\x5a\xcf\xad\x3f\xd0\xc5\x72\x3b\xa3\x1b\x2e\xf2\xdb\xe3\x43\x53\x4b\x86\x6d\x08\x39\x49\xed\x2c\xab\xb9\xda\x77\x7a\xa9\x59\x86\x36\x9b\xdb\x3b\x45\x54\xe0\x61\x6b\x92\x2b\xdf\xfb\x0d\x00\x00\xff\xff\x20\xd2\xb0\x1a\x49\x01\x00\x00")

func artifactsMysqlServiceReadYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsMysqlServiceReadYaml,
		"artifacts/mysql-service-read.yaml",
	)
}

func artifactsMysqlServiceReadYaml() (*asset, error) {
	bytes, err := artifactsMysqlServiceReadYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/mysql-service-read.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsMysqlServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xcd\x4e\x03\x21\x14\x85\xf7\x3c\xc5\x7d\x81\xa2\x4d\x13\x17\x6c\x75\xd3\x44\x1b\xb5\xd1\xfd\x2d\x3d\x56\x22\x7f\x5e\x98\x31\x4d\xd3\x77\x37\x30\xa3\x76\x23\x0b\x16\xe7\x87\xf3\xc1\xd9\xbd\x42\x8a\x4b\xd1\xd0\xb8\x54\x1f\x2e\xee\x0d\x6d\x21\xa3\xb3\x50\x01\x95\xf7\x5c\xd9\x28\xa2\xc8\x01\x86\x4e\xa7\x1f\x73\xc3\x01\xa4\xfb\x7d\x3e\x2b\x22\xcf\x3b\xf8\xd2\x92\x44\x9c\x73\x8f\x5e\xd8\xe9\x2b\x42\x9e\xf1\x06\x41\xb4\x98\x73\x0b\x9a\xf6\x1e\x8e\xdb\xa7\xfb\x5b\x3f\x94\x0a\xe9\x46\x7b\xe2\x8f\xcb\x8a\x0e\xc7\xf2\xe9\x53\x86\x70\x4d\xa2\x0f\x52\xbd\x3e\xb8\xfa\x3e\xec\xb4\x4d\xe1\x6a\x5c\xce\xad\x5f\xc8\x8b\xe5\x76\x06\xb7\x9f\xe4\x97\xf5\x5d\x53\x4b\x86\x6d\x08\x39\x49\xed\x2c\x8b\xb9\xda\x77\x7a\xa9\x59\x86\x56\xab\xeb\x1b\x45\x64\x27\xb6\xf5\xa3\xa1\x4d\x8a\x50\x44\x05\x1e\xb6\x26\xf9\xe7\xc3\xdf\x01\x00\x00\xff\xff\x9b\x05\x34\x42\x57\x01\x00\x00")

func artifactsMysqlServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsMysqlServiceYaml,
		"artifacts/mysql-service.yaml",
	)
}

func artifactsMysqlServiceYaml() (*asset, error) {
	bytes, err := artifactsMysqlServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/mysql-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsMysqlStatefulsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\x6d\x73\xda\x3a\x16\xfe\x9e\x5f\x71\x86\x66\x6e\xc8\x36\x22\xd0\x4e\xb7\xb3\xb4\x74\x27\x97\xd2\x97\xd9\x24\xe4\x02\xb7\xdd\x3b\x29\xcb\x15\xf2\x01\x34\xb1\x25\x47\x92\x49\xd8\x34\xfb\xdb\x77\x64\x19\x23\x1b\x93\xa6\x9d\xb9\x3b\x73\x97\x0f\x89\xd1\xcb\xd1\x73\x1e\x9d\x57\x43\x63\xfe\x09\x95\xe6\x52\xb4\x81\xc6\xb1\x3e\x5e\xb6\xf6\xae\xb8\x08\xda\x30\x34\xd4\xe0\x2c\x09\x87\x68\xf6\x22\x34\x34\xa0\x86\xb6\xf7\x00\x04\x8d\xb0\x0d\x77\x77\xfe\x82\x73\x1a\x21\x34\xd2\xbf\xf7\xf7\x7b\x00\xf2\x46\xa0\x1a\xe0\x0c\x15\x0a\x86\xda\x6e\x03\x20\xe0\x04\x9f\xad\x86\xbf\x9c\x76\xc3\x44\x1b\x54\xe9\x04\x80\x8f\x82\xa9\x46\xb4\xd2\xd7\xa1\x8c\x51\x51\x23\x55\x63\xae\x4c\xd8\x98\x73\xb3\x48\xa6\x0d\x26\x23\x8b\xd0\xed\xca\x91\x78\x27\xdb\x4f\xc2\x03\x37\xfc\xeb\xc7\xb7\x76\x54\xc7\xc8\x2c\x04\x8d\x21\x32\x23\x95\x83\x13\x51\xc3\x16\xa7\x74\x8a\x61\x86\xcf\xc2\x88\xcb\xf2\x34\xaa\x25\x67\x78\x5e\x71\x94\xc2\x38\xe4\x8c\x6a\x37\x3e\x8c\x91\x35\x06\xd9\x90\x5b\x60\x30\x8a\x43\x6a\x30\x3b\xcf\xe3\xd0\x7e\xc2\xc2\xd1\x95\x87\x03\xac\xa1\xdb\x0f\x17\xdc\x74\xa5\x30\x94\x0b\x54\xf9\x46\x92\xd1\x60\x67\x49\x4a\x5c\x2e\x91\x47\x74\x8e\x6d\x48\x07\xdb\x16\x88\x36\xf9\x1c\x93\x51\x44\x45\xb0\x39\x9e\xc0\x94\xea\x85\xf7\xb5\x46\x58\xcd\xfb\xfa\x35\x7f\xb6\xa4\x18\x20\x78\xeb\x8d\x3c\x81\xf7\x28\xec\x7d\xa1\x3b\x2e\xe5\x0d\x15\xe1\x01\xcc\x94\x8c\x20\x96\x01\x48\x15\x70\x41\x43\xe0\x22\xc0\xdb\x86\xb7\xf9\xf2\x12\x7e\x5f\x48\x6d\xac\x22\xbf\x43\xe7\x3f\x40\xea\x97\x4d\xf2\xb7\xf1\xd3\xc3\x7d\x18\x8f\xe1\xeb\x57\xc0\x5b\x6e\xa0\xe5\x6d\xc9\x64\x75\xf6\xef\x7e\x3e\x19\x7e\x98\x0c\x7a\x67\x27\xa3\xee\x87\xcb\xd6\xf8\xde\x5b\x84\x6c\x21\xe1\x32\xc5\x13\x8c\xe1\x0d\x1c\x47\xc2\x1c\x33\x29\x66\x8d\xe0\x38\xc7\xd7\x60\x62\x56\x50\xe4\x24\x08\x80\x0a\x90\xb3\x99\x55\xd3\x48\xa0\x4b\xc9\x03\x50\x98\x6e\x09\x36\x9a\x75\x9a\xb0\xa4\x61\x82\x8d\xf2\x91\x9b\x15\xfb\xf5\x7a\xab\xd9\x84\xa7\xb0\x9f\x01\x3e\x3c\x84\x37\x8f\x05\xd2\x95\xf1\xca\x5a\x85\x92\xb1\xe2\x96\x59\xb7\x03\x66\x3c\x44\xed\x68\xb5\x23\x7c\x4e\x22\x1a\x5b\xa0\x18\xc5\x66\xf5\x96\x2b\x1f\x10\x9f\x59\x7a\xd7\xc7\x03\xc1\x6b\x68\xc2\x78\xfc\x0a\xcc\x02\x85\xb7\x0e\x80\xc5\x1b\x5c\x4e\xe6\x71\x44\xad\x8f\x5a\x60\x05\xc8\xbe\xbe\xa1\xc6\x6f\x49\xd1\x21\x5d\xe2\x43\x42\x66\x3c\xff\xb2\x94\x61\x12\xe1\x99\x4c\x84\xd1\xbe\x6d\x3a\x1b\xb7\x5b\xbd\x7d\x91\x5d\x76\x41\xcd\xa2\xed\x4b\xae\xdc\xe5\xa0\x7c\x63\x6f\x61\x51\xbe\x3b\x94\x02\xab\x1d\xcb\x06\xa6\xe3\x5b\xa3\xe8\x94\xb2\xab\x24\xfe\xc3\x5d\x6c\x78\xc5\x63\x7b\x6f\x0e\x93\xbd\x5a\x1b\x4f\x80\x86\x0a\x69\xb0\xb2\x4e\xa2\x8d\x2e\x39\x16\x09\xe0\x78\x49\xd5\x71\xc8\xa7\xc7\xa9\x12\xee\xaf\xf5\xab\x9f\x7e\x72\x7e\xd5\xfc\x1f\xb8\xe2\x13\x18\xa0\x36\x52\xa1\x83\x9c\x1a\xaf\x63\x0d\xa4\x80\xcc\xce\x7e\xc8\x6c\xef\xee\xec\x62\x81\x59\x04\x7e\xa7\x64\xf4\x73\x2a\xb8\xe1\xfe\xa5\xd1\xb4\x56\xdb\x64\x87\x1f\x03\x04\xb0\xb9\x68\x20\x24\x56\x18\x53\x85\x40\x88\xa1\x6a\x8e\x86\x04\x5c\x75\x32\x82\xdd\xa2\xe3\xdd\x9b\x99\x8c\x57\xc4\x7e\x7b\xf4\xf6\xbb\x3b\x40\x11\x6c\xeb\x50\xb2\x88\x1c\x39\xd4\x0b\xc1\x16\x9a\x87\x45\x5d\xb6\xee\xdd\x73\xc1\x34\xf4\xa4\xe2\x36\xcc\xc4\x0a\x97\x5c\x26\x1a\x62\x2c\xd2\x22\x18\x35\x40\x88\x42\xb6\x24\x52\x84\x2b\x3f\x81\x91\xfd\x7a\x7d\x7d\x83\xa4\x75\x78\xd8\xf0\xe6\xe0\xf9\xf3\xe6\x4b\xf8\x0a\xb7\x53\x6d\x14\xd2\x08\xc8\x2d\x90\x6e\xc9\x54\x0b\x90\x2e\x32\xc6\xad\xb6\x8e\x22\x1f\xc7\x23\x2e\xa7\x5a\xf4\xc3\x01\xc7\x32\xb0\x23\x68\xec\x42\xaa\x93\xa9\x5b\x51\x9c\x78\x44\x08\x43\xc3\x32\x13\x28\x05\xb2\x1f\x30\xf1\xf5\x71\xca\x19\xf9\xae\xc0\xe7\x19\x5c\x95\x0a\x79\x5d\xe3\x1d\xf9\x51\x68\x43\x05\x43\xff\xb4\xb2\x79\xb2\x9d\x55\xca\xe3\x0b\x14\x14\xcb\xed\xfb\x38\xfb\x6d\xf8\xcb\xe9\x64\xd0\xef\x8f\x26\x17\x27\xc3\xe1\xe7\xfe\xe0\xad\x87\x3b\xcd\xc6\x16\x6a\xbb\x60\xec\x1a\x99\x42\xf3\x0f\x5c\x0d\x70\x56\x9c\xf1\x4b\xc8\x54\xcf\x0b\xaa\xf5\x8d\x54\x5b\x9e\x06\x70\x85\xab\x36\xc4\xd9\x74\x3e\x17\x4b\x55\x65\x36\x65\x9b\xc8\xf9\xb8\x90\xca\xb4\xad\xed\xff\xf5\x4f\x65\x80\x0a\xb5\x4c\x54\x5e\xc8\xaf\x07\xaf\x13\xd4\x46\xb7\x4b\xf9\x3f\x69\xc3\x8b\x66\x33\x2a\x8c\x46\x18\x49\xb5\x6a\x43\xeb\xfd\x26\xce\x84\x7c\x89\x02\xb5\xbe\x50\x72\x8a\xbe\x10\xbc\xdd\xd4\xbc\x6b\xfe\x5c\x26\x85\xcb\x9a\xcd\xa0\xb5\xa3\x34\x73\x1e\x41\x2d\x45\x4b\x83\x88\x0b\x88\xb9\x98\x03\x89\xbf\xd4\xf6\xef\x2a\xac\xe4\xfe\x4b\xad\x36\xf6\xf3\x8b\xe0\x86\xd3\xf0\x2d\x86\x74\x35\x44\x26\x45\xa0\xdb\xf0\xdc\x8f\x88\x31\x2a\x2e\x83\x7c\xae\xe5\xcf\x19\x1e\xa1\x4c\x4c\x3e\xf9\xc2\x63\x8a\x06\xfc\x91\x5a\x3d\x81\xee\x02\xd9\x15\xdc\x20\x30\x2a\xd2\x05\x89\x41\xb8\x4e\x50\x71\xd4\x20\x97\xa8\x60\xd4\xbd\x80\xba\xbe\xe2\x31\x11\x68\x6e\xa4\xba\xb2\x5a\x72\x6d\xab\xd3\x52\x48\x7f\x98\x23\x20\x0b\x68\x3d\x7b\xd9\x68\x36\x9a\x8d\xd6\x83\x34\x01\x41\x38\x18\xf6\x4e\x7b\xdd\x11\xb4\x0e\xbe\x4d\xda\x8b\xdd\x9c\x3d\x7b\x80\xb2\x56\x29\x30\x6c\x02\xf8\x77\x56\x59\x7f\xe6\x38\x51\xa1\x74\x55\xb0\x78\xf9\x47\x95\x94\xac\x5c\x1c\xee\x15\x52\xee\x5b\x34\xa8\x22\x2e\x10\xa6\x5c\x84\x72\x0e\xb1\xd4\xdc\x70\x69\x9b\x23\x57\x71\x04\x69\x80\x3a\xb2\xe9\x89\x8a\xd5\x76\x01\x47\x66\x9e\x8e\x93\xb4\x11\x98\x70\x31\x93\xd5\xa5\xdc\x13\xf8\xa7\x51\xd4\xe5\x99\xbc\xaa\x9d\x67\x7d\x65\x00\x14\x62\xaa\xac\x01\x42\xad\xfb\xe1\xe4\xfc\x7d\x0f\xce\x4e\x86\xa3\xde\x00\x46\xfd\x5a\xea\x35\xab\x92\xb4\x29\x32\x9a\x68\x84\x1b\x3c\x50\xae\x44\xb2\xde\x93\x16\x34\xa9\xbf\x71\x6d\xec\x80\xeb\x4f\x8a\xf1\x6a\xb9\x03\x37\x5b\x50\x31\xc7\x89\xab\xb2\x26\x46\x36\xf4\x75\xd8\xe0\x65\x35\x3e\xce\x85\xad\x2c\x3d\x11\x8e\x40\x27\x83\x0b\x30\x0b\xae\x81\x51\x8d\x50\xe7\xe6\x40\x43\xa2\x31\x44\xad\x4b\x2e\xad\xa2\x12\x81\x9e\x94\x42\x0f\x56\x45\xb6\x7f\xe2\x0e\xb6\x3f\x17\x78\x09\xb8\x42\x66\xc2\x95\x23\x28\xab\x80\xe1\x82\x2a\xbd\x75\xff\x45\x98\xb6\x67\xb0\x65\x60\xf5\xe1\x69\x17\xf1\xaf\x7a\xe3\x2f\x7f\x3f\xbc\xbc\x6c\xeb\x98\x32\x6c\x8f\xc7\x4f\xd3\x81\x9d\x3d\x45\xaa\xfb\x37\x15\xcf\xda\xed\x2d\x6b\xc8\x9e\x26\xa7\xfd\xf7\x93\x77\x1f\x4f\x7b\x9d\x83\xed\xde\xe4\xe0\xe8\x4b\xc9\x73\xed\xc7\xdb\x78\xd1\x1f\x96\x5b\x9a\x67\xe3\xfb\x1a\xbc\x79\x84\x0d\xcc\x78\xd1\x91\x5c\x9c\xe7\x33\x1b\xea\x05\x62\x60\x9b\x75\x26\xa3\x38\x44\x83\x40\xb3\xe2\x7d\xba\x02\x6d\xac\x81\x8b\xf9\xfa\xd5\x52\x99\xeb\xfc\x9e\x77\x40\xa8\xbe\x68\x47\xd2\x67\xca\x53\xd1\x33\xa9\x5c\x65\x90\xa2\x98\x22\x38\x37\xab\x53\xc6\x30\x4e\x57\x30\x29\x04\x32\x7b\xb6\x3e\xac\x15\x24\x25\xc2\xf0\x10\x2a\x13\xca\x8e\x7c\x92\xa6\x93\xda\x3a\x9d\xd4\x5e\x41\x20\x41\x87\x88\x31\xb4\xec\xb3\xc0\xbd\x0a\xa8\x1f\x5d\x9e\xe1\xff\x2e\x51\x91\xbd\xf3\x48\xd9\x5a\x9b\x62\xad\xec\x7a\xc2\xf9\x95\x8d\x50\xeb\x40\x9a\x56\xc1\x54\x99\x23\xa0\xc6\x60\x14\x1b\xe7\x7f\xd4\x90\x48\x6a\x43\xa4\x60\xdb\xee\xbf\x8b\xe1\xca\x71\xa9\xf8\xbc\x28\xe0\xbb\x28\x7a\xfd\xba\xd7\x7f\xe7\xed\xdf\xaf\xbf\xde\x79\xcc\xe1\xd1\x5e\x85\xc1\x7e\xe8\x0f\x47\x9d\x03\xbf\xf7\x6a\xfa\xdd\xd6\x41\xe5\xa6\x5f\x87\xbd\x41\xe7\x40\x49\x69\xaa\xe7\xd7\x08\xad\xfb\x54\x01\xaf\xde\xd5\xed\x9f\x9f\xf7\xba\xa3\xc9\xa0\x37\x1a\xfc\xd6\x69\x35\x5f\x79\x8b\x86\xa3\x93\xc1\x08\x86\xa7\x27\x9f\x7a\xfe\x70\x51\xfb\xb2\xf3\x0c\xed\xd5\x01\xcd\xde\xaa\x59\xa3\xd5\xb6\xd9\x70\x91\x41\xc3\xcd\x02\xc5\xba\x16\xc5\xc0\x7a\x91\xed\x51\x0b\xaf\x41\x6c\x65\xb5\xee\x54\x43\xae\x0d\x0a\x20\xe4\x0a\x31\x26\x32\x4e\x9f\xad\x40\xd7\xbe\x12\x12\xd1\x5b\x62\x3d\x40\x77\x5a\xae\x49\x25\x0c\x8a\xa1\xa2\x56\x68\x36\xf3\x87\x34\x4b\x90\x34\xde\x12\xe2\xba\xda\xce\xa6\xbd\x25\x0b\xa9\x4d\xc7\x33\x07\x92\x68\x54\x1d\xcb\xbe\xed\x58\xb3\x82\xa1\xf3\x40\xf9\xfa\x7f\xdd\x31\xb4\x76\x76\x0c\xcd\xe6\xd9\xba\x67\x70\x8a\x6f\xf5\x94\x05\x7c\xeb\x77\xa0\x6d\xb8\xbb\xaf\x58\x57\x7a\x0f\xe8\x86\xce\x68\xec\x43\xda\xee\xe1\xbe\xbb\x01\xdf\xd5\x7e\xc7\xa8\x74\x6a\x80\xe6\x53\xaa\x4b\x37\xa4\xbc\x50\x85\x32\x3b\x70\x5e\x2c\x36\xab\x8f\xcb\xcf\xf2\x9b\xef\xe5\x46\xea\x28\xfb\xf1\x21\xa5\x8b\x6c\xfd\xfa\x50\x32\x16\xff\x77\x06\x9b\x09\xb4\x3e\x93\x01\x6a\xdb\x4d\x0c\x90\x06\x9f\x15\x37\xd8\x17\x0c\xf3\x7e\xa0\xe2\x82\xab\xae\xd7\xaa\x9f\x56\xf1\xb9\x36\x43\x37\x62\xf1\xfe\x37\x00\x00\xff\xff\x1e\x46\x4c\x93\x7c\x1a\x00\x00")

func artifactsMysqlStatefulsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsMysqlStatefulsetYaml,
		"artifacts/mysql-statefulset.yaml",
	)
}

func artifactsMysqlStatefulsetYaml() (*asset, error) {
	bytes, err := artifactsMysqlStatefulsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/mysql-statefulset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"artifacts/backup-crd.yaml":         artifactsBackupCrdYaml,
	"artifacts/backup-cronjob.yaml":     artifactsBackupCronjobYaml,
	"artifacts/backup-pvc.yaml":         artifactsBackupPvcYaml,
	"artifacts/mysql-configmap.yaml":    artifactsMysqlConfigmapYaml,
	"artifacts/mysql-crd.yaml":          artifactsMysqlCrdYaml,
	"artifacts/mysql-service-read.yaml": artifactsMysqlServiceReadYaml,
	"artifacts/mysql-service.yaml":      artifactsMysqlServiceYaml,
	"artifacts/mysql-statefulset.yaml":  artifactsMysqlStatefulsetYaml,
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
	"artifacts": {nil, map[string]*bintree{
		"backup-crd.yaml":         {artifactsBackupCrdYaml, map[string]*bintree{}},
		"backup-cronjob.yaml":     {artifactsBackupCronjobYaml, map[string]*bintree{}},
		"backup-pvc.yaml":         {artifactsBackupPvcYaml, map[string]*bintree{}},
		"mysql-configmap.yaml":    {artifactsMysqlConfigmapYaml, map[string]*bintree{}},
		"mysql-crd.yaml":          {artifactsMysqlCrdYaml, map[string]*bintree{}},
		"mysql-service-read.yaml": {artifactsMysqlServiceReadYaml, map[string]*bintree{}},
		"mysql-service.yaml":      {artifactsMysqlServiceYaml, map[string]*bintree{}},
		"mysql-statefulset.yaml":  {artifactsMysqlStatefulsetYaml, map[string]*bintree{}},
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