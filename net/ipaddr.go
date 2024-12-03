package net

import (
	"os"
	"io/ioutil"
	"bufio"
	"strings"
	"errors"
	"bytes"
	extnet "net"
)

const (
	routerfile  = "/proc/net/route"
)


func parseLinuxProcNetRoute(f []byte) (string, error) {
        const (
                sep   = "\t" // field separator
                field = 2    // field containing hex gateway address
                flag = 3     // field containing hex gateway flag
                name = 0     // field containing interface name
        )

        scanner := bufio.NewScanner(bytes.NewReader(f))
        for scanner.Scan() {
                tokens := strings.Split(scanner.Text(), sep)
                if tokens[flag] == "0003" {
                        interfaceName := tokens[name]
                        return  interfaceName, nil
                }
        }

        return "", errors.New("Can not get gateway record")
}

func getGateInterFaceIPAddress(gatewayInterface string) (string, error) {
        interfaces, _ := extnet.Interfaces()

        for _, i := range interfaces {
                if i.Name == gatewayInterface {
                        byNameInterface, _ := extnet.InterfaceByName(i.Name)
                        addresses, _ := byNameInterface.Addrs()
                        if len(addresses) > 0 {
                                for _,addr := range addresses {
                                        ipv4AddrMask := addr.String()
                                        if strings.Count(ipv4AddrMask,":") < 2 {
                                                ipv4Addr, _, err := extnet.ParseCIDR(ipv4AddrMask)
                                                if err != nil {
                                                        return "", errors.New("Failed to change into IP")
                                                }
                                                return ipv4Addr.String(), nil
                                        }
                                }
                        }else{
                                return "", errors.New("Failed to get IP")
                        }
                }
        }
        return "", errors.New("Failed to get interface")
}

func GetLinuxIPaddress()(string, error) {
        f, _:= os.Open(routerfile)
        defer f.Close()
        bytes, _ := ioutil.ReadAll(f)
        interfaceName, err := parseLinuxProcNetRoute(bytes)
        if err != nil {
                return "", err
        }
        addr, err := getGateInterFaceIPAddress(interfaceName)
        if err != nil {
                return "", err
        }
        return addr, err
}

