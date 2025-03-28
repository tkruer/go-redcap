package redcap

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

/*
ImporArms imports arms into a REDCap project.

Args:

	None

Returns:

	A byte slice containing the response from the REDCap API.
*/
// func (r *RedCapClient) ImportArms() ([]byte, error) {
// 	// TODO: We need to come back to this and implement a loop to iterate over the parameters as a JSON builder
// 	formating := fmt.Sprintf(`token=%s&content=arm&action=import&override=0&format=%s&data=[{\"arm_num\":\"1\",\"name\":\"Arm%201\"}]`, r.Token, r.ResponseFormat)
//
// 	data := strings.NewReader(formating)
// 	req, err := http.NewRequest("POST", r.URL, data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Set("Accept", "application/json")
// 	resp, err := r.HTTPClient.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	defer resp.Body.Close()
//
// 	bodyText, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s\n", bodyText)
// 	return bodyText, nil
// }

// func (r *RedCapClient) ImportArms(fields map[string]string) ([]byte, error) {
// 	// Use builder to create form values
// 	form := buildForm(fields)
//
// 	// Build HTTP request
// 	req, err := http.NewRequest("POST", r.URL, strings.NewReader(form.Encode()))
// 	if err != nil {
// 		return nil, fmt.Errorf("creating request: %w", err)
// 	}
//
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Set("Accept", "application/json")
//
// 	resp, err := r.HTTPClient.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("sending request: %w", err)
// 	}
// 	defer resp.Body.Close()
//
// 	bodyText, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("reading response: %w", err)
// 	}
//
// 	return bodyText, nil
// }
//
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
