package redcap

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

/*
DeleteArms deletes arms from a REDCap project.

Args:

	arms: A list of arms to delete.

Returns:

	A byte slice containing the response from the REDCap API.
*/
func (r *RedCapClient) DeleteArms(arms []string) ([]byte, error) {
	client := &http.Client{}
	builderType := BuilderType("arms")
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
	builderType := BuilderType("dags")
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
	builderType := BuilderType("events")
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
	builderType := BuilderType("userRoles")
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
	builderType := BuilderType("users")
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
