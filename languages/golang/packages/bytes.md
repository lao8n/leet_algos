```
var buf bytes.Buffer
buf.WriteByte('[')
fmt.Fprintf(&buf, "%d", v)
buf.WriteRune(“是”) // better for unicode
```