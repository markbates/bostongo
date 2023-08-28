# Building Better CLIs

<figure id="mark.png">
<img src="assets/mark.png">
<figcaption>[@markbates](https://twitter.com/markbates)</figcaption>
</figure>

---

# Gopher Guides

<figure id="guides.png">
<img src="assets/guides.png">
<figcaption>[www.gopherguides.com](https://www.gopherguides.com)</figcaption>
</figure>

---

# Go Fundamentals

<img src="assets/book.png">

---

# Building Better CLIs

- Avoiding Globals
- Cleaner Code
- Easily Testable
- Composable CLI Apps
- Framework/3rd Party Free!

---

# The Library

<go doc="github.com/markbates/bostongo.Walker"></go>

---

# Avoid Globals!

- Command Line Arguments
- Current Working Directory
- Environment Variables
- File System
- I/O

---

# The `main` Function

<code src="cmd/walker/main.go#main"></code>

---

# Information Gathered

- Context
- Current Working Directory
- Command Line Arguments

---

# The Imports

<code src="cmd/walker/main.go#imports"></code>

---

# Directory Tree

<cmd exec='tree cmd/walker -I testdata' hide-cmd></cmd>

---

# The `cli.App` Type

<go doc="./cmd/walker/cli.App"></go>

---

# Standard I/O

<go doc="os.Stdout"></go>

---

# The `IO` Type

<go doc="github.com/markbates/iox.IO"></go>

---

# The `Stdout` Method

<go sym="github.com/markbates/iox.IO.Stdout"></go>

---

# The `cli.App` Type

<go doc="./cmd/walker/cli.App"></go>

---

# The `fs.FS` Interface

<go doc="io/fs.FS"></go>

---

# The `cli.App#Main` Function

<go doc="./cmd/walker/cli.App.Main"></go>

---

# The `Commander` Interface

<go doc="./cmd/bostongo/cli.Commander"></go>

---

# Web Application

<go doc="github.com/markbates/bostongo/web.App"></go>

---

# The ServeHTTP Function

<code src="web/app.go#serve"></code>

---

# Web Server's Main Function

<code src="cmd/server/main.go#main"></code>

---

# The Server `cli.App` Type

<go doc="./cmd/server/cli.App"></go>

---

# The `Env` Type

<code src="cmd/server/cli/env.go#type"></code>
<code src="cmd/server/cli/env.go#getenv"></code>

---

# Inside the Main Function

<code src="cmd/server/cli/main.go#main"></code>

---

# The SettableIO Interface

<go doc="./cmd/bostongo/cli.SettableIO"></go>

---

# Garlic

<go doc="github.com/markbates/garlic.Garlic"></go>

---

# Garlic

<go sym="github.com/markbates/garlic.Garlic.Main"></go>

# Final Folder Structure

<cmd exec="tree -I testdata -I assets"></cmd>
