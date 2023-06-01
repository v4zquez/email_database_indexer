package main

import (
    "os/exec"
    "bufio"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "strings"
    "time"
)

var (
    exitNow string = "F"
    i int = 1
    fileDirectory int = 1
    fileTotal int = 1
    f, s string

    Content_Header string
    Message_ID string
    Date string
    From string
    To string
    Subject string
    Mime_Version string
    Content_Type string
    Charset string
    Content_Transfer_Encoding string
    X_From string
    X_To string
    X_cc string
    X_bcc string
    X_Folder string
    X_Origin string
    X_FileName string
    Content string
)

type Header struct {
    Header                    string `json:"_index"`
}

type Index struct {
    //Index                     Header `json:"index"`
    Index                     Header `"index"`
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

func DeleteFile() {
    e := os.Remove("data1.ndjson")
    if e != nil {
        //log.Fatal(e)
        fmt.Printf("%s\n", e)
    }
}

func UploadJSON(cmd string, url string, arg1 string, arg2 string, userPassword string, arg3 string, file string) {

    out, err := exec.Command(cmd, url, arg1, arg2, userPassword, arg3, file).Output()

    if err != nil {
        fmt.Printf("%s", err)
    }

    output := string(out[:])
    fmt.Println(output)
}

func FormatString(originalString string) (string, string) {
    output := strings.Split(originalString, " ")

    secondElement := ""

    for i:=1; i<len(output); i++ {
        if i < 2 || i == (len(output) - 1) {
            secondElement += output[i]
        } else {
            secondElement += output[i] + " "
        }
    }

    if len(output) > 1 {
        return output[0], secondElement
    } else {
        return output[0], ""
    }
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

func PrepareTextFile(file string) {
    script := "python3 ./ReadFile.py"
    separateArguments := " "
    command := script + separateArguments + file

    cmd, err := exec.Command("/bin/bash", "-c", command).Output()

    if err != nil {
        fmt.Printf("error %s", err)
    }

    output := string(cmd[:])

    fmt.Printf(output)
}

func GetDirectories() string {
    cmd, err := exec.Command("/bin/bash", "-c", ". ./DirectoryList.sh").Output()

    if err != nil {
        fmt.Printf("error %s", err)
    }

    output := string(cmd[:])

    return output
}

func GetFiles(directory string) string {
    script := ". ./FilesList.sh"
    separateArguments := " "
    command := script + separateArguments + directory

    cmd, err := exec.Command("/bin/bash", "-c", command).Output()

    if err != nil {
        fmt.Printf("error %s", err)
    }

    output := string(cmd[:])

    return output
}

func main() {
    DeleteFile()

    fmt.Println("Loading directories...")
    data := GetDirectories()
    directories := strings.Split(data, " ")

    fmt.Printf("Uploaded directories: %d\n", len(directories))
    time.Sleep(2 * time.Second)

    out:
    for k:=0; k<len(directories); k++ {
        dir := directories[k]
        fmt.Println("Current directory:", dir)

        data1 := GetFiles(dir)
        files := strings.Split(data1, " ")

        fmt.Println("Loading files...")
        fmt.Printf("Uploaded files: %d\n\n", len(files))
        time.Sleep(2 * time.Second)

        for l:=0; l<len(files); l++ {
            Content_Header := "enron_mail_20110402"
            //Content_Header := files[l]
            fmt.Println("File:", files[l])

            PrepareTextFile(files[l])

            file, err := os.Open("file_from_python.txt")

            //handle errors while opening
            if err != nil {
                log.Fatalf("Error when opening file: %s", err)
            }

            fileScanner := bufio.NewScanner(file)

            //read line by line
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

	            i += 1
                }
            }

            AssignValue(16, Content)

            header := Header {
                Header:  Content_Header,
            }

            index := Index {
                Index:  header,
            }

            maildir := Maildir {
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

            fl, err := os.OpenFile(("data1.ndjson"),
	    os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	    if err != nil {
	        log.Println(err)
	    }
	    defer fl.Close()

            fl.WriteString(string(h))
            fl.WriteString(string(m))

            fmt.Printf("Processed files from the current directory: %d of %d", fileDirectory, len(files))
            fmt.Printf("\nTotal files processed: %d\n\n", fileTotal)
            fmt.Println(string(h))

            if fileTotal == 2 {
               exitNow = "T"
               break
            }

            m = nil
            fileDirectory += 1
            fileTotal += 1
            i = 1
        }

        fileDirectory = 1

        if exitNow == "T" {
            break out
        }
    }

    fmt.Print("Uploading the JSON file to the Zinclabs database...\n")
    UploadJSON("curl", "http://localhost:4080/api/_bulk", "-i", "-u", "admin:Complexpass#123", "--data-binary", "@data1.ndjson")
    fmt.Print("Server Address: http://localhost:4080\n")
}
