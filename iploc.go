package iploc

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
)

var COUNTRIES_ZH = map[string]string{
	"EU": "欧洲",
	"A1": "匿名代理",
	"A2": "卫星供应商",
	"AP": "亚洲/太平洋地区",
	"AQ": "南极洲",
	"BQ": "博内尔，圣尤斯特歇斯和萨巴",
	"SH": "圣赫勒拿",
	"FX": "法国，大都会",

	"AD": "安道​​尔",
	"AE": "阿拉伯联合大公国",
	"AF": "阿富汗",
	"AG": "安提瓜和巴布达",
	"AI": "安圭拉",
	"AL": "阿尔巴尼亚",
	"AM": "亚美尼亚",
	"AN": "荷属安的列斯群岛",
	"AO": "安哥拉",
	"AR": "阿根廷",
	"AS": "美属萨摩亚",
	"AT": "奥地利",
	"AU": "澳大利亚",
	"AW": "阿鲁巴",
	"AZ": "阿塞拜疆",
	"BA": "波斯尼亚和黑塞哥维那",
	"BB": "巴巴多斯",
	"BD": "孟加拉国",
	"BE": "比利时",
	"BF": "布基纳法索",
	"BG": "保加利亚",
	"BH": "巴林",
	"BI": "布隆迪",
	"BJ": "贝宁",
	"BM": "百慕大",
	"BN": "文莱",
	"BO": "玻利维亚",
	"BR": "巴西",
	"BS": "巴哈马",
	"BT": "不丹",
	"BV": "布维岛",
	"BW": "博茨瓦纳",
	"BY": "白俄罗斯",
	"BZ": "伯利兹",
	"CA": "加拿大",
	"CC": "科科斯群岛",
	"CD": "刚果",
	"CF": "中非共和国",
	"CG": "刚果",
	"CH": "瑞士",
	"CI": "科特迪瓦",
	"CK": "库克群岛",
	"CL": "智利",
	"CM": "喀麦隆",
	"CN": "中国",
	"CO": "哥伦比亚",
	"CR": "哥斯达黎加",
	"CU": "古巴",
	"CV": "佛得角",
	"CX": "圣诞岛",
	"CY": "塞浦路斯",
	"CZ": "捷克共和国",
	"DE": "德国",
	"DJ": "吉布提",
	"DK": "丹麦",
	"DM": "多米尼加",
	"DO": "多明尼加",
	"DZ": "阿尔及利亚",
	"EC": "厄瓜多尔",
	"EE": "爱沙尼亚",
	"EG": "埃及",
	"EH": "西撒哈拉",
	"ER": "厄立特里亚",
	"ES": "西班牙",
	"ET": "埃塞俄比亚",
	"FI": "芬兰",
	"FJ": "斐济",
	"FK": "福克兰群岛（马尔维纳斯群岛）",
	"FM": "密克罗尼西亚联邦",
	"FO": "法罗群岛",
	"FR": "法国",
	"GA": "加蓬",
	"GB": "联合王国",
	"GD": "格林纳达",
	"GE": "格鲁吉亚",
	"GF": "法属圭亚那",
	"GH": "加纳",
	"GI": "直布罗陀",
	"GL": "格陵兰",
	"GM": "冈比亚",
	"GN": "几内亚",
	"GP": "瓜德罗普岛",
	"GQ": "赤道几内亚",
	"GR": "希腊",
	"GS": "南乔治亚岛和南桑威奇群岛",
	"GT": "危地马拉",
	"GU": "关岛",
	"GW": "几内亚比绍",
	"GY": "圭亚那",
	"HK": "香港",
	"HM": "赫德岛和麦克唐纳群岛",
	"HN": "洪都拉斯",
	"HR": "克罗地亚",
	"HT": "海地",
	"HU": "匈牙利",
	"ID": "印尼",
	"IE": "爱尔兰",
	"IL": "以色列",
	"IN": "印度",
	"IO": "英属印度洋领地",
	"IQ": "伊拉克",
	"IR": "伊朗",
	"IS": "冰岛",
	"IT": "意大利",
	"JM": "牙买加",
	"JO": "约旦",
	"JP": "日本",
	"KE": "肯尼亚",
	"KG": "吉尔吉斯斯坦",
	"KH": "柬埔寨",
	"KI": "基里巴斯",
	"KM": "科摩罗",
	"KN": "圣基茨和尼维斯",
	"KP": "朝鲜",
	"KR": "韩国",
	"KW": "科威特",
	"KY": "开曼群岛",
	"KZ": "哈萨克斯坦",
	"LA": "老挝",
	"LB": "黎巴嫩",
	"LC": "圣卢西亚",
	"LI": "列支敦士登",
	"LK": "斯里兰卡",
	"LR": "利比里亚",
	"LS": "莱索托",
	"LT": "立陶宛",
	"LU": "卢森堡",
	"LV": "拉脱维亚",
	"LY": "利比亚",
	"MA": "摩洛哥",
	"MC": "摩纳哥",
	"MD": "摩尔多瓦",
	"MG": "马达加斯加",
	"MH": "马绍尔群岛",
	"MK": "马其顿",
	"ML": "马里",
	"MM": "缅甸",
	"MN": "蒙古",
	"MO": "澳门",
	"MP": "北马里亚纳群岛",
	"MQ": "马提尼克",
	"MR": "毛里塔尼亚",
	"MS": "蒙特塞拉特",
	"MT": "马耳他",
	"MU": "毛里求斯",
	"MV": "马尔代夫",
	"MW": "马拉维",
	"MX": "墨西哥",
	"MY": "马来西亚",
	"MZ": "莫桑比克",
	"NA": "纳米比亚",
	"NC": "新喀里多尼亚",
	"NE": "尼日尔",
	"NF": "诺福克岛",
	"NG": "尼日利亚",
	"NI": "尼加拉瓜",
	"NL": "荷兰",
	"NO": "挪威",
	"NP": "尼泊尔",
	"NR": "瑙鲁",
	"NU": "纽埃",
	"NZ": "新西兰",
	"OM": "阿曼",
	"PA": "巴拿马",
	"PE": "秘鲁",
	"PF": "法属波利尼西亚",
	"PG": "巴布亚新几内亚",
	"PH": "菲律宾",
	"PK": "巴基斯坦",
	"PL": "波兰",
	"PM": "圣皮埃尔和密克隆",
	"PN": "皮特凯恩群岛",
	"PR": "波多黎各",
	"PS": "巴勒斯坦领土",
	"PT": "葡萄牙",
	"PW": "帕劳",
	"PY": "巴拉圭",
	"QA": "卡塔尔",
	"RE": "团圆",
	"RO": "罗马尼亚",
	"RU": "俄罗斯联邦",
	"RW": "卢旺达",
	"SA": "沙特阿拉伯",
	"SB": "所罗门群岛",
	"SC": "塞舌尔",
	"SD": "苏丹",
	"SE": "瑞典",
	"SG": "新加坡",
	"SI": "斯洛文尼亚",
	"SJ": "斯瓦尔巴群岛和扬马延岛",
	"SK": "斯洛伐克",
	"SL": "塞拉利昂",
	"SM": "圣马力诺",
	"SN": "塞内加尔",
	"SO": "索马里",
	"SR": "苏里南",
	"ST": "圣多美和普林西比",
	"SV": "萨尔瓦多",
	"SY": "叙利亚",
	"SZ": "斯威士兰",
	"TC": "特克斯和凯科斯群岛",
	"TD": "乍得",
	"TF": "法国南部领土",
	"TG": "多哥",
	"TH": "泰国",
	"TJ": "塔吉克斯坦",
	"TK": "托克劳",
	"TM": "土库曼斯坦",
	"TN": "突尼斯",
	"TO": "汤加",
	"TL": "东帝汶",
	"TR": "土耳其",
	"TT": "特里尼达和多巴哥",
	"TV": "图瓦卢",
	"TW": "台湾",
	"TZ": "坦桑尼亚",
	"UA": "乌克兰",
	"UG": "乌干达",
	"UM": "美国本土外小岛屿",
	"US": "美国",
	"UY": "乌拉圭",
	"UZ": "乌兹别克斯坦",
	"VA": "罗马教廷（梵蒂冈城国）",
	"VC": "圣文森特和格林纳丁斯",
	"VE": "委内瑞拉",
	"VG": "英属维尔京群岛",
	"VI": "美属维京群岛，",
	"VN": "越南",
	"VU": "瓦努阿图",
	"WF": "瓦利斯群岛和富图纳群岛",
	"WS": "萨摩亚",
	"YE": "也门",
	"YT": "马约特岛",
	"RS": "塞尔维亚",
	"ZA": "南非",
	"ZM": "赞比亚",
	"ME": "黑山",
	"ZW": "津巴布韦",
	"AX": "奥兰群岛",
	"GG": "根西岛",
	"IM": "马恩岛",
	"JE": "新泽西州",
	"BL": "圣巴泰勒米",
	"MF": "圣马丁",
	"SS": "南苏丹",
}

