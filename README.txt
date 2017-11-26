应用名称 AppID      API Key                      Secret Key 
demo    10380310   XXXXXXXXXXXXXXXXXXXXXXXX     XXXXXXXXXXXXXXXXXXXXXXXX 

https://i.loli.net/2017/11/17/5a0e99056e9e4.png

http://ww1.sinaimg.cn/large/0060lm7Tly1fll5igdgstj30ij03g3z5.jpg
http://ww2.sinaimg.cn/large/0060lm7Tly1flpxzfpq31j30ny0853zd.jpg

https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic
curl -i -k 'https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic?access_token=24.7d5bcd947bb57ec54ad1169664b33464.2592000.1513493122.282335-10380310' --data 'img=""&url=http://ww1.sinaimg.cn/large/0060lm7Tly1fll5igdgstj30ij03g3z5.jpg&top_num=5' -H 'Content-Type:application/x-www-form-urlencoded'


curl -i -k 'https://aip.baidubce.com/rest/2.0/image-classify/v2/dish?access_token=【调用鉴权接口获取的token】' --data 'image=【图片Base64编码，需UrlEncode】&top_num=5' -H 'Content-Type:application/x-www-form-urlencoded'