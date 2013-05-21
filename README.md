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
![Alt text](http://github.com/jrenner/linux-directory-help/raw/master/dirhelp.png "screenshot")
