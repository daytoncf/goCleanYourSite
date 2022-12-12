# goCleanYourSite
goCleanYourSite is a simple Go program that will parse through your HTML files, create a set of class names, parse through your css files to find all class names that aren't withing that set, and then create new cleaned css files `new_<filename>.css`

# Background
There are multiple solutions out that do this exact thing, and in some cases more, but many are locked behind paywalls so I decided to just build my own.
A similar goal could be achieved with a pretty small bash script, but I decided to use Go for the expandability and strong standard library.
Given my reasoning you might be thinking, "Why not use Python for this?", and that is a fair question. Ultimately, I went with Go over Python for 4 reasons:

1. **Performance.**  
Since Go has built in first-class concurrency support and Go programs compile to a binary, creating this program in Go provides a performane edge over Pyhton.

2. **Expandable.**  
Writing this in Go allows me to give this project a wider feature set than a shell script and will be able to handle more complex tasks. 

3. **Reusability.**  
Inside the `/css` folder of this project you'll find a set of structures and functions that are used for CSS parsing and tokenization. I'm trying to make this sub-package as portable as possible, eventually hoping to build a CSS Tokenizer that meets W3C standards and is on par with the HTML parser in Go's standard library. I plan on publishing this as it's own Go package for other developers to use once it is complete. You'll also find within the `pkg/lib` directory a set of more niche but reusable function and data structures, like stack and queue implementations for the `rune` type and other small useful functions that I felt were ambiguous to the overall project.

4. **Gopher.**  
Go has a silly little guy. Python does not have a silly little guy. I mean just look at him!
![Go's cute little gopher guy](https://go.dev/blog/gopher/header.jpg)

# Usage
1. Install Go. You can download it [here](https://go.dev/dl/).
2. Clone the repository.  
`git clone https://github.com/daytoncf/goCleanYourSite.git`
3. Run `go build`. This will generate an executable `goCleanYourSite` or `goCleanYourSite.exe`
4. Create a directory called `/content` in the same folder as the executable. 
5. Dump all of your html and css files into the `content` folder.
6. Run `./goCleanYourSite` in your terminal.
---
Your new, nice, and polished CSS files will be generated in the `/content` folder. They will have the same names as the old CSS files just prefixed with `new_`.
