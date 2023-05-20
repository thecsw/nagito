# nagito

Shorten your URLs by sending requests to [monokuma](https://github.com/thecsw/monokuma).

Use simple commands like,

```sh
./nagito -url "https://sandyuraz.com" -auth "supersecretkey"

OUTPUT:
https://photos.sandyuraz.com/fBW
```

when `monokuma` was started with the same auth key,

```sh
./monokuma -auth "supersecretkey"
```

## Flags

There are a few flags that you can use to customize the behaviour of `nagito`,

```
Usage of ./nagito:
  -auth string
    	authentication token
  -key string
    	key to use when shortening a url
  -shortener string
    	monokuma shortener's url (default "https://photos.sandyuraz.com")
  -url string
    	url to shorten
  -urls string
    	urls to shorten (newline-separated)
```

So you can pass `-urls` to shorten multiple URLs at once,

```sh
https://sandyuraz.com
https://google.com,google
```

Where the last "google" is the key to use when shortening the URL, which will return
`https://photos.sandyuraz.com/google`.
