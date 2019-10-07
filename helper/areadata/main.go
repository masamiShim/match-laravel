package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"os"
)

type PrefData struct {
	prefCd string
	prefNm string
}

type MuniciData struct {
	muniCd string
	muniNm string
}

type Prefs struct {
	prefData   PrefData
	municiData []MuniciData
}

const (
	prefCodeCol = 1
	prefNameCol = 7
	muniCodeCol = 2
	muniNameCol = 9

	PrefTable        = "prefectures"
	MuniciTable      = "municipalities"
	RelPrefMuniTable = "rel_prefectures_municipalities"
	DatabaseName     = "matching"
)

/**
1. 市区町村ごとに分類する
2. 分類済みから県登録データ、市区町村登録データ、県と市区町村の紐付け登録データを準備する
*/
func main() {

	var prefs []Prefs = extract()

	// SQLのアウトプット
	createSqlPrefecutures(prefs)

	createSqlMunicipalities(prefs)

	createSqlForRel(prefs)

}

func createSqlPrefecutures(prefs []Prefs) {
	file, err := os.Create("./data/insert_pref.sql")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < len(prefs); i++ {
		query := "insert into " + DatabaseName + "." + PrefTable + " (code, name) values ('" + prefs[i].prefData.prefCd + "', '" + prefs[i].prefData.prefNm + "');\n"
		if _, err := writer.WriteString(query); err != nil {
			panic(err)
		}
	}
	writer.Flush()
}

func createSqlMunicipalities(prefs []Prefs) {
	file, err := os.Create("./data/insert_municipalities.sql")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < len(prefs); i++ {
		for j := 0; j < len(prefs[i].municiData); j++ {
			query := "insert into " + DatabaseName + "." + MuniciTable + " (code, name) values ('" + prefs[i].municiData[j].muniCd + "', '" + prefs[i].municiData[j].muniNm + "');\n"
			if _, err := writer.WriteString(query); err != nil {
				panic(err)
			}
		}
	}
	writer.Flush()
}

func createSqlForRel(prefs []Prefs) {
	file, err := os.Create("./data/insert_rel_prefecture_municipality.sql")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < len(prefs); i++ {
		for j := 0; j < len(prefs[i].municiData); j++ {
			query := "insert into " + DatabaseName + "." + RelPrefMuniTable + " (prefecture_code, municipality_code) values ('" + prefs[i].prefData.prefCd + "', '" + prefs[i].municiData[j].muniCd + "');\n"
			if _, err := writer.WriteString(query); err != nil {
				panic(err)
			}
		}
	}
	writer.Flush()
}

func extract() []Prefs {
	file, err := os.Open("./zenkoku.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	var line []string
	var items []Prefs
	for {

		line, err = reader.Read()

		if err != nil && err == io.EOF {
			break
		} else if err != nil && err != io.EOF {
			panic(err)
		}

		var prefCd = line[prefCodeCol]
		var prefName, _, _ = transform.String(japanese.ShiftJIS.NewDecoder(), line[prefNameCol])
		var municiCd = line[muniCodeCol]
		var municiName, _, _ = transform.String(japanese.ShiftJIS.NewDecoder(), line[muniNameCol])
		items = add(items, PrefData{prefCd: prefCd, prefNm: prefName}, MuniciData{muniCd: municiCd, muniNm: municiName})

	}
	// 先頭行はヘッダなので除去
	items = items[1:]
	fmt.Println(items)
	return items
}

func add(items []Prefs, pref PrefData, munici MuniciData) []Prefs {

	var prefExists int = -1
	var municiExists int = -1
	for i := 0; i < len(items); i++ {

		if municiExists > -1 {
			break
		}

		if items[i].prefData.prefCd == pref.prefCd {
			prefExists = i
			for j := 0; j < len(items[i].municiData); j++ {
				if items[i].municiData[j].muniCd == munici.muniCd {
					municiExists = j
					break
				}
			}
			break
		}

	}

	// 都道府県存在しない
	if prefExists == -1 {
		// 全追加
		var munis []MuniciData
		munis = append(munis, munici)
		items = append(items, Prefs{prefData: pref, municiData: munis})
		return items
	}

	// 都道府県が存在する場合
	if prefExists > -1 {
		// 市区町村ない場合追加
		if municiExists == -1 {
			items[prefExists].municiData = append(items[prefExists].municiData, munici)
			return items
		}
	}

	// どちらも存在する場合はそのまま返す
	return items
}
