// Package main 这个文件用于测试github.com/hallidave/mibtool提供的读取MIB文件到golang中的功能。
package main

import (
	"fmt"
	"testing"

	"github.com/hallidave/mibtool/smi"
)

func TestMIB(t *testing.T) {
	mib := smi.NewMIB("mibs")
	mib.Debug = true
	fmt.Printf("%+v\n", mib)
	err := mib.LoadModules("IF-MIB")
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Printf("%+v\n", mib)

	// Walk all symbols in MIB
	mib.VisitSymbols(func(sym *smi.Symbol, oid smi.OID) {
		fmt.Printf("%-40s %s\n", sym, oid)
	})

	// Look up OID for an OID string
	oid, err := mib.OID("IF-MIB::ifOperStatus.4")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(oid.String())
}
