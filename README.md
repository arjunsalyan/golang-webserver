## INTRODUCTION

This is a basic project based which was implemented as a part of learning some introductory concepts of Golang.

*The project might evolve over time as more and more concepts are learned and attempts to implement those concepts in this project are made. The definition of the Project may eveolve over time which might not just be limited to a web server.*

### Updates as on 27th April, 2020

A basic webserver which listens to the following requests:
- `/` - Static homepage, html template is rendered.
- `/read-file/{{ filepath }}` - Reads the file "filepath.txt" and displays the contents, the title of the web page is name of the time.
- `/view/?title="{{title}}&body={{body}}` - The title of the page and the body is read from get variabes.
- `404` - Not found page, also called when a view is not found.

The server handles the following erros gracefully:
- Page not found
- Invalid file path (/read-file/)
- Invalid Parameters (/view/)
