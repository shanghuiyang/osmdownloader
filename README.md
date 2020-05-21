[![build status](https://travis-ci.org/shanghuiyang/osmdownloader.svg?branch=master)](https://travis-ci.org/shanghuiyang/osmdownloader)

# osmdownloader
Download nodes or ways by specific ids from open street map

## Install
to use this tool, you need to install osmosis first.
please refer to https://wiki.openstreetmap.org/wiki/Osmosis/Installation for installing it.
for example, 
```shell
# for MacOS
$ brew install osmosis
```

install source codes
```shell
$ go get github.com/shanghuiyang/osmdownloader
```

## Build
```shell
$ go build -o osmdownloader main.go
```

## Usage
```
osmdownloader -f xxx.osm [-n list | -w list | -r list]
-f  the output file
-n  the id list of nodes
-w  the id list of ways
-r  the id list of relations
```

## Example
```shell
$ osmdownloader -f test.osm -n 111,222,333
# or
$ osmdownloader -f test.osm -w 444,555,666
# or
$ osmdownloader -f test.osm -n 111,111,222 -w 333,444 -r 555,666
```
