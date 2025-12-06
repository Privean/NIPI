package NIPI

import  "crypto/tls"
import  "errors"
import  "fmt"
import  "net"
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
	C=&Conn {
		atKy: AuthKey,
		cmps_addr: RemoteAddr,
		cmps_port: RemotePort,
		cmps_scrt: Security,
		cmps_auth: AuthKey,
	}
	/***3***/
	return
}
type    Conn struct {
	atKy string
	cmps_addr string
	cmps_port string
	cmps_scrt bool
	cmps_auth string
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
	xB11 := 0
	XB11  :
	xb11  , xb12:= net.DialTimeout (
		"tcp", c.cmps_addr+":"+c.cmps_port, time.Second*15,
	)
	xB11  = xB11 + 1
	if xb12 != nil && xB11 <= 5 { goto XB11 }
	if xb12 != nil {
		C= 500
		N= fmt.Sprintf (`Conn creation failed [%s]`, xb12.Error ())
		return
	}
	xb13 :=&tls.Config {ServerName:c.cmps_addr}
	if c.cmps_scrt == false { xb13.InsecureSkipVerify = true }
	xb15 := tls.Client (xb11, xb13)
	/***3***/
	xb20  , xb25 :=Forwarder (xb15, xb10)
	if xb20 != nil && xB11 < 5 { goto XB11 }
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
