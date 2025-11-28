package NIPI

import  "encoding/json"
import  "errors"
import  "fmt"
import  "io"
import  "net"
import  "time"

func    Forwarder (Conn net.Conn, Request map[string]any) (E error, Response map[string]any) {
	/***1***/
	xb05, _ := json.Marshal (Request)
	Conn.SetDeadline (time.Now ().Add (time.Minute*5))
	_, xb20 := Conn.Write (xb05)
	if xb20 != nil {
		E= errors.New (fmt.Sprintf (
			`Request write failed [%s]`, xb20.Error (),
		))
		return 
	}
	/***2***/
	xb25 := []byte { }
	for     {
		xc05 := make ([]byte, 1)
		_, xc15 := Conn.Read (xc05)
		if xc15 != nil && xc15 == io.EOF { continue }
		if xc15 != nil {
			E= errors.New (fmt.Sprintf (
				`Response read failed [%s]`, xc15.Error (),
			))
			return
		}
		xb25  = append (xb25, xc05... )
		if json.Valid  (xb25) { break }
	}
	/***3***/
	xb30 := json.Unmarshal (xb25,&Response)
	if xb30 != nil {
		E= errors.New (fmt.Sprintf (
			`Response unmarshal failed [%s]`, xb30.Error (),
		))
		return 
	}
	return
}
