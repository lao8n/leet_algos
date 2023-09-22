```
// \W matches non-word characters
re := regexp.MustCompile(`\W+`)
replaced := re.ReplaceAllString(str, " ")

// \w matches word characters
```