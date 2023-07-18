package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	addr string
	index string
	file  string
	regx = regexp.MustCompile(`(?s)siss_pay_id..(\d+)\s`)
	msg =`{"took":109,"timed_out":false,"_shards":{"total":4,"successful":4,"skipped":0,"failed":0},"hits":{"total":1837,"max_score":14.370081,"hits":[{"_index":"graylog_1205","_type":"message","_id":"8347c063-16ea-11ee-9f55-0242732e41f1","_score":14.370081,"_source":{"app":"pay","level":6,"line":0,"gl2_remote_ip":"172.18.0.215","gl2_remote_port":63496,"streams":["000000000000000000000001"],"gl2_message_id":"01H450QQV73M9X4Y2TC1ZMJK8Z","mch_no":"81020221014112001507","source":"思迅Pay","message":"移动支付","gl2_source_input":"5e662ea156182c00124c68c7","pay_way":"WX","file":"null","out_trade_no":"32064279000005230630S01207017","full_message":"2d84f629-a29f-4694-a48f-c7304c654a3e v2条码支付---\r\nDefaultConnectionLimit:2147483647\r\ngoods_detail.Count:1\r\n10:04:57.481获取IP开始\r\n10:04:57.481获取IP结束 101.229.111.45\r\n10:04:57.481解密开始\r\n10:04:57.481解密结束\r\n条码支付---\r\nDLL版本：23.06.26.16\r\n原IP101.229.111.45\r\ngoods_detail.Count:1\r\n10:04:57.481获取IP开始\r\n10:04:57.481获取IP结束 101.229.111.45\r\n手机条码:7536573115\r\n10:04:57.481判断思迅商户开始\r\n10:04:57.481判断思迅商户结束。\r\n siss_pay_id：81020221014112001507 siss_pay_store_id：0000 pay_way：WX total_amount：6.00 out_trade_no：32064279000005230630S01207017 产品名称：美食广场 支付客户端版本：3.20.04.20 客户端ip：101.229.111.45 branch_no：0000 siss_posid：05 siss_oper_id：6001 锁号：32064279\r\n是否上传商品：False\r\n pay_channel:YEAHKA,bank_channel:OTHER \r\n缓存中的值：乐刷直连\r\n匹配到 payTypeValue :乐刷直连\r\n交换前：[{\"goods_id\":\"05010\",\"goods_name\":\"绿豆百合粥\",\"quantity\":\"1.00\",\"price\":\"6.00\",\"item_subno\":null}]\r\nStart：10:04:57.497\r\n接口通道：YEAHKA\r\n10:04:57.497 请求报文：\r\namount=600&appid=wx47c8b96a32efb7de&auth_code=7536573115&body=条码支付-XXXX美食广场总部&merchant_id=8261414434&nonce_str=c6c610543bbb46bc9b6bce57ecd55a97&pos_no=00000005&service=upload_authcode&shop_no=0000&third_order_id=32064279000005230630S01207017&sign=E416E81CFB347515309C64067AB5BC48\r\n10:04:57.559 响应报文：\r\n<leshua>\n\t<resp_code><![CDATA[-5036]]></resp_code>\n\t<resp_msg><![CDATA[无效的付款码]]></resp_msg>\n</leshua>\n\r\n接口耗时：0.07秒\r\nStop：10:04:57.559\r\n接口耗时：0.07秒\r\nwpri.code:FAIL\r\n商户单号:32064279000005230630S01207017。交易单号:。三方交易单号:。","api_name":"CreatAndPay","server_ip":"47.106.231.132","pay_channel":"YEAHKA","gl2_source_node":"47d2947b-61d8-4aed-a2e2-b1bd852c6562","facility":"ECS","timestamp":"2023-06-30 02:04:57.000"}}]}}`
)

const (
	queryFmt =`{"query": {"match": {"full_message": "%s"}},"size": 1}`
)

func main() {
	matches := regx.FindStringSubmatch(msg)
	log.Println(matches,len(matches),matches[0])
	flag.StringVar(&addr,"host", "localhost:9200", "elasticsearch address")
	flag.StringVar(&index, "index", "graylog", "elasticsearch index name")
	flag.StringVar(&file, "file", "", "ip file")
	flag.Parse()
	if file == "" {
		log.Fatalln("file must be specified")
	}
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	url := fmt.Sprintf("http://%s/%s/_search", addr, index)
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), " ")
		log.Println(search(url, lines[len(lines) - 1]))
	}
}


func search(url string, key string) string {
	client := &http.Client {Timeout: time.Second * 2}
	payload := strings.NewReader(fmt.Sprintf(queryFmt, key))
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	matches := regx.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		log.Println("原始返回", string(body))
		return "未匹配" + key
	}
	return matches[1]
}