package backend

var (
	authEndpoint       = "https://readwise.io/api/v2/auth/"
	booksEndpoint      = "https://readwise.io/api/v2/books?page_size=1000&category=books&source=OctoberForKobo"
	coverEndpoint      = "https://readwise.io/api/v2/books/%d"
	highlightsEndpoint = "https://readwise.io/api/v2/highlights/"
	sourceCategory     = "books"
	sourceType         = "OctoberForKobo"
	userAgent          = "october/2.0.0 <https://github.com/marcus-crane/october>"
)
