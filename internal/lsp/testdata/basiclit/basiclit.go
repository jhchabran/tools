package basiclit

func _() {
	var a int // something for lexical completions

	_ = "hello." //@complete(".")

	_ = 1 //@complete(" //")

	_ = 1. //@complete(".")

	_ = 'a' //@complete("' ")

	_ = 'a' //@hovertooltip("'a'", "'a', U+0061, LATIN SMALL LETTER A")
	_ = 0x61 //@hovertooltip("0x61", "'a', U+0061, LATIN SMALL LETTER A")

	_ = '\u2211' //@hovertooltip("'\\u2211'", "'∑', U+2211, N-ARY SUMMATION")
	_ = 0x2211 //@hovertooltip("0x2211", "'∑', U+2211, N-ARY SUMMATION")
	_ = "foo \u2211 bar" //@hovertooltip("\\u2211", "'∑', U+2211, N-ARY SUMMATION")

	_ = '\a' //@hovertooltip("'\\a'", "U+0007, control")
	_ = "foo \a bar" //@hovertooltip("\\a", "U+0007, control")

	_ = '\U0001F30A' //@hovertooltip("'\\U0001F30A'", "'🌊', U+1F30A, WATER WAVE")
	_ = 0x0001F30A //@hovertooltip("0x0001F30A", "'🌊', U+1F30A, WATER WAVE")
	_ = "foo \U0001F30A bar" //@hovertooltip("\\U0001F30A", "'🌊', U+1F30A, WATER WAVE")

	_ = '\x7E' //@hovertooltip("'\\x7E'", "'~', U+007E, TILDE")
	_ = "foo \x7E bar" //@hovertooltip("\\x7E", "'~', U+007E, TILDE")

	_ = '\173' //@hovertooltip("'\\173'", "'{', U+007B, LEFT CURLY BRACKET")
	_ = "foo \173 bar" //@hovertooltip("\\173", "'{', U+007B, LEFT CURLY BRACKET")
	_ = "foo \173 bar \u2211 baz" //@hovertooltip("\\173", "'{', U+007B, LEFT CURLY BRACKET")
	_ = "foo \173 bar \u2211 baz" //@hovertooltip("\\u2211", "'∑', U+2211, N-ARY SUMMATION")
	_ = "foo\173bar\u2211baz" //@hovertooltip("\\173", "'{', U+007B, LEFT CURLY BRACKET")
	_ = "foo\173bar\u2211baz" //@hovertooltip("\\u2211", "'∑', U+2211, N-ARY SUMMATION")

	// search for runes in string only if there is an escaped sequence
	_ = "hello" //@hovertooltip("\"hello\"", "")

	// incorrect escaped rune sequences
	_ = '\0' //@hovertooltip("'\\0'", "")
	_ = '\u22111' //@hovertooltip("'\\u22111'", "")
	_ = '\U00110000' //@hovertooltip("'\\U00110000'", "")
	_ = '\u12e45'//@hovertooltip("'\\u12e45'", "")
	_ = '\xa' //@hovertooltip("'\\xa'", "")
	_ = 'aa' //@hovertooltip("'aa'", "")
	
	// other basic lits
	_ = 1 //@hovertooltip("1", "")
	_ = 1.2 //@hovertooltip("1.2", "")
	_ = 1.2i //@hovertooltip("1.2i", "")
	_ = 0123 //@hovertooltip("0123", "")
	_ = 0x1234567890 //@hovertooltip("0x1234567890", "")
}
