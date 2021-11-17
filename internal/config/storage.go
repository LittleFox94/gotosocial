/*
   GoToSocial
   Copyright (C) 2021 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package config

// StorageConfig contains configuration for storage and serving of media files and attachments
type StorageConfig struct {
	// Type of storage backend to use: currently only 'local' is supported.
	Backend string `yaml:"backend"`

	// Protocol to use when *serving* media files from storage
	ServeProtocol string `yaml:"serveProtocol"`
	// Host to use when *serving* media files from storage
	ServeHost string `yaml:"serveHost"`
	// Base path to use when *serving* media files from storage
	ServeBasePath string `yaml:"serveBasePath"`

	Local LocalStorageConfig `yaml:"local"`
	S3    S3StorageConfig    `yaml:"s3"`
}

type LocalStorageConfig struct {
	// The base path for storing things. Should be an already-existing directory.
	BasePath string `yaml:"basePath"`
}

type S3StorageConfig struct {
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	Endpoint        string `yaml:"endpoint"`
	UseSSL          bool   `yaml:"use_ssl"`
	Bucket          string `yaml:"bucket"`
	Region          string `yaml:"region"`
	UsePathStyle    bool   `yaml:"path_style"`
}

const (
	// StorageBackendLocal stores uploaded media in the local filesystem, below the BasePath.
	StorageBackendLocal = "local"

	// StorageBackendS3 stores uploaded media on a S3-like object storage service.
	StorageBackendS3 = "s3"
)
