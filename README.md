# Sendemail

A Go application meant to be run as a Docker container to send emails

## Usage example

```
  docker run \
  -e SMTP_PORT=465 \
  -e SMTP_HOST=smtp.example.com\
  -e SMTP_USERNAME=<username> \
  -e SMTP_PASSWORD=<password> \
  -e SMTP_TO=john@example.com \
  -e SMTP_FROM=no_reply@example.com \
  -e SMTP_SUBJECT="Email from Go" \
  -e SMTP_MESSAGE="Hello from Go" \
  moreillon/sendemail
```
