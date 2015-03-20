package screenshot

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import "screenshot-osx.objc"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// SaveFormat - Constants that specify what format to save the screenshot in
type SaveFormat int

// Save as a JPEG file
const FormatJpeg SaveFormat = 1

// Save as a PNG file
const FormatPng SaveFormat = 2

// SaveScreenshotToFile - Take a screenshot and save it to the file specified
func SaveScreenshotToFile(filename string, saveFormat SaveFormat) error {
	cPath := C.CString(filename)
	defer C.free(unsafe.Pointer(cPath))

	retval := C.saveScreenshotToFile(1, C.int(saveFormat), cPath)
	if retval != 0 {
		return fmt.Errorf("Failed to take a screenshot with error code %d", retval)
	}
	return nil
}
