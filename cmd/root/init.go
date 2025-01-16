package root

import "terramate-bootstrap/internal/tmcommands"

func Init() {
	tmcommands.CheckVersion()
}
