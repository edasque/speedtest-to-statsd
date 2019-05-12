package main

/*
while true; do
./main
sleep 60
done
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"gopkg.in/alexcesaro/statsd.v2"
)

func main() {

	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	type Config struct {
		Statsdhost string `json:"statsdhost"`
		Statsdport string `json:"statsdport"`
	}

	conf := Config{}
	json.Unmarshal(file, &conf)

	var uri = conf.Statsdhost + ":" + conf.Statsdport
	fmt.Println("StatsD URI:", uri)

	c, err := statsd.New(statsd.Address(uri)) // Connect to the UDP port 8125 by default.
	if err != nil {
		// If nothing is listening on the target port, an error is returned and
		// the returned client does nothing but is still usable. So we can
		// just log the error and go on.
		log.Print(err)
	}

	speedtestCmdResult := exec.Command("/usr/local/bin/speedtest-cli", "--csv")
	speedtestOut, err := speedtestCmdResult.Output()
	if err != nil {
		panic(err)
	}

	splitString := strings.Split(string(speedtestOut), ",")
	fmt.Println("Output:", splitString)
	fmt.Println("Length:", len(splitString))

	// Server ID,Sponsor,Server Name,Timestamp,Distance,Ping,Download,Upload
	if len(splitString) == 11 {
		distanceString := splitString[5]
		pingString := splitString[6]
		downloadString := splitString[7]
		uploadString := splitString[8]

		var distance, _ = strconv.ParseFloat(distanceString, 64)
		var ping, _ = strconv.ParseFloat(pingString, 64)
		var download, _ = strconv.ParseFloat(downloadString, 64)
		var upload, _ = strconv.ParseFloat(strings.TrimSpace(uploadString), 64)

		fmt.Println("Distance:", distance)
		fmt.Println("Ping:", ping)
		fmt.Println("Download:", download)
		fmt.Println("Upload:", upload)

		c.Gauge("speedtest.distance", distance)
		c.Gauge("speedtest.ping", ping)
		c.Gauge("speedtest.download", download)
		c.Gauge("speedtest.upload", upload)

	}

	defer c.Close()

}
