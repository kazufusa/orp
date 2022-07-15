# orp arranges the order of your PATH

and remove duplications.

```sh
$ echo $PATH
/opt/homebrew/bin:/home/kazufusa/go/bin:/usr/local/go/bin:/home/kazufusa/bin:/home/kazufusa/.asdf/shims:/home/kazufusa/.asdf/bin:/home/kazufusa/.cargo/bin:/home/kazufusa/go/bin:/usr/local/go/bin:/home/kazufusa/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/usr/lib/wsl/lib
$ eval $(orp asdf /go/ .cargo homebrew kazufusa/bin)
$ echo $PATH
/home/kazufusa/.asdf/bin:/home/kazufusa/.asdf/shims:/usr/local/go/bin:/home/kazufusa/go/bin:/home/kazufusa/.cargo/bin:/opt/homebrew/bin:/home/kazufusa/bin:/usr/bin:/sbin:/usr/local/games:/usr/lib/wsl/lib:/usr/sbin:/usr/games:/bin:/usr/local/sbin:/usr/local/bin
```
