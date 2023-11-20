Non-cryptographic hash function
```
hash := fnv.New32()
hash.Write([]byte(key))
shardNum := hash.Sum32() % shardCount
```