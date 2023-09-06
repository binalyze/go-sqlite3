package regexp

/*
#cgo CFLAGS: -I${SRCDIR}/../../ -DPCRE2_STATIC -DPCRE2_CODE_UNIT_WIDTH=8 -DSQLITE_CORE
#cgo linux,amd64 CFLAGS: -I${SRCDIR}/../build/pcre2-linux-x86_64/include
#cgo linux,amd64 LDFLAGS: -static -L${SRCDIR}/../build/pcre2-linux-x86_64/lib -lpcre2-8
#cgo linux,386 CFLAGS: -I${SRCDIR}/../build/pcre2-linux-i686/include
#cgo linux,386 LDFLAGS: -static -L${SRCDIR}/../build/pcre2-linux-i686/lib -lpcre2-8
#cgo linux,arm64 CFLAGS: -I${SRCDIR}/../build/pcre2-linux-aarch64/include
#cgo linux,arm64 LDFLAGS: -static -L${SRCDIR}/../build/pcre2-linux-aarch64/lib -lpcre2-8
#cgo darwin,amd64 CFLAGS: -I${SRCDIR}/../build/pcre2-darwin-x86_64/include
#cgo darwin,amd64 LDFLAGS: ${SRCDIR}/../build/pcre2-darwin-x86_64/lib/libpcre2-8.a
#cgo darwin,arm64 CFLAGS: -I${SRCDIR}/../build/pcre2-darwin-arm64/include
#cgo darwin,arm64 LDFLAGS: ${SRCDIR}/../build/pcre2-darwin-arm64/lib/libpcre2-8.a
#cgo windows,amd64 CFLAGS: -I${SRCDIR}/../build/pcre2-windows-x86_64/include
#cgo windows,amd64 LDFLAGS: -static ${SRCDIR}/../build/pcre2-windows-x86_64/lib/libpcre2-8.a
#cgo windows,386 CFLAGS: -I${SRCDIR}/../build/pcre2-windows-i686/include
#cgo windows,386 LDFLAGS: -static ${SRCDIR}/../build/pcre2-windows-i686/lib/libpcre2-8.a

#include "extension.h"

static int sqlite3_regexp_init(sqlite3* db, char** errmsg_ptr, const sqlite3_api_routines* api) {
  return regexp_init(db);
}

static void init() {
  sqlite3_auto_extension((void*)sqlite3_regexp_init);
}

*/
import "C"

func init() {
	C.init()
}
