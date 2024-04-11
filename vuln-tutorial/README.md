Open terminal and cd to this folder

Init module: `go mod init vuln.tutorial`

Download all dependencies: `go mod tidy`

Downgrade the version of golang.org/x/text to v0.3.5, which contains known vulnerabilities: `go get golang.org/x/text@v0.3.5`

Install govulncheck with the go install command: `go install golang.org/x/vuln/cmd/govulncheck@latest`

From the folder you want to analyze (in this case, vuln-tutorial). Run: `govulncheck ./...`

You should see this output:

```
Scanning your code and 50 packages across 2 dependent modules for known vulnerabilities...

=== Symbol Results ===

Vulnerability #1: GO-2021-0113
    Out-of-bounds read in golang.org/x/text/language
  More info: https://pkg.go.dev/vuln/GO-2021-0113
  Module: golang.org/x/text
    Found in: golang.org/x/text@v0.3.5
    Fixed in: golang.org/x/text@v0.3.7
    Example traces found:
      #1: main.go:12:29: vuln.main calls language.Parse

Your code is affected by 1 vulnerability from 1 module.
This scan also found 1 vulnerability in packages you import and 1 vulnerability
in modules you require, but your code doesn't appear to call these
vulnerabilities.
Use '-show verbose' for more details.
```

Upgrade golang.org/x/text to v0.3.8: `go get golang.org/x/text@v0.3.8`

Now run govulncheck again: `govulncheck ./...`

You will now see this output:

```
Scanning your code and 50 packages across 2 dependent modules for known vulnerabilities...

=== Symbol Results ===

No vulnerabilities found.

Your code is affected by 0 vulnerabilities.
This scan also found 0 vulnerabilities in packages you import and 1
vulnerability in modules you require, but your code doesn't appear to call these
vulnerabilities.
Use '-show verbose' for more details.
```