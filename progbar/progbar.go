package progbar

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Display(now int, max int) (string, error) {
	if max < now {
		return "", errors.New("input numer value err")
	}

	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	outs := strings.Replace(string(out), "\n", "", -1)
	outs = strings.Split(outs, " ")[1]
	allWidth, err := strconv.Atoi(outs)
	if err != nil {
		return "", errors.New("terminal width get error")
	}

	width := allWidth - (14 + len(strconv.Itoa(max))*2)
	cor := now * 100 / max
	bar := fmt.Sprintf("\r[%%-%vv] %%-3v", width)
	echo1 := fmt.Sprintf(bar, strings.Repeat("#", width*cor/100), cor)
	perc := fmt.Sprintf("(%v/%v)", now, max)
	pcmd := fmt.Sprintf("%%%vv", allWidth-width-7)
	echo2 := fmt.Sprintf(pcmd, perc)
	echo := "\r" + echo1 + "%" + echo2
	return echo, nil
}
