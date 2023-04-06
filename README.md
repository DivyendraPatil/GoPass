# GoPass

This is a CLI tool for generating random passwords with different levels of complexity. It allows users to specify the length of the password, the types of characters to include (uppercase letters, lowercase letters, numbers, and special characters), and the number of passwords to generate.

The tool also includes options for users to save the passwords to a file for future reference, check if the password has been breached before, and check the strength of the password.

## Usage

The tool is built using Go programming language, so you need to have Go installed on your machine to run the tool.

1. Clone the repository and navigate to the directory:
```
git clone https://github.com/divyendrapatil/password-generator.git
cd password-generator
```

2. Run the following command to build the tool:
```
go build .
```

3.
```
./gopass -length=16 -uppercase=true -lowercase=true -numbers=true -specials=true -num-passwords=10 -save-to-file=true -check-breaches=true -check-strength=true
```

The tool has the following options:
- length: the length of the password (default: 12)
- uppercase: include uppercase letters in the password (default: true)
- lowercase: include lowercase letters in the password (default: true)
- numbers: include numbers in the password (default: true)
- specials: include special characters in the password (default: true)
- num-passwords: the number of passwords to generate (default: 1)
- save-to-file: save the passwords to a file (default: false)
- check-breaches: check if the password has been breached before (default: false)
- check-strength: check the strength of the password (default: false)
  Note: You can use the shorthand version of the options, for example, - l instead of - length.

## Password Strength Meter

The tool also includes a password strength meter that analyzes the generated password and rates it based on its complexity. The strength meter uses a simple scoring system based on the following criteria:

- Length (1 point for every character)
- Uppercase letters (1 point for every uppercase letter)
- Lowercase letters (1 point for every lowercase letter)
- Numbers (1 point for every number)
- Special characters (2 points for every special character)

The password is rated on a scale from 0 to 10, where 0 is very weak and 10 is very strong.

## License

This code is released under the MIT license, which means you can use, copy, and modify it for any purpose, as long as you include the original license and copyright notice.
