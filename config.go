package fdfs_client

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type config struct {
	trackerAddr []string
	maxConns    int
}

func newConfig(configName string) (*config, error) {
	config := &config{}
	f, err := os.Open(configName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\r")
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			continue
		}
		str := strings.SplitN(line, "=", 2)
		switch str[0] {
		case "tracker_server":
			config.trackerAddr = append(config.trackerAddr, str[1])
		case "maxConns":
			config.maxConns, err = strconv.Atoi(str[1])
			if err != nil {
				return nil, err
			}
		}
		if err != nil {
			if err == io.EOF {
				return config, nil
			}
			return nil, err
		}
	}
	return config, nil
}
