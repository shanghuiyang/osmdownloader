[![build status](https://travis-ci.org/shanghuiyang/osmdownloader.svg?branch=master)](https://travis-ci.org/shanghuiyang/osmdownloader)

# osmdownloader
Download nodes or ways by specific ids from open street map

## Install
Osmdownloader is dependent on [osmosis](https://wiki.openstreetmap.org/wiki/Osmosis). To use osmdownloader, you need to install osmosis first. Please refer to [Install Osmosis](https://wiki.openstreetmap.org/wiki/Osmosis/Installation) for learning how to install it on different platforms. Take MacOS as an example,
```shell
# install osmosis for MacOS
$ brew install osmosis
```

You can download the osmdownloader binary from [here](https://github.com/shanghuiyang/osmdownloader/releases),
or build and install it from source codes.
```shell
# get the source codes
$ go get -u github.com/shanghuiyang/osmdownloader

# build and install
$ cd $GOPATH/src/github.com/shanghuiyang/osmdownloader
$ go install
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
