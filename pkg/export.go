package redcap

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

/*
ExportArms exports arms from a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportArms() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=arm&format=%s", r.Token, r.ResponseFormat)

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
ExportDags exports data access groups from a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportDags() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=dag&format=%s", r.Token, r.ResponseFormat)

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
ExportEvents exports events from a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportEvents() ([]byte, error) {
	// TODO: This looks like it will fail? What does it mean by `arms=`?
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=event&format=%s&arms=", r.Token, r.ResponseFormat)

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
ExportFeildNames exports field names from a REDCap project.

Args:

	Field: The field name to export.

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportFieldNames(feild string) ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=exportFieldNames&format=%s&field=%s", r.Token, r.ResponseFormat, feild)

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
ExportFile exports a file from a REDCap project.

Args:

	record: The record ID of the file to export.
	field: The field name of the file to export.
	event: The event name of the file to export.

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportFile(record string, feild string, event string) ([]byte, error) {
	client := &http.Client{}
	// TODO: Response format is not being used? Check API docs!
	formating := fmt.Sprintf("token=%s&content=file&action=export&record=%s&field=%s&event=%s", r.Token, record, feild, event)

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
ExportInstrumentEventMaps exports instrument event maps from a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportInstrumentEventMaps() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=formEventMapping&format=%s", r.Token, r.ResponseFormat)

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
ExportInstrumentPDF exports instrument PDFs from a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportInstrumentPDF() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=pdf&format=%s", r.Token, r.ResponseFormat)

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
ExportInstruments exports instruments from a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportInstruments() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=instrument&format=%s", r.Token, r.ResponseFormat)

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
	fmt.Printf("%s\n", bodyText)
	return bodyText, nil
}

/*
ExportLogging exports logging from a REDCap project.

Args:

	startTime: The start time of the log.
	endTime: The end time of the log.

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportLogging(startTime time.Time, endTime time.Time) ([]byte, error) {
	// TODO: COME BACK TO THIS. The logType ...string might not be the best way to handle this
	// TODO: IE: logType ...string, user ...string, record ...string
	// TODO: time.Time also needs to match the format of 10/06/2020 17:37
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=log&format=%s&logtype=&user=&record=&beginTime=%s&endTime=%s", r.Token, r.ResponseFormat, startTime, endTime)

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
ExportMetadata exports metadata from a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportMetadata() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=metadata&format=%s", r.Token, r.ResponseFormat)

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
ExportProjectXML exports project XML from a REDCap project.

Args:

	startTime: The start time of the log.
	endTime: The end time of the log.

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportProjectXML() ([]byte, error) {
	client := &http.Client{}
	// TODO: Right now we are not going to pass any additional parameters, we will have to come back to this.
	formating := fmt.Sprintf("token=%s&content=project_xml&returnMetadataOnly=false&exportSurveyFields=false&exportDataAccessGroups=false&returnformat=%s", r.Token, r.ResponseFormat)

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
ExportProject exports project from a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportProject() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=project&format=%s", r.Token, r.ResponseFormat)

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
	fmt.Printf("%s\n", bodyText)
	return bodyText, nil
}

/*
ExportRecords exports records from a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportRecords() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=record&format=%s&type=flat", r.Token, r.ResponseFormat)

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
ExportRedcapVersion exports logging from a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportRedcapVersion() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=version", r.Token)

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
ExportReports exports logging from a REDCap project.

Args:

	reportId: The report ID to export.

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportReports(reportID string) ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=report&format=%s&report_id=%s", r.Token, r.ResponseFormat, reportID)

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
ExportSurveyLink exports a survey link from a REDCap project.

Args:

	reportId: The report ID to export.

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportSurveyLink(recordID string, instrument string, event string) ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=surveyLink&record=%s&instrument=%s&event=%s&format=%s", r.Token, recordID, instrument, event, r.ResponseFormat)

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
ExportSurveyParticipants exports survey participants from a REDCap project.

Args:

	reportId: The report ID to export.

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportSurveyParticipants(instrument string, event string) ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=participantList&instrument=%s&event=%s&format=%s", r.Token, instrument, event, r.ResponseFormat)

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
ExportSurveyQueueLink exports survey queue links from a REDCap project.

Args:

	recordID: The record ID to export.
	instrument: The instrument to export.
	event: The event to export.

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportSurveyQueueLink(recordID string, instrument string, event string) ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=surveyQueueLink&record=%s&instrument=%s&event=%s&format=%s", r.Token, recordID, instrument, event, r.ResponseFormat)

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
ExportSurveyReturnCode exports survey return codes from a REDCap project.

Args:

	recordID: The record ID to export.
	instrument: The instrument to export.
	event: The event to export.

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportSurveyReturnCode(recordID string, instrument string, event string) ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=surveyReturnCode&record=%s&instrument=%s&event=%s&format=%s", r.Token, recordID, instrument, event, r.ResponseFormat)

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
	fmt.Printf("%s\n", bodyText)
	return bodyText, nil
}

/*
ExportDagMaps exports user DAG maps from a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportDagMaps() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=userDagMapping&format=%s", r.Token, r.ResponseFormat)

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
ExportUserRoles exports user roles from a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportUserRoles() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=userRole&format=%s", r.Token, r.ResponseFormat)

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
ExportUsers exports users from a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ExportUsers() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=user&format=%s", r.Token, r.ResponseFormat)

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
	fmt.Printf("%s\n", bodyText)
	return bodyText, nil
}
