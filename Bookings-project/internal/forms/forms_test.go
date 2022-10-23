package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whataver", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whataver", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whataver", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows that it does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	hasForm := form.Has("a")
	if hasForm {
		t.Error("form shows that has is not empty when it should be empty")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	x := form.Has("a")
	if !x {
		t.Error("form shows that fields are blank when they aren't")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whataver", nil)
	form := New(r.PostForm)
	form.MinLength("b", 2)
	if form.Valid() {
		t.Error("form shows that min lenght was respected when it wasn't")
	}

	isError := form.Errors.Get("b")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedValues := url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)
	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minlength of 100 when data is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")
	form = New(postedValues)
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows minlength of 1 was not attained when it was")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)
	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "me@here.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when it should be valid")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "x")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got a valid email when it should be invalid")
	}
}
