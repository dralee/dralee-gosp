/**
package main
2025.6.20 by dralee
抓包
*/

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	// 网卡
	device := "enp2s0" //"eth0"

	// 打开网卡进行抓包
	handler, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()

	// 设置过滤器，如只抓TCP包
	err = handler.SetBPFFilter("tcp")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("开始抓包，按Ctrl+C退出")

	// 开始抓包
	packetSource := gopacket.NewPacketSource(handler, handler.LinkType())
	for packet := range packetSource.Packets() {
		printPacketInfo(packet)
	}
}

func printPacketInfo(packet gopacket.Packet) {
	// 解析网络层TCP
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IP:", ipLayer)
		return
	}

	ip, _ := ipLayer.(*layers.IPv4)

	// 解析传输层 TCP
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer == nil {
		return
	}

	tcp, _ := tcpLayer.(*layers.TCP)

	flags := ""
	if tcp.SYN {
		flags += "SYN "
	}
	if tcp.ACK {
		flags += "ACK "
	}
	if tcp.FIN {
		flags += "FIN "
	}
	if tcp.RST {
		flags += "RST "
	}
	if tcp.PSH {
		flags += "PSH "
	}
	if tcp.URG {
		flags += "URG "
	}
	if tcp.ECE {
		flags += "ECE "
	}
	if tcp.CWR {
		flags += "CWR "
	}

	fmt.Printf("[%s] %s:%d -> %s:%d Flags:%s Seq:%d Ack:%d\n",
		time.Now().Format("2006-01-02 15:04:05"),
		ip.SrcIP, tcp.SrcPort, ip.DstIP, tcp.DstPort,
		flags, tcp.Seq, tcp.Ack)
}
