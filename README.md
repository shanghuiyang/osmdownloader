# osmdownloader
[![CI](https://github.com/shanghuiyang/osmdownloader/actions/workflows/ci.yml/badge.svg)](https://github.com/shanghuiyang/osmdownloader/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/shanghuiyang/osmdownloader/blob/master/LICENSE)

Download elements(nodes, ways or relations) using specific ids from [open street map](https://www.openstreetmap.org)

## Install
Osmdownloader is dependent on [osmosis](https://wiki.openstreetmap.org/wiki/Osmosis). To use osmdownloader, you need to install osmosis first. Please refer to [Install Osmosis](https://wiki.openstreetmap.org/wiki/Osmosis/Installation) for learning how to install it on different platforms. Take MacOS as an example,
```shell
# install osmosis for MacOS
$ brew install osmosis
```

Build and install it from source codes.
```shell
# get the source codes
$ go get -u github.com/shanghuiyang/osmdownloader

# build and install
$ cd $GOPATH/src/github.com/shanghuiyang/osmdownloader
$ go install
```

Or download the binary from [here](https://github.com/shanghuiyang/osmdownloader/releases).

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
