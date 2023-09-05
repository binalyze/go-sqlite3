#!/bin/bash

set -eo pipefail

cd "$(dirname "$0")"

PCRE2_VERSION=${PCRE2_VERSION:-10.42}
PREFIX_BASE=${PREFIX_BASE:-$(pwd)/../build}
ARCH=${ARCH:-$(uname -m)}

pcre2_extra_configure_args=""

function is_command() {
    command -v "$1" >/dev/null 2>&1
}

function is_macos() {
    [ "$(uname)" = "Darwin" ]
}

function is_linux() {
    [ "$(uname)" = "Linux" ]
}

function is_windows() {
    case "$(uname)" in
    MSYS* | MINGW*) return 0 ;;
    *) return 1 ;;
    esac
}

function get_pcre2() {
    local filename=pcre2-${PCRE2_VERSION}.tar.bz2
    local url=https://github.com/PCRE2Project/pcre2/releases/download/pcre2-${PCRE2_VERSION}/$filename
    rm -rf pcre2
    if [ ! -f "$filename" ]; then
        wget -O "./$filename" "$url"
    fi
    tar -xjf "./$filename"
    local move="mv pcre2-${PCRE2_VERSION} pcre2"
    $move || $move # retry 1 time for windows
}

function build_pcre2() {
    pushd pcre2
    #./autogen.sh
    ./configure --prefix="${PREFIX}" --disable-pcre2grep-jit --disable-pcre2grep-callout --enable-pcre2-8 \
        --disable-pcre2grep-callout-fork --disable-dependency-tracking --enable-static $pcre2_extra_configure_args

    # --enable-pcre2-16 --enable-pcre2-32 --enable-jit --disable-dependency-tracking  --disable-shared --enable-pcre2grep-libz --enable-pcre2grep-libbz2
    make
    make install
    popd
}

function build_macos() {
    if ! is_command brew; then
        echo "Homebrew is not installed"
        exit 1
    fi

    get_pcre2

    # Install dependencies using Homebrew.
    HOMEBREW_NO_AUTO_UPDATE=1 HOMEBREW_NO_INSTALLED_DEPENDENTS_CHECK=1 brew install autoconf automake libtool wget

    # Clean up.
    rm -rf "${PREFIX}"
    mkdir -p "${PREFIX}"

    export CC="clang"
    export CXX="clang"
    export CFLAGS=" -mmacosx-version-min=${MACOS_MIN_VERSION} -arch ${ARCH} -isysroot $SDKROOT "
    export LDFLAGS=" -mmacosx-version-min=${MACOS_MIN_VERSION} -arch ${ARCH} -isysroot $SDKROOT "
    export CPPFLAGS=" -mmacosx-version-min=${MACOS_MIN_VERSION} -arch ${ARCH} -isysroot $SDKROOT "
    export CXXFLAGS=" -mmacosx-version-min=${MACOS_MIN_VERSION} -arch ${ARCH} -isysroot $SDKROOT "

    build_pcre2
}

if is_macos; then
    MACOS_MIN_VERSION=${MACOS_MIN_VERSION:-11.0} # or 10.15
    PREFIX="${PREFIX_BASE}/pcre2-darwin-${ARCH}"
    pcre2_extra_configure_args=" --disable-shared "

    # Provide an appropriate SDK compatible with the macOS version we're building for.
    if [ "$SDKROOT" = "" ]; then
        SDKROOT=$(xcrun --sdk macosx --show-sdk-path)
    fi

    build_macos

elif is_linux; then
    PREFIX="${PREFIX_BASE}/pcre2-linux-${ARCH}"
    pcre2_extra_configure_args=" --disable-shared "

    get_pcre2
    build_pcre2

elif is_windows; then
    PREFIX="${PREFIX_BASE}/pcre2-windows-${MSYSTEM_CARCH}"
    pcre2_extra_configure_args=" --build=${MINGW_CHOST} --host=${MINGW_CHOST} --target=${MINGW_CHOST} "

    get_pcre2
    build_pcre2

else
    echo "Unsupported platform:$(uname)"
    exit 1
fi

cp pcre2/LICENCE "${PREFIX}/LICENSE"
cp pcre2/COPYING "${PREFIX}/COPYING"

# Clean up.
rm -rf "${PREFIX}/lib/pkgconfig" "${PREFIX}/share" "${PREFIX}/bin" "${PREFIX}/lib/"*.la
rm -rf pcre2
