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
    - One time setup for rocket chat access, domain, username and password is required

- `rchat update-status --status online|busy|away|offline {{status message}}`
    - This will update the status message and the status of the user
    - If status flag is optional one.
    - Sending empty message is supported for status is supported, it will reset the status message
