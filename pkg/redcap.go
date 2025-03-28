package redcap

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ResponseFormat string

const (
	JSON ResponseFormat = "json"
	XML  ResponseFormat = "xml"
	CSV  ResponseFormat = "csv"
)

func (rf ResponseFormat) String() string {
	return string(rf)
}

type BuilderType string

const (
	Arms      = "arms"
	Dags      = "dags"
	Events    = "events"
	Users     = "users"
	UserRoles = "userRoles"
)

type RedCapClient struct {
	Token          string
	URL            string
	ResponseFormat ResponseFormat
	HTTPClient     *http.Client
}

type RedCapResponse struct {
	Content    string
	StatusCode int
}

func parameterBuilder(parameters []string, builder BuilderType) string {
	var formating string
	switch builder {
	case Dags:
		for i, v := range parameters {
			formating += fmt.Sprintf("dags[%d]=%s", i, v)
		}
		return formating
	case Arms:
		for i, v := range parameters {
			formating += fmt.Sprintf("[%d]=%s", i, v)
		}
		return formating
	case Events:
		for i, v := range parameters {
			formating += fmt.Sprintf("events[%d]=%s", i, v)
		}
		return formating
	case UserRoles:
		for i, v := range parameters {
			formating += fmt.Sprintf("roles[%d]=%s", i, v)
		}
		return formating
	case Users:
		for i, v := range parameters {
			formating += fmt.Sprintf("users[%d]=%s", i, v)
		}
		return formating
	default:
		for i, v := range parameters {
			formating += fmt.Sprintf("[%d]=%s", i, v)
		}
		return formating
	}
}

func buildForm(fields map[string]string) url.Values {
	form := url.Values{}
	for k, v := range fields {
		form.Set(k, v)
	}
	return form
}

/*
ImportInstrumentEventMaps imports instrument event maps into a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) RenameRecord(record_id string, arm string, record_id_new string) ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&action=rename&content=record&record=%s&new_record_name=%s&arm=%s&returnFormat=%s", r.Token, record_id, record_id_new, arm, r.ResponseFormat)

	data := strings.NewReader(formating)
	req, err := http.NewRequest("POST", r.URL, data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return bodyText, nil
}

/*
ImportInstrumentEventMaps imports instrument event maps into a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) SwitchDag(dag string) ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=dag&action=switch&format=%s&dag=%s", r.Token, r.ResponseFormat, dag)

	data := strings.NewReader(formating)
	req, err := http.NewRequest("POST", r.URL, data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return bodyText, nil
}
