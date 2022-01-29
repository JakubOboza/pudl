# Pudl

Command line tool to help you stay on top of news on `pudelek.pl` (Polish gossip website)

# Usage

* `pudl list` - get list of hot topics
* `pudl top` - get only top of the list (default 3)
* `pudl cache` - get list of cached topics
* `pudl show ID` - open browser with selected topic id.

Typical usage would look something like this:

```
12:37 kuba@> pudl list
⭐⭐⭐ 😹 Lista tematow z pudla 🐩 😹 ⭐⭐⭐
1: 17-letnia siostra Julii Wieniawy chwali się nową fryzurą. Ma CZERWONE włosy (FOTO)
2: Justyna Gradek ujawnia, ile kosztowały ją NOWE POŚLADKI, na których przez dwa tygodnie nie mogła siadać...
...
12:37 kuba@> pudl show 5
```

## .pudl.yaml

Pudl is using a yaml file to store config and cache info. It is located in your home directory and named `.pudl.yaml`


# Install

### Build from source

You will need Go 1.16 at least.

1. Clone repo
2. make build
3. add to PATH <pudl_repo_path>/bin

If you build it in `/Users/kuba/pudl` you would need to add to your `.bashrc` or `.zshrc` line
```
PATH="/Users/kuba/pudl/bin:$PATH"
```

### Install from binary

To download ready binary use releases page [https://github.com/JakubOboza/pudl/releases](https://github.com/JakubOboza/pudl/releases)

1. Download binary for your operating system and arch eg. osx amd64
2. Untar it in a directory eg. `/Users/<user_name>/pudl` using `tar -zxvf pudl.osx-amd64.tar.gz`
3. Add it to the path var (if you are using bash ir will be ~./bashrc) `PATH="/Users/<user_name>/pudl:$PATH"`
