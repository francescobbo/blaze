package calculator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/francescobbo/blaze/store"
)

var weightUnits = map[string]float64{
	"kg": 1,
	"g":  1000,
	"mg": 1000000,
	"oz": 35.274,
	"lb": 2.20462,
	"st": 0.157473,
}

var lengthUnits = map[string]float64{
	"m":  1,
	"ft": 3.28084,
	"cm": 100,
	"mm": 1000,
	"km": 0.001,
	"mi": 0.000621371,
	"yd": 1.09361,
	"in": 39.3701,
}

var areaUnits = map[string]float64{
	"sqm":  1,
	"sqft": 10.7639,
	"sqcm": 10000,
	"sqmm": 1000000,
	"sqkm": 0.000001,
	"sqmi": 3.861e-7,
	"sqyd": 1.19599,
	"sqin": 1550,
	"ha":   0.0001,
	"acre": 0.000247105,
}

var volumeUnits = map[string]float64{
	"m3":   1,
	"l":    1000,
	"ml":   1000000,
	"cm3":  1000000,
	"ft3":  35.3147,
	"in3":  61023.7,
	"gal":  264.172,
	"qt":   1056.69,
	"pt":   2113.38,
	"cup":  4226.75,
	"tbsp": 67628,
	"tsp":  202884,
	"floz": 33814,
}

var timeUnits = map[string]float64{
	"h":   1,
	"min": 60,
	"s":   3600,
	"ms":  3600000,
	"d":   0.0416667,
	"w":   0.00595238,
	"mo":  0.00136986,
	"y":   0.000114155,
}

var currencyCodes = []string{
	"aed", "afn", "all", "amd", "ang", "aoa", "ars", "aud", "awg", "azn", "bam",
	"bbd", "bdt", "bgn", "bhd", "bif", "bmd", "bnd", "bob", "bov", "brl", "bsd",
	"btn", "bwp", "byn", "bzd", "cad", "cdf", "che", "chf", "chw", "clf", "clp",
	"cny", "cop", "cou", "crc", "cup", "cve", "czk", "djf", "dkk", "dop", "dzd",
	"egp", "ern", "etb", "eur", "fjd", "fkp", "gbp", "gel", "ghs", "gip", "gmd",
	"gnf", "gtq", "gyd", "hkd", "hnl", "htg", "huf", "idr", "ils", "inr", "iqd",
	"irr", "isk", "jmd", "jod", "jpy", "kes", "kgs", "khr", "kmf", "kpw", "krw",
	"kwd", "kyd", "kzt", "lak", "lbp", "lkr", "lrd", "lsl", "lyd", "mad", "mdl",
	"mga", "mkd", "mmk", "mnt", "mop", "mru", "mur", "mvr", "mwk", "mxn", "mxv",
	"myr", "mzn", "nad", "ngn", "nio", "nok", "npr", "nzd", "omr", "pab", "pen",
	"pgk", "php", "pkr", "pln", "pyg", "qar", "ron", "rsd", "rub", "rwf", "sar",
	"sbd", "scr", "sdg", "sek", "sgd", "shp", "sle", "sll", "sos", "srd", "ssp",
	"stn", "svc", "syp", "szl", "thb", "tjs", "tmt", "tnd", "top", "try", "ttd",
	"twd", "tzs", "uah", "ugx", "usd", "usn", "uyi", "uyu", "uyw", "uzs", "ved",
	"ves", "vnd", "vuv", "wst", "xaf", "xag", "xau", "xba", "xbb", "xbc", "xbd",
	"xcd", "xdr", "xof", "xpd", "xpf", "xpt", "xsu", "xts", "xua", "xxx", "yer",
	"zar", "zmw", "zwl",
}

var cryptocurrencyCodes = []string{
	"btc", "eth", "ltc", "doge",
}

// Function to convert a value to another unit
func convert(value Value, toUnit string) (Value, error) {
	if value.Unit == "" {
		value.Unit = toUnit
		return value, nil
	}

	value.Unit = strings.ToLower(value.Unit)
	toUnit = strings.ToLower(toUnit)

	// determine the source unit type
	sourceUnit := value.Unit

	if _, ok := weightUnits[sourceUnit]; ok {
		return convertWith(value, toUnit, weightUnits)
	}

	if _, ok := lengthUnits[sourceUnit]; ok {
		return convertWith(value, toUnit, lengthUnits)
	}

	if _, ok := areaUnits[sourceUnit]; ok {
		return convertWith(value, toUnit, areaUnits)
	}

	if _, ok := volumeUnits[sourceUnit]; ok {
		return convertWith(value, toUnit, volumeUnits)
	}

	if _, ok := timeUnits[sourceUnit]; ok {
		return convertWith(value, toUnit, timeUnits)
	}

	if sourceUnit == "f" || sourceUnit == "c" || sourceUnit == "k" {
		return convertTemperature(value, toUnit)
	}

	if slices.Contains(currencyCodes, sourceUnit) || slices.Contains(cryptocurrencyCodes, sourceUnit) {
		if slices.Contains(currencyCodes, toUnit) || slices.Contains(cryptocurrencyCodes, toUnit) {
			return convertCurrency(value, toUnit)
		}
	}

	return Value{}, fmt.Errorf("cannot convert from %s to %s", value.Unit, toUnit)
}

