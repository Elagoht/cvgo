package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

// Data is a generic type for unmarshaling JSON.
type Data any

const (
	defaultJSONPath     = "./data.json"
	defaultTemplatePath = "./template.html"
	defaultOutputPath   = "./cv.html"
)

// Reads a JSON file, unmarshals its content, loads an HTML template,
// and executes the template with the JSON data.
func renderTemplate(w io.Writer, jsonPath, templatePath string) error {
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return fmt.Errorf("error opening JSON file %s: %w", jsonPath, err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("error reading JSON file %s: %w", jsonPath, err)
	}

	var data Data
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON from %s: %w", jsonPath, err)
	}

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("error parsing template file %s: %w", templatePath, err)
	}

	return tmpl.Execute(w, data)
}

// renderToFile renders the template to a specified file.
func renderToFile(outputPath, jsonPath, templatePath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file %s: %w", outputPath, err)
	}
	defer file.Close()
	log.Printf("Rendering to %s...", outputPath)
	return renderTemplate(file, jsonPath, templatePath)
}

func main() {
	app := &cli.App{
		Name:  "cvgen",
		Usage: "Generate a CV from JSON data and a Go template",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "data",
				Aliases: []string{"d"},
				Value:   defaultJSONPath,
				Usage:   "Path to the JSON data file",
			},
			&cli.StringFlag{
				Name:    "template",
				Aliases: []string{"t"},
				Value:   defaultTemplatePath,
				Usage:   "Path to the HTML template file",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output filename. If not specified, 'cv.html' is used.",
			},
			&cli.BoolFlag{
				Name:    "watch",
				Aliases: []string{"w"},
				Usage:   "Start a server in watch mode to automatically re-render on file changes.",
			},
		},
		Action: func(context *cli.Context) error {
			jsonPath := context.String("data")
			templatePath := context.String("template")
			outputPath := context.String("output") // Get the value of the -o flag
			watchMode := context.Bool("watch")     // Check if the -w flag is set

			// Check for conflicts
			if outputPath != "" && watchMode {
				return fmt.Errorf(
					"cannot use --output (-o) and --watch (-w) flags simultaneously. Please choose one",
				)
			}

			// Handle watch mode
			if watchMode {
				log.Println("Starting server in watch mode on http://localhost:8080...")

				var lastJSONMod time.Time
				var lastTemplateMod time.Time

				http.HandleFunc("/", func(
					writer http.ResponseWriter,
					request *http.Request,
				) {
					jsonStat, err := os.Stat(jsonPath)
					if err != nil {
						http.Error(
							writer,
							fmt.Sprintf("Error accessing JSON file: %v", err),
							http.StatusInternalServerError,
						)
						log.Printf("Error stating JSON file: %v", err)
						return
					}

					templateStat, err := os.Stat(templatePath)
					if err != nil {
						http.Error(
							writer,
							fmt.Sprintf("Error accessing template file: %v", err),
							http.StatusInternalServerError,
						)
						log.Printf("Error stating template file: %v", err)
						return
					}

					if jsonStat.ModTime().After(lastJSONMod) ||
						templateStat.ModTime().After(lastTemplateMod) {
						log.Println("Detected file change. Re-rendering...")
						if err := renderTemplate(
							writer,
							jsonPath,
							templatePath,
						); err != nil {
							http.Error(
								writer,
								fmt.Sprintf("Error rendering template: %v", err),
								http.StatusInternalServerError,
							)
							log.Printf("Error rendering template in watch mode: %v", err)
							return
						}
						lastJSONMod = jsonStat.ModTime()
						lastTemplateMod = templateStat.ModTime()
					} else {
						// If no changes, re-render to ensure the page is always served.
						if err := renderTemplate(
							writer,
							jsonPath,
							templatePath,
						); err != nil {
							http.Error(
								writer,
								fmt.Sprintf("Error rendering template: %v", err),
								http.StatusInternalServerError,
							)
							log.Printf("Error rendering template in watch mode: %v", err)
							return
						}
					}
				})

				return http.ListenAndServe(":8080", nil)
			}

			// Handle output to file (either specified by -o or default)
			targetOutputPath := defaultOutputPath
			if outputPath != "" {
				targetOutputPath = outputPath
			}

			return renderToFile(targetOutputPath, jsonPath, templatePath)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
