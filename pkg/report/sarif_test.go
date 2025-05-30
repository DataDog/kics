package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Checkmarx/kics/pkg/model"
	reportModel "github.com/Checkmarx/kics/pkg/report/model"
	"github.com/Checkmarx/kics/test"
	"github.com/stretchr/testify/require"
)

type reportTestCase struct {
	caseTest       jsonCaseTest
	expectedResult model.Summary
}

type sarifReport struct {
	basePath     string                 `json:"-"`
	Schema       string                 `json:"$schema"`
	SarifVersion string                 `json:"version"`
	Runs         []reportModel.SarifRun `json:"runs"`
}

var sarifTests = []reportTestCase{
	{
		caseTest: jsonCaseTest{
			summary:  test.SummaryMock,
			path:     "./testdir",
			filename: "testout",
		},
		expectedResult: test.SummaryMock,
	},
	{
		caseTest: jsonCaseTest{
			summary:  test.SummaryMockCritical,
			path:     "./testdir",
			filename: "testout2",
		},
		expectedResult: test.SummaryMockCritical,
	},
}

// TestPrintSarifReport tests the functions [PrintSarifReport()] and all the methods called by them
func TestPrintSarifReport(t *testing.T) {
	for idx, test := range sarifTests {
		t.Run(fmt.Sprintf("Sarif File test case %d", idx), func(t *testing.T) {
			if err := os.MkdirAll(test.caseTest.path, os.ModePerm); err != nil {
				t.Fatal(err)
			}
			err := PrintSarifReport(test.caseTest.path, test.caseTest.filename, test.caseTest.summary, model.SCIInfo{}, false)
			checkFileExists(t, err, &test, "sarif")
			jsonResult, err := os.ReadFile(filepath.Join(test.caseTest.path, test.caseTest.filename+".sarif"))
			require.NoError(t, err)
			var resultSarif sarifReport
			err = json.Unmarshal(jsonResult, &resultSarif)
			require.NoError(t, err)
			require.Equal(t, "", resultSarif.basePath)
			require.Equal(t, "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/master/Schemata/sarif-schema-2.1.0.json", resultSarif.Schema)
			require.Equal(t, "2.1.0", resultSarif.SarifVersion)
			require.Len(t, resultSarif.Runs, len(test.expectedResult.Queries))
			os.RemoveAll(test.caseTest.path)
		})
	}
}

func TestPrintSarifReportWithToolVersionFullScan(t *testing.T) {
	for idx, test := range sarifTests {
		t.Run(fmt.Sprintf("Sarif File test case %d", idx), func(t *testing.T) {
			if err := os.MkdirAll(test.caseTest.path, os.ModePerm); err != nil {
				t.Fatal(err)
			}
			err := PrintSarifReport(test.caseTest.path, test.caseTest.filename, test.caseTest.summary, model.SCIInfo{RunType: "full_scan"}, false)
			checkFileExists(t, err, &test, "sarif")
			jsonResult, err := os.ReadFile(filepath.Join(test.caseTest.path, test.caseTest.filename+".sarif"))
			require.NoError(t, err)
			var resultSarif sarifReport
			err = json.Unmarshal(jsonResult, &resultSarif)
			require.NoError(t, err)
			require.Len(t, resultSarif.Runs, len(test.expectedResult.Queries))
			require.Equal(t, resultSarif.Runs[0].Tool.Driver.ToolVersion, "full_scan")
			os.RemoveAll(test.caseTest.path)
		})
	}
}

func TestPrintSarifReportWithToolVersionCodeUpdate(t *testing.T) {
	for idx, test := range sarifTests {
		t.Run(fmt.Sprintf("Sarif File test case %d", idx), func(t *testing.T) {
			if err := os.MkdirAll(test.caseTest.path, os.ModePerm); err != nil {
				t.Fatal(err)
			}
			err := PrintSarifReport(test.caseTest.path, test.caseTest.filename, test.caseTest.summary, model.SCIInfo{RunType: "code_update"}, false)
			checkFileExists(t, err, &test, "sarif")
			jsonResult, err := os.ReadFile(filepath.Join(test.caseTest.path, test.caseTest.filename+".sarif"))
			require.NoError(t, err)
			var resultSarif sarifReport
			err = json.Unmarshal(jsonResult, &resultSarif)
			require.NoError(t, err)
			require.Len(t, resultSarif.Runs, len(test.expectedResult.Queries))
			require.Equal(t, resultSarif.Runs[0].Tool.Driver.ToolVersion, "code_update")
			os.RemoveAll(test.caseTest.path)
		})
	}
}

func checkFileExists(t *testing.T, err error, tc *reportTestCase, extension string) {
	require.NoError(t, err)
	require.FileExists(t, filepath.Join(tc.caseTest.path, tc.caseTest.filename+fmt.Sprintf(".%s", extension)))
}
