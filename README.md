# Ghlangs
## About
Ghlangs is a command-line tool utilizing the [GitHub GraphQL API v4](https://developer.github.com/v4/) to tell you what programming languages are dominant in any user's or oganization's GitHub repositories. You can either view a short summary for each repository or display the total amount of code written in each language across all repositories. Only counts repositories that the target user/organization owns and that are not forks.

## Getting started
### 1. GitHub API token
Github GraphQL API v4 requires authentication so you'll need to [generate a personal access token](https://github.com/settings/tokens/new) to use with this tool. No scopes are required but if you want to see data from your private repositories, you'll have to select the _repo (Full control of private repositories)_ scope. Save your token to an environment variable called `GITHUB_AUTH_TOKEN`.

### 2. Install ghlangs
#### Using the go tool
Simply install the project with `go get`
```
go get github.com/hajekjiri/ghlangs
```

#### Building from source
Clone the repository
```
git clone https://github.com/hajekjiri/ghlangs
```
and then build the project
```
go build
```

### 3. Running ghlangs
Run `ghlangs -h` or `ghlangs -help` to display a short description of the parameters.
```
$ ghlangs -help
Usage: ghlangs [-user USER] [-org ORGANIZATION] [-format FORMAT] [-sort-by KEY] [-sort-order ORDER] [-unit UNIT]
  -format string
    	(detail|total) display format (default "total")
  -h	show help (shorthand)
  -help
    	show help
  -org string
    	login of the organization whose repositories you want to query, cannot combine with "-user"
  -sort-by string
    	(name|size) sort key for sorting languages (default "size")
  -sort-order string
    	(asc|desc) sort order for sorting languages (default "desc")
  -unit string
    	(auto|B|kB|MB|GB|TB|PB|EB) unit used for displaying sizes (default "auto")
  -user string
    	login of the user whose repositories you want to query, cannot combine with "-org"
```

Use `-key VALUE` or `-key=VALUE` to pass arguments. Both of these syntaxes are acceptable.
```
# these are equivalent
ghlangs -format total -sort-by size -sort-order desc -unit auto
ghlangs -format=total -sort-by=size -sort-order=desc -unit=auto
```

Running `ghlangs` without any parameters is possible. In that case, it will default to
```
ghlangs -format=total -sort-by=size -sort-order=desc -unit=auto
```

#### Targets
You can run `ghlangs` on 3 kinds of targets, no more than 1 target at a time.

##### Viewer
Displays viewer's repositories. Viewer is the owner of the provided API token. Counts private repositories if the provided API token has access to the _repo_ scope. This is the default target.
```
# Display data from repos owned by the API token owner
ghlangs
```

##### User
You can run `ghlangs` on any GitHub user by passing their login with the `-user` parameter. Counts private repositories if the viewer has access to them and the token has access to the _repo_ scope.
```
ghlangs -user=USER
```

##### Organization
You can run `ghlangs` on any GitHub organization by passing its login with the `-org` parameter. Counts private repositories if the viewer has access to them and the token has access to the _repo_ scope.
```
ghlangs -org=ORGANIZATION
```

#### Examples
Display details for each repository owned by viewer, sort languages by name in ascending order, and use bytes for displaying sizes.
```
$ ghlangs -format=detail -sort-by=name -sort-order=asc -unit=B
Progress: 8/8 repositories (API Rate Limit   84/5000)
hajekjiri/duckshooter-vr
---------------------------------
|Total size|45253.000  B|100.00%|
---------------------------------
|CSS       | 1288.000  B|  2.85%|
|HTML      |  369.000  B|  0.82%|
|JavaScript|43596.000  B| 96.34%|
---------------------------------

hajekjiri/ghlangs
---------------------------------
|Total size|27279.000  B|100.00%|
---------------------------------
|Go        |27279.000  B|100.00%|
---------------------------------

hajekjiri/hajekjiri.github.io
--------------------------------
|Total size|1970.000  B|100.00%|
--------------------------------
|CSS       |1287.000  B| 65.33%|
|HTML      | 683.000  B| 34.67%|
--------------------------------

hajekjiri/lenovo-deals-web
----------------------------------
|Total size|234046.000  B|100.00%|
----------------------------------
|HTML      |  6066.000  B|  2.59%|
|JavaScript|227980.000  B| 97.41%|
----------------------------------

hajekjiri/lenovo-outlet-scraper
--------------------------------
|Total size|4821.000  B|100.00%|
--------------------------------
|JavaScript|4821.000  B|100.00%|
--------------------------------

hajekjiri/neanderthals-vr
---------------------------------
|Total size|52954.000  B|100.00%|
---------------------------------
|JavaScript|52954.000  B|100.00%|
---------------------------------

hajekjiri/pacman
---------------------------------
|Total size|82726.000  B|100.00%|
---------------------------------
|C++       |81632.000  B| 98.68%|
|Makefile  | 1094.000  B|  1.32%|
---------------------------------
```

Display the total size of code for each language across all repositories owned by user "torvalds", sort languages by size in descending order, and automatically select the most suitable unit for each size entry.
```
# equivalent to
# ghlangs -user=torvalds

$ ghlangs -user=torvalds -format=total -sort-by=size -sort-order=desc -unit=auto
Progress: 4/4 repositories (API Rate Limit   85/5000)
All repositories:
---------------------------------
|Total size  |852.072 MB|100.00%|
---------------------------------
|C           |822.009 MB| 96.47%|
|C++         | 11.236 MB|  1.32%|
|Assembly    |  9.461 MB|  1.11%|
|Objective-C |  2.258 MB|  0.27%|
|Shell       |  2.237 MB|  0.26%|
|Makefile    |  2.110 MB|  0.25%|
|Perl        |  1.138 MB|  0.13%|
|Python      |  1.104 MB|  0.13%|
|Roff        |132.269 kB|  0.02%|
|SmPL        |126.460 kB|  0.01%|
|Yacc        |119.794 kB|  0.01%|
|Lex         | 62.887 kB|  0.01%|
|Awk         | 43.290 kB|  0.01%|
|UnrealScript| 17.732 kB|  0.00%|
|Gherkin     |  8.328 kB|  0.00%|
|M4          |  3.325 kB|  0.00%|
|Clojure     |  1.450 kB|  0.00%|
|XS          |  1.239 kB|  0.00%|
|Raku        |  1.176 kB|  0.00%|
|Batchfile   |944.000  B|  0.00%|
|sed         |379.000  B|  0.00%|
---------------------------------
```

Display the total size of code for each language across all repositories owned by organization "google", sort languages by name in ascending order, and automatically select the most suitable unit for each size entry.
```
# equivalent to
# ghlangs -org=google -sort-by=name -sort-order=asc

$ ghlangs -org=google -format=total -sort-by=name -sort-order=asc -unit=auto
Progress:  100/1727 repositories (API Rate Limit   86/5000)
Progress:  200/1727 repositories (API Rate Limit   87/5000)
Progress:  300/1727 repositories (API Rate Limit   88/5000)
Progress:  400/1727 repositories (API Rate Limit   89/5000)
Progress:  500/1727 repositories (API Rate Limit   90/5000)
Progress:  600/1727 repositories (API Rate Limit   91/5000)
Progress:  700/1727 repositories (API Rate Limit   92/5000)
Progress:  800/1727 repositories (API Rate Limit   93/5000)
Progress:  900/1727 repositories (API Rate Limit   94/5000)
Progress: 1000/1727 repositories (API Rate Limit   95/5000)
Progress: 1100/1727 repositories (API Rate Limit   96/5000)
Progress: 1200/1727 repositories (API Rate Limit   97/5000)
Progress: 1300/1727 repositories (API Rate Limit   98/5000)
Progress: 1400/1727 repositories (API Rate Limit   99/5000)
Progress: 1500/1727 repositories (API Rate Limit  100/5000)
Progress: 1600/1727 repositories (API Rate Limit  101/5000)
Progress: 1700/1727 repositories (API Rate Limit  102/5000)
Progress: 1727/1727 repositories (API Rate Limit  103/5000)
All repositories:
---------------------------------------------
|Total size              |  4.806 GB|100.00%|
---------------------------------------------
|ActionScript            | 36.347 kB|  0.00%|
|AMPL                    | 33.503 kB|  0.00%|
|AngelScript             | 14.827 kB|  0.00%|
|ANTLR                   | 56.628 kB|  0.00%|
|ApacheConf              |917.000  B|  0.00%|
|AppleScript             | 14.189 kB|  0.00%|
|Arduino                 |  4.918 kB|  0.00%|
|Assembly                | 95.917 MB|  2.00%|
|AutoIt                  |917.000  B|  0.00%|
|Awk                     |285.892 kB|  0.01%|
|Batchfile               |490.721 kB|  0.01%|
|Bison                   |237.699 kB|  0.00%|
|Brainfuck               |  2.192 kB|  0.00%|
|C                       |  2.419 GB| 50.33%|
|C#                      | 23.525 MB|  0.49%|
|C++                     |820.737 MB| 17.08%|
|Cap'n Proto             |  2.726 kB|  0.00%|
|Clojure                 | 23.103 kB|  0.00%|
|CMake                   |  6.343 MB|  0.13%|
|CoffeeScript            | 42.835 kB|  0.00%|
|Common Lisp             |322.853 kB|  0.01%|
|Cool                    | 33.164 kB|  0.00%|
|Coq                     | 16.167 kB|  0.00%|
|CSS                     | 10.663 MB|  0.22%|
|Cuda                    |  1.706 MB|  0.04%|
|Dart                    | 14.815 MB|  0.31%|
|DIGITAL Command Language|336.715 kB|  0.01%|
|Dockerfile              |858.500 kB|  0.02%|
|DTrace                  | 21.370 kB|  0.00%|
|Eagle                   |463.972 kB|  0.01%|
|eC                      |  5.252 kB|  0.00%|
|Elixir                  | 50.689 kB|  0.00%|
|Elm                     | 28.269 kB|  0.00%|
|Emacs Lisp              |276.683 kB|  0.01%|
|Erlang                  |467.127 kB|  0.01%|
|F#                      | 15.029 kB|  0.00%|
|Filebench WML           |849.000  B|  0.00%|
|Forth                   |  3.722 kB|  0.00%|
|Fortran                 |784.023 kB|  0.02%|
|GAP                     |104.274 kB|  0.00%|
|GDB                     | 22.263 kB|  0.00%|
|Gherkin                 | 33.216 kB|  0.00%|
|GLSL                    |  1.041 MB|  0.02%|
|Gnuplot                 |  4.687 kB|  0.00%|
|Go                      |114.682 MB|  2.39%|
|Gosu                    | 30.192 kB|  0.00%|
|Groff                   |  7.264 MB|  0.15%|
|Groovy                  |185.486 kB|  0.00%|
|Hack                    | 21.241 kB|  0.00%|
|Handlebars              |  6.682 kB|  0.00%|
|Haskell                 |  2.879 MB|  0.06%|
|HCL                     |128.990 kB|  0.00%|
|HLSL                    | 41.820 kB|  0.00%|
|HTML                    |120.616 MB|  2.51%|
|Idris                   | 39.280 kB|  0.00%|
|Inno Setup              | 18.486 kB|  0.00%|
|Java                    |372.577 MB|  7.75%|
|JavaScript              |176.518 MB|  3.67%|
|Jsonnet                 |  1.157 MB|  0.02%|
|Julia                   | 49.676 kB|  0.00%|
|Jupyter Notebook        | 60.613 MB|  1.26%|
|KiCad Layout            |524.438 kB|  0.01%|
|Kotlin                  |  5.905 MB|  0.12%|
|Lasso                   |283.287 kB|  0.01%|
|Lex                     |438.739 kB|  0.01%|
|LiveScript              | 12.206 kB|  0.00%|
|LLVM                    |163.827 MB|  3.41%|
|Logos                   |  4.482 kB|  0.00%|
|Lua                     |828.067 kB|  0.02%|
|M                       |  6.903 kB|  0.00%|
|M4                      |733.360 kB|  0.02%|
|Makefile                | 11.841 MB|  0.25%|
|Mako                    |  6.997 kB|  0.00%|
|Mathematica             |  5.584 kB|  0.00%|
|MATLAB                  |932.368 kB|  0.02%|
|Max                     |  2.677 kB|  0.00%|
|Mercury                 |  1.193 kB|  0.00%|
|Meson                   |  8.566 kB|  0.00%|
|MLIR                    |  3.708 MB|  0.08%|
|NASL                    |121.092 kB|  0.00%|
|Nginx                   |  2.364 kB|  0.00%|
|Nix                     | 28.053 kB|  0.00%|
|NSIS                    |  6.612 kB|  0.00%|
|Objective-C             | 73.955 MB|  1.54%|
|Objective-C++           |  2.872 MB|  0.06%|
|OCaml                   |416.165 kB|  0.01%|
|OpenEdge ABL            |  4.118 MB|  0.09%|
|OpenSCAD                | 29.426 kB|  0.00%|
|Pascal                  |221.232 kB|  0.00%|
|Pawn                    | 29.935 kB|  0.00%|
|Perl                    | 11.416 MB|  0.24%|
|Perl 6                  | 32.711 kB|  0.00%|
|PHP                     | 11.642 MB|  0.24%|
|PLpgSQL                 |  1.243 MB|  0.03%|
|PLSQL                   | 33.456 kB|  0.00%|
|PostScript              | 15.560 kB|  0.00%|
|POV-Ray SDL             |  4.775 kB|  0.00%|
|PowerShell              |160.609 kB|  0.00%|
|Processing              | 68.747 kB|  0.00%|
|Prolog                  | 30.600 kB|  0.00%|
|Protocol Buffer         |  2.028 MB|  0.04%|
|Puppet                  | 22.667 kB|  0.00%|
|Pure Data               |964.000  B|  0.00%|
|PureBasic               |298.382 kB|  0.01%|
|Python                  |180.685 MB|  3.76%|
|QMake                   |  6.477 kB|  0.00%|
|R                       |  1.951 MB|  0.04%|
|Ragel                   |  8.013 kB|  0.00%|
|Ragel in Ruby Host      |128.173 kB|  0.00%|
|Raku                    | 18.923 kB|  0.00%|
|Red                     |188.000  B|  0.00%|
|RenderScript            |  2.493 kB|  0.00%|
|Rich Text Format        |  7.379 kB|  0.00%|
|Roff                    |888.716 kB|  0.02%|
|Ruby                    |830.530 kB|  0.02%|
|Rust                    |  7.465 MB|  0.16%|
|SaltStack               |  2.743 kB|  0.00%|
|Scala                   |322.292 kB|  0.01%|
|Scheme                  |201.294 kB|  0.00%|
|Scilab                  |173.650 kB|  0.00%|
|sed                     |  1.137 kB|  0.00%|
|ShaderLab               | 19.275 kB|  0.00%|
|Shell                   | 22.957 MB|  0.48%|
|Smali                   |  1.164 MB|  0.02%|
|Smalltalk               | 36.333 kB|  0.00%|
|Smarty                  | 74.539 kB|  0.00%|
|SmPL                    |351.044 kB|  0.01%|
|SourcePawn              |189.759 kB|  0.00%|
|SQLPL                   |132.428 kB|  0.00%|
|Standard ML             |  9.140 kB|  0.00%|
|Starlark                |  8.799 MB|  0.18%|
|SuperCollider           |  3.884 kB|  0.00%|
|Swift                   |  7.557 MB|  0.16%|
|SWIG                    |462.098 kB|  0.01%|
|SystemVerilog           |676.258 kB|  0.01%|
|Tcl                     |591.000  B|  0.00%|
|TeX                     |805.438 kB|  0.02%|
|Thrift                  |  7.875 kB|  0.00%|
|TSQL                    |632.546 kB|  0.01%|
|TypeScript              | 12.043 MB|  0.25%|
|UnrealScript            | 52.179 kB|  0.00%|
|VBA                     | 34.355 kB|  0.00%|
|VBScript                |  4.696 kB|  0.00%|
|Verilog                 |289.266 kB|  0.01%|
|VHDL                    |194.252 kB|  0.00%|
|Vim script              |516.923 kB|  0.01%|
|Visual Basic            | 33.849 kB|  0.00%|
|Vue                     |551.525 kB|  0.01%|
|XML                     | 17.504 kB|  0.00%|
|XS                      | 24.035 kB|  0.00%|
|XSLT                    |154.648 kB|  0.00%|
|Yacc                    |  2.229 MB|  0.05%|
---------------------------------------------
```

## Used external modules
* [shurcooL/graphql](https://github.com/shurcooL/graphql) is licensed under the [MIT license](https://github.com/shurcooL/graphql/blob/master/LICENSE)
