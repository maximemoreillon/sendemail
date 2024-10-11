# moreillon/sendemail

A Docker container to send emails. Basically Nodemailer containerized

## Usage

```bash
docker run \
  -e SMTP_PORT=465 \
  -e SMTP_HOST=your.smtp.server \
  -e SMTP_USERNAME=username \
  -e SMTP_PASSWORD=password \
  -e SMTP_TO=john@example.com \
  -e SMTP_FROM=no_reply@example.com \
  -e SMTP_SUBJECT="Container test" \
  -e SMTP_MESSAGE="Hello from nodemailer" \
  moreillon/sendemail
```
