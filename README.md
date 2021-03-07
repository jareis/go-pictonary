[![Codacy Badge](https://app.codacy.com/project/badge/Grade/7964b44404d54118a248faa086c01041)](https://www.codacy.com/gh/jareis/go-pictonary/dashboard?utm_source=github.com&utm_medium=referral&utm_content=jareis/go-pictonary&utm_campaign=Badge_Grade)

Word Generation Game Made in GOLANG for the web

Intended to be a exercise for learning GO and more

## How to Build

_Requirement_

-   GNU make
-   GOLANG

```
$make build
$cd dist && ./go-pictonary
```

Alternatively you can build with docker

```
$docker build -t go-pictonary -f Dockerfile .
$docker run -it -p 8080:8080 go-pictonary
```
