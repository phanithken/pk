# pk-cli

I found myself lost everytime I try to google for commands when I work on something different. So I made this as a productivity tools for myself.
I scaffold my project by reading

# System configuration
Mac, *Nix

# Reference
* https://www.thorsten-hans.com/lets-build-a-cli-in-go-with-cobra/

```bash
go build -o ./dist/pk -ldflags="-X 'github.com/phanithken/pk/cmd/pk.version=0.0.2'" main.go
```
