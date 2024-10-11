import nodemailer from "nodemailer"

const {
  SMTP_HOST,
  SMTP_PORT,
  SMTP_USERNAME,
  SMTP_PASSWORD,
  SMTP_FROM,
  SMTP_TO,
  SMTP_SUBJECT,
  SMTP_MESSAGE,
} = process.env

const transporterOptions: any = {
  host: SMTP_HOST,
  port: SMTP_PORT,
  secure: true, // upgrade later with STARTTLS
  auth: {
    user: SMTP_USERNAME,
    pass: SMTP_PASSWORD,
  },
}

const transporter = nodemailer.createTransport(transporterOptions)

const email = {
  from: SMTP_FROM,
  to: SMTP_TO,
  subject: SMTP_SUBJECT,
  text: SMTP_MESSAGE,
}

transporter
  .sendMail(email)
  .then(() => {
    console.log("Success")
  })
  .catch((error) => {
    console.error(error)
  })
