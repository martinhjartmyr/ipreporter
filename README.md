# ipReporter

ipReporter is a small serverless api that helps keeping track of machines with dynamic IP addresses.
The main purpose of the project is to try some AWS technologies.

### ipReporter is using the following AWS technologies:

- Lambda (Golang)
- DynamoDB
- API Gateway
- Certificate Manager

## Examples

Environment variables:
```
REGION=eu-north-1 // AWS Stockholm
SECRET=changeMe // x-api-key header used to GET and PUT
TABLE_NAME=ipReporter // DynamoDB table name
```

Crontab updating the IP every other hour:
```
0 */2 * * * curl -H "x-api-key:changeMe" -X PUT https://xxxx.xxxxxx.se/{ALIAS} > /dev/null 2>&1
```

Retrieving the IP:
```
curl -H "x-api-key:changeMe" https://xxxx.xxxxxx.se/{ALIAS}
Output: {"alias":"{ALIAS}","ip":"xx.xxx.xxx.xx","timestamp":"2019-07-17T07:55:10Z"}
```

