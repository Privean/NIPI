package NIPI

import  "container/list"
import  "crypto/tls"
import  "errors"
import  "fmt"
import  "net"
import  "sync"
import  "time"

func    Conn_Create (
	RemoteAddr, RemotePort string, Security bool, AuthKey string, ChannelSize int,
	)       (
	E error,C *Conn,
	)       {
	/***1***/
	if ChannelSize < 1 {
		E = errors.New (fmt.Sprintf (`Channel size can not be less than '1'`) )
		return
	}
	/***2***/
	xb05 := ""
	C=&Conn {
		tknn: make(chan bool, ChannelSize) ,
		pool: list.New (),
		mtxx: &sync.Mutex {},
		atKy: AuthKey,
	}
	for xc05 := 1; xc05 <= ChannelSize; xc05++ {
		xc15 , xc20 := net.DialTimeout (
			"tcp", RemoteAddr+":"+RemotePort, time.Minute*1,
		)
		if xc20 != nil { xb05 = xc20.Error (); continue }
		xc25:=&tls.Config {ServerName:RemoteAddr}
		if Security == false { xc25.InsecureSkipVerify = true }
		xc15 = tls.Client (xc15, xc25)
		C.pool.PushBack (xc15)
		C.tknn  <- true
	}
	if C.pool.Len() == 0  {
		E = errors.New (fmt.Sprintf (
			`Channel creation failed [Last Error: %s]`, xb05,
		) )
		return
	}
	/***3***/
	return
}
type    Conn struct {
	pool *list.List
	mtxx *sync.Mutex
	atKy string
	tknn chan bool
}
func(c *Conn) Write (MssgType string, MssgSeed map[string]any) (
	C int, N string, Y map[string]any,
	)   {
	/***1***/
	xb05 := ""
	if        MssgType == "15" { xb05 = "Entt:rgst"
	} else if MssgType == "20" { xb05 = "Entt:dltt"
	} else if MssgType == "25" { xb05 = "Entt[pswr]:updt"
	} else if MssgType == "30" { xb05 = "Entt[pswr]:athn"
	} else if MssgType == "35" { xb05 = "Entt/Sssn:athn"
	} else if MssgType == "40" { xb05 = "Entt/Sssn:dltt"
	}
	xb10 := map[string] any {  }
	xb10["Service"]  = xb05
	xb10["Seed"] = MssgSeed
	/***2***/
	_  = <- c.tknn
	c.mtxx.Lock ()
	xb15 := c.pool.Front ()
	c.pool.Remove (xb15)
	conn := xb15.Value.(net.Conn)
	c.mtxx.Unlock ()
	defer func () {
		c.mtxx.Lock ()
		c.pool.PushBack(conn)
		c.mtxx.Unlock ()
		c.tknn <- true
	} ( )
	/***3***/
	xb20  , xb25 := Forwarder (conn, xb10)
	if xb20 != nil {
		C= 500
		N= fmt.Sprintf (`Request forwarding failed [%s]`, xb20.Error ())
		return 
	}
	/***4***/
	C=int ( xb25["ExctnOutcomeCode"].(float64) )
	xb30 := xb25["ExctnOutcomeNote"]
	if xb30 != nil {N=xb25["ExctnOutcomeNote"].(string)}
	xb35 := xb25["Yield"]
	if xb35 != nil {Y=xb25["Yield" ].(map[string]any)}
	return
}
func(c *Conn) Water () (E error) {
	c.mtxx.Lock ()
	xb05 := [ ]net.Conn {}
	defer  func ()  {
		for _, xd05 := range xb05 { c.pool.PushBack (xd05) }
		c.mtxx.Unlock  ( )
	} ( )
	xb10 := map[string]any { }
	xb10["Service"] = ""
	xb15 := c.pool.Len()
	for  xc05 := 1; xc05 <= xb15; xc05++ {
		xc10 := c.pool.Front ()
		c.pool.Remove (xc10)
		xc15 := xc10.Value.(net.Conn)
		Forwarder(xc15,xb10)
		xb05  = append(xb05,xc15)
	}
	return
}
