package service

import "net"
import "io/ioutil"
import "os"
import "io"
import "bufio"
//import "fmt"

var texts []string
var f_read int = 0
func GetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ":("
	}
	for i := range addrs {
		ipNet, isIpNet := addrs[i].(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ":("
}

func GetFiles() []string {
	files, _ := ioutil.ReadDir("./static/files")
	var tmp []string
	for _, file :=  range files {
		tmp = append(tmp, file.Name())
	}
	return tmp
}

func GetText() []string {
	if f_read == 0 {
		file, err := os.Open("./static/text")
		if err != nil {
			return nil
		}
		reader := bufio.NewReader(file)
		for {
			text, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			text = text[0:len(text)-1]
			texts = append(texts, text)
		}
		f_read = 1
	}
	return texts
}

func FileType(filename string) string {
	return ":)"
}

func CleanAll() {
	
}

func WriteText(text string) int {
	texts = append(texts, text)
	file, err := os.OpenFile("./static/text", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		return 0
	}
	defer file.Close()
	file.WriteString(text)
	return 1
}
