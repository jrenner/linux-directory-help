#linux-directory-help

dirhelp - a command line tool to give information about the linux directory structures (FHS)

### Installing with Installer:
1. `wget -qO - https://github.com/giodamelio/linux-directory-help/raw/master/installer.sh | /bin/bash`

### Installing Manually
1.  Grab the appropriate binary for your system: 
    - [amd64 - 64 bit](https://github.com/jrenner/linux-directory-help/raw/master/bin/dirhelp-linux-amd64)
    - [386 - 32 bit](https://github.com/jrenner/linux-directory-help/raw/master/bin/dirhelp-linux-386)
    - [arm - Raspberry Pi, etc](https://github.com/jrenner/linux-directory-help/raw/master/bin/dirhelp-linux-arm)
2. Copy it somewhere in your PATH environment variable, /usr/local/bin is a good place.
3. Rename it to something easy to type like "dirhelp".
4. Make it executable "chmod 755 (filename)".

### Usage
- run "dirhelp" in some typical directories like "/", "/media", "/var/log"
- run "dirhelp [path]" to get help on a specific path
- run "dirhelp -a" to see all the help strings

### Example Session
```
jrenner@main:/$ dirhelp
[/] Primary hierarchy root and root directory of the entire file system hierarchy.
jrenner@main:/$ dirhelp opt
[/opt] Optional application software packages.
jrenner@main:/$ cd /var/log
jrenner@main:/var/log$ dirhelp
[/var/log] Log files. Various logs.
jrenner@main:/var/log$ dirhelp /usr/share
[/usr/share] Architecture-independent (shared) data. This directory contains subdirectories with specific application data, that can be shared among different architectures of the same OS.  Often one finds stuff here  that  used  to live in /usr/doc or /usr/lib or /usr/man.
jrenner@main:/var/log$ dirhelp /usr/local
[/usr/local] Tertiary hierarchy for local data, specific to this host. Typically has further subdirectories, e.g., bin/, lib/, share/.
```
