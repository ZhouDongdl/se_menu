package linklist

import (
	"fmt"
	"strings"
)

func FindCmd(head *DataNode, cmd string) *DataNode {
	if head == nil || cmd == "" {
		return nil
	}
	var p *DataNode = head
	for p != nil {
		if strings.EqualFold(cmd, p.Cmd) {
			return p
		}
		p = p.Next
	}
	return nil
}

func ShowAllCmd(head *DataNode) {
	p := head
	for p != nil {
		if p.Cmd == "help" {
			p = p.Next
			continue
		}
		fmt.Printf("%s - %s\n", p.Cmd, p.Desc)
		p = p.Next
	}
}