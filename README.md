# ASCII Art Generator (Go)

A simple **Go CLI program** that converts text into **ASCII Art** using a predefined template font.

The program reads a text input from the command line and prints the ASCII-art representation of that text in the terminal.

---

## Features

- Convert text to ASCII Art
- Supports multiple lines using `\n`
- Automatically downloads the ASCII template if it doesn't exist locally
- Works with all printable ASCII characters (`32–126`)
- Handles empty lines correctly

---

## Project Structure

```
.
├── main.go
├── standard.txt
└── README.md
```

- **main.go** → Main application logic  
- **standard.txt** → ASCII art font template  
- **README.md** → Project documentation  

---

## How It Works

1. The program accepts **one command line argument**.
2. It checks if the ASCII font template (`standard.txt`) exists locally.
3. If the file does not exist, it downloads it automatically from:

```
https://acad.learn2earn.ng/api/content/root/01-edu_module/content/ascii-art/standard.txt
```

4. The template is parsed into a structured format.
5. Each character from the input text is mapped to its ASCII art representation.
6. The result is printed to the terminal.

---


## Usage

Run the program with:

```bash
go run . "Hello"
```

Example Output:

```
  _    _      _ _       
 | |  | |    | | |      
 | |__| | ___| | | ___  
 |  __  |/ _ \ | |/ _ \ 
 | |  | |  __/ | | (_) |
 |_|  |_|\___|_|_|\___/
 ```

 ---

 ## Multi-line Input

 You can add line breaks using `\n`.

 Example:

 ```bash
 go run . "Hello\nWorld"
 ```

 ---

 ## Example

 Input:

 ```bash
 go run . "Go"
 ```

 Output:

 ```
     ____      
    / ___| ___ 
    | |  _ / _ \
    | |_| | (_) |
     \____|\___/
     ```

     ---

     ## Error Handling

     The program checks for:

     - Missing arguments
     - File read/write errors
     - HTTP download errors
     - Invalid template structure

     If the program is run incorrectly, it will show:

     ```
     Usage: go run . <text>
     ```

     ---

     ## Requirements

     - Go **1.18+**
     - Internet connection (only required for first run if `standard.txt` is missing)

     ---

     ## Author

     Built with Go for learning **ASCII art rendering and CLI tools**.