func convertWith(value Value, toUnit string, table map[string]float64) (Value, error) {
	if factor, ok := table[toUnit]; ok {
		value.Number *= factor / table[value.Unit]
		value.Unit = toUnit
		return value, nil
	}

	return Value{}, fmt.Errorf("unknown unit %s", toUnit)
}

func convertTemperature(value Value, toUnit string) (Value, error) {
	if value.Unit == toUnit {
		return value, nil
	}

	if toUnit != "f" && toUnit != "c" && toUnit != "k" {
		return Value{}, fmt.Errorf("unknown unit %s", toUnit)
	}

	var kValue float64
	if value.Unit == "f" {
		kValue = (value.Number + 459.67) * 5 / 9
	} else if value.Unit == "c" {
		kValue = value.Number + 273.15
	} else {
		kValue = value.Number
	}

	var toValue float64
	if toUnit == "f" {
		toValue = kValue*9/5 - 459.67
	} else if toUnit == "c" {
		toValue = kValue - 273.15
	} else {
		toValue = kValue
	}

	return Value{Number: toValue, Unit: toUnit}, nil
}

func convertCurrency(value Value, toCurrency string) (Value, error) {
	// Convert the value to USD
	usdValue, err := convertCurrencyToUSD(value)
	if err != nil {
		return Value{}, err
	}

	// Convert the USD value to the target currency
	targetValue, err := convertUSDToCurrency(usdValue, toCurrency)
	if err != nil {
		return Value{}, err
	}

	return targetValue, nil
}

func convertCurrencyToUSD(value Value) (float64, error) {
	if value.Unit == "usd" {
		return value.Number, nil
	}

	isCryptocurrency := slices.Contains(cryptocurrencyCodes, strings.ToLower(value.Unit))
	updateTolerance := 3600

	if isCryptocurrency {
		updateTolerance = 60
	}

	usd, lastUpdate, err := store.GetCurrencyUsdValue(value.Unit)
	if err != nil || time.Now().Unix()-int64(lastUpdate) > int64(updateTolerance) {
		// Currency is not up-to-date, fetch the latest value
		usd, err = refreshCurrencyRate(value.Unit)
		if err != nil {
			return 0, err
		}
		store.SaveCurrencyUsdValue(value.Unit, usd, int(time.Now().Unix()))
	}

	return value.Number / usd, nil
}

func convertUSDToCurrency(usdValue float64, toCurrency string) (Value, error) {
	if toCurrency == "usd" {
		return Value{Number: usdValue, Unit: toCurrency}, nil
	}

	isCryptocurrency := slices.Contains(cryptocurrencyCodes, strings.ToLower(toCurrency))
	updateTolerance := 3600

	if isCryptocurrency {
		updateTolerance = 60
	}

	usd, lastUpdate, err := store.GetCurrencyUsdValue(toCurrency)
	if err != nil || time.Now().Unix()-int64(lastUpdate) > int64(updateTolerance) {
		// Currency is not up-to-date, fetch the latest value
		usd, err = refreshCurrencyRate(toCurrency)
		if err != nil {
			return Value{}, err
		}
		store.SaveCurrencyUsdValue(toCurrency, usd, int(time.Now().Unix()))
	}

	return Value{Number: usdValue * usd, Unit: toCurrency}, nil
}

func refreshCurrencyRate(toCurrency string) (float64, error) { // source is always 1 USD
	isCryptocurrency := slices.Contains(cryptocurrencyCodes, strings.ToLower(toCurrency))

	if isCryptocurrency {
		// Fetch the latest cryptocurrency rate
		return coinbaseUsdRate(toCurrency)
	} else {
		// Fetch the latest currency rate
		return erapiUsdRate(toCurrency)
	}
}

func erapiUsdRate(sourceCurrency string) (float64, error) {
	url := fmt.Sprintf("https://open.er-api.com/v6/latest/%s", sourceCurrency)
	body, err := httpsRequest(url)
	if err != nil {
		return 0, err
	}

	type erapiResponse struct {
		Rates map[string]float64 `json:"rates"`
	}

	var response erapiResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return 0, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if rate, ok := response.Rates["USD"]; ok {
		return 1.0 / rate, nil
	}

	return 0, nil
}

func coinbaseUsdRate(sourceCurrency string) (float64, error) {
	url := fmt.Sprintf("https://api.coinbase.com/v2/exchange-rates?currency=%s", sourceCurrency)
	body, err := httpsRequest(url)
	if err != nil {
		return 0, err
	}

	type coinbaseResponse struct {
		Data struct {
			Rates map[string]string `json:"rates"`
		} `json:"data"`
	}

	var response coinbaseResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return 0, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if rate, ok := response.Data.Rates["USD"]; ok {
		fl, err := strconv.ParseFloat(rate, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse rate: %v", err)
		}

		return 1.0 / fl, nil
	}

	return 0, nil
}

func httpsRequest(url string) ([]byte, error) {
	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create the HTTPS request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set appropriate headers if needed
	req.Header.Set("Accept", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Check for HTTP status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}
