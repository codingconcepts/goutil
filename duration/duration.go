package duration

import "time"

// Round a given time.Duration d to a precision p.
func Round(d, p time.Duration) time.Duration {
	if p <= 0 {
		return d
	}

	neg := d < 0
	if neg {
		d = -d
	}

	if m := d % p; m+m < p {
		d = d - m
	} else {
		d = d + p - m
	}

	if neg {
		return -d
	}
	return d
}
