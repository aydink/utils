package utils

import (
	//"encoding/binary"
	"bytes"
	"errors"
	"fmt"
	"os"
	//"path/filepath"
	"regexp"
)

var re, re_filename *regexp.Regexp
var EXIF_HEADER_SIGNATURE = []byte{0xff, 0xd8, 0xff, 0xe1}
var JPEG_SIGNATURE = []byte{0xff, 0xd8}
var EXIF_SIGNATURE = []byte{0x45, 0x78, 0x69, 0x66, 0x00, 0x00}
var DATETIME_SIGNATURE_BE = []byte{0x90, 0x03, 0x00, 0x02}
var DATETIME_SIGNATURE_LE = []byte{0x03, 0x90, 0x02, 0x00}

var BIG_ENDIAN = true

func ReadExifDateTime(path string) (string, error) {

	datetime := ""

	f, err := os.Open(path)
	if err != nil {
		return datetime, err
	}

	// create 128k buffer to read file header
	buf := make([]byte, 131072)
	n, err := f.Read(buf)
	n++

	//fmt.Println("read %d bytes", n)

	if err != nil {
		return datetime, err
	}

	jpegIndex := bytes.Index(buf, JPEG_SIGNATURE)
	if jpegIndex != 0 {
		fmt.Println(path)
		fmt.Println("not JPEG file")
	} else {
		index := bytes.Index(buf, EXIF_SIGNATURE)
		if jpegIndex == 0 && index < 0 {
			fmt.Println(path)
			fmt.Println("Exif data not found")
			return datetime, errors.New("Error: Exif data not found")
		} else {

			//fmt.Println("Exif file")

			if buf[index+6] == 73 {
				BIG_ENDIAN = false
				//fmt.Println("Little Endian")
			} else {
				BIG_ENDIAN = true
				//fmt.Println("Big Endian")
			}

			if BIG_ENDIAN {
				datetimeIndex := bytes.Index(buf, DATETIME_SIGNATURE_BE)

				if datetimeIndex < 0 {
					fmt.Println(path)
					fmt.Println("Invalid Exif file")
					return datetime, errors.New("Error: Invalid Exif file")
				}

				datetimeLength := bytesToUint32be(buf[int(datetimeIndex+4):int(datetimeIndex+8)])
				datetimeOffset := bytesToUint32be(buf[int(datetimeIndex+8):int(datetimeIndex+12)])

				datetime = string(buf[int(datetimeOffset)+index+6 : int(datetimeOffset+datetimeLength-1)+index+6])
			} else {
				datetimeIndex := bytes.Index(buf, DATETIME_SIGNATURE_LE)
				if datetimeIndex < 0 {
					fmt.Println(path)
					fmt.Println("Invalid Exif file")
					return datetime, errors.New("Error: Invalid Exif file")
				}

				datetimeLength := bytesToUint32le(buf[int(datetimeIndex+4):int(datetimeIndex+8)])
				datetimeOffset := bytesToUint32le(buf[int(datetimeIndex+8):int(datetimeIndex+12)])

				datetime = string(buf[int(datetimeOffset)+index+6 : int(datetimeOffset+datetimeLength-1)+index+6])
			}
		}
	}

	return datetime, nil

}

/*
Chec if datetime string formatted like this "2014:12:31 23:56:09"
*/
func CheckDateTimeFormat(datetime string) bool {
	s := re.FindString(datetime)
	if s == "" {
		return false
	} else {
		return true
	}
}

/*
Chec if datetime string formatted like this "2014:12:31 23:56:09"
*/
func CheckFilenameDateTime(datetime string) bool {
	s := re.FindString(datetime)
	if s == "" {
		return false
	} else {
		return true
	}
}

func bytesToUnt16le(b []byte) uint16 {
	return uint16(b[0]) | uint16(b[1])<<8
}

func bytesToUnt16be(b []byte) uint16 {
	return uint16(b[1]) | uint16(b[0])<<8
}

func bytesToUint32le(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24

}

func bytesToUint32be(b []byte) uint32 {
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func init() {
	re = regexp.MustCompile("[0-9]{4}:[0-9]{2}:[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}")
	re_filename = regexp.MustCompile("[0-9]{8}_[0-9]{6}")
}
