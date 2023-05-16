package gorgeous

// Document returns a new HTML document structure.
// eg:
// 	<!DOCTYPE html>
// 	<html>
// 	  <head>
// 	    <!-- head -->
// 	  </head>
// 	  <body>
// 	    <!-- body -->
// 	  </body>
// 	</html>
func Document(
	head *HTMLElement,
	body *HTMLElement,
) *HTMLElement {
	return Html(
		EB{
			Children: CE{
				head,
				body,
			},
		},
	)
}
