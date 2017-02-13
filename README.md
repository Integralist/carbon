# Carbon

HTTP Response Header Sorting and Filtering

> carbon filtering: method of filtering impurities

## Build

To build and install the `carbon` executable run:

```bash
make install
```

> this will install to `/usr/local/bin`

## Usage

Help...

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
carbon -filter X-Cache https://www.buzzfeed.com/

  X-Cache:
    [HIT]

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

## Bash?

Yeah sure I could've used Bash.

A simple oneliner would have looked like this:

```bash
curl -D ./headers.txt -o /dev/null -s https://www.buzzfeed.com/?country=us; cat ./headers.txt | sort; rm ./headers.txt
```

> Uses `-D, --dump-headers` to access headers  
> That's not possible via a pipe or when using `-v` verbose mode

Or if you wanted to abstract it away into a nice reusable Bash function:

```bash
function headers {
  if [ -z "$1" ]; then
    printf "\n\tExample: headers https://www.buzzfeed.com/?country=us 'x-cache|x-timer'\n"
    return
  fi

  local url=$1
  local pattern=${2:-''}
  local response=$(curl -H Fastly-Debug:1 -D - -o /dev/null -s "$url" | sort) # -D - will dump to stdout

  echo "$response" | egrep -i "$pattern"
}
```

Bash is super simple to write, and the above snippet should suffice for anyone not interested in Go.

But ultimately, I wanted to write some Go and this seemed like a fun/small thing to play around with
