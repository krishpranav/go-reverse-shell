# go-reverse-shell
A polyglot payload generator

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

# Installation:
```
git clone https://github.com/krishpranav/go-reverse-shell
cd go-reverse-shell
sudo chmod +x *
bash install.sh
./gorevshell
```

# Example Payload Lists
``` golang
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
```

# Reference
<img src="">

- Author: krishpranav | https://github.com/krishpranav

