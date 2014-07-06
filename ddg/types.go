package ddg

type Info struct {
	// DefinitionSource: name of Definition source
	DefinitionSource string
	// Heading: name of topic that goes with Abstract
	Heading    string
	ImageWidth int
	// RelatedTopics: array of internal links to related topics associated with Abstract
	Related []Result
	// Type: response category, i.e. A (article), D (disambiguation),
	// C (category), N (name), E (exclusive), or nothing.
	Type string
	// Redirect: !bang redirect URL
	Redirect string
	// DefinitionURL: deep link to expanded definition page in DefinitionSource
	DefinitionURL string
	// AbstractURL: deep link to expanded topic page in AbstractSource
	AbstractURL string
	// Definition: dictionary definition (may differ from Abstract)
	Definition string
	// AbstractSource: name of Abstract source
	AbstractSource string
	InfoBox        InfoBox
	// Image: link to image that goes with Abstract
	Image       string
	ImageIsLogo int
	// Abstract: topic summary (can contain HTML, e.g. italics)
	Abstract string
	// AbstractText: topic summary (with no HTML)
	AbstractText string
	// AnswerType: type of Answer, e.g. calc, color, digest, info,
	// ip, iploc, phone, pw, rand, regexp, unicode, upc, or zip
	// (see goodies & tech pages for examples).
	AnswerType  string
	ImageHeight int
	Results     []Result
	// Answer: instant answer
	Answer string
}

type Result struct {
	// Result: HTML link(s) to related topic(s)
	Result string
	// Icon: icon associated with related topic(s)
	Icon struct {
		// URL: URL of icon
		URL string
		// Height: height of icon (px)
		Height int
		// Width: width of icon (px)
		Width int
	}
	// FirstURL: first URL in Result
	FirstURL string
	// Text: text from first URL
	Text string
}

type InfoBox struct {
	Content []Tag `json:"content"`
	Meta    []Tag `json:"meta"`
}

type Tag struct {
	Type  string `json:"data_type"`
	Value string `json:"value"`
	Label string `json:"label"`
}
