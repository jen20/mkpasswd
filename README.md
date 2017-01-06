[![Mozilla Public License](https://img.shields.io/badge/license-MPL-blue.svg)](https://www.mozilla.org/MPL/)
# mkpasswd

## Fork Notes

This forked version only works with SHA512 hashes and will generate a salt if one is not provided.

Forked from: https://github.com/myENA/mkpasswd, but since functionality has been removed it is unlikely I will ever try to upstream these changes.

## Summary

Simple mkpasswd utility written in golang for platform portability.

## Installing

With a proper Go environment simply run:

```
go get -u github.com/jen20/mkpasswd
```

## Usage

### Summary

```
ahurt$ ./mkpasswd -h
Usage of mkpasswd:
  -password string
        Optional password argument
  -salt string
        Optional salt argument without prefix
```

### Example

```
$ ./mkpasswd
Password: ****
Confirm:  ****
$6$amUMrbDAEvqAdrtz$Jg0xMnIVeRR2IrZExX3AJj/IIMkfqDGGebIiUFRM2A376d8rbIJYBMOQGjoLeHu3mPlq//0Awc55zEtBNH43m.
```
