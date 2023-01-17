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
			Dst        string          `json:"dst"`
			PrefixWrap int             `json:"prefixWrap"`
			Result     [][]interface{} `json:"result"`
			Src        string          `json:"src"`
		} `json:"data"`
		From     string `json:"from"`
		Status   int    `json:"status"`
		To       string `json:"to"`
		Type     int    `json:"type"`
		Phonetic []struct {
			SrcStr string `json:"src_str"`
			TrgStr string `json:"trg_str"`
		} `json:"phonetic"`
	} `json:"trans_result"`
	DictResult struct {
		Edict struct {
			Item []struct {
				TrGroup []struct {
					Tr          []string `json:"tr"`
					Example     []string `json:"example"`
					SimilarWord []string `json:"similar_word"`
				} `json:"tr_group"`
				Pos string `json:"pos"`
			} `json:"item"`
			Word string `json:"word"`
		} `json:"edict"`
		Collins struct {
			Entry []struct {
				EntryID string `json:"entry_id"`
				Type    string `json:"type"`
				Value   []struct {
					MeanType []struct {
						InfoType string `json:"info_type"`
						InfoID   string `json:"info_id"`
						Example  []struct {
							ExampleID string `json:"example_id"`
							TtsSize   string `json:"tts_size"`
							Tran      string `json:"tran"`
							Ex        string `json:"ex"`
							TtsMp3    string `json:"tts_mp3"`
						} `json:"example,omitempty"`
						Posc []struct {
							Tran    string `json:"tran"`
							PoscID  string `json:"posc_id"`
							Example []struct {
								ExampleID string `json:"example_id"`
								Tran      string `json:"tran"`
								Ex        string `json:"ex"`
								TtsMp3    string `json:"tts_mp3"`
							} `json:"example"`
							Def string `json:"def"`
						} `json:"posc,omitempty"`
					} `json:"mean_type"`
					Gramarinfo []struct {
						Tran  string `json:"tran"`
						Type  string `json:"type"`
						Label string `json:"label"`
					} `json:"gramarinfo"`
					Tran   string `json:"tran"`
					Def    string `json:"def"`
					MeanID string `json:"mean_id"`
					Posp   []struct {
						Label string `json:"label"`
					} `json:"posp"`
				} `json:"value"`
			} `json:"entry"`
			WordName      string `json:"word_name"`
			Frequence     string `json:"frequence"`
			WordEmphasize string `json:"word_emphasize"`
			WordID        string `json:"word_id"`
		} `json:"collins"`
		From        string `json:"from"`
		SimpleMeans struct {
			WordName  string   `json:"word_name"`
			From      string   `json:"from"`
			WordMeans []string `json:"word_means"`
			Exchange  struct {
				WordPl []string `json:"word_pl"`
			} `json:"exchange"`
			Tags struct {
				Core  []string `json:"core"`
				Other []string `json:"other"`
			} `json:"tags"`
			Symbols []struct {
				PhEn  string `json:"ph_en"`
				PhAm  string `json:"ph_am"`
				Parts []struct {
					Part  string   `json:"part"`
					Means []string `json:"means"`
				} `json:"parts"`
				PhOther string `json:"ph_other"`
			} `json:"symbols"`
		} `json:"simple_means"`
		Lang   string `json:"lang"`
		Oxford struct {
			Entry []struct {
				Tag  string `json:"tag"`
				Name string `json:"name"`
				Data []struct {
					Tag  string `json:"tag"`
					Data []struct {
						Tag  string `json:"tag"`
						Data []struct {
							Tag  string `json:"tag"`
							Data []struct {
								Tag  string `json:"tag"`
								Data []struct {
									Tag    string `json:"tag"`
									EnText string `json:"enText,omitempty"`
									ChText string `json:"chText,omitempty"`
									G      string `json:"g,omitempty"`
									Data   []struct {
										Text      string `json:"text"`
										HoverText string `json:"hoverText"`
									} `json:"data,omitempty"`
								} `json:"data"`
							} `json:"data"`
						} `json:"data,omitempty"`
						P     string `json:"p,omitempty"`
						PText string `json:"p_text,omitempty"`
						N     string `json:"n,omitempty"`
						Xt    string `json:"xt,omitempty"`
					} `json:"data"`
				} `json:"data"`
			} `json:"entry"`
			Unbox []struct {
				Tag  string `json:"tag"`
				Type string `json:"type"`
				Name string `json:"name"`
				Data []struct {
					Tag     string `json:"tag"`
					Text    string `json:"text,omitempty"`
					Words   string `json:"words,omitempty"`
					Outdent string `json:"outdent,omitempty"`
					Data    []struct {
						Tag    string `json:"tag"`
						EnText string `json:"enText"`
						ChText string `json:"chText"`
					} `json:"data,omitempty"`
				} `json:"data"`
			} `json:"unbox"`
		} `json:"oxford"`
		BaiduPhrase []struct {
			Tit   []string `json:"tit"`
			Trans []string `json:"trans"`
		} `json:"baidu_phrase"`
		QueryExplainVideo struct {
			ID           int    `json:"id"`
			UserID       string `json:"user_id"`
			UserName     string `json:"user_name"`
			UserPic      string `json:"user_pic"`
			Query        string `json:"query"`
			Direction    string `json:"direction"`
			Type         string `json:"type"`
			Tag          string `json:"tag"`
			Detail       string `json:"detail"`
			Status       string `json:"status"`
			SearchType   string `json:"search_type"`
			FeedURL      string `json:"feed_url"`
			Likes        string `json:"likes"`
			Plays        string `json:"plays"`
			CreatedAt    string `json:"created_at"`
			UpdatedAt    string `json:"updated_at"`
			DuplicateID  string `json:"duplicate_id"`
			RejectReason string `json:"reject_reason"`
			CoverURL     string `json:"coverUrl"`
			VideoURL     string `json:"videoUrl"`
			ThumbURL     string `json:"thumbUrl"`
			VideoTime    string `json:"videoTime"`
			VideoType    string `json:"videoType"`
		} `json:"queryExplainVideo"`
	} `json:"dict_result"`
	LijuResult struct {
		Double string   `json:"double"`
		Tag    []string `json:"tag"`
		Single string   `json:"single"`
	} `json:"liju_result"`
	Logid int `json:"logid"`
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
