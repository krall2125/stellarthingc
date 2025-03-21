mkdir build

echo "Compiling to Linux amd64..."

mkdir build/stellarthing-linux
mkdir build/stellarthing-linux/assets
cp -r assets/* build/stellarthing-linux/assets
env GOOS=linux GOARCH=amd64 go build -o build/stellarthing-linux/stellarthing
# so i can cd
function archive_linux() {
    cd build
    tar -czvf stellarthing-linux.tar.gz stellarthing-linux
    cd ..
}
archive_linux

echo "Compiling to Windows amd64..."

mkdir build/stellarthing-windows
mkdir build/stellarthing-windows/assets
cp -r assets/* build/stellarthing-windows/assets
env CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o build/stellarthing-windows/stellarthing.exe
# so i can cd
function archive_windows() {
    cd build
    zip -r stellarthing-windows.zip stellarthing-windows
    cd ..
}
archive_windows

echo "Compiling completed."