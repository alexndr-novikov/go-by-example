package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
	Extra  int
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
	Extra  int      `json:",omitempty"`
}

func RunSample() {
	fmt.Println("json package output:")
	jBool, _ := json.Marshal(true)
	fmt.Println(string(jBool))

	jNil, _ := json.Marshal(nil)
	fmt.Println(string(jNil))

	arr := []string{"Hello", "World"}
	jArr, _ := json.Marshal(arr)
	fmt.Println(string(jArr))

	mapEx := map[string]int{"apple": 1, "orange": 8}
	jMap, _ := json.Marshal(mapEx)
	fmt.Println(string(jMap))

	jStruct1 := &response1{
		Page:   1,
		Fruits: []string{"Apple", "Banana"},
	}

	jStruct2 := &response2{
		Page:   1,
		Fruits: []string{"Apple", "Banana"},
	}

	jStructResp1, _ := json.Marshal(jStruct1)
	fmt.Println(string(jStructResp1))

	jStructResp2, _ := json.Marshal(jStruct2)
	fmt.Println(string(jStructResp2))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var bytStr map[string]interface{}
	json.Unmarshal(byt, &bytStr)
	fmt.Println(bytStr)
	fmt.Println(bytStr["num"].(float64))
	strngs := bytStr["strs"].([]interface{})
	fmt.Println(strngs[0])

	js := `{"Page":1,"Fruits":["Apple","Banana"],"Extra":0}`
	res1 := response1{}
	res2 := response1{}
	json.Unmarshal([]byte(js), &res1)
	json.Unmarshal([]byte(js), &res2)
	fmt.Println(res1)
	fmt.Println(res2)

	enc := json.NewEncoder(os.Stderr)
	dat := map[string]string{"a": "b", "c": "d"}
	enc.Encode(dat)
	fmt.Println("")
}
