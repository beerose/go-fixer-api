# go-fixer-api

go-fixer-api is a simple web server that converts particular amount of money in some currency to other currencies. Both amount and base currency should be provided in request url.
Server returns respons in json or xml format depending on request header (eg. "Accept: application/json"). The default format is json.

Usage example:
* JSON:
``` curl -i "https://fixerapi.appspot.com/convert?currency=PLN&amount=111"
```
* XML:
``` curl -H "Accept: application/xml" -i "https://fixerapi.appspot.com/convert?currency=PLN&amount=111"
```

### Google Cloud Server link: https://fixerapi.appspot.com/

### Testing
``` go test fixerapi...
```




