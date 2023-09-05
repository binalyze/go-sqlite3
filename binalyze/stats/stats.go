package stats

/*
#cgo CFLAGS: -I${SRCDIR}/../../ -DSQLITE_CORE
#cgo linux windows LDFLAGS: -static -lm
#cgo darwin LDFLAGS: -lm

#include "extension.h"

static int sqlite3_stats_init(sqlite3* db, char** errmsg_ptr, const sqlite3_api_routines* api) {
  return stats_init(db);
}

static void init() {
  sqlite3_auto_extension((void*)sqlite3_stats_init);
}

*/
import "C"

func init() {
	C.init()
}
