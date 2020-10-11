package tpcGo

import (
	"encoding/xml"
	"io/ioutil"
)

type XMLTest struct {
	XMLName   xml.Name       `xml:"Test"`
	TestName  string         `xml:"TestName,attr"`
	TestCases []XMLTestCases `xml:"TestCase"`
}

type XMLTestCases struct {
	XMLName       xml.Name          `xml:"TestCase"`
	CaseID        string            `xml:"CaseID,attr"`
	NumOfWorker   int               `xml:"NumOfWorker,attr"`
	Votes         VoteSlice     `xml:"TestValues>Vote"`
	Decisions     DecisionSlice `xml:"TestOracles>Decision"`
	FinalDecision DecisionEnum  `xml:"TestOracles>FinalDecision"`
}

func ParseXMLTestCase(file string, xmlTestCaseType interface{}) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return xml.Unmarshal(b, &xmlTestCaseType)
}

