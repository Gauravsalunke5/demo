package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

//=============== Util Methods ===========================

var testLog = shim.NewLogger("securecert-chaincode_test")

//checkState
func checkState(t *testing.T, stub *shim.MockStub, name string) {
	bytes := stub.State[name]
	if bytes == nil {
		testLog.Info("State", name, "failed to get value")
		t.FailNow()
	}
	testLog.Info("State value", name, "is", string(bytes), "as expected")
}

//checkQuery
func checkQuery(t *testing.T, stub *shim.MockStub, function string, argument string) {
	res := stub.MockInvoke("1", [][]byte{[]byte(function), []byte(argument)})
	if res.Status != shim.OK {
		testLog.Info("Query", function, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		testLog.Info("Query", function, "failed to get value")
		t.FailNow()
	}
	payload := string(res.Payload)

	testLog.Info("Query value", function, "is", payload, "as expected")
}

//checkInvoke
func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
	fmt.Println("Invoke successfully", string(res.Message))
}

//checkInvokeError
func checkInvokeError(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.ERROR {
		fmt.Println("Invoke", args, "success", string(res.Message))
		t.FailNow()
	}
	fmt.Println("Invoke failed", string(res.Message))
}

//================= test cases =======================

//test initLedger
func Test_initLedger(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex01", scc)
	fmt.Println("================Test initLedger==========================")
	fmt.Println("invoke initLedger:")
	checkInvoke(t, stub, [][]byte{[]byte("initLedger")})
	fmt.Println("check state:")
	checkState(t, stub, "201916721")
	fmt.Println("================End ======================================")
	fmt.Println("")
}

//test Student
func Test_Student(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex01", scc)
	fmt.Println("")
	fmt.Println("===================Test Student==============")
	// addStudent
	fmt.Println("create student:")
	checkInvoke(t, stub, [][]byte{[]byte("addStudent"), []byte("201516383"), []byte("201516383"),
		[]byte("Gaurav"), []byte("U"), []byte("Salunke"), []byte("PCCE"), []byte("IT"), []byte("2015"),
		[]byte("gauravsal15@gmail.com"), []byte("8007067665")})
	// readStudent
	fmt.Println("query student:")
	checkQuery(t, stub, "readStudent", "201516383")
	fmt.Println("=====================end========================= ")
	fmt.Println("")
}

//test Certificate
func Test_Certificate(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex01", scc)
	fmt.Println("")
	fmt.Println("===================Test Certificate==============")
	// addCert
	fmt.Println("create Certificate:")
	checkInvoke(t, stub, [][]byte{[]byte("addCert"), []byte("201516383"), []byte("PCCE"), []byte("101"),
		[]byte("nov/dec"), []byte("2018"), []byte("abc")})
	// readCert
	fmt.Println("query Certificate:")
	checkQuery(t, stub, "readCert", "101")
	//transferCert
	fmt.Println("transfer Certificate:")
	checkInvoke(t, stub, [][]byte{[]byte("transferCert"), []byte("101"), []byte("Gaurav")})
	// readCert
	fmt.Println("query Certificate:")
	checkQuery(t, stub, "readCert", "101")
	fmt.Println("======================End========================= ")
	fmt.Println("")
}

//test Invoke
func Test_Invoke(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex01", scc)
	fmt.Println("")
	fmt.Println("=====================Test Invoke=================")
	fmt.Println("create student with key \"201516383\":")
	checkInvoke(t, stub, [][]byte{[]byte("addStudent"), []byte("201516383"), []byte("201516383"),
		[]byte("Gaurav"), []byte("U"), []byte("Salunke"), []byte("PCCE"), []byte("IT"), []byte("2015"),
		[]byte("gauravsal15@gmail.com"), []byte("8007067665")})
	fmt.Println("create student with key \"201516383\":")
	checkInvokeError(t, stub, [][]byte{[]byte("addStudent"), []byte("201516383"), []byte("201516383"),
		[]byte("Gaurav"), []byte("U"), []byte("Salunke"), []byte("PCCE"), []byte("IT"), []byte("2015"),
		[]byte("gauravsal15@gmail.com"), []byte("8007067665")})
	fmt.Println("create certificate with key \"101\":")
	checkInvoke(t, stub, [][]byte{[]byte("addCert"), []byte("201516383"), []byte("PCCE"), []byte("101"),
		[]byte("nov/dec"), []byte("2018"), []byte("abc")})
	fmt.Println("create certificate with key \"101\":")
	checkInvokeError(t, stub, [][]byte{[]byte("addCert"), []byte("201516383"), []byte("PCCE"), []byte("101"),
		[]byte("nov/dec"), []byte("2018"), []byte("abc")})
	fmt.Println("transfer certificate with invalid key \"100\":")
	checkInvokeError(t, stub, [][]byte{[]byte("transferCert"), []byte("100"), []byte("Gaurav")})
	fmt.Println("=====================end========================= ")
	fmt.Println("")
}
