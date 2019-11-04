# getpassword
password generator

## Install
    $ git clone https://github.com/kursadbilgin/getpassword.git
    $ cd getpassword
    $ go build getpassword.go

## Run
    $ -/getpassword (Default)
    or
    $ -/getpassword -difficulty hard -length 9


| Flag        | Type   | Choice        | Default |
|-------------|--------|---------------|---------|
| -length     | int    | Max 25        | 7       |
| -difficulty | string | basic or hard | basic   |
