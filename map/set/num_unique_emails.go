package data_structures

import (
	"fmt"
	"strings"
)

// clarifying questions
// * local name = remove dots + afterwards ignore
// * domain name = no mapping rules

// approaches
// * set to track uniqueness -> just map each one
func numUniqueEmails(emails []string) int {
	set := make(map[string]struct{})
	for _, email := range emails {
		split := strings.Split(email, "@")
		forward := []rune{}
		// local
		for _, c := range split[0] {
			if c == '.' {
				continue
			}
			if c == '+' {
				break
			}
			forward = append(forward, c)
		}
		forwardStr := fmt.Sprintf("%s@%s", string(forward), split[1])
		set[forwardStr] = struct{}{}
	}
	return len(set)
}
