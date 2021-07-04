package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func getMD5Hash(file io.Reader) (string, error) {
	h := md5.New()

	_, err := io.Copy(h, file)
	if err != nil {
		fmt.Println("getFileMD5Hash", err)
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// SaveToDisk - save file to local
func SaveToDisk(path string, file multipart.File, extension string) (string, int64, error) {
	// duplicate file reader, one for copying and other for generating md5 hash
	var buf bytes.Buffer
	tee := io.TeeReader(file, &buf)

	// get md5 hash of file
	fileName, err := getMD5Hash(tee)
	if err != nil {
		return "", 0, err
	}

	fileName = "./" + path + "/" + fileName + extension
	// create local file with md5 hash name
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("saveToDisk", err)
		return "", 0, err
	}
	defer f.Close()

	// copy file data
	_, err = io.Copy(f, &buf)
	if err != nil {
		fmt.Println("saveToDisk", err)
		return "", 0, err
	}

	fi, err := f.Stat()
	if err != nil {
		return "", 0, err
	}

	return fileName, fi.Size(), nil
}

// GetFileMIMEType - get mime type of file
func GetFileMIMEType(extension string) string {
	extension = strings.ToLower(extension)
	switch extension {
	// image
	case ".bmp":
		return "image/bmp"
	case ".cod":
		return "image/cis-cod"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".ief":
		return "image/ief"
	case ".jpe":
		return "image/jpeg"
	case ".jpeg":
		return "image/jpeg"
	case ".jpg":
		return "image/jpeg"
	case ".jfif":
		return "image/pipeg"
	case ".svg":
		return "image/svg+xml"
	case ".tif":
		return "image/tiff"
	case ".tiff":
		return "image/tiff"
	case ".ras":
		return "image/x-cmu-raster"
	case ".cmx":
		return "image/x-cmx"
	case ".ico":
		return "image/x-icon"
	case ".pnm":
		return "image/x-portable-anymap"
	case ".pbm":
		return "image/x-portable-bitmap"
	case ".pgm":
		return "image/x-portable-graymap"
	case ".ppm":
		return "image/x-portable-pixmap"
	case ".rgb":
		return "image/x-rgb"
	case ".xbm":
		return "image/x-xbitmap"
	case ".xpm":
		return "image/x-xpixmap"
	case ".xwd":
		return "image/x-xwindowdump"
	case ".png":
		return "image/png"
	default:
		return ""
	}
}
