# sharding

## Install

```bash
go get -u github.com/WindomZ/sharding
```

## Usage

Hash crc32:

```go
sd := sharding.NewHashShard("_", 0, 64)  // create hash formatter
fmt.Println(sd.Format("hello", "world")) // hello_3
```

Time:

```go
sd := sharding.NewTimeShard("_", "20060102") // create time formatter
fmt.Println(sd..Format("hello"))             // hello_20180705
```
