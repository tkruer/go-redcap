package redcap

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type ResponseFormat string

const (
	JSON ResponseFormat = "json"
	XML  ResponseFormat = "xml"
	CSV  ResponseFormat = "csv"
)

type BuilderType string

const (
	Arms   = "arms"
	Dags   = "dags"
	Events = "events"
	Users = "users"
	UserRoles = "userRoles"
)

type RedCapClient struct {
	Token          string
	URL            string
	ResponseFormat ResponseFormat
}

type RedCapResponse struct {
	Content string
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

/*
	DeleteArms deletes arms from a REDCap project.
	
	Args:
		arms: A list of arms to delete.
	
	Returns:
		A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) DeleteArms(arms []string) ([]byte, error) {
	
	client := &http.Client{}
	var builderType = BuilderType("arms")
	params := parameterBuilder(arms, builderType)
	formating := fmt.Sprintf("token=%s&content=arm&action=delete&format=%s&arms=%s", r.Token, r.ResponseFormat, params)

	data := strings.NewReader(formating)
	req, err := http.NewRequest("POST", r.URL, data)
	if err != nil {
		log.Fatal("Error creating HTTP request: ", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending HTTP request: ", err)
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	return bodyText, nil
}


/*
	DeleteDags deletes data access groups from a REDCap project.
	
	Args:
		dags: A list of arms to delete.
	
	Returns:
		A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) DeleteDags(dags []string) ([]byte, error) {
	client := &http.Client{}
	var builderType = BuilderType("dags")
	params := parameterBuilder(dags, builderType)
	formating := fmt.Sprintf("token=%s&content=dag&action=delete&format=%s&%s", r.Token, r.ResponseFormat, params)

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
	DeleteEvents deletes events from a REDCap project.
	
	Args:
		dags: A list of arms to delete.
	
	Returns:
		A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) DeleteEvents(events []string) ([]byte, error) {
	client := &http.Client{}
	var builderType = BuilderType("events")
	params := parameterBuilder(events, builderType)
	formating := fmt.Sprintf("token=%s&content=event&action=delete&format=%s&%s", r.Token, r.ResponseFormat, params)

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
	DeleteFile deletes a file from a REDCap project.
	
	Args:
		record: The record ID of the file to delete.
		field: The field name of the file to delete.
		event: The event name of the file to delete.
	
	Returns:
		A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) DeleteFile(record string, field string, event string) ([]byte, error) {
	client := &http.Client{}
	// TODO: Response format is not being used? Check API docs!
	formating := fmt.Sprintf("token=%s&content=file&action=delete&record=%s&field=%s&event=%s", r.Token, record, field, event)

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
	DeleteRecords deletes records from a REDCap project.
	
	Args:
		record: The record ID of the file to delete.
		field: The field name of the file to delete.
		event: The event name of the file to delete.
	
	Returns:
		A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) DeleteRecords(records string, arms string, instrument string) ([]byte, error) {
	// TODO: For simplicity implementing, we can for now just use strings for the parameters
	// TODO: We will need to come back and implement a **kwargs, (I think it's ...) in Go for this
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&action=delete&content=record&records[0]=%s&arm=%s&instrument=%s&event=visit_1_arm_1&returnformat=%s", r.Token, records, arms, instrument, r.ResponseFormat)
	// TODO: ^ The problem is for deleting multiple records or multiple arms, we need to implement a loop to iterate over the parameters
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
	DeleteUserRoles deletes user roles from a REDCap project.
	
	Args:
		roles: A list of roles to delete.
	
	Returns:
		A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) DeleteUserRoles(roles []string) ([]byte, error) {
	client := &http.Client{}
	var builderType = BuilderType("userRoles")
	params := parameterBuilder(roles, builderType)
	formating := fmt.Sprintf("token=%s&content=userRole&action=delete&format=%s&%s", r.Token, r.ResponseFormat, params)

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
	DeleteUsers deletes users from a REDCap project.
	
	Args:
		users: A list of users to delete.
	
	Returns:
		A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) DeleteUsers(users []string) ([]byte, error) {
	client := &http.Client{}
	var builderType = BuilderType("users")
	params := parameterBuilder(users, builderType)
	formating := fmt.Sprintf("token=%s&content=user&action=delete&format=%s&%s", r.Token, r.ResponseFormat, params)

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

/*
	ImporArms imports arms into a REDCap project.
	
	Args:
		None
	
	Returns:
		A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ImportArms() ([]byte, error) {
	// TODO: We need to come back to this and implement a loop to iterate over the parameters as a JSON builder
	client := &http.Client{}
	formating := fmt.Sprintf(`token=%s&content=arm&action=import&override=0&format=%s&data=[{\"arm_num\":\"1\",\"name\":\"Arm%201\"}]`, r.Token, r.ResponseFormat)

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
	ImportDags imports data access groups into a REDCap project.
	
	Args:
		None
	
	Returns:
		A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ImportDags() ([]byte, error) {
	// TODO: We need to come back to this and implement a loop to iterate over the parameters as a JSON builder
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=dag&action=import&format=%s&data=[{\"data_access_group_name\":\"Group%20API\",\"unique_group_name\":\"\"}]", r.Token, r.ResponseFormat)

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
	ImportEvents imports events into a REDCap project.
	
	Args:
		None
	
	Returns:
		A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ImportEvents() ([]byte, error) {
	// TODO: We need to come back to this and implement a loop to iterate over the parameters as a JSON builder
	client := &http.Client{}
	formating := fmt.Sprintf(`token=%s&content=event&action=import&override=0&format=%s&data=[{\"event_name\":\"Event%201\",\"arm_num\":\"1\",\"day_offset\":\"0\",\"offset_min\":\"0\",\"offset_max\":\"0\",\"unique_event_name\":\"event_1_arm_1\"}]`, r.Token, r.ResponseFormat)

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

// TODO: FIX THIS - FILE IMPORT IS A BINARY FILE
func (r *RedCapClient) ImportFile() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("", r.Token, r.ResponseFormat)

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
	ImportInstrumentEventMaps imports instrument event maps into a REDCap project.
	
	Args:
		None
	
	Returns:
		A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) ImportInstrumentEventMaps() ([]byte, error) {
	// TODO: We need to come back to this and implement a loop to iterate over the parameters as a JSON builder
	client := &http.Client{}
	formating := fmt.Sprintf(`token=%s&content=formEventMapping&format=%s&data=[{\"arm\":{\"number\":\"1\",\"event\":[{\"unique_event_name\":\"event_1_arm_1\",\"form\":[\"instr_1\",\"instr_2\"]}]}},{\"arm\":{\"number\":\"2\",\"event\":[{\"unique_event_name\":\"event_2_arm_1\",\"form\":[\"instr_1\"]}]}}]`, r.Token, r.ResponseFormat)

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

func (r *RedCapClient) ImportProject() ([]byte, error) {
	// TODO: We need to come back to this and implement a loop to iterate over the parameters as a JSON builder
	client := &http.Client{}
	formating := fmt.Sprintf("token=$API_SUPER_TOKEN&content=project&format=%s&data=[{\"project_title\":\"New%20Project%20via%20API\",\"purpose\":0,\"purpose_other\":\"\",\"project_note\":\"Some%20notes%20about%20the%20project\"}]", r.Token, r.ResponseFormat)

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

// TODO: FIX THIS - RECORD IMPORT IS A MESS
func (r *RedCapClient) ImportRecords() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("", r.Token, r.ResponseFormat)

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

func (r *RedCapClient) ImportUserDagMaps() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=userDagMapping&action=import&format=%s&data=[{\"username\":\"testuser\",\"redcap_data_access_group\":\"api_testing_group\"}]", r.Token, r.ResponseFormat)

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

func (r *RedCapClient) ImportUserRoles() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=userRole&format=%s&data=[{\"unique_role_name\":\"U-2119C4Y87T\",\"role_label\":\"Project Manager\",\"data_access_group\":\"1\",\"data_export\":\"0\",\"mobile_app\":\"0\",\"mobile_app_download_data\":\"0\",\"lock_records_all_forms\":\"0\",\"lock_records\":\"0\",\"lock_records_customization\":\"0\",\"record_delete\":\"0\",\"record_rename\":\"0\",\"record_create\":\"1\",\"api_import\":\"1\",\"api_export\":\"1\",\"api_modules\":\"1\",\"data_quality_execute\":\"1\",\"data_quality_create\":\"1\",\"file_repository\":\"1\",\"logging\":\"1\",\"data_comparison_tool\":\"1\",\"data_import_tool\":\"1\",\"calendar\":\"1\",\"stats_and_charts\":\"1\",\"reports\":\"1\",\"user_rights\":\"1\",\"design\":\"1\"}]", r.Token, r.ResponseFormat)

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

func (r *RedCapClient) ImportUsers() ([]byte, error) {
	client := &http.Client{}
	formating := fmt.Sprintf("token=%s&content=user&format=%s&data=[{\"username\":\"test_user_47\",\"expiration\":\"\",\"data_access_group\":\"1\",\"data_export\":\"0\",\"mobile_app\":\"0\",\"mobile_app_download_data\":\"0\",\"lock_record_multiform\":\"0\",\"lock_record\":\"0\",\"lock_record_customize\":\"0\",\"record_delete\":\"0\",\"record_rename\":\"0\",\"record_create\":\"1\",\"api_import\":\"1\",\"api_export\":\"1\",\"api_modules\":\"1\",\"data_quality_execute\":\"1\",\"data_quality_design\":\"1\",\"file_repository\":\"1\",\"data_logging\":\"1\",\"data_comparison_tool\":\"1\",\"data_import_tool\":\"1\",\"calendar\":\"1\",\"graphical\":\"1\",\"reports\":\"1\",\"user_rights\":\"1\",\"design\":\"1\"}]", r.Token, r.ResponseFormat)

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
