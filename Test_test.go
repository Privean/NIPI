package NIPI

import  "fmt"
import _"github.com/google/uuid"
import  "sync"
import  "testing"

func    TestMain (T *testing.T) {
	/***1***/
//	xb05  , xb10 := Conn_Create ("127.0.0.1", "8558", false, "", 1)
	xb05  , xb10 := Conn_Create ("am.ams.lytup.qbqevell.ng", "10021", true, "", 10)
	if xb05 != nil {
		xc05 := fmt.Sprintf (`Conn creation failed [%s]`, xb05.Error ())
		fmt.Println (xc05)
		return
	}
	/***2***/
	xb12 := "c6baff6c-a401-4ac7-aa21-a2cdf1d3aa73"
	xb15 := 20
	xb20 := &sync.WaitGroup {}
	xb20.Add (xb15)
	for xc05 := 1; xc05 <= xb15; xc05 ++ {
		go func () {
			defer  xb20.Done  ()
			xd15 , xd20, xd25 := SR15 (xb10, xb12)
			if xd15 != 200 {
				xe05 := fmt.Sprintf (`Request 15 failed [%d : %s]`, xd15, xd20)
				fmt.Println (xe05)
				return
			}
			fmt.Println (xd25)
		} ()
	}
	xb20.Wait ()
	/***4***/
	/*
	xb25 , xb30 := SR20 (xb10, xb12)
	if xb25 != 200 {
		xc05 := fmt.Sprintf (`Request 20 failed [%d : %s]`, xb25, xb30)
		fmt.Println (xc05)
		return
	}
	*/
	/***5***/
	/*
	xb35 , xb40 := SR25 (xb10, xb12, "UGFzcyEh")
	if xb35 != 200 {
		xc05 := fmt.Sprintf (`Request 25 failed [%d : %s]`, xb35, xb40)
		fmt.Println (xc05)
		return
	}
	*/
	/***6***/
	/*
	xb35 , xb40 , xb45 , xb50 , xb55 :=SR30 (xb10, xb12, "UGFzcyEh", 30)
	if xb35 != 200 {
		xc05 := fmt.Sprintf (`Request 30 failed [%d : %s]`, xb35, xb40)
		fmt.Println (xc05)
		return
	}
	fmt.Println ( xb45 , xb50 , xb55 )
	*/
	/***7***/
	/*
	xb75 , xb80 :=SR40 ( xb10 , xb12 , xb50, xb55)
	if xb75 != 200 {
		xc05 := fmt.Sprintf (`Request 40 failed [%d : %s]`, xb75, xb80)
		fmt.Println (xc05)
		return
	}
	*/
	/***8***/
	/*
	xb60 , xb65 , xb70 :=SR35 ( xb10 , xb12, xb50, xb55)
	if xb60 != 200 {
		xc05 := fmt.Sprintf (`Request 35 failed [%d : %s]`, xb60, xb65)
		fmt.Println (xc05)
		return
	}
	fmt.Println ( xb70 )
	*/
}
