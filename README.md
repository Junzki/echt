# Echt

### Tips for cannot download packages
As is widely known, `golang.org` is blocked in China so `go get` cannot fetch packages from it.  
Fortunately, Glide supports a `mirror` command which can help a little.

For example:
```
$ glide mirror set https://golang.org/x/sys https://github.com/golang/sys --vcs git
```

And than, Glide will fetch `sys` from Github instead of the official source.

_So, fuck GFW._