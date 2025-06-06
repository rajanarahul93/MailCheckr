#  MailCheckr

MailCheckr is a simple Go-based CLI tool that checks a domain's email configuration by verifying the presence of **MX**, **SPF**, and **DMARC** DNS records.

##  Features

- Checks if a domain has MX records (used for receiving emails)
-  Extracts and verifies SPF records (used for sender policy framework)
-  Looks for DMARC records (used to prevent email spoofing)
-  Outputs a clean summary of each domain's email DNS status

##  Installation

Make sure you have [Go installed](https://golang.org/dl/) on your machine.

```bash
git clone https://github.com/rajanarahul93/MailCheckr.git
cd MailCheckr
go build -o mailcheckr
```

##  Usage

```bash
./mailcheckr
```

You'll be prompted to enter domain names one by one. Press `Ctrl+C` to exit.

###  Example

![image](https://github.com/user-attachments/assets/620c8634-41dd-40f0-b660-e34ecdadfe6f)



##  Author

Made with ❤️ by [Vara Rahul Rajana](https://github.com/rajanarahul93)
