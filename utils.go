package utils

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type FileMeta struct {
	Path     string
	Size     int64
	ModTime  int64
	DateTime int64
	Hash     []byte
}

type FolderMeta []FileMeta

/*
Creates hash->path map of duplicates files in a directory tree

*/
func (fm FolderMeta) Duplicates() map[string][]string {

	duplicates := make(map[string][]string)
	hashMap := make(map[string]string)

	for i := range fm {
		hash := hex.EncodeToString(fm[i].Hash)
		if val, ok := hashMap[hash]; ok {
			//fmt.Println(val)
			//fmt.Println(fm[i].Path)
			if _, ok = duplicates[hash]; ok {
				duplicates[hash] = append(duplicates[hash], fm[i].Path)
			} else {
				duplicates[hash] = []string{val}
				duplicates[hash] = append(duplicates[hash], fm[i].Path)
			}
		} else {
			hashMap[hash] = fm[i].Path
		}
	}
	return duplicates
}

/*
Remove those files that have "sting s in it"
*/
func (fm FolderMeta) RemoveFiles(s string) FolderMeta {
	folderMeta := make(FolderMeta, 0)

	for i := range fm {
		if strings.Index(fm[i].Path, s) == -1 {
			folderMeta = append(folderMeta, fm[i])
		}
	}

	return folderMeta
}

/*
Remove those files that starts with "prefix"
*/
func (fm FolderMeta) RemoveFilesHasPrefix(prefix string) FolderMeta {
	folderMeta := make(FolderMeta, 0)

	for i := range fm {
		if strings.HasPrefix(fm[i].Path, prefix) {
			folderMeta = append(folderMeta, fm[i])
		}
	}

	return folderMeta
}

/*
Remove those files that ends with "suffix"
*/
func (fm FolderMeta) RemoveFilesHasSuffix(suffix string) FolderMeta {
	folderMeta := make(FolderMeta, 0)

	for i := range fm {
		if strings.HasSuffix(fm[i].Path, suffix) {
			folderMeta = append(folderMeta, fm[i])
		}
	}

	return folderMeta
}

/*
Just show files that have "sting s in it"
*/
func (fm FolderMeta) FilterFiles(s string) FolderMeta {
	folderMeta := make(FolderMeta, 0)

	for i := range fm {
		if strings.Index(fm[i].Path, s) > -1 {
			folderMeta = append(folderMeta, fm[i])
		}
	}

	return folderMeta
}

/*
Just show files that starts with "prefix"
*/
func (fm FolderMeta) FilterFilesHasPrefix(prefix string) FolderMeta {
	folderMeta := make(FolderMeta, 0)

	for i := range fm {
		if strings.HasPrefix(fm[i].Path, prefix) {
			folderMeta = append(folderMeta, fm[i])
		}
	}

	return folderMeta
}

/*
Just show files that ends with "suffix"
*/
func (fm FolderMeta) FilterFilesHasSuffix(suffix string) FolderMeta {
	folderMeta := make(FolderMeta, 0)

	for i := range fm {
		if strings.HasPrefix(fm[i].Path, suffix) {
			folderMeta = append(folderMeta, fm[i])
		}
	}

	return folderMeta
}

type ByHash []FileMeta

func (v ByHash) Len() int      { return len(v) }
func (v ByHash) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v ByHash) Less(i, j int) bool {
	if bytes.Compare(v[i].Hash, v[j].Hash) < 0 {
		return true
	} else {
		return false
	}
}

type ByPath []FileMeta

func (v ByPath) Len() int           { return len(v) }
func (v ByPath) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v ByPath) Less(i, j int) bool { return v[i].Path < v[j].Path }

type ByModTime []FileMeta

func (v ByModTime) Len() int           { return len(v) }
func (v ByModTime) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v ByModTime) Less(i, j int) bool { return v[i].ModTime < v[j].ModTime }

type ByDateTime []FileMeta

func (v ByDateTime) Len() int           { return len(v) }
func (v ByDateTime) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v ByDateTime) Less(i, j int) bool { return v[i].DateTime < v[j].DateTime }

func Intersection(A, B []FileMeta) []FileMeta {
	intersection := make([]FileMeta, 0)

	n1 := len(A)
	n2 := len(B)

	i := 0
	j := 0

	for i < n1 && j < n2 {
		if bytes.Compare(A[i].Hash, B[j].Hash) > 0 {
			j++
		} else if bytes.Compare(B[j].Hash, A[i].Hash) > 0 {
			i++
		} else {
			intersection = append(intersection, A[i])
			i++
			j++
		}
	}
	return intersection
}

func IntersectionByDateTime(A, B []FileMeta) []FileMeta {
	intersection := make([]FileMeta, 0)

	n1 := len(A)
	n2 := len(B)

	i := 0
	j := 0

	for i < n1 && j < n2 {
		if A[i].DateTime > B[j].DateTime {
			j++
		} else if B[j].DateTime > A[i].DateTime {
			i++
		} else {
			intersection = append(intersection, A[i])
			i++
			j++
		}
	}
	return intersection
}

