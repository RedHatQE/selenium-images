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
		"30000",
	)
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

func seleniumStandalone() *exec.Cmd {
	return exec.Command(
		"java",
		"-jar",
		os.Getenv("SELENIUM_PATH"),
		"standalone",
		"--host",
		"0.0.0.0",
		"--port",
		os.Getenv("SELENIUM_PORT"),
	)
}

func seleniumNode() *exec.Cmd {
	return exec.Command(
		"java",
		"-jar",
		os.Getenv("SELENIUM_PATH"),
		"node",
		"--publish-events",
		fmt.Sprintf("tcp://%s:%s", os.Getenv("SE_EVENT_BUS_HOST"), os.Getenv("SE_EVENT_BUS_PUBLISH_PORT")),
		"--subscribe-events",
		fmt.Sprintf("tcp://%s:%s", os.Getenv("SE_EVENT_BUS_HOST"), os.Getenv("SE_EVENT_BUS_SUBSCRIBE_PORT")),
	)
}

func startSelenium() *exec.Cmd {
	fmt.Printf("Starting selenium %s\n", os.Getenv("SELENIUM_MODE"))
	var funcMap = map[string]func() *exec.Cmd{
		"standalone": seleniumStandalone,
		"node":       seleniumNode,
	}
	selenium := funcMap[os.Getenv("SELENIUM_MODE")]()
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

func checkVars() {
	if os.Getenv("SELENIUM_MODE") == "node" {
		_, seEventBusHostPresent := os.LookupEnv("SE_EVENT_BUS_HOST")
		if !seEventBusHostPresent {
			fmt.Println("SE_EVENT_BUS_HOST not set, exiting!")
			os.Exit(1)
		}
		_, seEventBusPublishPort := os.LookupEnv("SE_EVENT_BUS_PUBLISH_PORT")
		if !seEventBusPublishPort {
			fmt.Println("SE_EVENT_BUS_PUBLISH_PORT not set, exiting!")
			os.Exit(1)
		}
		_, seEventBusSubscribePort := os.LookupEnv("SE_EVENT_BUS_SUBSCRIBE_PORT")
		if !seEventBusSubscribePort {
			fmt.Println("SE_EVENT_BUS_SUBSCRIBE_PORT not set, exiting!")
			os.Exit(1)
		}
	}
}

func main() {
	checkVars()
	xvnc, fluxbox, selenium := startProcesses()
	waitForSignals()
	stopProcesses(xvnc, fluxbox, selenium)
	fmt.Println("Bye bye")
	os.Exit(0)
}
