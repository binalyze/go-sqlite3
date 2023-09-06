# Binalyze SQLite3 Extensions

We provide a set of SQLite3 extensions. As of now, we have the following C extensions: regexp and stats. These extensions are statically compiled and automatically loaded by the SQLite3 library. Extensions are enabled using build tags.

**To enable all extensions, use `binalyze_sqlite3_all` build tag.**

## Regular Expression Functions
- `regexp` statement
- `regexp_like`
- `regexp_substr`
- `regexp_capture`
- `regexp_replace`

See regexp package unit tests for more examples in regexp_test.go.

To enable `regexp` extension, use `binalyze_sqlite3_regexp` build tag.

## Statistics Functions
- `stddev`
- `stddev_samp`
- `stddev_pop`
- `variance`
- `var_samp`
- `var_pop`
- `median`
- `percentile`
- `percentile_25`
- `percentile_75`
- `percentile_90`
- `percentile_95`
- `percentile_99`

See stats package unit tests for more examples in stats_test.go.

To enable `stats` extension, use `binalyze_sqlite3_stats` build tag.

## Maintainance

github.com/binalyze/go-sqlite module is a fork of github.com/mattn/go-sqlite3 module. We will try to keep it up to date with the original module. If you see that we are behind, please open an issue.

We will assign the same tag to the forked module as the original module. For example, if the original module has a tag `v1.14.0`, we will assign the same tag to the forked module. This way, you can use the forked module as a drop-in replacement for the original module. But if there is a change after tagging, we will assign a new tag and only increment the patch version. For example, if the original module has a tag `v1.14.0` and there is a change after the tag, we will assign a new tag `v1.14.1` to the forked module.

### Building C Libraries

As of know, only pcre2 library is required to build regexp extension. To build pcre2 library, you should use the build scripts in pcre2-formula folder for supported platforms.

Building C libraries statically is important for us. Because we want to distribute a single binary. We don't want to distribute a binary with a dependency on a shared library.

As of now, we support Linux (386, amd64, arm64), Windows (386, amd64), and macOS (amd64, arm64) platforms. We will add more platforms in the future so update the build scripts accordingly. After building the libraries, you should copy them to the related folders in the project.

Lastly, you should build required C libraries like pcre2 on native platforms with minimum supported versions.

## License

This project is licensed under the MIT License - see the [LICENSE](../LICENSE) file for details. C libraries are licensed under their own licenses. See the related folders for more information.
