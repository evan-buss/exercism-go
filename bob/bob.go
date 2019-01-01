package bob

import (
	"strings"
)

// Hey returns Bob's response to the remark as a string
func Hey(remark string) string {
	remark = strings.TrimSpace(remark) // Remove all whitespace and special chars
	if remark == "" {                  // Nothing said
		return "Fine. Be that way!"
	} else if strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark {
		if strings.HasSuffix(remark, "?") {
			// Yelling a question
			return "Calm down, I know what I'm doing!"
		}
		// Yelling
		return "Whoa, chill out!"
	} else if strings.HasSuffix(remark, "?") {
		// General question
		return "Sure."
	} else {
		// Any other remark
		return "Whatever."
	}
}
