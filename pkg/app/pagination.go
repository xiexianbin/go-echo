package app

// SearchOptions define the fields expect in the URL query string
type SearchOptions struct {
	Query any      `url:"q,omitempty"`    // maps to "q" parameter in query string
	Sort  []string `url:"sort,omitempty"` // Slice of sorting criteria
	Page  int      `url:"page,omitempty"` // maps to "page" parameter
	Size  int      `url:"size,omitempty"` // maps to "size" parameter
}

// DecodeQuery decode the query string
// func DecodeQuery(c echo.Context) (*SearchOptions, error) {
// 	var options SearchOptions
// 	err := c.Bind(&options)
// 	if err != nil {
// 		return nil, err
// 	}

// 	options.Sort = []string{}
// 	for _, sortParam := range values["sort"] {
// 		parts := strings.Split(sortParam, ",")
// 		if len(parts) < 2 {
// 			continue
// 		}
// 		field := parts[0]
// 		order := parts[1]
// 		options.Sort = append(options.Sort, field+","+order)
// 	}

// 	return &options, nil
// }
