## TASKS:
- [x] find the current shell
- [x] find the location of said shell
- [x] print the contents of the rcfiles

# Shell Configuration Reader

This Go program reads and prints the user's shell configuration file based on the shell they are using (`zsh`, `bash`, or `fish`).

## Features
- Automatically detects the current shell by checking the `SHELL` environment variable.
- Executes the `whoami` command to determine the current user.
- Reads and displays the configuration file (`.zshrc`, `.bashrc`, or `config.fish`) based on the detected shell.

## How It Works
1. **Detect the Current Shell:**
   The program fetches the user's shell by reading the `SHELL` environment variable.
   
2. **Determine the User:**
   The `whoami` command is executed to get the current user's username.

3. **Read Configuration File:**
   The program checks if the shell is:
   - **zsh**: Reads `/home/<username>/.zshrc`
   - **bash**: Reads `/home/<username>/.bashrc`
   - **fish**: Reads `/home/<username>/.config/fish/config.fish`

4. **Display the Content:**
   The content of the relevant configuration file is printed to the console.

## Usage
- Run the program in a terminal. It will automatically detect your shell and print the content of the corresponding configuration file.
