# Notifi

This is an example end point deisng to send notifications from a relay smtp server service or sendgrids excellent system

to be use you need to map it out to a volume

`./path/to/you/dir:/var`

and a file called `.env`

.env file for server relay:
```sh
EMAIL_HOST=<email_server>
EMAIL_HOST_USER=<email_user>
EMAIL_HOST_PASSWORD=<email_password>
EMAIL_PORT=<email_port>
SENDGRID_API_KEY=<send_grid_api_key>
```

.env file for sendgrid api:
```sh
SENDGRID_API_KEY=<send_grid_api_key>
```
