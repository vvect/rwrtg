# rwrtg 

This is a companion tool to [rwrt](https://github.com/vvect/rwrt) used to generate hooks for Android applications.

## Getting Started

Personally, I build the binary and assign an alias in my zshrc, but feel free to use whatever works for you.

```sh
# Clone the repository
git clone https://github.com/vvect/rwrtg $HOME/rwrtg
# Build the binary
cd $HOME/rwrtg
go build .
# Setup an alias in your rc file
echo 'alias rwrtg=$HOME/rwrtg/rwrtg' >> $HOME/.zshrc
```

Which will generate a Frida script based on the template used, the `-t` flag can be used to change the template. You can also add your own templates!
