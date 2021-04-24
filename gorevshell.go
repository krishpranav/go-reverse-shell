package main
import (
	"github.com/akamensky/argparse"
	"os"
	"fmt"
    "github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
   	"github.com/chzyer/readline"
    "github.com/common-nighthawk/go-figure"
	"net"
	"io"
	"strings"
	"encoding/base64"
	"strconv"
	"bufio"
	"reflect"
	"math/rand"
	"github.com/gobuffalo/packr"
	"time"
)

// function for showing the banner
func banner(){
	banner := figure.NewFigure("GO REV SHELL", "", true)
	color.Set(color.Bold)
	banner.Print()
	color.Unset()
		fmt.Println("")
		fmt.Println(cyan("\t -- A polygot payload generator -- "))
	fmt.Println("")
}

// function for showing the payload list
func list(){
	actions_data := [][]string{
        []string{"reverse_shell", "Spawn a reverse shell"},
        []string{"cmd_exec", "Execute a command"},
		[]string{"forkbomb", "Run a forkbomb"},
		[]string{"memexec", "Embed and execute a binary"},
		[]string{"download_exec", "Download and execute a file"},
		[]string{"shutdown", "Shutdown computer"},
        []string{"custom", "Use custom Bash and Powershell scripts"},
    }
	actions_table := tablewriter.NewWriter(os.Stdout)
	actions_table.SetAutoWrapText(false)
	actions_table.SetHeader([]string{"NAME", "DESCRIPTION"})
    actions_table.SetColumnColor(
        tablewriter.Colors{tablewriter.FgGreenColor},
        tablewriter.Colors{}, 
    )
	for v := range actions_data {
		actions_table.Append(actions_data[v])
	}
	fmt.Println("")
	fmt.Println("[*] Payloads: ")
    actions_table.Render()
    fmt.Println("")
}

// colors
var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var cyan = color.New(color.FgBlue).SprintFunc()
var bold = color.New(color.Bold).SprintFunc()

// print message
func print(msg string){
	fmt.Printf("%s %s", green("[+]"), msg)
}

// info
func print_info(msg string){
	fmt.Println("[+]", msg)
}

// print error
func print_error(msg string){
	fmt.Printf("%s %s", red("[x]"), msg)
}

// print header
func print_header(message string) {
	color.Set(colors.Bold)
	fmt.Printf("-- %s --", message)
	color.Unset()
	fmt.Println("")
}

func contains(s interface{}, elem interface{}) bool {
    arrV := reflect.ValueOf(s)
	if arrV.Kind() == reflect.Slice {
        for i := 0; i < arrV.Len(); i++ {
            if arrV.Index(i).Interface() == elem {
                return true
            }
        }
    }
	return false
}

// convert payload from str to int
func str_to_int(string_integer string) int {
	i, _ := strconv.Atoi(string_integer)
	return i
}

// gap interval
func interval_to_seconds(interval string) int64{ 
    period_letter := string(interval[len(interval)-1])
    intr := string(interval[:len(interval)-1]) //Check this
    i, _ := strconv.ParseInt(intr, 10, 64) 
    switch period_letter{
        case "s":
            return i
        case "m":
            return i*60
        case "h":
            return i*3600
    }
    return i
}

// input function
func input(name string, message string, default_value string) string{
    if default_value == ""{
        default_value = "none"
    }
    final_prompt := fmt.Sprintf("%s %s (default: %s): ", red(name), message, default_value)
    p, _ := readline.NewEx(&readline.Config{
        Prompt:              final_prompt,
        InterruptPrompt:     "^C",
    })
    line, _ := p.Readline()
    if (len(line) == 0 || contains([]string{"y", "yes"}, line)){
        return default_value
    } else {
        return line
    }
}

// write file function
func write_to_file(filename string, data string) error {
	file, err := os.Create(filename)
	exit_on_error("[FILE CREATION ERROR]", err)
	defer file.Close()

	_, err = io.WriteString(file, data)
	exit_on_error("[FILE WRITE ERROR]", err)
	return file.Sync()
}

// read file
func read_file(filename string) string {
	contents := ""
	file, err := os.Open(filename)
	exit_on_error("{FILE READ ERROR}")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents += scanner.Text()
	}
	return contents
}

// exit connection on error
func exit_on_error(message string, err error){
	if err != nil{
		fmt.Printf("%s %v", red(message+":"),err)
		os.Exit(0)
	}
}

// decode base64 files
func base64_decode(str string) string {
	raw, _ := base64.StdEncoding.DecodeString(str)
	return fmt.Sprintf("%s", raw)
}

// encode base64 files
func base64_encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// function for templates
func get_template(template_name string) string{
	template, err := packr.NewBox("./").FindString(template_name)
	exit_on_error("[PACKR ERROR]", err)
	return template
}

