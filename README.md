# sharding

> Simple sharding tables fo Go.

[![Build Status](https://travis-ci.org/WindomZ/sharding.svg?branch=master)](https://travis-ci.org/WindomZ/sharding)
[![codecov](https://codecov.io/gh/WindomZ/sharding/branch/master/graph/badge.svg)](https://codecov.io/gh/WindomZ/sharding)

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
