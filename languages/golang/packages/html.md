```
import "html/template"
var issueList = template.Must(template.New("issuelist").Parse(`...`))
doc, err := html.Parse(os.Stdin)
```