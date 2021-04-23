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

}