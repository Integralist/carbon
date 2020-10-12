# Carbon

HTTP Response Header Sorting and Filtering

> carbon filtering: method of filtering impurities

## Install

```bash
go get -u github.com/integralist/carbon
```

## Build

To build and install the `carbon` executable locally, then run:

```bash
make install
```

> this will install to `/usr/local/bin` where as `go get` installs to the `~/go/bin`.

## Usage

```bash
carbon -help

Usage of carbon:
  -filter string
        comma-separated list of headers to be displayed
        e.g. X-Siterouter-Upstream,X-Cache
  -help
        show available flags
```

With filter...

```bash
carbon -filter cache,vary https://www.buzzfeed.com

Cache-Control:
  [no-cache, no-store, must-revalidate]

Vary:
  [X-BF-User-Edition, Accept-Encoding]

X-Cache:
  [MISS]

X-Cache-Hits:
  [0]

Status Code: 200 OK
```

No filter...

```bash
carbon https://www.buzzfeed.com/

Accept-Ranges:
  [bytes]

Age:
  [47]

Cache-Control:
  [no-cache, no-store, must-revalidate]

...lots of stuff...

X-Timer:
  [S1486813451.981669,VS0,VE0]

Status Code: 200 OK
```

## Tests?

Nope. This was just a quick hack

## Bash Alternative?

Here's an official repo for the Bash version](https://github.com/Integralist/Bash-Headers) (if you're interested).
