```
input := bufio.NewScanner(os.Stdin)
for input.Scan() {
    fmt.Println(input.Text())
}
```