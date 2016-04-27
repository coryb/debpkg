# debpkg

debpkg is a native [Go](https://golang.org) implementation with zero dependencies 
 to create [Debian](https://www.debian.org) packages.

[![License][License-Image]][License-Url]
![Stability][Stability-Status-Image]
[![Godoc][Godoc-Image]][Godoc-Url]
[![ReportCard][ReportCard-Image]][ReportCard-Url]
[![Build][Build-Status-Image]][Build-Status-Url]
[![BuildAppVeyor][BuildAV-Status-Image]][BuildAV-Status-Url]
[![Coverage][Coverage-Image]][Coverage-Url]

**Features**

- [ ] Create simple debian packages from files and folders
- [ ] Add custom control files (preinst, postinst, prerm, postrm etcetera)
- [ ] [dpkg](http://manpages.ubuntu.com/manpages/precise/man1/dpkg.1.html) like tool with a subset of commands (--contents, --control, --extract, --info)
- [ ] Create package from debpkg.yml specfile (like [packager.io](https://packager.io) without cruft)
- [ ] GPG sign package
- [ ] GPG verify package

## Why this package was created

This package was created due to the lack to debianize from other platforms (windows/mac/*bsd). Because
 the whole debian package management system is glued together with Perl scripts and uses a bunch of Perl
 modules.

And converting a directory and files into a debian package is a pain without knowing the `deb`-file internals.

This package is heavily inspired by [godeb](https://github.com/niemeyer/godeb) and
 [CPackDeb](https://cmake.org/cmake/help/v3.5/module/CPackDeb.html). It is very royal [licensed](LICENSE) to
 be used wherever your hearts have their desires to debianize stuffs.

## Installation

`go get -u github.com/xor-gate/debpkg`

## Status

The package is currently in experimental state. The API is unstable and it has not throughout
 been tested and many things are unfinished or not implemented at all. **USE AT YOUR OWN RISK!**

# debpkg.yml specfile

The specfile is written in the [YAML markup format](http://yaml.org/). It controls
 the package information and data.

A simple example is listed below:

```
# debpkg.yml specfile
description:
  - version: 0.0.1
  - short: This package is just a test
  - long:
      This package tests the working of debpkg.
      And can wrap multiple
      lines.

      And multiple paragraphs.
```

# License

[MIT](LICENSE)

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[Stability-Status-Image]: http://badges.github.io/stability-badges/dist/experimental.svg
[Build-Status-Url]: http://travis-ci.org/xor-gate/debpkg
[Build-Status-Image]: https://travis-ci.org/xor-gate/debpkg.svg?branch=master
[BuildAV-Status-Url]: https://ci.appveyor.com/project/xor-gate/debpkg
[BuildAV-Status-Image]: https://ci.appveyor.com/api/projects/status/iuw1j84l33ynxs32?svg=true
[Godoc-Url]: https://godoc.org/github.com/xor-gate/debpkg
[Godoc-Image]: https://godoc.org/github.com/xor-gate/debpkg?status.svg
[ReportCard-Url]: http://goreportcard.com/report/xor-gate/debpkg
[ReportCard-Image]: http://goreportcard.com/badge/xor-gate/debpkg
[Coverage-Url]: https://coveralls.io/r/xor-gate/debpkg?branch=master
[Coverage-image]: https://img.shields.io/coveralls/xor-gate/debpkg.svg
