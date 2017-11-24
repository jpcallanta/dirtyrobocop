# dirtyrobocop
Offensive Twitter bot

See it in action [here](https://twitter.com/dirtyrobocop)

## Requirements
Needs access to Twitter API via a valid Twitter account with app access

The twitter API uses SSL which requires root certificates installed on the docker image, or you can
link the host /etc/ssl/certs folder into the image when running the container (see example below)

## Usage
Can be run in a container with the Dockerfile:

```docker build -t=dirtyrobocop:latest ./```

```docker run -d --name dirtyrobocop -v /etc/ssl/certs:/etc/ssl/certs dirtyrobocop:latest```

To run once and exit:

```docker run -it -v /etc/ssl/certs:/etc/ssl/certs dirtyrobocop:latest dirtyrobocop /etc/config.json runonce```

Or do a ```go get && go build``` and run ```./dirtyrobocop config.json runonce```

## Sample config.json
```json
{
	"consumerKey": "",
	"consumerSecret": "",
	"accessToken" : "",
	"accessSecret" : "",
	"foruntePath": "/path/to/fortune",
	"sleepMin" : 21600,
	"sleepMax" : 259200
}
```
