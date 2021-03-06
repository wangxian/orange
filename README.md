orange
======

A development tool, automatically refresh the browser when file changed Like LiveEdit

Later any further need not call it a day to refresh your browser.

Support: WEB dev, PHP, Node.js, Python, etc...

# Usage

```shell

Usage of orange:
  -http=":4000": Static server port, The port must>1024, default :4000
  -ignores="": Not watch files, split width `,` Not regexp eg: `.go,.git/`, default no ignores
  -precmd="": Before refresh browser, execute precmd command. eg: `ls {0}`, {0} is the changed file
  -proxy="": Proxy webserver when file saved refresh browser, like :80
  -rootdir="./": Server root dir, default current dir
  -watchdir="./": Watch dir which change will refresh the browser, default current dir
  
```

eg:

```sh
cd ~/Sites/
orange -proxy 80 -watchdir ~/Sites/blog/ -ignores .cache,.db
```

### Notes

1. If you don't want to monitor files, set `orange -ignores .`
2. If you want ignore some dirs, eg: `orange -ignores ".git,.svn,dirs"`
3. If you want run command when file change, set `-precmd` eg: `orange -precmd "du -sh {0}"
4. For PHP/Node.js/etc... , You can proxy your proxy, eg: proxy apache `orange -portproxy 80`
5. If not port proxy, cd your dir, run orange, view a html file when current dir files change, browser will refresh automatically
6. When start, `Automatically Open URL http://localhost:${yourport}` in your Browser

# Downloads

Binaries

- [Mac OSX 10.8+](https://www.dropbox.com/s/t1ewa0wavmfuixt/orange-osx-2.4)
- [Windows 64bit](https://www.dropbox.com/s/prkzjecs7o76r87/orange-win64-2.4.exe)
- [Windows 32bit](https://www.dropbox.com/s/o8uk0fi4gtw2tph/orange-win32-2.4.exe)
- [Linux](https://www.dropbox.com/s/8gvg7b3fq3anru1/orange-linux-2.4)

Rename `orange-xxx-$VERSION` to `orange`, And move orange file to your `$PATH`

If you have already installed the golang

```sh
go get -u github.com/wangxian/orange
```

THEN add your `$GOPATH/bin` to your `$PATH`

# License

`orange's` code uses the MIT license, see our `LICENSE` file.
