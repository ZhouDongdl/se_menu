/**************************************************************************************************/
/* Copyright (C) zhoujinlong, SSE@USTC, 2023                                                      */
/*                                                                                                */
/*  FILE NAME             :  menu.go                                                              */
/*  PRINCIPAL AUTHOR      :  Zhoujinlong                                                          */
/*  SUBSYSTEM NAME        :  menu                                                                 */
/*  MODULE NAME           :  menu                                                                 */
/*  LANGUAGE              :  golang                                                               */
/*  TARGET ENVIRONMENT    :  ANY                                                                  */
/*  DATE OF FIRST RELEASE :  2023/03/31                                                           */
/*  DESCRIPTION           :  This is a menu program                                               */
/**************************************************************************************************/

/*
version:2.0
*/



package main

import (
	"fmt"
	"os"
	"se_menu/linklist"
)

const (
	CMD_MAX_LEN = 128
	DESC_LEN = 1024
 	CMD_NUM = 20
)

var head = &linklist.DataNode{}
var cmds = []*linklist.DataNode{
    {"help", "this is help cmd!", Help, nil},
    {"version", "menu v2.0", nil, nil},
    {"quit", "exit menu", Exit, nil},
}

func Help() {
	linklist.ShowAllCmd(head)
}

func Exit() {
	os.Exit(0);
}

func main() {
	tail := head
	for _, v := range cmds {
		tail.Next = v
		tail = tail.Next
	}
	head = head.Next
	var cmd string
	for true {
		fmt.Printf("Please enter your command:\n")
		fmt.Scanln(&cmd)
		p := linklist.FindCmd(head, cmd)
		if p == nil {
			fmt.Println("This cmd is not exist")
		} else {
			fmt.Printf("%s - %s\n", p.Cmd, p.Desc)
			if p.Handler != nil {
				p.Handler()
			}
		}
	}
}
