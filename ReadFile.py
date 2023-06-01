# -*- coding: utf-8 -*-

import os
import sys

def read_file(filename):
    headers = [
        "Message-ID:",
        "Date:",
        "From:",
        "To:",
        "Subject:",
        "Mime-Version:",
        "Content-Type:",
        "Content-Transfer-Encoding:",
        "X-From:",
        "X-To:",
        "X-cc:",
        "X-bcc:",
        "X-Folder:",
        "X-Origin:",
        "X-FileName:"
    ]

    file_content = []
    content_line = ""
    composite_line = ""
    header = ""
    body_email = False

    with open(filename) as file:
        for line in file:
            words=line.split()

            if header != "X-FileName:":
                if len(words) > 1:
                    composite_line = ""

                    # more than 2 element
                    if len(words) > 2:
                        header = str(words[0])
                        content_line = str(words[1:len(words)])

                        string_to_be_saved = ''
                        for w in words[1:len(words)]:
                            #string_to_be_saved += ' ' + repr(w)
                            string_to_be_saved += ' ' + w
                            string_to_be_saved = string_to_be_saved.replace("'", "") 

                        content_line = string_to_be_saved

                        #file_content.append(header+content_line+'\n')
                        file_content.append(header+(content_line+'\n'))
                    else:
                        string_to_be_saved = ' '
                        header = str(words[0])
                        content_line = string_to_be_saved + str(words[1])
                        #file_content.append(header+content_line+'\n')
                        file_content.append(header+(content_line+'\n'))

                    string_to_be_saved = ''
                    for w in words[1:len(words)]:
                        string_to_be_saved += ' ' + repr(w)
                        string_to_be_saved = string_to_be_saved.replace("'", "") 

                    composite_line += string_to_be_saved

                elif len(words) == 1:
                    if words[0] in headers:
                        composite_line = " "
                        string_to_be_saved = ''

                        for w in words[0]:
                            string_to_be_saved += repr(w)
                            string_to_be_saved = string_to_be_saved.replace("'", "") 

                        header = string_to_be_saved

                        file_content.append(header+'\n')
                    else:
                        string_to_be_saved = ' '

                        for w in words[0]:
                            string_to_be_saved += repr(w)
                            #string_to_be_saved = string_to_be_saved.replace("'", " ") 

                        string_to_be_saved = string_to_be_saved.replace("'", "") 
                        composite_line += string_to_be_saved

                        del file_content[-1]
                        #file_content.append(header+composite_line+'\n')
                        file_content.append(header+(composite_line+'\n'))
                        composite_line = " "

                if header == "X-FileName:":
                    body_email = True
            else:
                file_content.append(line)

    return file_content

if __name__ == '__main__':
    os.system("rm file_from_python.txt 2> /dev/null")

    content_to_be_saved = read_file(sys.argv[1])

    with open("file_from_python.txt", 'w') as file:
        for row in content_to_be_saved:
            file.write(str(row))
    file.close()
