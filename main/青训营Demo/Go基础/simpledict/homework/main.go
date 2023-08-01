package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type DictResponse struct {
	TransResult struct {
		Data []struct {
			Dst        string          `json_test:"dst"`
			PrefixWrap int             `json_test:"prefixWrap"`
			Result     [][]interface{} `json_test:"result"`
			Src        string          `json_test:"src"`
		} `json_test:"data"`
		From     string `json_test:"from"`
		Status   int    `json_test:"status"`
		To       string `json_test:"to"`
		Type     int    `json_test:"type"`
		Phonetic []struct {
			SrcStr string `json_test:"src_str"`
			TrgStr string `json_test:"trg_str"`
		} `json_test:"phonetic"`
	} `json_test:"trans_result"`
	DictResult struct {
		Edict struct {
			Item []struct {
				TrGroup []struct {
					Tr          []string `json_test:"tr"`
					Example     []string `json_test:"example"`
					SimilarWord []string `json_test:"similar_word"`
				} `json_test:"tr_group"`
				Pos string `json_test:"pos"`
			} `json_test:"item"`
			Word string `json_test:"word"`
		} `json_test:"edict"`
		Collins struct {
			Entry []struct {
				EntryID string `json_test:"entry_id"`
				Type    string `json_test:"type"`
				Value   []struct {
					MeanType []struct {
						InfoType string `json_test:"info_type"`
						InfoID   string `json_test:"info_id"`
						Example  []struct {
							ExampleID string `json_test:"example_id"`
							TtsSize   string `json_test:"tts_size"`
							Tran      string `json_test:"tran"`
							Ex        string `json_test:"ex"`
							TtsMp3    string `json_test:"tts_mp3"`
						} `json_test:"example,omitempty"`
						Posc []struct {
							Tran    string `json_test:"tran"`
							PoscID  string `json_test:"posc_id"`
							Example []struct {
								ExampleID string `json_test:"example_id"`
								Tran      string `json_test:"tran"`
								Ex        string `json_test:"ex"`
								TtsMp3    string `json_test:"tts_mp3"`
							} `json_test:"example"`
							Def string `json_test:"def"`
						} `json_test:"posc,omitempty"`
					} `json_test:"mean_type"`
					Gramarinfo []struct {
						Tran  string `json_test:"tran"`
						Type  string `json_test:"type"`
						Label string `json_test:"label"`
					} `json_test:"gramarinfo"`
					Tran   string `json_test:"tran"`
					Def    string `json_test:"def"`
					MeanID string `json_test:"mean_id"`
					Posp   []struct {
						Label string `json_test:"label"`
					} `json_test:"posp"`
				} `json_test:"value"`
			} `json_test:"entry"`
			WordName      string `json_test:"word_name"`
			Frequence     string `json_test:"frequence"`
			WordEmphasize string `json_test:"word_emphasize"`
			WordID        string `json_test:"word_id"`
		} `json_test:"collins"`
		From        string `json_test:"from"`
		SimpleMeans struct {
			WordName  string   `json_test:"word_name"`
			From      string   `json_test:"from"`
			WordMeans []string `json_test:"word_means"`
			Exchange  struct {
				WordPl []string `json_test:"word_pl"`
			} `json_test:"exchange"`
			Tags struct {
				Core  []string `json_test:"core"`
				Other []string `json_test:"other"`
			} `json_test:"tags"`
			Symbols []struct {
				PhEn  string `json_test:"ph_en"`
				PhAm  string `json_test:"ph_am"`
				Parts []struct {
					Part  string   `json_test:"part"`
					Means []string `json_test:"means"`
				} `json_test:"parts"`
				PhOther string `json_test:"ph_other"`
			} `json_test:"symbols"`
		} `json_test:"simple_means"`
		Lang   string `json_test:"lang"`
		Oxford struct {
			Entry []struct {
				Tag  string `json_test:"tag"`
				Name string `json_test:"name"`
				Data []struct {
					Tag  string `json_test:"tag"`
					Data []struct {
						Tag  string `json_test:"tag"`
						Data []struct {
							Tag  string `json_test:"tag"`
							Data []struct {
								Tag  string `json_test:"tag"`
								Data []struct {
									Tag    string `json_test:"tag"`
									EnText string `json_test:"enText,omitempty"`
									ChText string `json_test:"chText,omitempty"`
									G      string `json_test:"g,omitempty"`
									Data   []struct {
										Text      string `json_test:"text"`
										HoverText string `json_test:"hoverText"`
									} `json_test:"data,omitempty"`
								} `json_test:"data"`
							} `json_test:"data"`
						} `json_test:"data,omitempty"`
						P     string `json_test:"p,omitempty"`
						PText string `json_test:"p_text,omitempty"`
						N     string `json_test:"n,omitempty"`
						Xt    string `json_test:"xt,omitempty"`
					} `json_test:"data"`
				} `json_test:"data"`
			} `json_test:"entry"`
			Unbox []struct {
				Tag  string `json_test:"tag"`
				Type string `json_test:"type"`
				Name string `json_test:"name"`
				Data []struct {
					Tag     string `json_test:"tag"`
					Text    string `json_test:"text,omitempty"`
					Words   string `json_test:"words,omitempty"`
					Outdent string `json_test:"outdent,omitempty"`
					Data    []struct {
						Tag    string `json_test:"tag"`
						EnText string `json_test:"enText"`
						ChText string `json_test:"chText"`
					} `json_test:"data,omitempty"`
				} `json_test:"data"`
			} `json_test:"unbox"`
		} `json_test:"oxford"`
		BaiduPhrase []struct {
			Tit   []string `json_test:"tit"`
			Trans []string `json_test:"trans"`
		} `json_test:"baidu_phrase"`
		QueryExplainVideo struct {
			ID           int    `json_test:"id"`
			UserID       string `json_test:"user_id"`
			UserName     string `json_test:"user_name"`
			UserPic      string `json_test:"user_pic"`
			Query        string `json_test:"query"`
			Direction    string `json_test:"direction"`
			Type         string `json_test:"type"`
			Tag          string `json_test:"tag"`
			Detail       string `json_test:"detail"`
			Status       string `json_test:"status"`
			SearchType   string `json_test:"search_type"`
			FeedURL      string `json_test:"feed_url"`
			Likes        string `json_test:"likes"`
			Plays        string `json_test:"plays"`
			CreatedAt    string `json_test:"created_at"`
			UpdatedAt    string `json_test:"updated_at"`
			DuplicateID  string `json_test:"duplicate_id"`
			RejectReason string `json_test:"reject_reason"`
			CoverURL     string `json_test:"coverUrl"`
			VideoURL     string `json_test:"videoUrl"`
			ThumbURL     string `json_test:"thumbUrl"`
			VideoTime    string `json_test:"videoTime"`
			VideoType    string `json_test:"videoType"`
		} `json_test:"queryExplainVideo"`
	} `json_test:"dict_result"`
	LijuResult struct {
		Double string   `json_test:"double"`
		Tag    []string `json_test:"tag"`
		Single string   `json_test:"single"`
	} `json_test:"liju_result"`
	Logid int `json_test:"logid"`
}

