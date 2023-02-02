# pk-cli
* create-react-app build & deploy command
* set slack status

# System configuration
Mac, *Nix

# Reference
* https://www.thorsten-hans.com/lets-build-a-cli-in-go-with-cobra/

```bash
go build -o ./dist/pk -ldflags="-X 'github.com/phanithken/pk/cmd/pk.version=0.0.2'" main.go
```
