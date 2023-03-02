package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func startXvnc() *exec.Cmd {
	xvnc := exec.Command(
		"Xvnc",
		os.Getenv("DISPLAY"),
		"-alwaysshared",
		"-depth",
		"16",
		"-geometry",
		os.Getenv("VNC_GEOMETRY"),
		"-securitytypes",
		"none",
		"-auth",
		fmt.Sprintf("%s/.Xauthority", os.Getenv("HOME")),
		"-fp",
		"catalogue:/etc/X11/fontpath.d",
		"-pn",
		"-rfbport",
		os.Getenv("VNC_PORT"))
	fmt.Println("Starting Xvnc")
	xvnc.Start()
	return xvnc
}

func waitForPort() {
	n := 1
	address := net.JoinHostPort("localhost", os.Getenv("VNC_PORT"))
	for n < 50 {
		conn, _ := net.Dial("tcp", address)
		if conn != nil {
			conn.Close()
			break
		}
		n++
		time.Sleep(10 * time.Millisecond)
	}
}

func startFluxbox() *exec.Cmd {
	fluxbox := exec.Command("fluxbox")
	fmt.Println("Starting fluxbox")
	fluxbox.Start()
	return fluxbox
}

func printSeleniumCombinedOutput(seleniumStdout io.ReadCloser) {
	scanner := bufio.NewScanner(seleniumStdout)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func startSelenium() *exec.Cmd {
	fmt.Println("Starting selenium standalone")
	selenium := exec.Command(
		"java",
                "-Dcom.redhat.fips=false",
                "-Djava.security.disableSystemPropertiesFile=true",
		"-jar",
		os.Getenv("SELENIUM_PATH"),
		"standalone",
		"--port",
		os.Getenv("SELENIUM_PORT"),
		"--session-timeout",
		os.Getenv("SELENIUM_SESSION_TIMEOUT"),
	)
	seleniumStdout, _ := selenium.StdoutPipe()
	selenium.Stderr = selenium.Stdout
	go printSeleniumCombinedOutput(seleniumStdout)
	selenium.Start()
	return selenium
}

func startProcesses() (*exec.Cmd, *exec.Cmd, *exec.Cmd) {
	xvnc := startXvnc()
	waitForPort()
	fluxbox := startFluxbox()
	selenium := startSelenium()
	return xvnc, fluxbox, selenium
}

func stopProcesses(xvnc *exec.Cmd, fluxbox *exec.Cmd, selenium *exec.Cmd) {
	fmt.Println("Stopping selenium")
	selenium.Process.Kill()
	selenium.Wait()
	fmt.Println("Stopping fluxbox")
	fluxbox.Process.Kill()
	fluxbox.Wait()
	fmt.Println("Stopping Xvnc")
	xvnc.Process.Kill()
	xvnc.Wait()
}

func main() {
	// start procs, then allow graceful shutdown with HTTP API or signal handler
	xvnc, fluxbox, selenium := startProcesses()

	// shutdown handler based on:
	// https://medium.com/@int128/shutdown-http-server-by-endpoint-in-go-2a0e2d7f9b8c
	// using signal.NotifyContext instead of context.WithCancel
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	m := http.NewServeMux()
	address := net.JoinHostPort("localhost", os.Getenv("API_PORT"))
	s := http.Server{Addr: address, Handler: m}
	m.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received shutdown request via HTTP")
		w.Write([]byte("OK"))
		stop()
	})
	go func() {
		fmt.Printf("HTTP server listening on '%s'\n", address)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	select {
	case <-ctx.Done():
		// Shutdown the HTTP server if context is cancelled
		// Context can either be cancelled by '/shutdown' handler or by SIGINT/SIGTERM
		s.Shutdown(ctx)
	}

	stopProcesses(xvnc, fluxbox, selenium)
	fmt.Println("Bye bye")
	os.Exit(0)
}
