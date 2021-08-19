package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
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
		os.Getenv("VNC_PORT"),
		"-rfbwait",
		"30000")
	fmt.Println("Starting Xvnc")
	xvnc.Start()
	return xvnc
}

func waitForPort() {
	timeout := 5 * time.Second
	n := 1
	for n < 5 {
		conn, _ := net.DialTimeout("tcp", net.JoinHostPort("localhost", os.Getenv("VNC_PORT")), timeout)
		if conn != nil {
			conn.Close()
			break
		}
		n++
		time.Sleep(time.Second)
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
	selenium := exec.Command("java", "-jar", os.Getenv("SELENIUM_PATH"), "-port", os.Getenv("SELENIUM_PORT"))
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

func waitForSignals() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
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
	xvnc, fluxbox, selenium := startProcesses()
	waitForSignals()
	stopProcesses(xvnc, fluxbox, selenium)
	fmt.Println("Bye bye")
	os.Exit(0)
}