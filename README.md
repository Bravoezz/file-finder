# File Search CLI in Go

This CLI tool allows you to search for files by name within a directory and its subdirectories. It also provides an option to open the file in the Windows file explorer if found.

## Features

- Searches for a file by name recursively in all subdirectories.
- Excludes common directories like node_modules, dist, build, .git, and .idea.
- If a file is found, you can open it directly in the Windows file explorer.

## Prerequisites

- [Go](https://golang.org/doc/install) (1.17 or higher)

## How to Build

1. Clone the repository or download the Go file:

\`\`\`bash
git clone <repository_url>
cd <repository_folder>
\`\`\`

2. Build the Go project:

\`\`\`bash
go build -o searchfile.exe
\`\`\`

This will generate an executable named searchfile.exe in the current directory.

## How to Run

1. Run the executable from the terminal:

\`\`\`bash
./searchfile.exe -name=<filename>
\`\`\`

Replace <filename> with the name of the file you're searching for.

Example:

\`\`\`bash
./searchfile.exe -name=myFile.txt
\`\`\`

If the file is found, you will be prompted to open it in the file explorer.

## Install to Program Files (Windows)

To make the CLI accessible from any location in your Windows Command Prompt, follow these steps:

1. **Move the Executable to Program Files**:

   Copy the searchfile.exe to a folder inside C:\\Program Files\\.

   Example:

   \`\`\`bash
   move searchfile.exe "C:\\Program Files\\SearchFileCLI\\"
   \`\`\`

2. **Add to Path**:

   Add the folder C:\\Program Files\\SearchFileCLI\\ to your Windows environment variables' Path:

    - Right-click on **This PC** or **Computer** and select **Properties**.
    - Click on **Advanced system settings**.
    - Under the **System Properties** window, click the **Environment Variables** button.
    - Under **System variables**, scroll down and select the Path variable, then click **Edit**.
    - In the **Edit Environment Variable** dialog, click **New**, and paste the path C:\\Program Files\\SearchFileCLI\\.
    - Click **OK** to close all windows.

3. **Test**:

   Open a new Command Prompt and type:

   \`\`\`bash
   searchfile.exe -name=myFile.txt
   \`\`\`

   You should now be able to run the tool from any location.

## License

This project is licensed under the MIT License.

