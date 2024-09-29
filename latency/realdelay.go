package latency

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/galavpncom/xrayping/config"
	"github.com/galavpncom/xrayping/utils"
	"github.com/galavpncom/xrayping/xray"

	"golang.org/x/net/proxy"
)

type Tester struct {
	URL     string
	Socks5  string
	Verbose bool
}

type IPResult struct {
	IP      string
	Latency time.Duration
}

func NewTester(url, socks5 string, verbose bool) *Tester {
	return &Tester{URL: url, Socks5: socks5, Verbose: verbose}
}

func (lt *Tester) MeasureLatency() (time.Duration, error) {
	socks5Dialer, err := proxy.SOCKS5("tcp", lt.Socks5, nil, proxy.Direct)
	if err != nil {
		return 0, fmt.Errorf("error setting up SOCKS5 proxy: %v", err)
	}

	client := &http.Client{
		Transport: &http.Transport{Dial: socks5Dialer.Dial},
		Timeout:   10 * time.Second,
	}

	start := time.Now()
	resp, err := client.Get(lt.URL)
	if err != nil {
		return 0, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	latency := time.Since(start)
	if lt.Verbose {
		fmt.Printf("%sLatency test completed: %d ms%s\n", utils.ColorYellow, latency.Milliseconds(), utils.ColorReset)
	}

	return latency, nil
}

func ProcessIP(ip string, configManager *config.Manager, xrayManager *xray.Manager, latencyTester *Tester, retries int, semaphore chan struct{}, wg *sync.WaitGroup, successIPs *[]IPResult, mu *sync.Mutex) {
	defer wg.Done()

	if err := configManager.UpdateConfig(ip); err != nil {
		fmt.Println(utils.WrapTextWithColor(fmt.Sprintf("Failed to update config for IP %s: %v", ip, err), utils.ColorRed))
		<-semaphore
		return
	}

	cmd, err := xrayManager.Start(configManager.Config)
	if err != nil {
		fmt.Println(utils.WrapTextWithColor(fmt.Sprintf("Failed to start Xray for IP %s: %v", ip, err), utils.ColorRed))
		<-semaphore
		return
	}

	var bestLatency time.Duration
	success := false

	for i := 1; i <= retries; i++ {
		if latency, err := latencyTester.MeasureLatency(); err == nil {
			success = true
			if bestLatency == 0 || latency < bestLatency {
				bestLatency = latency
			}
		}
	}

	if success {
		mu.Lock()
		*successIPs = append(*successIPs, IPResult{IP: ip, Latency: bestLatency})
		mu.Unlock()
		if latencyTester.Verbose {
			fmt.Println(utils.WrapTextWithColor(fmt.Sprintf("IP: %s, Best Latency: %d ms", ip, bestLatency.Milliseconds()), utils.ColorGreen))
		}
	}

	if err := xrayManager.Stop(cmd); err != nil {
		fmt.Println(utils.WrapTextWithColor(fmt.Sprintf("Failed to stop Xray for IP %s: %v", ip, err), utils.ColorRed))
	}

	<-semaphore
}
