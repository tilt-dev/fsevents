// +build darwin,go1.10

package fsevents

/*
#include <CoreServices/CoreServices.h>
*/
import "C"

const (
	nullCFAllocatorRef = C.CFAllocatorRef(0)
	nullCFStringRef    = C.CFStringRef(0)
	nullCFUUIDRef      = C.CFUUIDRef(0)
)
