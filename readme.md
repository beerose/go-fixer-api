# go-fixer-api

go-fixer-api is a simple web server that converts particular amount of money in some currency to other currencies. Both amount and base currency should be provided in request url.
Server returns respons in json or xml format depending on request header (ex. "Accept: application/json"). The default format is json.

### Google Cloud Server
Server is hosted on Google Cloud Server.

link: https://fixerapi.appspot.com/

### Usage examples:
* JSON:
``` 
curl -i "https://fixerapi.appspot.com/convert?currency=PLN&amount=111"
```
* XML:
``` 
curl -H "Accept: application/xml" -i "https://fixerapi.appspot.com/convert?currency=PLN&amount=111"
```
### Sample responses:
* JSON:
```
{
  "amount": 111,
  "currency": "PLN",
  "converted": {
    "AUD": 39.33,
    "BGN": 51.01,
    "BRL": 97.86,
    "CAD": 38.5,
    "CHF": 30.08,
    "CNY": 203.02,
    "CZK": 673.3,
    ...
```

* XML
```
<validResponse amount="111" currency="PLN">
    <converted>
        <SGD value="41.77">
        </SGD>
        <DKK value="194.17">
        </DKK>
        <INR value="2001">
        </INR>
        <HRK value="195.85">
        </HRK>
        <ZAR value="412.25">
        ...
```

### Testing
Tests can be run with command:
``` 
go test fixerapi...
```




