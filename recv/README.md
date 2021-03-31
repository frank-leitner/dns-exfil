# dns-exfil: recv

Run recv on the DNS server to re-assemble the exfiltrated files.

## How to use

```bash
$ recv -d -qlog ../dns-logs/example.log -marker wsxdec
```

## Options

```bash
$ recv
  -d	enable debugging.
  -help
    	show help.
  -marker string
    	a unique marker to identify the file in the dns logs. (default "jzp")
  -qlog string
    	path to the dns query log file.
```
