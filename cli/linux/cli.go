package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//Valid attributes are:
//    Reset, Bold, Faint, Italic, Underline, BlinkSlow, BlinkRapid, ReverseVideo, Concealed, CrossedOut
//
//Valid color are:
//    Black, Red, Green, Yellow, Blue, Magenta, Cyan, White
func EditPrompt(pLabel string, pDefault string) string {
	strflag := true
	matched, err := regexp.MatchString(`^[0-9]+\.[0-9]+$`, pDefault)
	if err != nil {
		fmt.Println("reg match err1")
	}
	if matched {
		// float
		//		fmt.Println("Edit Float", matched)
		strflag = false
	}
	matched, err = regexp.MatchString(`^[0-9]+$`, pDefault)

	if err != nil {
		fmt.Println("reg match err2")
	}
	if matched {
		// int
		//		fmt.Println("Edit Int", matched)
		strflag = false
	}
	pLower := strings.ToLower(pDefault)
	if "true" == pLower || "false" == pLower {
		_, selV := SelectPrompt(pLower, []string{"true", "false"})
		return selV
	}

	prompt := promptui.Prompt{
		Label:   pLabel,
		Default: pLower,
	}

	ChangedValue, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	matched, err = regexp.MatchString(`^[a-z]+$`, pDefault)
	if strflag == false {
		return ChangedValue
	} else if matched {
		return "\"" + ChangedValue + "\""
	} else {
		return "\"" + ChangedValue + "\""
	}
}

func SelectPrompt(pLabel string, pList []string) (int, string) {
	Open := ""
	Close := ""
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   ">> {{ . | magenta | bold }}",
		Inactive: "   {{ . | cyan }}",
		Selected: "> {{ . | magenta | cyan }}",
		Details:  Close,
	}
	searcher := func(input string, index int) bool {
		aValue := pList[index]
		bValue := strings.Replace(strings.ToLower(aValue), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(bValue, input)
	}
	prompt := promptui.Select{
		Label:     pLabel + " " + Open,
		Items:     pList,
		Templates: templates,
		Size:      8,
		Searcher:  searcher,
	}
	SelIndex, SelValue, err := prompt.Run()
	if err != nil {
		panic(err)
	}
	return SelIndex, SelValue
}

func JsonRPC(CoinName string, URL string, FunctionName string, Parameters string, printlog bool) {

	id := 1
	ID := strconv.Itoa(id)

	file, err := ioutil.ReadFile(CoinName + ".template")
	if err != nil {
		fmt.Println("File Error : ", err)
		os.Exit(1)
	}
	CoinTemplate := string(file)

	tmpl, err := template.New("CoinTemplate").Parse(CoinTemplate)
	if err != nil {
		panic(err)
	}

	DataType := struct {
		FunctionName interface{}
		Parameters   interface{}
		ID           string
	}{
		FunctionName: template.HTML(FunctionName),
		Parameters:   template.HTML(Parameters),
		ID:           ID,
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, DataType)
	if err != nil {
		panic(err)
	}

	/*	testjson := "{\"jsonrpc\":\"2.0\","
		testjson = testjson + "\"method\":\"" + FunctionName + "\","
		if Parameters != "" {
			testjson = testjson + "\"params\":" + Parameters + ", "
		}
		testjson = testjson + "\"id\":" + strconv.Itoa(id) + "}"
	*/
	ParamBytes := bytes.NewBuffer([]byte(tpl.String()))

	req, err := http.NewRequest("POST", URL, ParamBytes)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if printlog {
		fmt.Println("tpl.String", tpl.String())
		fmt.Println("Send:\n", JsonPrint(tpl.String()))
		fmt.Printf("Receive:\n")
		fmt.Println(JsonPrint(string(data)))
	}
}

