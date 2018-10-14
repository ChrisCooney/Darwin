# Darwin

I realised there were a lot of things that I do manually on my machine. At the same exact moment, I realised that my Golang knowledge was awful. It's not a big leap to work out what happened next.

## Installation

In the scripts folder is a `build-and-install.sh`. Make sure you run it from within the scripts folder!

## Secrets

The one file that is missing from the source is the `secret` module. That module needs to provide some values, exposed as functions. An example of a correct secret go file is something like this:

```go
package secret

func GuardianApiKey() string {
	return "MY_GUARDIAN_API_KEY"
}

func WolframAlphaAppId() string {
	return "MY_WOLFRAM_ALPHA_APP_ID"
}
```

This is somewhat basic and it would be preferable to define a JSON format. I'm just writing this incredibly jet lagged and I didn't fancy it. Will fix in the future.

## Usage

Darwin is talkative, once installed just ask him 

```
darwin "What can you do?"
```

He'll tell you everything you need to know.

## Disclaimer

There is absolutely no warranty or liability with this technology. I built it to make my life easier but I thought I'd put it up on Github because... well why not. If this is the Facebook for developers, this is a bit like me uploading a picture of my brunch. Trashy, but everyone's doing it.

## Limitations

It relies on Applescript for some of the functionality so this is only going to play nice with MacOS.
