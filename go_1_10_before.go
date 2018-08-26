// +build darwin,!go1.10

package fsevents

/*
#include <CoreServices/CoreServices.h>
*/
import "C"

var (
	nullCFAllocatorRef = C.CFAllocatorRef(nil)
	nullCFStringRef    = C.CFStringRef(nil)
	nullCFUUIDRef      = C.CFUUIDRef(nil)
)
