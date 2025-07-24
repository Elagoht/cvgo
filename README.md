# CVGO

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![HTML](https://img.shields.io/badge/HTML-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![JSON](https://img.shields.io/badge/JSON-black?style=for-the-badge&logo=json&logoColor=white)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/Elagoht/cvgo/release.yaml?style=for-the-badge)
[![GitHub Stars](https://img.shields.io/github/stars/Elagoht/cvgo.svg?style=for-the-badge)](https://github.com/Elagoht/cvgo/stargazers)
![GitHub License](https://img.shields.io/github/license/Elagoht/cvgo?style=for-the-badge)

A command-line tool written in Go that generates a professional CV (Curriculum Vitae) from a JSON data file and an HTML template. This project is designed to be easily customizable, allowing users to create their own CVs by simply modifying the `data.json` file and `template.html`.

## ✨ Features

* **Flexible Data Input:** Uses a JSON file to define all your CV content (personal info, experience, education, skills, projects, etc.).
* **Customizable Template:** Renders your CV using a standard Go HTML template, giving you full control over the layout and styling.
* **Multiple Output Modes:**
  * **Default:** Generates `cv.html` in the current directory.
  * **Output to File:** Specify a custom output filename using the `-o` or `--output` flag.
  * **Watch Mode:** Starts a local web server (`:8080`) that automatically re-renders your CV whenever you save changes to your `data.json` or `template.html` files, perfect for live development.
* **Configurable Paths:** Easily specify custom paths for your JSON data and HTML template files using the `-d`/`--data` and `-t`/`--template` flags.
* **Clean CLI:** Built with `urfave/cli` for a user-friendly and robust command-line interface.

## 🚀 Getting Started

To get started, simply download the appropriate executable for your device from the [Releases page](https://github.com/Elagoht/cvgo/releases). Once downloaded, unarchive the file (if it's a `.zip` or `.tar.gz`) and place your `data.json` and `template.html` files in the same directory as the executable. Double-click the executable to generate your `cv.html`!

## 🛠️ Usage

Run the `main.go` file with various flags to control its behavior.

### For Developers (Command-Line Usage)

If you have Go installed and prefer using the command line, you can run the tool directly:

* **Default Output:** Generates `cv.html` in the current directory using `data.json` and `template.html`.
* **Output to a Specific File:** Use the `-o` or `--output` flag to specify a different output filename.
* **Watch Mode (Development Server):** Use the `-w` or `--watch` flag to start a local web server on `:8080`. This server will automatically re-render and serve your CV whenever you make changes to `data.json` or `template.html`.
    Open your browser to `http://localhost:8080` and refresh the page after saving changes to your data or template files.
* **Custom Data and Template Paths:** Use the `-d` or `--data` flag for your JSON file and `-t` or `--template` flag for your HTML template file. These can be combined with any other mode.
* **Help Information:** To see all available flags and their descriptions:

```bash
go run main.go --help
```

### For End-Users (Executable Usage - No Command Line Needed!)

You don't need any command-line knowledge or Go installation to use this tool! Executable binaries for various operating systems will be available in the [Releases](https://github.com/Elagoht/cvgo/releases) section of this GitHub repository.

Here's how to generate your CV:

1. **Download the Executable:** Go to the [Releases page](https://github.com/Elagoht/cvgo/releases) and download the executable file for your operating system (e.g., `cvgo-windows-amd64.exe` for Windows, `cvgo-linux-amd64` for Linux, `cvgo-darwin-arm64` for macOS).
2. **Place Data and Template:** Ensure your `data.json` and `template.html` files are in the **same directory** as the downloaded executable.
   * **Need `data.json`?** You can use the anonymized example provided in this README. Fill it with your own information. If you need help generating the content, you can even ask an LLM (Large Language Model) to help you structure your CV data in JSON format!
3. **Run the Executable:**
   * **Windows:** Double-click the `cvgo-windows-amd64.exe` file.
   * **macOS/Linux:** You might need to give it executable permissions first: `chmod +x cvgo-darwin-amd64` (or `cvgo-linux-amd64`), then double-click or run the executable from your terminal.
4. **CV Generated!** A file named `cv.html` will be rendered in the same directory.

### 🖨️ Generating a PDF CV

Once `cv.html` is generated, you can easily convert it to a professional PDF document:

1. **Open `cv.html`:** Double-click `cv.html` to open it in your preferred web browser (Chrome, Firefox, Edge, Safari, etc.).
2. **Print to PDF:**
   * Press `Ctrl + P` (Windows/Linux) or `Cmd + P` (macOS) to open the print dialog.
   * In the printer selection, choose "Save as PDF" or "Microsoft Print to PDF" (or a similar option provided by your OS/browser).
   * **Crucially, make sure the "Print backgrounds" or "Background graphics" option is checked.** This ensures that any colors, images, or styling defined in your `template.html` (or its CSS) are included in the PDF.
   * Click "Save" or "Print" to save your CV as a PDF file.

Your professional CV in PDF format is now ready!

## ⚙️ Data Structure (`data.json`)

The `data.json` file is the heart of your CV content. It's a simple JSON object that maps directly to the fields you'll access in your HTML template. The provided anonymized example shows the expected structure.

**Key points:**

* Arrays (e.g., `titles`, `links`, `experience`) are used for repeatable sections.

* Objects within arrays (e.g., items in `experience`, `projects`) allow for structured data.

* The field names in your JSON (e.g., `name`, `email`, `experience`, `title`, `company`) must match the names you use in your Go template.

## 🎨 Template Structure (`template.html`)

The `template.html` file uses [Go's `html/template` package](https://pkg.go.dev/html/template). This allows you to embed Go template syntax directly into your HTML to dynamically render data from your `data.json`.

**Basic Template Syntax:**

* `{{.FieldName}}`: Access a field directly from the top-level JSON object (e.g., `{{.name}}`, `{{.email}}`).

* `{{range .ArrayName}} ... {{end}}`: Iterate over an array. Inside the `range` block, `.` refers to the current item in the array (e.g., `{{.title}}` when ranging over `links`).

* `{{if .Condition}} ... {{end}}`: Conditional rendering.

Feel free to add any CSS, or external libraries to your `template.html` to achieve the desired look and feel for your CV.

## 🤝 Contributing

Contributions are welcome! If you have suggestions for improvements, bug fixes, or new features, please open an issue or submit a pull request.

## 📄 License

This project is licensed under the GNU GPLv3 - see the [LICENSE](/LICENSE) file for details.
