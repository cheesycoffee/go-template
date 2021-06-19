package logger

// Log text color
var (
	Green   = []byte{27, 91, 57, 55, 59, 52, 50, 109}
	White   = []byte{27, 91, 57, 48, 59, 52, 55, 109}
	Yellow  = []byte{27, 91, 57, 48, 59, 52, 51, 109}
	Red     = []byte{27, 91, 57, 55, 59, 52, 49, 109}
	Blue    = []byte{27, 91, 57, 55, 59, 52, 52, 109}
	Magenta = []byte{27, 91, 57, 55, 59, 52, 53, 109}
	Cyan    = []byte{27, 91, 57, 55, 59, 52, 54, 109}
	Reset   = []byte{27, 91, 48, 109}
)
