package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestErrorHandling(t *testing.T) {
	testCases := []struct {
		banner           string
		input            string
		expectedResponse int
	}{
		{
			// when valid banner and input is passed
			banner:           "shadow",
			input:            "this is history",
			expectedResponse: http.StatusOK,
		},
		{
			// when an invalid input file is passed
			banner:           "shadow",
			input:            "",
			expectedResponse: http.StatusBadRequest,
		},
		{
			// when an invalid banner file
			banner:           "shadwo",
			input:            "this is history",
			expectedResponse: http.StatusInternalServerError,
		},
	}
	for _, tc := range testCases {
		form := url.Values{}
		form.Add("banner", tc.banner)
		form.Add("input-name", tc.input)

		request := httptest.NewRequest("POST", "/ascii-art", strings.NewReader(form.Encode()))
		request.PostForm = form
		responseRecorder := httptest.NewRecorder()

		Process(responseRecorder, request)
		if responseRecorder.Code != tc.expectedResponse {
			t.Errorf("Want status '%d', got '%d'", tc.expectedResponse, responseRecorder.Code)
		}
		// assert.Equal(t, responseRecorder.Code, tc.expectedResponse)
	}
}

func TestPostEndpoint(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(Process))
	defer svr.Close()
	form := url.Values{}
	form.Add("banner", "thinkertoy")
	form.Add("input-name", "hello sad world")

	request, err := http.NewRequest("POST", svr.URL+"/ascii-art", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Want status '%d', got '%d'", 200, resp.StatusCode)
	}
}
func TestGetEndpoint(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(Process))
	defer svr.Close()
	resp, err := http.Get(svr.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Want status '%d', got '%d'", 200, resp.StatusCode)
	}
}
func ReadTestFile(arg string) []string {
	data, err := ioutil.ReadFile(arg)
	if err != nil {
		// panic(err)
		fmt.Print("lol")
	}
	contentString := string(data)
	contentSplit := strings.Split(contentString, "#")
	return contentSplit
}
func TestAsciiArt(t *testing.T) {
	testCases := []string{
		"hello",
		"HELLO",
		"HeLlo HuMaN",
		"1Hello 2There",
		"Hello\\nThere",
		"{Hello & There #}",
		"hello There 1 to 2!",
		"MaD3IrA&LiSboN",
		"1a\"#FdwHywR&/()=",
		"{|}~",
		"[\\]^_ 'a",
		"RGB",
		":;<=>?@",
		"\\!\" #$%&'()*+,-./",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghijklmnopqrstuvwxyz",
		"nice 2 meet you",
		"It's Working",
	}
	ban := []string{
		"standard",
		"shadow",
		"thinkertoy",
	}
	filename := []string{
		"standardoutput.txt",
		"shadowoutput.txt",
		"thinkertoyoutput.txt",
	}
	for i := range ban {
		expectedOutput := ReadTestFile(filename[i])
		for index, testCase := range testCases {
			output, err := AsciiArt(testCase, ban[i]) // test the AsciiArt(input, ban) function
			if err != nil {
				panic(err)
			}
			if output != expectedOutput[index] {
				t.Errorf("\nTest fails when given case:\n\t\"%s\","+"\nThe test should show:\n%s\nInstead it shows:\n%s", testCase, expectedOutput[index], output)
			}
		}
	}
}
