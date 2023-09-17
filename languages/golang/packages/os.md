```
for _, arg := range os.Args {
    fmt.Println(arg)
}
input := bufio.NewScanner(os.Stdin) // type *os.File

f, err := os.Open(file) // type *os.File
if err != nil {
    fmt.Fprintf(os.Stderr, "dup: %v\n", err)
    os.Exit(1)
}
...
defer f.Close()
```
```
data, err := os.ReadFile("example.txt") // byte slice

err := os.WriteFile("example.txt", data, 0644)
```