var RESERVED_IP_RANGE = []struct {
	ipa   uint32
	ipb   uint32
	title string
}{
	{0, 16777215, "IANA保留作为特殊地址"},
	{167772160, 184549375, "IANA保留用于内部网络地址"},
	{1681915904, 1686110207, "IANA保留用于共享地址空间"},
	{2130706432, 2147483647, "IANA保留用于本机回环地址"},
	{2851995648, 2852061183, "IANA保留作为链路本地地址"},
	{2886729728, 2887778303, "IANA保留作为私有地址"},
	{3232235520, 3232301055, "IANA保留用于局域网地址"},
}

// IANA保留地址
const FLAG_RESERVED = 1 << 0

// 已分配地址
const FLAG_INUSE = 1 << 1

// IANA未分配地址
const FLAG_NOTUSE = 1 << 2

func ip2long(ipAddr string) (uint32, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0, errors.New("wrong ipAddr format")
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip), nil
}

func long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

var ipData = &struct {
	repo       []byte
	indexStart uint32
	indexNums  uint32
	dataStart  uint32
	loaded     bool
}{}

type IpInfo struct {
	Ip                               string
	IpLong                           uint32
	Flag                             uint8
	Code, Country, Region, City, Isp string
	Note                             string
}

type Store struct {
	data   []byte
	length uint32
	cursor uint32
}

