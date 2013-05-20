echo "Installing linux-directory-help"
echo ""


# Download function to remove all the extra crap from wget
download() {
    local url=$1
    printf "    "
    wget -O "dirhelp" --progress=dot  $url 2>&1 | grep --line-buffered "%" | sed -u -e "s,\.,,g" | awk '{printf("\b\b\b\b%4s", $2)}'
    echo ""
}

printf "Detecting system arcitecture... "
case `uname -m` in

armv*)  
    echo `uname -m`
    printf "Downloading dirhelp-linux-arm... "
    download "https://github.com/jrenner/linux-directory-help/raw/master/bin/dirhelp-linux-arm"
    ;;
i386|i486|i586|i686)
    echo `uname -m`
    printf "Downloading dirhelp-linux-386... "
    download "https://github.com/jrenner/linux-directory-help/raw/master/bin/dirhelp-linux-386"
    ;;
x86_64) 
    echo `uname -m`
    printf "Downloading dirhelp-linux-amd64... "
    download "https://github.com/jrenner/linux-directory-help/raw/master/bin/dirhelp-linux-amd64"
    ;;
*)
    echo ""
    echo "Can not detect system arcitecture, you should do a manual install."
    exit 1
    ;;
esac

echo ""
echo "Installing to /usr/local/bin (requires sudo)"
sudo mv dirhelp /usr/local/bin/dirhelp
sudo chmod 755 /usr/local/bin/dirhelp

echo ""
echo "linux-directory-help sucessfully installed"
