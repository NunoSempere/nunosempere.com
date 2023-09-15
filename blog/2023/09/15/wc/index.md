Count words in <50 lines of C
===

The Unix utility wc counts words. You can make simple, non-POSIX compatible version of it that solely counts words in [159 words and 42 lines of C](https://git.nunosempere.com/personal/wc/src/branch/master/src/wc.c). Or you can be like GNU and take 3615 words and 1034 lines to do something more complex.

## Desiderata

- Simple: Just count words as delimited by spaces, tabs, newlines.
- Allow: reading files, piping to the utility, and reading from stdin—concluded by pressing Ctrl+D.
- Separate utilities for counting different things, like lines and characters, into their own tools.
- Avoid off-by-one errors.
- Linux only.
- Small.

## Comparison with other versions of wc

The [version of wc.c](https://git.nunosempere.com/personal/wc/src/branch/master/src/wc.c) in this repository sits at 44 lines. It decides to read from stdin if the number of arguments fed to it is otherwise zero, and uses the linux read function to read character by character. It doesn't have flags, instead, there are further utilities in the src/extra/ folder for counting characters and lines, sitting at 32 and 35 lines of code, respectively. This version also has little error checking.

[Here](https://github.com/dspinellis/unix-history-repo/blob/Research-V7-Snapshot-Development/usr/src/cmd/wc.c) is a version of wc from UNIX V7, at 86 lines. It allows for counting characters, words and lines. I couldn't find a version in UNIX V6, so I'm guessing this is one of the earliest versions of this program. It decides to read from stdin if the number of arguments fed to it is zero, and reads character by character using the standard C getc function.

The busybox version ([git.busybox.net](https://git.busybox.net/busybox/tree/coreutils/wc.c)) of wc sits at 257 lines (162 with comments stripped), while striving to be [POSIX-compliant](https://pubs.opengroup.org/onlinepubs/9699919799/), meaning it has a fair number of flags and a bit of complexity. It reads character by character by using the standard getc function, and decides to read from stdin or not using its own fopen_or_warn_stdin function. It uses two GOTOs to get around, and has some incomplete Unicode support.

The [plan9](https://9p.io/sources/plan9/sys/src/cmd/wc.c) version implements some sort of table method in 331 lines. It uses plan9 rather than Unix libraries and methods, and seems to read from stdin if the number of args is 0.

The plan9port version of wc ([github](https://github.com/9fans/plan9port/blob/master/src/cmd/wc.c)) also implements some sort of table method, in 352 lines. It reads from stdin if the number of args is 0, and uses the Linux read function to read character by character.

The [OpenBSD](https://github.com/openbsd/src/blob/master/usr.bin/wc/wc.c) version is just *nice*. It reads from stdin by default, and uses a bit of buffering using read to speed things up. It defaults to using fstat when counting characters. It is generally pleasantly understandable, nice to read. I'm actually surprised at how pleasant it is to read.

The [FreeBSD version](https://cgit.freebsd.org/src/tree/usr.bin/wc/wc.c) sits at 367 lines. It has enough new things that I can't parse all that it's doing: in lines 137-143, what is capabilities mode? what is casper?, but otherwise it decides whether to read from stdin by the number of arguments, in line 157. It uses a combination of fstat and read, depending on the type of file.

Finally, the GNU utils version ([github](https://github.com/coreutils/coreutils/tree/master/src/wc.c), [savannah](http://git.savannah.gnu.org/gitweb/?p=coreutils.git;a=blob;f=src/wc.c;hb=HEAD)) is a bit over 1K lines of C. It does many things and checks many possible failure modes. I think it detects whether it should be reading from stdin using some very wrapped fstat, and it reads character by character using its own custom function.

So this utility started out reasonably small, then started getting more and more complex. [The POSIX committee](https://pubs.opengroup.org/onlinepubs/9699919799/) ended up codifying that complexity, and now we are stuck with it because even implementations like busybox which strive to be quite small try to keep to POSIX.

## Installation

```
git clone
make
sudo make install 
## ^ installs to /bin/ww if there isn't a /bin/ww already
```

## Usage examples

```
echo "En un lugar de la Mancha" | ww
cat README.md | ww
ww README.md 
ww # write something, then exit with Ctrl+D
```

## Relationship with cat-v

Does one really need to spend 1k lines of C code to count characters, words and lines? There are many versions of this rant one could give, but the best and probably best known is [this one](https://harmful.cat-v.org/cat-v/unix_prog_design.pdf) on cat -v. Busybox itself has given up here, and its [version of cat](https://git.busybox.net/busybox/tree/coreutils/cat.c) has the following comment:

> Rob had "cat -v" implemented as a separate applet, catv.
> See "cat -v considered harmful" at
> http://cm.bell-labs.com/cm/cs/doc/84/kp.ps.gz
> From USENIX Summer Conference Proceedings, 1983
> 
> &#34;&#34;&#34;
> 
> The talk reviews reasons for UNIX's popularity and shows, using UCB cat
> as a primary example, how UNIX has grown fat. cat isn't for printing
> files with line numbers, it isn't for compressing multiple blank lines,
> it's not for looking at non-printing ASCII characters, it's for
> concatenating files.
> We are reminded that ls isn't the place for code to break a single column
> into multiple ones, and that mailnews shouldn't have its own more
> processing or joke encryption code.
> 
> &#34;&#34;&#34;
> 
> I agree with the argument. Unfortunately, this ship has sailed (1983...).
> There are dozens of Linux distros and each of them has "cat" which supports -v.
> It's unrealistic for us to "reeducate" them to use our, incompatible way
> to achieve "cat -v" effect. The actual effect would be "users pissed off
> by gratuitous incompatibility".

I'm not sure that gratuitous incompatibility is so bad if it leads to utilities that are much simpler and easier to understand and inspect. That said, other projects aiming in this direction that I could find, like [tiny-core](https://github.com/keiranrowan/tiny-core/tree/master/src) or [zig-coreutils](https://github.com/leecannon/zig-coreutils) don't seem to be making much headway.

## To do

- [ ] Possible follow-up: Write simple versions for other coreutils. Would be a really nice project.
- [ ] Get some simple version of this working on a DuskOS/CollapseOS machine?
  - [ ] Or, generally find a minimalistic kernel that could use some simple coreutils.
- [ ] Add man pages?
- [ ] Pitch to lwn.net as an article?
- [ ] Come back to writting these in zig.

## Done or discarded

- [x] Look into how C utilities both read from stdin and from files.
- [x] Program first version of the utility
- [x] Compare with other implementations, see how they do it, after I've created my own version
  - [x] Compare with gnu utils.
  - [x] Compare with busybox implementation
  - [x] Compare with other versions
  - [x] Compare with other projects: <https://github.com/leecannon/zig-coreutils>, <https://github.com/keiranrowan/tiny-core/tree/master>.
- [x] Install to ww, but check that ww is empty (installing to wc2 or smth would mean that you don't save that many keypresses vs wc -w)
- [x] Look specifically at how other versions do stuff.
  - [x] Distinguish between reading from stdin and reading from a file
    - [x] If it doesn't have arguments, read from stdin. 
  - [x] Open files, read characters.
- [x] Write version that counts lines (lc)
  - [x] Take into account what happens if file doesn't end in newline.
  - [ ] ~~Count EOF as word & line separator~~
  - [x] Document it
- [x] Document reading from user-inputed stdin (end with Ctrl+D)
- [x] add chc, or charcounter (cc is "c compiler")
- [x] Add licenses to historical versions before redistributing.
- [ ] ~~Could use zig? => Not for now~~
- [ ] ~~Maybe make some pull requests, if I'm doing something better? => doesn't seem like it~~
- [ ] ~~Write man files?~~
