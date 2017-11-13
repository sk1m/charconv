package charconv

import (
	"testing"
)

type testCase struct {
	utf8, other, otherEncoding string
}

// testCases for encoding and decoding.
var testCases = []testCase{
	{"Résumé", "Résumé", "utf8"},
	{"Résumé", "R\xe9sum\xe9", "latin1"},
	{"これは漢字です。", "S0\x8c0o0\"oW[g0Y0\x020", "UTF-16LE"},
	{"これは漢字です。", "0S0\x8c0oo\"[W0g0Y0\x02", "UTF-16BE"},
	{"Hello, world", "Hello, world", "ASCII"},
	{"Gdańsk", "Gda\xf1sk", "ISO-8859-2"},
	{"Ââ Čč Đđ Ŋŋ Õõ Šš Žž Åå Ää", "\xc2\xe2 \xc8\xe8 \xa9\xb9 \xaf\xbf \xd5\xf5 \xaa\xba \xac\xbc \xc5\xe5 \xc4\xe4", "ISO-8859-10"},
	{"สำหรับ", "\xca\xd3\xcb\xc3\u047a", "ISO-8859-11"},
	{"latviešu", "latvie\xf0u", "ISO-8859-13"},
	{"Seònaid", "Se\xf2naid", "ISO-8859-14"},
	{"€1 is cheap", "\xa41 is cheap", "ISO-8859-15"},
	{"românește", "rom\xe2ne\xbate", "ISO-8859-16"},
	{"nutraĵo", "nutra\xbco", "ISO-8859-3"},
	{"Kalâdlit", "Kal\xe2dlit", "ISO-8859-4"},
	{"русский", "\xe0\xe3\xe1\xe1\xda\xd8\xd9", "ISO-8859-5"},
	{"ελληνικά", "\xe5\xeb\xeb\xe7\xed\xe9\xea\xdc", "ISO-8859-7"},
	{"Kağan", "Ka\xf0an", "ISO-8859-9"},
	{"Résumé", "R\x8esum\x8e", "macintosh"},
	{"hello_мир", "hello_\xec\xe8\xf0", "windows-1251"},
	{"hello_мир", "hello_\xdc\xd8\xe0", "ISO-8859-5"},
	{"hello_мир", "hello_\xcd\xc9\xd2", "KOI8-R"},
	{"hello_мир", "hello_\xac\xa8\xe0", "cp866"},
	{"hello_мир", "hello_\xac\xa8\xe0", "CP866"},
	{"hello_мир", "hello_\xac\xa8\xe0", "IBM866"},
	{"Gdańsk", "Gda\xf1sk", "windows-1250"},
	{"русский", "\xf0\xf3\xf1\xf1\xea\xe8\xe9", "windows-1251"},
	{"Résumé", "R\xe9sum\xe9", "windows-1252"},
	{"ελληνικά", "\xe5\xeb\xeb\xe7\xed\xe9\xea\xdc", "windows-1253"},
	{"Kağan", "Ka\xf0an", "windows-1254"},
	{"עִבְרִית", "\xf2\xc4\xe1\xc0\xf8\xc4\xe9\xfa", "windows-1255"},
	{"العربية", "\xc7\xe1\xda\xd1\xc8\xed\xc9", "windows-1256"},
	{"latviešu", "latvie\xf0u", "windows-1257"},
	{"Việt", "Vi\xea\xf2t", "windows-1258"},
	{"สำหรับ", "\xca\xd3\xcb\xc3\u047a", "windows-874"},
	{"русский", "\xd2\xd5\xd3\xd3\xcb\xc9\xca", "KOI8-R"},
	{"українська", "\xd5\xcb\xd2\xc1\xa7\xce\xd3\xd8\xcb\xc1", "KOI8-U"},
	{"Hello 常用國字標準字體表", "Hello \xb1`\xa5\u03b0\xea\xa6r\xbc\u0437\u01e6r\xc5\xe9\xaa\xed", "big5"},
	{"Hello 常用國字標準字體表", "Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed", "gbk"},
	{"Hello 常用國字標準字體表", "Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed", "gb18030"},
	{"עִבְרִית", "\x81\x30\xfb\x30\x81\x30\xf6\x34\x81\x30\xf9\x33\x81\x30\xf6\x30\x81\x30\xfb\x36\x81\x30\xf6\x34\x81\x30\xfa\x31\x81\x30\xfb\x38", "gb18030"},
	{"㧯", "\x82\x31\x89\x38", "gb18030"},
	{"これは漢字です。", "\x82\xb1\x82\xea\x82\xcd\x8a\xbf\x8e\x9a\x82\xc5\x82\xb7\x81B", "SJIS"},
	{"Hello, 世界!", "Hello, \x90\xa2\x8aE!", "SJIS"},
	{"ｲｳｴｵｶ", "\xb2\xb3\xb4\xb5\xb6", "SJIS"},
	{"これは漢字です。", "\xa4\xb3\xa4\xec\xa4\u03f4\xc1\xbb\xfa\xa4\u01e4\xb9\xa1\xa3", "EUC-JP"},
	{"Hello, 世界!", "Hello, \x1b$B@$3&\x1b(B!", "ISO-2022-JP"},
	{"다음과 같은 조건을 따라야 합니다: 저작자표시", "\xb4\xd9\xc0\xbd\xb0\xfa \xb0\xb0\xc0\xba \xc1\xb6\xb0\xc7\xc0\xbb \xb5\xfb\xb6\xf3\xbe\xdf \xc7մϴ\xd9: \xc0\xfa\xc0\xdb\xc0\xdaǥ\xbd\xc3", "EUC-KR"},
}

func TestDecode(t *testing.T) {
	for _, tc := range testCases {
		got := Decode(tc.other, tc.otherEncoding)
		if got != tc.utf8 {
			t.Errorf("%s: got %q, want %q", tc.otherEncoding, got, tc.utf8)
		}
	}
}

type encTest struct {
	base64   string
	encoding string
}

var encTests = []encTest{
	{base64: "aGVsbG9f7Ojw", encoding: "windows-1251"},
	{base64: "agvsbg9f3njg", encoding: "iso-8859-5"},
	{base64: "agvsbg9f0lzqunga", encoding: "utf-8"},
	{base64: "agvsbg9fzcns", encoding: "koi8-r"},
	{base64: "agvsbg9frkjg", encoding: "ibm866"},
}
