package config

import (
	"os"
	"strings"
)

const (
	test_file = `
	{
		"以下为通用配置":"--------------------------------------------",
		"serial_number":"1.0",
		"service_name":"jewelryserver",
		"service_displayname":"jewelryserver",
		"sercice_desc":"jewelryserver",
		"http_port":"7009",
		"https_port":"8012",
		"db_url":"root:qwer@tcp(192.168.0.99:3306)/jewelry_db?charset=utf8&parseTime=True&loc=Local",
		"leveldb_dir":"./database",
		"token_type":"nomal",
		"app_id":"wwwthings",
		"app_secret":"4EE0A9A43B9B911C067BEE5CC50A9972",
		"is_dev":true,
		"oauth2_url":"http://192.168.0.152:7001/oauth2/api/v1",
		"apiserver_url":"http://192.168.0.152:7031/apiserver/api/v1",
		"register_url":"http://192.168.0.152:7021/register/api/v1",
		"es_addr_url":"http://192.168.0.152:9200",
		"domain_name":"http://192.168.0.40",
		"以下为逻辑配置":"--------------------------------------------",
			
		"wx_appid":"wx60ca0a6f3b6b5d1b",
		"wx_secret":"1ad30e53264144079de0998d9ec6e330",
		"wx_apiKey":"AIYE20180920WechatSuccess6981206",
		"wx_mchId":"1514466201",
		"sign_api_key":"apiserver",
		"sms_appid":"C40190998",
		"sms_secret":"0f5d6e69abb15e600c6ad169f8681a98",
		"paper_qrcode_url":"www.baidu.com",
		"register_app_id":"apiserver",
		"baidu_key":"pRISNUttEDgIl8G8tQqrIDCk",
		"baidu_secret":"AKbXfGVvM3RNcHB6HD01DUrGyz6HRPvW",
		"getui_appid":"Nd0EgNJdZL7CXu4shQane4",
		"getui_appkey":"VKGI1VBHjtA2AVeevTmiV5",
		"getui_mastersecret":"aoFZ1SvzOZ6tQ2KgDW9Du8",
		"cos_bucket":"jewelry",
		"cos_appid":"1257634258",
		"cos_secretid":"AKIDHTj3KUrNYjjZHnjFG5i8Rl4GG9uVwJQL",
		"cos_secretkey":"sb9aVJ5rTt7A5qTu9C3kHfrrwSp01qUw",
		"cos_region":"ap-chengdu",
		"cos_domain":"https://jewelry-1257634258.file.myqcloud.com",
		"content_count":10,
		"qrcode_wx_url":"pages/scanIn/index",
		"qrcode_url":"pages/scanIn/index?scene=",
		"qrcode_width":430,
		"top_count":5,
		"wx_public_name":"公众号暂无",
		"wx_notifyUrl":"http://www.xxjwxc.cn"
		}
		
`
)

//判断是否在测试环境下使用
func IsRunTesting() bool {
	if len(os.Args) > 1 {
		return strings.HasPrefix(os.Args[1], "-test")
	}
	return false
}
