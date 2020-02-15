# Geo Lookup IPs

I needed to lookup geolocation info for a list of IPs, and came upon [IP
Location Lite](https://lite.ip2location.com/file-download). You will need to
download the latest `IP2LOCATION-LITE-DB11.BIN` file.

I haven't made any pretty command-line flags or anything either. It reads stdin
and writes csv to stdout. Running this:

```
cat << EOF | go run main.go
8.8.8.8
4.4.4.4
EOF
```

should return:

```
ip,country_short,country_long,region,city,latitude,longitude
8.8.8.8,US,United States of America,California,Mountain View,37.405991,-122.078514
4.4.4.4,US,United States of America,Minnesota,Minneapolis,44.979969,-93.263840
```

What you'll more likely do in practice is

```
cat some-ips.txt | go run main.go > some-ips-with-geo.csv
```
