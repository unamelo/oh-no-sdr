 ██████╗ ██╗  ██╗    ███╗   ██╗ ██████╗     ███████╗██████╗ ██████╗ 
██╔═══██╗██║  ██║    ████╗  ██║██╔═══██╗    ██╔════╝██╔══██╗██╔══██╗
██║   ██║███████║    ██╔██╗ ██║██║   ██║    ███████╗██║  ██║██████╔╝
██║   ██║██╔══██║    ██║╚██╗██║██║   ██║    ╚════██║██║  ██║██╔══██╗
╚██████╔╝██║  ██║    ██║ ╚████║╚██████╔╝    ███████║██████╔╝██║  ██║
 ╚═════╝ ╚═╝  ╚═╝    ╚═╝  ╚═══╝ ╚═════╝     ╚══════╝╚═════╝ ╚═╝  ╚═╝

            o h - n o - s d r

---Purpose---
- Intended to convert legacy SDR (Single Return Data) text files to comma-separated CSV with headers.
- No more painful VLOOKUP and copy-paste operations.


---Project Snapshot---
- Language: Go
- TUI: Bubbletea https://github.com/charmbracelet/bubbletea

---How To Run---
1. Install Go 1.21+
`curl -sS https://webi.sh/golang | sh`
2. From the project root, download dependencies: `go mod tidy`.
3. Build the binary: `go build ./...`
4. Run the app: `go run ./...`

---Troubleshooting---
- If Go complains about missing modules, re-run `go mod tidy`.

