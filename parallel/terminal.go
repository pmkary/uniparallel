package parallel

import (
	"bufio"
	"os"
)

// ReadLine reads from the terminal screen's
// stream buffer and implements a read line
func ReadLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return text
}
