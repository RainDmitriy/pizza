package triable

import (
	"time"
)

func ExecuteWithTries(f func() error, tries int, delay time.Duration) (er error) {
	for tries > 0 {
		if er = f(); er != nil {
			time.Sleep(delay)

			continue
		}
		tries--

		return nil
	}

	return
}
