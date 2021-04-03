package osm

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type OverpassResult struct {
	tipe string
	id   int64
	lat  float64
	lon  float64
}

func Search(searchType string, searchQuery string, searchArea string, timeout uint) {
	link := "https://overpass.kumi.systems/api/interpreter?data="
	queryBuf := new(bytes.Buffer)
	_, err := fmt.Fprintf(queryBuf,
		`[out:json][timeout:%d];
				area[name="%s"]->.searchArea;
				(
					node[%s="%s"](area.searchArea);
					way[%s="%s"](area.searchArea);
					relation[%s="%s"](area.searchArea);
				);
				out;
				>;
				out skel qt;`,
		timeout, searchArea, searchType, searchQuery, searchType,
		searchQuery, searchType, searchQuery)
	query := url.QueryEscape(queryBuf.String())
	if err != nil {
		panic(err)
	}
	resp, err := http.Get(link + query)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(link + query)
	fmt.Println(queryBuf.String())
}

func SearchIntersectionInArea(areaName string, timeout uint, target interface{}) error {
	link := "https://overpass.kumi.systems/api/interpreter?data="
	queryBuf := new(bytes.Buffer)
	_, err := fmt.Fprintf(queryBuf, `[out:json][timeout:%d];
area[name~".*%s.*",i]->.searchArea;
way[highway~"^(motorway|trunk|primary|secondary|tertiary|(motorway|trunk|primary|secondary)_link)$"](area.searchArea)->.major;
way[highway~"^(motorway|trunk|primary|secondary|tertiary|(motorway|trunk|primary|secondary)_link)$"](area.searchArea)->.minor;
node(w.major)(w.minor);
out;`,
		timeout, areaName)
	query := url.QueryEscape(queryBuf.String())
	if err != nil {
		return err
	}
	resp, err := http.Get(link + query)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	//fmt.Println(link + query)
	//fmt.Println(queryBuf.String())

	return json.NewDecoder(resp.Body).Decode(target)
}
