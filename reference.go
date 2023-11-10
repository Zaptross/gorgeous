package gorgeous

import (
	_ "embed"
	"fmt"

	"github.com/google/uuid"
)

type Reference struct {
	name    string
	element *HTMLElement
}

const (
	ReferenceCacheServiceName = "referenceCache"
	ReferenceCacheService     = `
const referenceCache = {};

const observer = new MutationObserver(function (mutations) {
  mutations.forEach(function (mutation) {
    if (mutation.type === "childList" && mutation.removedNodes.length > 0) {
      const referencedElements = Object.keys(referenceCache).map(
        (el) => referenceCache[el].id
      );
      if (referencedElements.includes(mutation.removedNodes[0].id)) {
        delete referenceCache[mutation.removedNodes[0].id];
      }
    }
  });
});

// Get and cache a reference to the element with the given id
// if it is not already cached, otherwise return the cached element.
// Also setup an observer to remove the element from the cache if it is removed from the DOM.
function referenceCacheGet(elementId) {
  if (referenceCache[elementId] === undefined) {
    referenceCache[elementId] = document.getElementById(elementId);
    observer.observe(referenceCache[elementId].parentNode, {childList: true});
  }
  return referenceCache[elementId];
}
`
)

// CreateRef creates a reference to an element in the DOM.
// This is useful for creating interactive components, such as modals, forms, and popovers.
func CreateRef(name string) *Reference {
	Service(ReferenceCacheServiceName, ReferenceCacheService)

	return &Reference{
		name: name,
	}
}

// Get stores a reference to the element in the DOM, for later use, returning it
// so that it can be used in the component.
//
// // Create a reference to a form field for use in a script.
//
//	func LoginForm() *g.HTMLElement {
//		unRef := g.CreateRef("username")
//		pwRef := g.CreateRef("password")
//
//		return g.Form(g.EB{
//			Children: g.CE{
//				unRef.Get(g.Input(g.EB{
//					Props: g.Props{
//						"type":        "email",
//						"placeholder": "Username",
//					},
//				})),
//				pwRef.Get(g.Input(g.EB{
//					Props: g.Props{
//						"type":        "password",
//						"placeholder": "Password",
//					},
//				})),
//			},
//			Script: fmt.Sprintf(`
//				const un = %s;
//				const pw = %s;
//
//				fetch("/login", {
//					method: "POST",
//					body: JSON.stringify({
//						username: un.value,
//						password: pw.value,
//					}),
//				}).then(/* ... */)
//			`, unRef.Element(), pwRef.Element()),
//		})
//	}
func (r *Reference) Get(el *HTMLElement) *HTMLElement {
	if el.Id == "" || el.EB.Id == "" {
		el.Id = uuid.New().String()
	}

	r.element = el

	return r.element
}

// Javascript returns a javascript expression that can be used to reference the element.
// This is cached on the first lookup to avoid unnecessary DOM lookups.
//
// Note: this will return undefined if the element is not in the DOM.
// It is recommended to use nullish chaining to avoid errors.
//
// eg:
//
//	const el = referenceCacheGet("myElementId")?.value;
//	if (el === undefined) {
//		// ...
func (r *Reference) Javascript() JavaScript {
	return JavaScript(fmt.Sprintf(`referenceCacheGet("%s")`, r.element.Id))
}

// Element returns the element that the reference points to.
// This is useful for passing the element to a function in Go.
//
// Note: this will return nil if the element has not been set.
// If you are using the ref within the same component, you should use the element directly,
// or consider moving the ref to the higher level component.
//
// eg:
//
//			func MyHigherComponent() *g.HTMLElement {
//					ref := g.CreateRef("myElement")
//			   // ...
//			   return g.Div(g.EB{
//		       Children: g.CE{
//	 	       MyComponent(ref)
//					 },
//			     Script: fmt.Sprintf(`
//			       const el = %s;
//			       // ...
//			     `, ref.Element()),
//		    }
//			}
//
//			func MyComponent(ref *g.Reference) *g.HTMLElement {
//			   // ...
//			   return g.Div(g.EB{
//			     Children: g.CE{
//			       ref.Get(g.Input(g.EB{})),
//			     }
//			   })
//			}
func (r *Reference) Element() *HTMLElement {
	return r.element
}
