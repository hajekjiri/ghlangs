# Ghlangs
## About
Ghlangs is a command-line tool utilizing the [GitHub GraphQL API v4](https://developer.github.com/v4/) to tell you what programming languages are dominant in your GitHub repositories. You can either view a short summary for each repository or display the total amount of code written in each language across all of your repositories.

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
You can run `ghlangs` without any parameters. In that case, it will default to
```
ghlangs -format total -sort-by size -sort-order desc -unit auto
```
You can also replace `-key VALUE` with `-key=VALUE`. Both of these syntaxes are acceptable.
```
ghlangs -format=total -sort-by=size -sort-order=desc -unit=auto
```

Run `ghlangs -h` or `ghlangs -help` to display a short description of the parameters.
```
$ ghlangs -help
Usage: ghlangs [-format FORMAT] [-sort-by KEY] [-sort-order ORDER]
  -format string
    	(detail|total) display format (default "total")
  -h	show help (shorthand)
  -help
    	show help
  -sort-by string
    	(name|size) sort key for sorting languages (default "size")
  -sort-order string
    	(asc|desc) sort order for sorting languages (default "desc")
  -unit string
    	(auto|B|kB|MB|GB|TB|PB|EB) unit used for displaying sizes (default "auto")
```

#### Examples
Display details for each repository, sort languages by name in ascending order and use bytes for displaying sizes. I skipped some repositories in this output to keep it short.
```
$ ghlangs -format=detail -sort-by=name -sort-order=asc -unit=B
Progress: 8/8 repositories (API Rate Limit 2/5000)
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
|Total size|26424.000  B|100.00%|
---------------------------------
|Go        |26424.000  B|100.00%|
---------------------------------

hajekjiri/pacman
---------------------------------
|Total size|82726.000  B|100.00%|
---------------------------------
|C++       |81632.000  B| 98.68%|
|Makefile  | 1094.000  B|  1.32%|
---------------------------------

...
```

```
# equivalent to running ghlangs without parameters
$ ghlangs -format=total -sort-by=size -sort-order=desc -unit=auto
Progress: 8/8 repositories (API Rate Limit 3/5000)
All repositories:
-------------------------------
|Total size|441.402 kB|100.00%|
-------------------------------
|JavaScript|321.632 kB| 72.87%|
|C++       | 79.719 kB| 18.06%|
|Go        | 25.805 kB|  5.85%|
|HTML      |  6.951 kB|  1.57%|
|Python    |  3.686 kB|  0.83%|
|CSS       |  2.515 kB|  0.57%|
|Makefile  |  1.096 kB|  0.25%|
-------------------------------
```

## Used external modules
* [shurcooL/githubv4](https://github.com/shurcooL/githubv4) is licensed under the [MIT license](https://github.com/shurcooL/githubv4/blob/master/LICENSE)
* [golang/oauth2](https://github.com/golang/oauth2) is licensed under the [BSD 3-Clause license](https://github.com/golang/oauth2/blob/master/LICENSE)
