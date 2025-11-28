package NIPI

import  "fmt"

func    SR15 (Conn  *Conn, ID string) (
	ExctnOtcmCode int, ExctnOtcmNote string, Yield string,
	)    {
	/***1***/
	xb05 := map[string]any {}
	xb05 ["AuthKey"] = Conn.atKy
	xb05 ["ID"] = ID
	/***2***/
	ExctnOtcmCode, ExctnOtcmNote, xb15 := Conn.Write ("15", xb05)
	if ExctnOtcmCode == 200 {
		Yield = xb15["ID"].(string)
	}
	return 
}
func    SR20 (Conn  *Conn, ID string) (
	ExctnOtcmCode int, ExctnOtcmNote string,
	)    {
	/***1***/
	xb05 := map[string]any {}
	xb05 ["AuthKey"] = Conn.atKy
	xb05 ["ID"] = ID
	/***2***/
	ExctnOtcmCode, ExctnOtcmNote, xb15 := Conn.Write ("20", xb05)
	_=xb15
	return
}
func    SR25 (Conn  *Conn, ID, Password string) (
	ExctnOtcmCode int, ExctnOtcmNote string,
	)    {
	/***1***/
	xb05 := map[string]any { }
	xb05 ["AuthKey" ] = Conn.atKy
	xb05 ["ID" ] = ID
	xb05 ["Password"] = Password
	/***2***/
	ExctnOtcmCode, ExctnOtcmNote, xb15 := Conn.Write ("25", xb05)
	_=xb15
	return
}
func    SR30 (Conn  *Conn, ID , Password string, SVD int) (
	ExctnOtcmCode int, ExctnOtcmNote string, Stts bool, SsID, SsKy string,
	)    {
	/***1***/
	xb05 := map[string]any { }
	xb05 ["AuthKey" ] = Conn.atKy
	xb05 ["ID" ] = ID
	xb05 ["Password"] = Password
	xb05 ["SessionValidityDuration"] = fmt.Sprintf ("%d", SVD)
	/***2***/
	ExctnOtcmCode, ExctnOtcmNote, xb15 := Conn.Write ("30", xb05)
	if ExctnOtcmCode == 200 {
		Stts = xb15["Status"].(bool)
	}
	if ExctnOtcmCode == 200 && Stts {
		SsID = xb15["SessionID"].(string)
		SsKy = xb15["SessionKy"].(string)
	}
	return 
}
func    SR35 (Conn  *Conn, Entity, SessionID, SessionKy string) (
	ExctnOtcmCode int, ExctnOtcmNote string, Yield bool,
	)    {
	/***1***/
	xb05 := map[string]any { }
	xb05 ["AuthKey  "] = Conn.atKy
	xb05 ["Entity"   ] = Entity
	xb05 ["SessionID"] = SessionID
	xb05 ["SessionKy"] = SessionKy
	/***2***/
	ExctnOtcmCode, ExctnOtcmNote, xb15 := Conn.Write ("35", xb05)
	if ExctnOtcmCode == 200 {
		Yield = xb15["Status"].(bool)
	}
	return 
}
func    SR40 (Conn  *Conn, Entity, SessionID, SessionKy string) (
	ExctnOtcmCode int, ExctnOtcmNote string,
	)    {
	/***1***/
	xb05 := map[string]any { }
	xb05 ["AuthKey"  ] = Conn.atKy
	xb05 ["Entity"   ] = Entity
	xb05 ["SessionID"] = SessionID
	/***2***/
	ExctnOtcmCode, ExctnOtcmNote, xb15 := Conn.Write ("40", xb05)
	_=xb15
	return
}
