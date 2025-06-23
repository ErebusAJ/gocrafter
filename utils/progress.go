package utils

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

// shows cli progress animations for a command
func ProgressTask(task string, work func() error) error {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithColor("cyan"))
	s.Suffix = fmt.Sprintf("  %v", task)
	s.Start()

	err := work()
	if err != nil {
		s.FinalMSG = fmt.Sprintf("❌ %s failed \n", task)
		s.Stop()
		return err
	}

	s.FinalMSG = fmt.Sprintf("✅ %s complete\n", task)
	s.Stop()

	return nil
}
