package redcap

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type RedCapClient struct {
	Token          string
	Url            string
	ResponseFormat string
}

func (r RedCapClient) DeleteArms() {
	client := &http.Client{}
	var formating = fmt.Sprintf("token=%d&content=arm&action=delete&format=%d&arms[0]=1", r.Token, r.ResponseFormat)

	var data = strings.NewReader(formating)
	req, err := http.NewRequest("POST", "http://example.com/redcap/api/", data)
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
}

func (r RedCapClient) DeleteDags() {}

func (r RedCapClient) DeleteEvents() {}

func (r RedCapClient) DeleteFile() {}

func (r RedCapClient) DeleteRecords() {}

func (r RedCapClient) DeleteUserRoles() {}

func (r RedCapClient) DeleteUsers() {}

func (r RedCapClient) ExportArms() {}

func (r RedCapClient) ExportDags() {}

func (r RedCapClient) ExportEvents() {}

func (r RedCapClient) ExportFieldNames() {}

func (r RedCapClient) ExportFile() {}

func (r RedCapClient) ExportInstrumentEventMaps() {}

func (r RedCapClient) ExportInstrumentPDF() {}

func (r RedCapClient) ExportInstruments() {}

func (r RedCapClient) ExportLogging() {}

func (r RedCapClient) ExportMetadata() {}

func (r RedCapClient) ExportProjectXML() {}

func (r RedCapClient) ExportProject() {}

func (r RedCapClient) ExportRecords() {}

func (r RedCapClient) ExportRedcapVersion() {}

func (r RedCapClient) ExportReports() {}

func (r RedCapClient) ExportSurveyLink() {}

func (r RedCapClient) ExportSurveyParticipants() {}

func (r RedCapClient) ExportSurveyQueueLink() {}

func (r RedCapClient) ExportSurveyReturnCode() {}

func (r RedCapClient) ExportDagMaps() {}

func (r RedCapClient) ExportUserRoles() {}

func (r RedCapClient) ExportUsers() {}

func (r RedCapClient) ImportArms() {}

func (r RedCapClient) ImportDags() {}

func (r RedCapClient) ImportEvents() {}

func (r RedCapClient) ImportFile() {}

func (r RedCapClient) ImportInstrumentEventMaps() {}

func (r RedCapClient) ImportProject() {}

func (r RedCapClient) ImportRecords() {}

func (r RedCapClient) ImportUserDagMaps() {}

func (r RedCapClient) ImportUserRoles() {}

func (r RedCapClient) ImportUsers() {}

func (r RedCapClient) RenameRecord() {}

func (r RedCapClient) SwitchDag() {}
