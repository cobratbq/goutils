package log

import "log"

// Debug writes a line to os.Stderr with prefix 'debug'.
func Debug(line string) {
	log.Println("[debug] " + line)
}

// Info writes a line to os.Stderr with prefix 'info'.
func Info(line string) {
	log.Println("[info] " + line)
}

// Warn writes a line to os.Stderr with prefix 'warn'.
func Warn(line string) {
	log.Println("[warn] " + line)
}

// Error writes a line to os.Stderr with prefix 'ERROR'.
func Error(line string) {
	log.Println("ERROR: " + line)
}

func Errorf(format string, args ...interface{}) {
	log.Printf("ERROR: "+format, args...)
}