func (s *Store) Tell() uint32 {
	return s.cursor
}

func (s *Store) Seek(n uint32) {
	if n > s.length {
		s.cursor = s.length
	} else {
		s.cursor = n
	}
}

func (s *Store) Read(n uint32) (result []byte) {
	if n > 0 && s.cursor < s.length {
		result = s.data[s.cursor : s.cursor+n]
		s.Seek(s.cursor + n)
	}
	return
}

func (s *Store) ReadUntil(char byte, params ...uint8) (result []byte) {
	var i, max uint8
	max = 255
	if len(params) > 0 {
		max = params[0]
	}
	start := s.cursor
	for {
		if i == max {
			break
		}
		chars := s.Read(1)
		if chars == nil {
			break
		}
		if chars[0] == char {
			result = s.data[start : s.cursor-1]
			return
		}
		i++
	}
	return
}

func NewStore(data []byte) *Store {
	return &Store{data, uint32(len(data)), 0}
}

type ipLoc struct {
	store      *Store
	indexStart uint32
	dataStart  uint32
	indexNums  uint32
}

func (ip *ipLoc) findOffset(ipLong uint32, start uint32, end uint32) uint32 {
find:
	if end-start <= 1 {
		return start*9 + ip.dataStart
	}

	middle := (end-start)/2 + start
	middleOffset := ip.indexStart + middle*4

	ip.store.Seek(middleOffset)
	ipLong2 := binary.LittleEndian.Uint32(ip.store.Read(4))

	if ipLong < ipLong2 {
		end = middle
	} else {
		start = middle
	}
	goto find
	return 0
}

func (ip *ipLoc) readAsString(str *string, data []byte, bits uint8) {
	var index uint32
	switch bits {
	case 16:
		index = uint32(binary.LittleEndian.Uint16(data))
	case 32:
		index = binary.LittleEndian.Uint32(data)
	}
	if index > 0 {
		ip.store.Seek(index)
		s := ip.store.ReadUntil(0x0)
		if s != nil {
			*str = string(s)
		}
	}
}

func (ip *ipLoc) readBlock(ipAddr string, ipLong uint32, offset uint32) (info *IpInfo) {
	store := ip.store
	store.Seek(offset)

	info = &IpInfo{}

	info.Ip = ipAddr
	info.IpLong = ipLong

	data := store.Read(9)

	flag := uint8(data[0])

	info.Flag = flag

	if flag == FLAG_INUSE {
		Code := data[1:3]
		info.Code = string(Code)

		if v, ok := COUNTRIES_ZH[info.Code]; ok {
			info.Country = v
		}

		ip.readAsString(&info.Region, data[3:5], 16)
		ip.readAsString(&info.City, data[5:7], 16)
		ip.readAsString(&info.Isp, data[7:9], 16)
	} else {
		switch flag {
		case FLAG_RESERVED:
			info.Note = "IANA保留地址"
			for _, r := range RESERVED_IP_RANGE {
				if ipLong >= r.ipa && ipLong <= r.ipb {
					info.Note = r.title
				}
			}
		case FLAG_NOTUSE:
			info.Note = "IANA未分配地址"
		}
	}
	return
}

func (ip *ipLoc) getInfo(ipAddr string) (info *IpInfo, err error) {
	if inited == false {
		return nil, errors.New("iploc file not inited")
	}

	ipLong, er := ip2long(ipAddr)
	if er != nil {
		err = er
		return
	}

	offset := ip.findOffset(ipLong, 0, ip.indexNums)
	info = ip.readBlock(ipAddr, ipLong, offset)
	return
}

func NewIpLoc() *ipLoc {
	return &ipLoc{NewStore(ipData.repo), ipData.indexStart, ipData.dataStart, ipData.indexNums}
}

var (
	inited bool
)

func IpLocInit(IPDatPath string) {
	if inited {
		return
	}

	fi, err := os.Stat(IPDatPath)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "your IpLoc file %s not exist, plz check\n", IPDatPath)
		os.Exit(1)
	}

	ipData.repo = make([]byte, fi.Size())
	file, _ := os.Open(IPDatPath)
	file.Read(ipData.repo)

	// read 12 bytes of file head
	ipData.indexStart = binary.LittleEndian.Uint32(ipData.repo[:4])
	ipData.dataStart = binary.LittleEndian.Uint32(ipData.repo[4:8])
	ipData.indexNums = binary.LittleEndian.Uint32(ipData.repo[8:12])

	inited = true
}

func GetIpInfo(ipAddr string) (*IpInfo, error) {
	ip := NewIpLoc()
	return ip.getInfo(ipAddr)
}
