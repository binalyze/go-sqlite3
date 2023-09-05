//go:build binalyze_ignore
// +build binalyze_ignore

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	goarch := flag.String("goarch", runtime.GOARCH, "go arch to test")
	workdir := flag.String("w", "", "base test directory path (default: current working directory)")
	timeout := flag.Duration("timeout", 0, "timeout for test execution (default: no timeout)")
	tags := flag.String("tags", "binalyze_sqlite3_all", "test build tags")
	flag.Parse()

	if *workdir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		*workdir = cwd
	}
	if p, e := filepath.Abs(*workdir); e == nil {
		*workdir = p
	}

	log.Println("workdir:", *workdir)

	goExePath, err := exec.LookPath("go")
	if err != nil {
		log.Fatalf("failed to find go: %v", err)
	}
	goExePath, err = filepath.Abs(goExePath)
	if err != nil {
		log.Fatalf("failed to get absolute path of go: %v", err)
	}

	ctx := context.Background()
	if *timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	switch runtime.GOOS {
	case "windows":
		err = testWindows(ctx, goExePath, *workdir, *goarch, *tags)
	case "darwin", "linux":
		err = testUnix(ctx, goExePath, *workdir, *goarch, *tags)
	default:
		err = fmt.Errorf("unsupported GOOS: %s", runtime.GOOS)
	}
	if err != nil {
		log.Fatalf("test failed: %v", err)
	}
}

func testUnix(ctx context.Context, goExePath, workDir, goarch, tags string) error {
	args := []string{"test", "-v", "-failfast", "-count=1", "-tags", tags, "./..."}

	cmd := newCmd(ctx, goExePath, args, workDir, goarch, nil)
	envs := cmd.Env
	if err := printGoEnvs(ctx, goExePath, envs); err != nil {
		return fmt.Errorf("failed to print go envs: %w", err)
	}

	log.Printf("running command: %s", cmd)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run go test command: %w", err)
	}
	return nil
}

func testWindows(ctx context.Context, goExePath, workDir, goarch, tags string) error {
	_, err := os.Lstat(`C:\msys64\usr\bin\bash.exe`)
	if err != nil {
		return fmt.Errorf("failed to find MSYS2 bash.exe: %w", err)
	}

	var msystem string
	switch goarch {
	case "amd64", "":
		msystem = "MINGW64"
	case "386":
		msystem = "MINGW32"
	default:
		return fmt.Errorf("unsupported GOARCH: %s", goarch)
	}

	cmd := newCmd(ctx, `C:\msys64\usr\bin\bash.exe`, []string{"-l"}, workDir, goarch,
		[]string{"CHERE_INVOKING=yes", "MSYSTEM=" + msystem})

	cmd.Stdin = strings.NewReader(fmt.Sprintf(`
set -exo pipefail
goexe=%q
"$goexe" env
"$goexe" test -v -failfast -count=1 -tags %q ./...
`, goExePath, tags))

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run go test command using MSYS2: %w", err)
	}
	return nil
}

func newCmd(ctx context.Context, name string, args []string, workDir, goarch string, addenvs []string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = workDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(append(os.Environ(), "CGO_ENABLED=1", "GOARCH="+goarch), addenvs...)
	return cmd
}

func printGoEnvs(ctx context.Context, goExePath string, envs []string) error {
	cmd := exec.CommandContext(ctx, goExePath, "env")
	cmd.Env = envs
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("running command: %s", cmd)
	return cmd.Run()
}
