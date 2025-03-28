package redcap

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Arm struct {
	ArmNum string `json:"arm_num"`
	Name   string `json:"name"`
}

//
// func (r *RedCapClient) ExportArms(arms []string) ([]Arm, error) {
// 	form := url.Values{}
// 	form.Set("token", r.Token)
// 	form.Set("content", "arm")
// 	form.Set("format", r.ResponseFormat.String())
// 	form.Set("action", "export")
//
// 	for i, arm := range arms {
// 		form.Set(fmt.Sprintf("arms[%d]", i), arm)
// 	}
//
// 	req, err := r.HTTPClient.Do("POST", r.URL, strings.NewReader(form.Encode()))
// 	if err != nil {
// 		return nil, fmt.Errorf("creating request: %w", err)
// 	}
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//
// 	resp, err := r.HTTPClient.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("sending request: %w", err)
// 	}
// 	defer resp.Body.Close()
//
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("reading response: %w", err)
// 	}
//
// 	var armsResponse []Arm
// 	if err := json.Unmarshal(body, &armsResponse); err != nil {
// 		return nil, fmt.Errorf("decoding response: %w", err)
// 	}
// 	return armsResponse, nil
// }

func (r *RedCapClient) ImportArms(arms []string) ([]Arm, error) {
	form := url.Values{}
	form.Set("token", r.Token)
	form.Set("content", "arm")
	form.Set("format", r.ResponseFormat.String())
	form.Set("action", "import")

	for i, arm := range arms {
		form.Set(fmt.Sprintf("arms[%d]", i), arm)
	}

	req, err := http.NewRequest("POST", r.URL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := r.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}
	fmt.Printf("Response Body: %s\n", string(body)) // Debugging line to see the response
	var armsResponse []Arm
	if err := json.Unmarshal(body, &armsResponse); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}
	return armsResponse, nil
}

func (r *RedCapClient) DeleteArms(arms []string) ([]Arm, error) {
	form := url.Values{}
	form.Set("token", r.Token)
	form.Set("content", "arm")
	form.Set("format", r.ResponseFormat.String())
	form.Set("action", "delete")

	for i, arm := range arms {
		form.Set(fmt.Sprintf("arms[%d]", i), arm)
	}

	req, err := http.NewRequest("POST", r.URL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	var armsResponse []Arm
	if err := json.Unmarshal(body, &armsResponse); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}
	return armsResponse, nil
}
