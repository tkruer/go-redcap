package redcaptest

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	redcap "github.com/tkruer/go-redcap/pkg"
)

type testConfig struct {
	ProjectName string
	Token       string
	ServerURL   string
	PID         int
	Server      string
	ReadOnly    bool
	Operational bool
	Notes       string
}

func initConfig() []testConfig {
	// Send a get request to get some test data
	resp, err := http.Get("https://raw.githubusercontent.com/redcap-tools/redcap-test-datasets/master/connections/oklahoma-bbmc.csv")
	if err != nil {
		log.Fatal("Error fetching CSV:", err)
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	reader.Comma = ',' // CSV field delimiter

	_, err = reader.Read()
	if err != nil {
		log.Fatal("Error reading headers:", err)
	}

	var configs []testConfig
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading record:", err)
		}

		if len(record) < 8 {
			log.Fatal("Record does not contain enough data")
		}

		pid, err := strconv.Atoi(record[3])
		if err != nil {
			log.Fatalf("Invalid PID: %v", err)
		}

		readOnly, err := strconv.ParseBool(strings.ToLower(record[5]))
		if err != nil {
			log.Fatalf("Invalid ReadOnly value: %v", err)
		}

		operational, err := strconv.ParseBool(strings.ToLower(record[6]))
		if err != nil {
			log.Fatalf("Invalid Operational value: %v", err)
		}

		config := testConfig{
			ProjectName: record[0],
			Token:       record[1],
			ServerURL:   record[2],
			PID:         pid,
			Server:      record[4],
			ReadOnly:    readOnly,
			Operational: operational,
			Notes:       record[7],
		}
		configs = append(configs, config)
	}
	return configs
}
func TestRedcap(t *testing.T) {
	config := initConfig()	

	var err error

	client := redcap.RedCapClient{
		URL:            config[1].ServerURL,
		Token:          config[1].Token,
		ResponseFormat: "json",
	}

	var testArr = []string{"1", "2"}

	_, err = client.DeleteArms(testArr)

	if err != nil {
		t.Error("\n", err)
	}
	
	var testDags = []string{"group_api", "group_api2"}
	_, err = client.DeleteDags(testDags)

	if err != nil {
		t.Error(err)
	}

	var testEvents = []string{"event_1", "event_2"}
	_, err = client.DeleteEvents(testEvents)

	if err != nil {
		t.Error(err)
	}


	_, err = client.DeleteFile("1", "test_field", "test_event")

	if err != nil {
		t.Error(err)
	}

	_, err = client.DeleteRecords("1", "example_field", "example_event")

	if err != nil {
		t.Error(err)
	}
	var testUserRoles = []string{"admin", "user"}
	_, err = client.DeleteUserRoles(testUserRoles)

	if err != nil {
		t.Error(err)
	}
	var testUsers = []string{"user1", "user2"}
	_, err = client.DeleteUsers(testUsers)

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportArms()

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportDags()

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportEvents()

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportFieldNames("some_field")

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportFile("1", "some_field", "some_event")

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportInstrumentEventMaps()

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportInstrumentPDF()

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportInstruments()

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportLogging(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC))

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportMetadata()

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportProjectXML()

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportProject()

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportRecords()

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportRedcapVersion()

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportReports("1")

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportSurveyLink("1", "some_instrument", "some_event")

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportSurveyParticipants("some_instrument", "some_event")

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportSurveyQueueLink("1", "some_instrument", "some_event")

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportSurveyReturnCode("1", "some_instrument", "some_event")

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportDagMaps()

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportUserRoles()

	if err != nil {
		t.Error(err)
	}

	_, err = client.ExportUsers()

	if err != nil {
		t.Error(err)
	}

	// _, err = client.ImportArms()

	// if err != nil {
	// 	t.Error(err)
	// }

	// _, err = client.ImportDags()

	// if err != nil {
	// 	t.Error(err)
	// }

	// _, err = client.ImportEvents()

	// if err != nil {
	// 	t.Error(err)
	// }

	// _, err = client.ImportFile()

	// if err != nil {
	// 	t.Error(err)
	// }

	// _, err = client.ImportInstrumentEventMaps()

	// if err != nil {
	// 	t.Error(err)
	// }

	// _, err = client.ImportProject()

	// if err != nil {
	// 	t.Error(err)
	// }

	// _, err = client.ImportRecords()

	// if err != nil {
	// 	t.Error(err)
	// }

	// _, err = client.ImportUserDagMaps()

	// if err != nil {
	// 	t.Error(err)
	// }

	// _, err = client.ImportUserRoles()

	// if err != nil {
	// 	t.Error(err)
	// }

	// _, err = client.ImportUsers()

	// if err != nil {
	// 	t.Error(err)
	// }

	_, err = client.RenameRecord(
		"1",
		"some_field",
		"2",
	)

	if err != nil {
		t.Error(err)
	}

	_, err = client.SwitchDag("some_dag")

	if err != nil {
		t.Error(err)
	}
}
