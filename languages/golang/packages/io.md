`io/ioutil`
```
data, err := ioutil.ReadFile(filename)
...
for _, line := range strings.Split(string(data), "\n") { ... }
```
`iotuil` has been deprecated with functionality moved to `os.WriteFile` and `os.ReadFile`
```
resp. err := http.Get(url)
...
b, err := ioutil.ReadAll(resp.Body) // b is a string
resp.Body.Close()
if err != nil { ... }
```