func JsonPrint(RawJson string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(RawJson), "", "  ")
	if err != nil {
		panic(err)
	}
	return out.String()
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func GJsonRunFunction(Title string, pItem interface{}) interface{} {
	RunCmd := gjson.Result{gjson.String, "실행", "실행", 0, 0}
	BackCmd := gjson.Result{gjson.String, "뒤로", "뒤로", 0, 0}

	//	fmt.Println("inGJsonRunFunction", pItem)
	jItem := gjson.Parse(fmt.Sprintf("%v", pItem))

	if jItem.IsArray() {
		//		fmt.Println("inArray")
		jItemArrayLen := len(jItem.Array())
		if jItemArrayLen == 0 {
			return jItem.String()
		}
		jItemArray := jItem.Array()

		//		fmt.Println("jItemArray:", jItemArray)

		if Title == "parameters" {
			jItemArray = append(jItemArray, RunCmd)
		} else {
			jItemArray = append(jItemArray, BackCmd)
		}

		SelIndex := 0
		for SelIndex < jItemArrayLen {
			var jItemStrArray []string
			for key, val := range jItemArray {
				jItemStrArray = append(jItemStrArray, fmt.Sprintf("%v:%v", key, val))
			}
			SelIndex, _ = SelectPrompt(Title, jItemStrArray)
			if SelIndex >= jItemArrayLen {
				break
			}
			RetContainer := GJsonRunFunction(strconv.Itoa(SelIndex), jItemArray[SelIndex])
			//			fmt.Println("RetContainer:", RetContainer, " : ", reflect.TypeOf(RetContainer))

			retval := ""
			//			matched, _ := regexp.MatchString(`^[0-9]+$`, RetContainer.(string))
			matched, _ := regexp.MatchString(`^\{`, RetContainer.(string))
			pLower := strings.ToLower(RetContainer.(string))
			//			fmt.Println("pLower", pLower, matched)
			if "true" == pLower || "false" == pLower || matched {
				//				fmt.Println("set string")
				retval, _ = sjson.SetRaw(jItem.String(), strconv.Itoa(SelIndex), RetContainer.(string))
			} else {
				//				fmt.Println("set raw")
				retval, _ = sjson.Set(jItem.String(), strconv.Itoa(SelIndex), RetContainer.(string))
			}

			//			fmt.Println("Array Edited Json:", retval)
			jItem = gjson.Parse(retval)
			jItemArray = jItem.Array()

			if Title == "parameters" {
				jItemArray = append(jItemArray, RunCmd)
			} else {
				jItemArray = append(jItemArray, BackCmd)
			}
		}
		return jItem.String()

	} else if jItem.IsObject() {
		//		fmt.Println("inObject")

		jItemMapLen := len(jItem.Map())
		if jItemMapLen == 0 {
			return jItem.String()
		}
		jItemMap := jItem.Map()

		//		fmt.Println("jItemMap:", jItemMap)

		SelIndex := 0
		for SelIndex < jItemMapLen {
			var jItemStrArray []string
			for key, val := range jItemMap {
				jItemStrArray = append(jItemStrArray, fmt.Sprintf("%v:%v", key, val))
			}
			if Title == "parameters" {
				jItemStrArray = append(jItemStrArray, "실행")
			} else {
				jItemStrArray = append(jItemStrArray, "뒤로")
			}
			SelIndex, _ = SelectPrompt(Title, jItemStrArray)
			if SelIndex >= jItemMapLen {
				break
			}
			Select := strings.Split(jItemStrArray[SelIndex], ":")
			aaa := Select[0]
			RetContainer := GJsonRunFunction(aaa, jItemMap[aaa])
			//fmt.Println("RetContainer:", RetContainer)
			ttt, _ := sjson.Set(jItem.String(), aaa, RetContainer)
			//fmt.Println("Map Edited Json:", ttt)
			jItem = gjson.Parse(ttt)
			jItemMap = jItem.Map()
		}
		return jItem.String()

	} else {
		//fmt.Println("inText")
		EditedItem := EditPrompt(Title, jItem.String())
		//fmt.Println("EditedItem", EditedItem, "pItem type:", reflect.TypeOf(pItem), "val:", pItem)
		return EditedItem
	}
	return jItem.String()
}

func main() {
	//RunCmd := "실행"

	ExitCmd := "EXIT"

	if len(os.Args) < 2 {
		fmt.Println("error : cli INPUTJSON")
		return
	}
	JsonFile := os.Args[1]
	file, err := ioutil.ReadFile(JsonFile)
	if err != nil {
		fmt.Println("File Error : ", err)
		os.Exit(1)
	}
	JsonRoot := string(file)

	if !gjson.Valid(JsonRoot) {
		errors.New("invalid json file")
	}

	CoinName := gjson.Get(JsonRoot, "CoinName")
	URL := gjson.Get(JsonRoot, "URL")
	RPCTYPE := gjson.Get(JsonRoot, "RPCTYPE")
	SSL := gjson.Get(JsonRoot, "SSL")

	fmt.Println("")
	fmt.Println("CoinName: ", CoinName)
	fmt.Println("URL: ", URL)
	fmt.Println("RPCTYPE: ", RPCTYPE)
	fmt.Println("SSL: ", SSL)

	var FunctionNameList []string
	gjson.Get(JsonRoot, "functions.#.function").ForEach(func(_, val gjson.Result) bool { FunctionNameList = append(FunctionNameList, val.String()); return true })
	FunctionList := gjson.Get(JsonRoot, "functions").Array()

	FunctionNameList = append(FunctionNameList, ExitCmd)

	FunctionIndex := 0
	SelectedFunctionName := ""
	for FunctionIndex < len(FunctionList) {
		FunctionIndex, SelectedFunctionName = SelectPrompt("Select Function", FunctionNameList)
		if FunctionIndex >= len(FunctionList) {
			break
		}
		//		fmt.Println("SelectedFunctionName", SelectedFunctionName)

		//		fmt.Println("FunctionAll : ", FunctionList[FunctionIndex].String())

		SendParams := ""
		Parameters := gjson.Get(FunctionList[FunctionIndex].String(), "parameters")
		if Parameters.Exists() {
			RetParam := GJsonRunFunction("parameters", Parameters.String())

			SendParams = RetParam.(string)

		}
		JsonRPC(CoinName.String(), URL.String(), SelectedFunctionName, SendParams, true)
	}
}
