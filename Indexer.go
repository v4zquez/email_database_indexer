package main

import (
	"os/exec"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	//"flag"
)

var ParentDirectory string

var i int = 1
var f, s string

var Message_ID string
var Date string
var From string
var To string
var Subject string
var Mime_Version string
var Content_Type string
var Charset string
var Content_Transfer_Encoding string
var X_From string
var X_To string
var X_cc string
var X_bcc string
var X_Folder string
var X_Origin string
var X_FileName string
var Content string

type Header struct {
	Header                    string `json:"_index"`
}

type Index struct {
	Index                     Header
}

type Maildir struct {
	Message_ID                string `json:"Message-ID"`
	Date                      string `json:"Date"`
	From                      string `json:"From"`
	To                        string `json:"To"`
	Subject                   string `json:"Suject"`
	Mime_Version              string `json:"Mime-Version"`
	Content_Type              string `json:"Content-Type"`
	Charset                   string `json:"Charset"`
	Content_Transfer_Encoding string `json:"Content-Transfer-Encoding"`
	X_From                    string `json:"X-From"`
	X_To                      string `json:"X-To"`
	X_cc                      string `json:"X-cc"`
	X_bcc                     string `json:"X-bcc"`
	X_Folder                  string `json:"X-Folder"`
	X_Origin                  string `json:"X-Origin"`
	X_FileName                string `json:"X-FileName"`
	Content                   string `json:"Content"`
}

func FormatString(originalString string) (firstElement string, secondElement string) {
	vector := strings.Split(originalString, ":")

	firstElement = vector[0]

		if len(vector) > 1 {
			for i := 1; i < len(vector); i++ {
				if i == 1 {
					secondElement = strings.Replace(vector[i], " ", "", 1)
				} else {
					secondElement += ":"
					secondElement += vector[i]
				}
			}
		} else {
			secondElement = vector[1]
		}

		return
	}

	func AssignValue(i int, j string) {
		switch i {
		case 1:
			Message_ID = j
		case 2:
			Date = j
		case 3:
			From = j
		case 4:
			To = j
		case 5:
			Subject = j
		case 6:
			Mime_Version = j
		case 7:
			Content_Type = j
		case 8:
			Content_Transfer_Encoding = j
		case 9:
			X_From = j
		case 10:
			X_To = j
		case 11:
			X_cc = j
		case 12:
			X_bcc = j
		case 13:
			X_Folder = j
		case 14:
			X_Origin = j
		case 15:
			X_FileName = j
		case 16:
			Content = j
		}
	}

/*func LocateDirectory() {
	ParentDirectory := flag.String("ParentDirectory", "", "")
        
	out, err := exec.Command("")  

	flag.Parse()
	fmt.Println("PD:", *ParentDirectory)
}*/

func GetDirectories() string {
    cmd, err := exec.Command("/bin/bash", "-c", ". ./DirectoryList.sh").Output()
    if err != nil {
        fmt.Printf("error %s", err)
    }
    output := string(cmd[:])
    return output
}

/*func ObtainResources(resource string, directory string) string {
    script := ". ./Controller.sh"
    separateArguments := " "
    command := script + separateArguments + resource + directory
    cmd, err := exec.Command("/bin/bash", "-c", command).Output()
    if err != nil {
        fmt.Printf("error %s", err)
    }
    output := string(cmd[:])
    return output
}*/

func GetFiles(directory string) string {
    script := ". ./FilesList.sh"
    separateArguments := " "
    command := script + separateArguments + directory
    //fmt.Printf(command)
    cmd, err := exec.Command("/bin/bash", "-c", command).Output()
    if err != nil {
        fmt.Printf("error %s", err)
    }
    output := string(cmd[:])
    return output
}

	func main() {
		/*data := ObtainResources("GetDirectories", "")
		directories := strings.Split(data, " ")
		fmt.Println(directories)

		dir := directories[10]
		data1 := ObtainResources("GetFiles", dir)
		files := strings.Split(data1, " ")
		fmt.Println(files)*/

		data := GetDirectories()
		directories := strings.Split(data, " ")
		fmt.Println(directories)
		//fmt.Println(directories[0])
	
		dir := directories[10]
		data1 := GetFiles(dir)
		files := strings.Split(data1, " ")
		fmt.Println(files)

		file, err := os.Open("1.")

		//handle errors while opening
		if err != nil {
			log.Fatalf("Error when opening file: %s", err)
		}

		fileScanner := bufio.NewScanner(file)

		// read line by line
		for fileScanner.Scan() {
			if fileScanner.Text() != "" {

				if i < 16 {
					f, s = FormatString(fileScanner.Text())
					AssignValue(i, s)
				}

				if i >= 16 {
					if fileScanner.Text() != "" {
						Content += "\n"
						Content += fileScanner.Text()
					}
				}

				i = i + 1
			}
		}

		AssignValue(16, Content)

		header := Header {
			Header:  "LATIERRAESPLANA",
		}

		index := Index {
			Index:  header,
		}

		maildir := &Maildir{
			Message_ID:                Message_ID,
			Date:                      Date,
			From:                      From,
			To:                        To,
			Subject:                   Subject,
			Mime_Version:              Mime_Version,
			Content_Type:              Content_Type,
			Content_Transfer_Encoding: Content_Transfer_Encoding,
			X_From:                    X_From,
			X_To:                      X_To,
			X_cc:                      X_cc,
			X_bcc:                     X_bcc,
			X_Folder:                  X_Folder,
			X_Origin:                  X_Origin,
			X_FileName:                X_FileName,
			Content:                   Content,
		}

		h, err := json.Marshal(index)
		m, err := json.Marshal(maildir)

		fmt.Println(string(h))
		fmt.Println(string(m))
	}

