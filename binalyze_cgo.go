//go:build binalyze_sqlite3_all || binalyze_sqlite3_regexp || binalyze_sqlite3_stats
// +build binalyze_sqlite3_all binalyze_sqlite3_regexp binalyze_sqlite3_stats

package sqlite3

import (
	_ "github.com/mattn/go-sqlite3/binalyze"
)

/*
#ifdef USE_LIBSQLITE3
#error "USE_LIBSQLITE3 is not supported with Binalyze sqlite3 extensions"
#endif
*/
import "C"
