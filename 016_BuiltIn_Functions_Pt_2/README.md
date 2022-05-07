# #15 BuiltIn Functions Part 2

## String

String functions are important to manipulate with string data. They can be accessed via the `strings` package using `strings.func(args)`. Here's a list:

| Function   | Syntax                  | Returns                                                      |
| ---------- | ----------------------- | ------------------------------------------------------------ |
| ToLower    | ToLower(s)              | Lowercase string of s                                        |
| ToUpper    | ToUpper(s)              | Uppercase string of s                                        |
| ToTitle    | ToTitle(s)              | Titlecase string of s                                        |
| Contains   | Contains(s, sub)        | Boolean based on whether sub in present in s                 |
| Count      | Count(s, sub)           | Count of sub in s                                            |
| Compare    | Compare(s1,s2)          | integer comparing two strings lexicographically              |
| HasPrefix  | HasPrefix(s, sub)       | Boolean based on whether s starts with sub                   |
| HasSuffix  | HasSuffix(s, sub)       | Boolean based on whether s ends with sub                     |
| Fields     | Fields(s)               | Slice of s by separating based on whitespace                 |
| Split      | Split(s,del)            | Slice of s by separating based on delimeter del              |
| Index      | Index(s, sub)           | Index of first occurence of sub in s                         |
| LastIndex  | LastIndex(s, sub)       | Index of last occurence of sub in s                          |
| Join       | Join(ss, del)           | String by joining slice ss with delimeter del                |
| Repeat     | Repeat(s,n)             | String s repeated n times                                    |
| Replace    | Replace(s, old, new, n) | String replacing n occurences of old with new in s           |
| ReplaceAll | ReplaceAll(s, old, new) | String replacing all occurences old with new in s            |
| Map        | Map(func(),s)           | String s after modifying each character of s by mapping func |

## Time

Time related functions are provided by the `time` package. You can access it using `time.func(args)`.

| Function | Syntax                   | Returns                                          |
| -------- | ------------------------ | ------------------------------------------------ |
| Now      | Now()                    | current timestamp                                |
| Format   | Format(layout)           | format time based on given layout                |
| Day      | Day()                    | day of the month                                 |
| Month    | Month()                  | month of the year                                |
| Year     | Year()                   | year of the time                                 |
| Date     | Date(y,m,d,h,m,s,ns,loc) | date in format yyyy-mm-dd hh:mm:ss+ns loc        |
| Sleep    | Sleep(t)                 | pauses current goroutine execution by duration t |

That's a wrap on day 16.