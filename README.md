puppetquery
===========

Query puppetdb sane and dependency free.

[![Build Status][1]][2]

[1]: https://secure.travis-ci.org/nightlyone/puppetquery.png
[2]: http://travis-ci.org/nightlyone/puppetquery


LICENSE
-------
BSD

documentation
-------------
[package documentation at go.pkgdoc.org](http://go.pkgdoc.org/github.com/nightlyone/puppetquery)


quick usage
-----------
* put URL to your puppetdb in `$HOME/.config/puppetquery/config.ini`
  with a line like `url=http://localhost:8080`
* query all active puppet nodes by calling `nq` without parameters
* query active puppet nodes having `fact1=foo` and `fact2=bar` (implicit and)

	nq fact1=foo fact2=bar

build and install
=================

install from package
--------------------
Just install the package puppetquery.

build the package
-----------------
Works like any other debian source package
 * Install build tools from debian via `apt-get install devscripts fakeroot`
 * Install build dependencies as reported with `dpkg-checkbuilddeps`
 * Run `fakeroot debian/rules binary`

Note: Please don't forget to increase the version number and adding your changes
to the debian/changelog via `dch -i` before building a package you plan to release!

install from source
-------------------

Install [Go 1][3], either [from source][4] or [with a prepackaged binary][5].

Then run

	go get github.com/nightlyone/puppetquery
	go get github.com/nightlyone/puppetquery/cmd/nq

List all active puppet nodes

	$GOPATH/bin/nq

List all active puppet nodes have 2 processors and running Debian

	$GOPATH/bin/nq processorcount=2 osfamily=Debian

[3]: http://golang.org
[4]: http://golang.org/doc/install/source
[5]: http://golang.org/doc/install

LICENSE
-------
BSD

documentation
-------------

contributing
============

Contributions are welcome. Please open an issue or send me a pull request for a dedicated branch.
Make sure the git commit hooks show it works.

git commit hooks
-----------------------
enable commit hooks via

        cd .git ; rm -rf hooks; ln -s ../git-hooks hooks ; cd ..

