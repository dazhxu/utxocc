/*
Author: xuyuzhuang
Date: 2018-02-27
*/
package main

import (
    fmt

    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//Utxo Chaincode 
type UtxoChaincode struct {
}

func (u *UtxoChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Utxocc Init")
	_, args := stub.GetFunctionAndParameters()
	//TODO 添加初始化逻辑

	return shim.Success(nil)
}

func (u *UtxoChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Utxocc Invoke")
	function, args := stub.GetFunctionAndParameters()

    //TODO 添加处理方法逻辑
	switch function {
	case "":
		//TODO
	}

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(UtxoChaincode))
	if err != nil {
		fmt.Println("Error when starting Utxo chaincode: %s", err)
	}
}
