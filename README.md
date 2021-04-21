# Rocket Chat CLI
Cli to interact with rocket chat APIs

### Requirement
 - Golang v1.16+

### Installation
Install the cli through following command

```
go get github.com/samit22/rchat

go install github.com/samit22/rchat

rchat -h
```


### Commands

- `rchat init`
    - One time setup for rocket chat access, domain, username and password is required.
    - Username can be either email of rchat user name.
    - Your user name and password is not saved, the token generated from the credentials is encoded and saved in local file, which you can delete by `rm ~/.rchat-config`

- `rchat update-status --status online|busy|away|offline {{status message}}`
    - This will update the status message and the status of the user.
    - The status flag is optional one, only status message is updated if the flag is not sent.
    - Sending empty message is supported for status message, it will reset the status message.
