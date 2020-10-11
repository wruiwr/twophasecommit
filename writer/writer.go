package main

import (
	"encoding/xml"
	"fmt"
	tpc "github.com/selabhvl/tpcGo"
	"os"
	"path/filepath"
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
	Votes         tpc.VoteSlice     `xml:"TestValues>Votes"`
	Decisions     tpc.DecisionSlice `xml:"TestOracles>Decisions"`
	FinalDecision tpc.DecisionEnum  `xml:"TestOracles>FinalDecision"`
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

// xmlWriter can write test cases for ReadQF, WriteQF or System tests into xml files.
func xmlWriter(dir string) {

	var err error
	var output []byte

	test := xmlTestStruct()
	output, err = xml.MarshalIndent(test, " ", "  ")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// Show xml format in the Terminal
	os.Stdout.Write(output)

	// Start to write to the xml file
	xmlOutput := []byte(string(output))

	testFilePath, err := filepath.Abs(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(testFilePath)
	}

	f, err := os.Create(testFilePath)
	checkErr(err)

	defer f.Close()

	_, err = f.Write(xmlOutput)
	checkErr(err)

}

func xmlTestStruct() *XMLTest {

	result := &XMLTest{
		TestName: "TPCTest",
		TestCases: []XMLTestCases{
			{
				CaseID:      "1",
				NumOfWorker: 3,
				Votes: tpc.VoteSlice{
					{
						WorkerID:  tpc.WorkerID(1),
						VoteValue: tpc.Yes,
					},
					{
						WorkerID:  tpc.WorkerID(2),
						VoteValue: tpc.Yes,
					},
					{
						WorkerID:  tpc.WorkerID(3),
						VoteValue: tpc.Yes,
					},
				},
				Decisions: tpc.DecisionSlice{
					{
						WorkerID:      tpc.WorkerID(1),
						DecisionValue: tpc.Commit,
					},
					{
						WorkerID:      tpc.WorkerID(2),
						DecisionValue: tpc.Commit,
					},
					{
						WorkerID:      tpc.WorkerID(3),
						DecisionValue: tpc.Commit,
					},
				},
				FinalDecision: tpc.Commit,
			},
			{
				CaseID:      "2",
				NumOfWorker: 3,
				Votes: tpc.VoteSlice{
					{
						WorkerID:  tpc.WorkerID(1),
						VoteValue: tpc.Yes,
					},
					{
						WorkerID:  tpc.WorkerID(2),
						VoteValue: tpc.No,
					},
					{
						WorkerID:  tpc.WorkerID(3),
						VoteValue: tpc.Yes,
					},
				},
				Decisions: tpc.DecisionSlice{
					{
						WorkerID:      tpc.WorkerID(1),
						DecisionValue: tpc.Commit,
					},
					{
						WorkerID:      tpc.WorkerID(2),
						DecisionValue: tpc.Abort,
					},
					{
						WorkerID:      tpc.WorkerID(3),
						DecisionValue: tpc.Commit,
					},
				},
				FinalDecision: tpc.Abort,
			},
			{
				CaseID:      "3",
				NumOfWorker: 3,
				Votes: tpc.VoteSlice{
					{
						WorkerID:  tpc.WorkerID(1),
						VoteValue: tpc.Yes,
					},
					{
						WorkerID:  tpc.WorkerID(2),
						VoteValue: tpc.No,
					},
					{
						WorkerID:  tpc.WorkerID(3),
						VoteValue: tpc.No,
					},
				},
				Decisions: tpc.DecisionSlice{
					{
						WorkerID:      tpc.WorkerID(1),
						DecisionValue: tpc.Commit,
					},
					{
						WorkerID:      tpc.WorkerID(2),
						DecisionValue: tpc.Abort,
					},
					{
						WorkerID:      tpc.WorkerID(3),
						DecisionValue: tpc.Abort,
					},
				},
				FinalDecision: tpc.Abort,
			},
			{
				CaseID:      "4",
				NumOfWorker: 3,
				Votes: tpc.VoteSlice{
					{
						WorkerID:  tpc.WorkerID(1),
						VoteValue: tpc.No,
					},
					{
						WorkerID:  tpc.WorkerID(2),
						VoteValue: tpc.No,
					},
					{
						WorkerID:  tpc.WorkerID(3),
						VoteValue: tpc.No,
					},
				},
				Decisions: tpc.DecisionSlice{
					{
						WorkerID:      tpc.WorkerID(1),
						DecisionValue: tpc.Abort,
					},
					{
						WorkerID:      tpc.WorkerID(2),
						DecisionValue: tpc.Abort,
					},
					{
						WorkerID:      tpc.WorkerID(3),
						DecisionValue: tpc.Abort,
					},
				},
				FinalDecision: tpc.Abort,
			},
		},
	}

	return result
}

func main() {
	xmlWriter("./xml/tests.xml")
}
