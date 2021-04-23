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

