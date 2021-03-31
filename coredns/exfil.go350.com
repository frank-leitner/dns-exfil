$TTL 900
@   900 IN	SOA gen.go350.com. brad.go350.com. (
        1           ;Serial
        7200        ;Refresh
        3600        ;Retry
        1209600     ;Expire
        3600        ;Negative response caching TTL
)


	900 IN NS gen.go350.com.

www 900 IN A     127.0.0.1
www 900 IN AAAA  ::1
