# sharding

> Simple sharding tables fo Go.

[![Build Status](https://travis-ci.org/WindomZ/sharding.svg?branch=master)](https://travis-ci.org/WindomZ/sharding)
[![codecov](https://codecov.io/gh/WindomZ/sharding/branch/master/graph/badge.svg)](https://codecov.io/gh/WindomZ/sharding)

## Install

```bash
go get -u github.com/WindomZ/sharding
```

## Usage

Hash(crc32):

```go
shard := sharding.NewHashShard("_", 0, 64)  // create hash formatter
fmt.Println(shard.Format("hello", "world")) // hello_3
```

Time:

```go
shard := sharding.NewTimeShard("_", "20060102") // create time formatter
fmt.Println(shard..Format("hello"))             // hello_20180705
```

## Contributing

Welcome to pull requests, report bugs, suggest ideas and discuss, 
i would love to hear what you think on [issues page](https://github.com/WindomZ/sharding/issues).

If you like it then you can put a :star: on it.

## License

[MIT](https://github.com/WindomZ/sharding/blob/master/LICENSE)
