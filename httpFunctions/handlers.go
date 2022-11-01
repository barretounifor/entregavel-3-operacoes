package httpFunctions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func HandleSum(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	var requestObject Request

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
	}

	err = json.Unmarshal(requestBody, &requestObject)

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
	}

	operationResult := requestObject.First + requestObject.Second

	var responseResult Result

	responseResult.Result = operationResult

	_, err = json.Marshal(responseResult)

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
	}

	w.Write([]byte("O resultado da soma entre " + strconv.FormatFloat(requestObject.First, 'G', -1, 64) + " e " + strconv.FormatFloat(requestObject.Second, 'G', -1, 32) + " é " + strconv.FormatFloat(operationResult, 'G', -1, 32)))
}

func HandleSub(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	var requestObject Request

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
	}

	err = json.Unmarshal(requestBody, &requestObject)

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
	}

	operationResult := requestObject.First - requestObject.Second

	var responseResult Result

	responseResult.Result = operationResult

	_, err = json.Marshal(responseResult)

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
	}

	w.Write([]byte("O resultado da subtração entre " + strconv.FormatFloat(requestObject.First, 'G', -1, 64) + " e " + strconv.FormatFloat(requestObject.Second, 'G', -1, 32) + " é " + strconv.FormatFloat(operationResult, 'G', -1, 32)))
}

func HandleDiv(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	var requestObject Request

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
	}

	err = json.Unmarshal(requestBody, &requestObject)

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
	}

	operationResult := requestObject.First / requestObject.Second

	var responseResult Result

	responseResult.Result = operationResult

	response, err := json.Marshal(responseResult)

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
	}

	fmt.Println(response)
	w.Write([]byte("O resultado da divisão entre " + strconv.FormatFloat(requestObject.First, 'G', -1, 64) + " e " + strconv.FormatFloat(requestObject.Second, 'G', -1, 32) + " é " + strconv.FormatFloat(operationResult, 'G', -1, 32)))
}

func HandleMult(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	var requestObject Request

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
	}

	err = json.Unmarshal(requestBody, &requestObject)

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
	}

	operationResult := requestObject.First * requestObject.Second

	var responseResult Result

	responseResult.Result = operationResult

	response, err := json.Marshal(responseResult)

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
	}

	fmt.Println(response)
	w.Write([]byte("O resultado da multiplicação entre " + strconv.FormatFloat(requestObject.First, 'G', -1, 64) + " e " + strconv.FormatFloat(requestObject.Second, 'G', -1, 32) + " é " + strconv.FormatFloat(operationResult, 'G', -1, 32)))
}

type Request struct {
	First  float64 `json: "first"`
	Second float64 `json: "second"`
}

type Result struct {
	Result float64 `json: "result"`
}
