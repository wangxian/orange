orange
======

A development tool, automatically refresh the browser when file changed Like LiveEdit

Later any further need not call it a day to refresh your browser.

Support: WEB dev, PHP, Node.js, Python, etc...

```shell
Usage of orange:
  -dir="./": Watch dir which change will refresh the browser, default current dir
  -ignores="": Not watch files, split width `,` Not regexp like `.go,.git/`, default no ignores
  -port=4000: Static server port, The port must>1024, default 4000
  -portproxy=0: Proxy http://localhost:{{port}}/ when file saved refresh browser, set 0 not proxy
  ------------------------------------------------------------
```