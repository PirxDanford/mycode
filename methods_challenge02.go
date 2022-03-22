package main

import "fmt"

type Virtmach struct {
	name string
	environment string // e.g. DEV, QA, STAGING, PROD
	ipv4 string
	ipv6 string
	cpu string
	mem string
}

func (t Virtmach)PrintIpv4(){
	fmt.Println(t.ipv4)
}

func (t *Virtmach)SetIpv4(newval string){
	t.ipv4 = newval
}

func main(){
	vm1 := Virtmach{"VM1","DEV","192.168.0.1","0000:0000:0000:0000:0000:FFFF:C0A8:0001","AMD EPYC 7452","128GB"}
	fmt.Printf("%+v\n",vm1)
	vm1.PrintIpv4()
	vm1.SetIpv4("127.0.0.1")
	vm1.PrintIpv4()
}