func query(word string) {
	client := &http.Client{}
	payload := url.Values{}
	payload.Set("from", "en")
	payload.Set("to", "zh")
	payload.Set("query", word)
	payload.Set("transtype", "translang")
	payload.Set("simple_means_flag", "3")
	payload.Set("sign", "431039.159886")
	payload.Set("token", "44c8d9a28bca7742211e3d21316d8dbf")
	payload.Set("domain", "common")
	req, err := http.NewRequest(http.MethodPost, "https://fanyi.baidu.com/v2transapi?from=en&to=zh", strings.NewReader(payload.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Acs-Token", "1673856314151_1673863522951_5nfnvniCGuXKn17dtB6RTyrRRI9R/oBRE0HDY6EvG3Dj35U1s9t37gbU7Nec5lgZi7FMcCYhoiz/GLHPT27sHyTqHJfDm4im8j1ELWemPJV4/DBlVF1n268/AumjZR56NdVydZuEfB9R9MPkscIvjjz/nULoFwpOAcISbQdAtB3feoecf0aXR01XuqvSk7fTKsvIVwRoLF5XeKvsv+/UdcypZE92mBv4Sbrn/LeQ//jzNH3oD4W81C2ZpkidtrwqIC33cIQwyTfriGsMSIERSEp3F5EuxhnqqpMeePqNrOStKqcuqvanUh2vjIixzO9DozttxTc0eGq0Yn4pF3SZAw==")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "BIDUPSID=A765EF4FEE9C0384F73E01FDF23E513B; PSTM=1645949791; __yjs_duid=1_7cec9d33c2f2cec43b8ba808296b9bf01646010338025; REALTIME_TRANS_SWITCH=1; SOUND_SPD_SWITCH=1; HISTORY_SWITCH=1; FANYI_WORD_SWITCH=1; SOUND_PREFER_SWITCH=1; APPGUIDE_10_0_2=1; BDUSS_BFESS=NzWHdqbXpqcXY5ZkNIemxjVnF4NWNkdTYxUldqMTZIVkFpZy1ISG1PdlJLVEpqSUFBQUFBJCQAAAAAAAAAAAEAAADYmChtyczT2squxt8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAANGcCmPRnApjQ2; BAIDUID=740BEA58EBC9C853BC1E476A7085B276:FG=1; Hm_lvt_64ecd82404c51e03dc91cb9e8c025574=1671350932; delPer=0; PSINO=7; ZFY=GpU:AkE4gAy1F:AHItUx:BnNQ3rjZ8UHnHiAyyAkvfR1hc:C; BAIDUID_BFESS=740BEA58EBC9C853BC1E476A7085B276:FG=1; BDRCVFR[PGlCcVoRm90]=mk3SLVN4HKm; H_PS_PSSID=36546_37971_37645_37559_37515_38023_37906_36921_38035_37990_36807_37924_37904_26350_37285_22157_37881; Hm_lpvt_64ecd82404c51e03dc91cb9e8c025574=1673856687; Hm_lvt_246a5e7d3670cfba258184e42d902b31=1673858137; Hm_lpvt_246a5e7d3670cfba258184e42d902b31=1673858138; ab_sr=1.0.1_NmRlYTQyNDBlZjcxN2IwNjc0MWExMTRkNTBmMTQxZTY4MGMyOTQ3N2ZkOWUxMjE0NGE0YTUzZTY1NmM1MGFkYmVkMjlkY2U4NDQ3YzVmNGFiNzJmZTJhYjA2YTdjOTkwZTc5OTY3NDlhMDA0NDVmNDYwOWYwNmM3MDgzYTY1MzVkZmI0YWE1ZjhhYzBjNWViMmJmZDY5MTA5ZDBlNzk4MmNlMzY1NzNkNTBkMzRkMGIyMzY1MTBiNTRlMDA0OWVk")
	req.Header.Set("Origin", "https://fanyi.baidu.com")
	req.Header.Set("Referer", "https://fanyi.baidu.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.46")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Microsoft Edge";v="108"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	//获取响应
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	//json序列化
	var dictResponse DictResponse

	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dictResponse.TransResult.Data[0].Dst)
}

func main() {
	//if len(os.Args) != 2 {
	//	fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
	//example: simpleDict hello`)
	//	os.Exit(1)
	//}
	//word := os.Args[1]
	word := "test"
	query(word)
}
