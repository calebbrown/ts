# ts

Add date and time to stdout.

## Overview

Ever wanted to add timestamps to the output of another command?
Then `ts` is your friend.

It simply prepends every line it receives from stdin with the current
time and date.

And yes, you could just run something like
`alias ts='xargs -Iln bash -c "echo \$(date) ln"'`,
but this is faster and easier to use. Plus it's written in [Go](http://golang.org/)
so it's way better.


## Installation

Make sure Go 1.1 is installed and `$GOPATH` is set.

Then run:

    go get github.com/calebbrown/ts

## Usage

    $ long_running_cmd | ts
    [2013-07-04 16:06:57] some output
    [2013-07-04 16:08:30] some more output
    ...


### Custom time format

Using the `-f` argument:

    $ long_running_cmd | ts -f="Jan 02, 2006 3:04pm"
    [Jul 04, 2013 4:06pm] some output
    [Jul 04, 2013 4:08pm] some more output
    ...

Using the `TS_TIME_FORMAT` environment variable:

    $ export TS_TIME_FORMAT="Jan 02, 2006 3:04pm"
    $ long_running_cmd | ts
    [Jul 04, 2013 4:06pm] some output
    [Jul 04, 2013 4:08pm] some more output
    ...

For more information about what custom time formats look like the Go's [time package docs](http://golang.org/pkg/time/#pkg-constants) have some notes.