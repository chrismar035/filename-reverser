# filename-reverser

`filename-reverser` is a command line utility written in go to reverse the
characters in a filename while preserving the extension.

## Examples

```
$ ls
sirhc.txt
ykeens.eb.ot.yrt.mkv

$ filename-reverser
sirhc.txt => chris.txt
ykeens.eb.ot.yrt.mkv => try.to.be.sneaky.mkv

$ ls
chris.txt
try.to.be.sneaky.mkv
```

You can also specify a path command line argument.

```
$ filename-reverser /path/to/reversed/files
sirhc.txt => chris.txt
ykeens.eb.ot.yrt.mkv => try.to.be.sneaky.mkv

$ ls /path/to/reversed/files
chris.txt
try.to.be.sneaky.mkv
```

`filename-reverser` does not modify the contents of the files, just renames them.
