package data_structures

func removeComments(source []string) []string {
	var res []string

	in_comment := false
	var cur string

	for _, s := range source {
		i := 0
		for i < len(s) {
			var peek string
			if i+1 <= len(s)-1 {
				peek = s[i : i+2]
			} else {
				peek = s[i:]
			}

			if peek == "//" && !in_comment {
				break
			}
			if peek == "/*" && !in_comment {
				i += 2
				in_comment = true
				continue
			}
			if peek == "*/" && in_comment {
				i += 2
				in_comment = false
				continue
			}
			if in_comment {
				i++
				continue
			}
			cur += string(s[i])
			i++
		}
		if len(cur) != 0 && !in_comment {
			res = append(res, cur)
			cur = ""
		}
	}
	return res
}
