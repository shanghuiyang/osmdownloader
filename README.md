# osmdownloader
Download nodes or ways by specific ids from open street map

## Install
```shell
$ go get github.com/shanghuiyang/osmdownloader
```

## Build
```shell
$ go build -o osmdownloader main.go
```

## Usage
```
-t  element type: node or way
-i  node or way ids
```

## Example
```shell
$ osmdownloader -t node -i 123,456,789
# or
$ osmdownloader -t way -i 123,456,789
```
the map data will be saved to merge.osm
