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

// banner for revshell
func banner(){
	banner := figure.NewFigure("GO REV SHELL", "", true)
	color.Set(color.Bold)
	banner.Print()
	color.Unset()
		fmt.Println("")
		fmt.Println(cyan("\t -- A polygot payload generator -- "))
	fmt.Println("")
}