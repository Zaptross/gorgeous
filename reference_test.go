package gorgeous

import "testing"

func TestCreateRef(t *testing.T) {
	ref := CreateRef("test")

	if ref.name != "test" {
		t.Errorf("CreateRef(\"test\") did not set name to \"test\", got: \"%s\"", ref.name)
	}
}

func TestReferenceGet(t *testing.T) {
	ref := CreateRef("test")
	el := Div(EB{})
	ref.Get(el)

	if ref.element.Id != el.Id {
		t.Errorf("Get(el) did not set elementId to el.Id, got: \"%s\"", ref.element.Id)
	}
}

func TestReferenceJavascript(t *testing.T) {
	ref := CreateRef("test")
	el := Div(EB{})
	ref.Get(el)

	if ref.Javascript().String() != "referenceCacheGet(\""+el.Id+"\")" {
		t.Errorf("Element() did not return \"referenceCacheGet(\"+el.Id+\")\", got: \"%s\"", ref.Javascript())
	}
}

func TestReferenceCacheServiceAdded(t *testing.T) {
	ref := CreateRef("test")
	el := Div(EB{})
	ref.Get(el)

	if services[ReferenceCacheServiceName] != ReferenceCacheService {
		t.Errorf("ReferenceCacheService was not added to services map")
	}
}

func TestReferenceElement(t *testing.T) {
	ref := CreateRef("test")
	el := Div(EB{})
	ref.Get(el)

	if ref.Element() != el {
		t.Errorf("Element() did not return el, got: \"%v\"", ref.Element())
	}
}
