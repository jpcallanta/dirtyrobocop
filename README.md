# dirtyrobocop
Offensive Twitter bot

See it in action [here](https://twitter.com/dirtyrobocop)

## Requirements
Needs access to Twitter API via a valid Twitter account with app access

## Usage
Can be run in a container with the Dockerfile 

```docker build -t=dirtyrobocop:latest ./```

```docker run -d dirtyrobocop:latest```

Or do a ```go get && go build``` and run ```./dirtyrobocop config.json```

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
