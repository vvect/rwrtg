# rwrtg 

This is a companion tool to [rwrt](https://github.com/vvect/rwrt) used to generate hooks for Android applications.

## Getting Started

Personally, I build the binary and assign an alias in my zshrc, but feel free to use whatever works for you.

- git clone https://github.com/vvect/rwrtg
- cd rwrtg
- go build .

Once you've setup an alias or moved the binary to `/usr/local/sbin`, you can use it as follows: 

```sh
rwrtg -p static.rwrt.json
```

Which will generate a Frida script based on the template used, the `-t` flag can be used to change the template. You can also add your own templates!