/*
This function handles Intersection with duplicates.
Example: folder A has duplicates files (different file names but hash of the files match)
and folder B has also duplicate files.

A = 2 	3	5	5	5	5 	8	10
B = 5	5	6	8	8	9


Intersection of A and B
I = 5	5	5	5	5	5	8	8	8
*/
func IntersectionWithDuplicates(A, B []FileMeta) []FileMeta {
	intersection := make([]FileMeta, 0)

	n1 := len(A)
	n2 := len(B)

	i := 0
	j := 0

	for i < n1 && j < n2 {
		if bytes.Compare(A[i].Hash, B[j].Hash) > 0 {
			j++
		} else if bytes.Compare(B[i].Hash, A[j].Hash) > 0 {
			i++
		} else {
			intersection = append(intersection, A[i])

			tempA := A[i].Hash
			tempB := B[i].Hash

			i++
			j++

			for bytes.Compare(A[i].Hash, tempA) == 0 {
				i++
			}

			for bytes.Compare(B[j].Hash, tempB) == 0 {
				j++
			}

		}
	}
	return intersection
}

func ResizeImage(src, dst string) error {

	file, err := os.Open(src)
	if err != nil {
		return err
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}
	file.Close()

	// to fit the thumbnail inside a rectange
	// Find which is smaller: MaxWidth/w or MaxHeight/h Then multiply w and h by that number
	scalex := float64(260) / float64(img.Bounds().Size().X)
	scaley := float64(200) / float64(img.Bounds().Size().Y)

	var thumbx, thumby uint

	if scalex < scaley {
		thumbx, thumby = uint(float64(img.Bounds().Size().X)*scalex), uint(float64(img.Bounds().Size().Y)*scalex)
	} else {
		thumbx, thumby = uint(float64(img.Bounds().Size().X)*scaley), uint(float64(img.Bounds().Size().Y)*scaley)
	}
	//fmt.Println(thumbx, thumby)

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(thumbx, thumby, img, resize.Bilinear)

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, &jpeg.Options{85})

	return nil
}

func GetFileMeta(path string) (FileMeta, error) {

	meta := FileMeta{}
	meta.Path = path

	//fmt.Println(path)
	infile, err := os.Open(path)
	if err != nil {
		return meta, err
	}

	stat, err := infile.Stat()
	if err != nil {
		return meta, err
	}

	hash := sha1.New()
	io.Copy(hash, infile)

	meta.Hash = hash.Sum([]byte(""))
	meta.Size = stat.Size()
	meta.ModTime = stat.ModTime().Unix()

	_, f := filepath.Split(path)

	if CheckFilenameDateTime(f) {
		ts, err := ParseDateTime(f[0:15])
		if err != nil {
			meta.DateTime = meta.ModTime
		} else {
			meta.DateTime = ts
		}
	} else {

		if strings.ToLower(filepath.Ext(path)) == ".jpg" {

			dt, err := ReadExifDateTime(path)
			if err != nil {
				meta.DateTime = meta.ModTime
			} else {
				ts, err := ParseExifDateTime(dt)
				if err != nil {
					meta.DateTime = meta.ModTime
				} else {
					meta.DateTime = ts
				}
			}

			meta.DateTime = meta.ModTime
		} else {
			meta.DateTime = meta.ModTime
		}
	}
	return meta, nil

}

func SaveFolderMeta(folderMeta FolderMeta, path string) {
	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}

	for i := range folderMeta {
		fmt.Fprintf(f, "%x\t", folderMeta[i].Hash)
		fmt.Fprintf(f, "%d\t", folderMeta[i].ModTime)
		fmt.Fprintf(f, "%d\t", folderMeta[i].DateTime)
		fmt.Fprintf(f, "%d\t", folderMeta[i].Size)
		fmt.Fprintf(f, "%s\n", folderMeta[i].Path)
	}
}

func LoadFolderMeta(filename string) (FolderMeta, error) {

	folderMeta := make(FolderMeta, 0, 1000)

	f, err := os.Open(filename)
	if err != nil {
		//fmt.Println(err)
		return folderMeta, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "\t")
		//fmt.Println(s)

		if len(s) == 5 {
			meta := FileMeta{}
			meta.Hash, _ = hex.DecodeString(s[0])
			meta.ModTime, _ = strconv.ParseInt(s[1], 10, 64)
			meta.DateTime, _ = strconv.ParseInt(s[2], 10, 64)
			meta.Size, _ = strconv.ParseInt(s[3], 10, 64)
			meta.Path = s[4]

			folderMeta = append(folderMeta, meta)
		}
	}

	return folderMeta, nil
}

func FileHash(path string) ([]byte, error) {

	var s []byte

	infile, err := os.Open(path)
	if err != nil {
		return s, err
	}

	hash := sha1.New()
	io.Copy(hash, infile)
	s = hash.Sum([]byte(""))

	return s, nil
}

/*
ParseDateTime takes filename without extension and converts it into Unix Timestamp
Accectep filename format is "20060102_150405"
*/
func ParseDateTime(filename string) (int64, error) {
	t, err := time.Parse("20060102_150405", filename)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

/*
ParseExifDateTime takes filename without extension and converts it into Unix Timestamp
Accectep filename format is "2006:01:02 15:04:05"
*/
func ParseExifDateTime(filename string) (int64, error) {
	t, err := time.Parse("2006:01:02 15:04:05", filename)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
