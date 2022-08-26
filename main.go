package main

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
	"log"
	"os"
	"net/http"
	"strconv"
	"strings"
	"time"
	"github.com/joho/godotenv"
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

type Record struct {
	Link string `json:"string_letters_numbers"`
}

func init() {
	godotenv.Load()
}

/*
func makePart_1() string {
	POSSIBLE_SYMBOLS, _ := os.LookupEnv("POSSIBLE_SYMBOLS")
	LENGTH_1, _ := os.LookupEnv("LENGTH_1")
	length, _ := strconv.Atoi(LENGTH_1)
	possible_symbols := []rune(POSSIBLE_SYMBOLS)
	for {
		var temp strings.Builder
		for i := 0; i < length; i++ {
			temp.WriteRune(possible_symbols[rand.Intn(len(possible_symbols))])
		}
		if strings.Contains(temp.String(), "b") {
			return temp.String()
		} else {
			continue
		}
	}
}
*/

func makePart_1() string {
	POSSIBLE_SYMBOLS, _ := os.LookupEnv("SYMBOLS_FNZ")
	LENGTH_1, _ := os.LookupEnv("LENGTH_1")
	length, _ := strconv.Atoi(LENGTH_1)
	possible_symbols := []rune(POSSIBLE_SYMBOLS)
	var temp strings.Builder
	for i := 0; i < length; i++ {
		temp.WriteRune(possible_symbols[rand.Intn(len(possible_symbols))])
	}
	return temp.String()
}

func makePart_2() string {
	POSSIBLE_SYMBOLS, _ := os.LookupEnv("SYMBOLS_BN")
	LENGTH_2, _ := os.LookupEnv("LENGTH_2")
	length, _ := strconv.Atoi(LENGTH_2)
	possible_symbols := []rune(POSSIBLE_SYMBOLS)	
	var temp strings.Builder
	for i := 0; i < length; i++ {
		temp.WriteRune(possible_symbols[rand.Intn(len(possible_symbols))])
	}
	return temp.String()
}

func makePart_3() string {
	POSSIBLE_SYMBOLS, _ := os.LookupEnv("SYMBOLS_DN")
	possible_symbols := []rune(POSSIBLE_SYMBOLS)
	var temp strings.Builder
	temp.WriteRune(possible_symbols[rand.Intn(len(possible_symbols))])
	return temp.String()
}

func makePart_4() string {
	POSSIBLE_SYMBOLS, _ := os.LookupEnv("SYMBOLS_ACN")
	LENGTH_4, _ := os.LookupEnv("LENGTH_4")
	length, _ := strconv.Atoi(LENGTH_4)
	possible_symbols := []rune(POSSIBLE_SYMBOLS)	
	var temp strings.Builder
	for i := 0; i < length; i++ {
		temp.WriteRune(possible_symbols[rand.Intn(len(possible_symbols))])
	}
	return temp.String()
}

func makePart_5() string {
	POSSIBLE_SYMBOLS, _ := os.LookupEnv("SYMBOLS_N2")
	possible_symbols := []rune(POSSIBLE_SYMBOLS)
	var temp strings.Builder
	temp.WriteRune(possible_symbols[rand.Intn(len(possible_symbols))])
	return temp.String()
}

func makePart_6() string {
	POSSIBLE_SYMBOLS, _ := os.LookupEnv("POSSIBLE_SYMBOLS")
	LENGTH_6, _ := os.LookupEnv("LENGTH_6")
	length, _ := strconv.Atoi(LENGTH_6)
	possible_symbols := []rune(POSSIBLE_SYMBOLS)	
	var temp strings.Builder
	for i := 0; i < length; i++ {
		temp.WriteRune(possible_symbols[rand.Intn(len(possible_symbols))])
	}
	return temp.String()
}

func createUrl(part_1 string, part_2 string, part_3 string, part_4 string, part_5 string, part_6 string) string {
	DOMAIN, _ := os.LookupEnv("DOMAIN")
	url := DOMAIN + part_1  + part_2 + part_3 + "f" + part_4 + "f" + part_5 + "2215" + part_6 + ".jpg"
	return url
}

func createRecordDB(url string) {
	d_key, _ := os.LookupEnv("D_KEY")
	db, error := deta.New(deta.WithProjectKey(d_key))
	if error != nil {
		log.Print("DB connect error:", error)
		return
	}
	responses200_db, _ := base.New(db, "responses200")
	responses200_db.Insert(&Record{
		Link: url,
	})
}


func validationAmount(str string) bool {
	symbols := [14]string{"11", "33", "44", "55", "66", "77", "88", "99", "aa", "bb", "cc", "dd", "ee", "ff" }

	for i := 0; i < len(symbols); i++ {
		if strings.Contains(str, symbols[i]) {
			return false
		}
	}
	return true
}


func validationZero(str string) bool {
	symbols := [10]string{"10", "20", "30", "40", "50", "60", "70", "80", "90", "00"}

	for i := 0; i < len(symbols); i++ {
		if strings.Contains(str, symbols[i]) {
			return true
		}
	}
	return false
}

func makeRequestCodeStatus(url string) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}
	request, err := client.Get(url)
	if err != nil {
		log.Print("GET error:", err)
		return
	}
	if request.StatusCode == 200 {
		createRecordDB(url)
	}
	defer request.Body.Close()
}

func thread() {

	for {
		safeNum, err := crypto.Int(crypto.Reader, big.NewInt(999999999999999999))
		if err != nil {
			log.Print(err)
		}
		rand.Seed(safeNum.Int64())

		for {
			part_1 := makePart_1()
			part_2 := makePart_2()
			part_3 := makePart_3()
			part_4 := makePart_4()
			part_5 := makePart_5()
			part_6 := makePart_6()

			str := part_1  + part_2 + part_3 + "f" + part_4 + "f" + part_5 + "2218" + part_6

			if validationZero(str) {
				break
			}

			if validationAmount(str) {
				url := createUrl(part_1, part_2, part_3, part_4, part_5, part_6)
				makeRequestCodeStatus(url)
				//log.Print(url)
			}
		}
	}		
}

func main() {
	
	for i := 0; i < 520; i++ {
		go thread()
	}
	thread()
}
