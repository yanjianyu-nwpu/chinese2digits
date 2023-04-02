package chinese2digits

// NumberResult 解析的结果
type NumberResult struct {
	InputText       string   `json:"input_text,omitempty"`    // 是input text
	ReplacedText    string   `json:"replaced_text,omitempty"` // 改成chinese之后的记过
	CHNumberStrList []string `json:"origin_ch_number_list"`
	DigitsStrList   []string `json:"digits_str_list"`
}
