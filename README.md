## Introduction

This GO-application can send mails with attachments to addressed people.

## Table of contents

- Requirements
- Usage
- Configuration

## Requirements

This module requires the following modules:

- ["gomail"](https://gopkg.in/gomail.v2)
- ["encoding/json"](https://pkg.go.dev/encoding/json)
- ["flag"](https://pkg.go.dev/flag)
- ["fmt"](https://pkg.go.dev/fmt)
- ["os"](https://pkg.go.dev/os)

## Usage



``` text
Usage:
  -filePath string
        Give the full system path to the attachment.
  -toAddress string
        Give the e-mail address of the recipient.
```

NOTE: BOTH PARAMETERS NEED TO BE GIVEN!

## Configuration

These is the contents from conf.json file.
These configurations will be used to send the e-mail.

```json
{
    "smtpServer": "SMTP-SERVER",
    "port": PORT,
    "mailAddress": "YOUR-MAIL-ADDRESS",
    "mailPassword": "YOUR-PASSWORD"
}
```
NOTE: BE SURE TO HAVE A MAIL ACCOUNT THAT ALLOWS 3TH-PARTY CLIENTS OR UNSAVE APPLICATIONS!
