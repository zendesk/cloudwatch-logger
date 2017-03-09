cloudwatch-logger
=================
A simple tool to connect the standard input stream to Amazon CloudWatch.

Usage
-----
```
cloudwatch-logger [-t] LOG-GROUP-NAME LOG-STREAM-NAME
```

If the `-t` option is specified, the input will also be copied to standard
output.

The log group and log stream will automatically be created if necessary.

**NOTE:** The named log stream must not exist prior to running this program.

Build instructions
------------------
```
$ make
```

If you want to cross-compile, set the `GOOS` and `GOARCH` environment variables
first.  See https://golang.org/doc/install/source for details on the possible
values.

Author
------
Michael S. Fischer, <mfischer@zendesk.com>

Thanks
------
* [Eric Holmes](https://github.com/ejholmes) for his dead-simple [CloudWatch
  Logs stream library](https://github.com/ejholmes/cloudwatch).
