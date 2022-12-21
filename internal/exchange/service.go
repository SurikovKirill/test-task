package exchange

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.cbr-xml-daily.ru/daily_json.js"

func GetExchange(cur string) (map[string]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var ci map[string]interface{}
	if err := json.Unmarshal(body, &ci); err != nil {
		return nil, err
	}
	if ci["Valute"].(map[string]interface{})[cur] == nil {
		return nil, http.ErrAbortHandler
	}
	n := ci["Valute"].(map[string]interface{})[cur].(map[string]interface{})["Name"].(string)
	nom := ci["Valute"].(map[string]interface{})[cur].(map[string]interface{})["Nominal"].(float64)
	val := ci["Valute"].(map[string]interface{})[cur].(map[string]interface{})["Value"].(float64)
	cc := ci["Valute"].(map[string]interface{})[cur].(map[string]interface{})["CharCode"].(string)
	rub := setValidName(int(val))
	d := map[string]string{}
	d[cc] = fmt.Sprintf("%d %s равен %f %s", int(nom), n, val, rub)
	return d, nil
}

func setValidName(rem int) string {
	if rem%10 == 1 {
		return "рублю"
	} else {
		return "рублям"
	}
}
