# Carbon

HTTP Response Header Sorting and Filtering

> carbon filtering: method of filtering impurities

![Example output of running the carbon executable](./carbon.png)

## Install

```bash
go install github.com/integralist/carbon@latest
```

## Build

To build and install the `carbon` executable locally, then run:

```bash
make install
```

## Usage

```bash
carbon -help

Usage of carbon:
  -filter string
        comma-separated list of headers to be displayed
        e.g. cache,vary
  -help
        show available flags
  -json
        output is formatted into JSON for easy parsing
  -plain
    	output is formatted for easy piping
```

No filter...

```bash
carbon https://www.fastly.com/

Accept-Ranges:
  [bytes]

Cache-Control:
  [max-age=0, private, must-revalidate]

Content-Type:
  [text/html]

Date:
  [Tue, 27 Oct 2020 09:55:39 GMT]

Etag:
  ["c248491ee6293167e071523b47b4625e"]

Server:
  [Artisanal bits]

Strict-Transport-Security:
  [max-age=31536000]

Vary:
  [Accept-Encoding]

X-Cache:
  [HIT]

X-Content-Type-Options:
  [nosniff]

X-Frame-Options:
  [DENY]

X-Xss-Protection:
  [1; mode=block]

Status Code: 200 OK
```

With filter...

```bash
carbon -filter cache,vary https://www.fastly.com

Cache-Control:
  [max-age=0, private, must-revalidate]

Vary:
  [Accept-Encoding]

X-Cache:
  [HIT]

Status Code: 200 OK
```

With filter AND json...

```bash
carbon -filter cache,vary --json https://www.fastly.com

[{"Key":"Cache-Control","Val":["max-age=0, private, must-revalidate"]},{"Key":"Vary","Val":["Accept-Encoding"]},{"Key":"X-Cache","Val":["HIT"]},{"Key":"X-Cache-Hits","Val":["2"]}]
```

Plain...

```bash
carbon -filter cache,vary -plain https://www.fastly.com

Cache-Control: max-age=0, private, must-revalidate
Vary: Accept-Encoding
X-Cache: HIT
Status Code: 200 OK
```

## Tests?

Nope. This was just a quick hack

## Bash Alternative?

Here's an official repo for the Bash version](https://github.com/Integralist/Bash-Headers) (if you're interested).
