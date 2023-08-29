# A Simple CLI

## The `main` Function

<figure id="cmd/walker/main.go#main">

<code src="cmd/walker/main.go#main"></code>

<figcaption>[cmd/walker/main.go](./cmd/walker/main.go)</figcaption>

</figure>

## Information Gathered

- Command Line Arguments
- Current Working Directory
- Context

## The Imports

<figure id="imports">

<code src="cmd/walker/main.go#imports"></code>

<figcaption>[cmd/walker/main.go](./cmd/walker/main.go)</figcaption>

</figure>

## Directory Tree

<figure id="walker.tree">

<cmd exec='tree cmd/walker -I testdata'></cmd>

<figcaption>Directory structure of the `cmd/walker` command.</figcaption>

</figure>

## The `cli.App` Type

<figure id="cmd/walker/cli.App">

<go doc="./cmd/walker/cli.App"></go>

<figcaption>[cmd/walker/cli/app.go](./cmd/walker/cli/app.go)</figcaption>

</figure>

## Standard I/O

<figure id="os.Stdout">

<go doc="os.Stdout"></go>

<figcaption><godoc>os#Stdout</godoc></figcaption>

</figure>

### The `IO` Type

<figure id="iox.IO">

<go doc="github.com/markbates/iox.IO"></go>

<figcaption><godoc>github.com/markbates/iox#IO</godoc></figcaption>

</figure>

#### The `Stdout` Method

<figure id="iox.IO.Stdout">

<go sym="github.com/markbates/iox.IO.Stdout"></go>

<figcaption><godoc>github.com/markbates/iox#IO.Stdout</godoc></figcaption>

</figure>

### In Testing

<figure id="in.testing">

<code src="cmd/bostongo/cli/garlic_test.go#io"></code>

<figcaption>[cmd/bostongo/cli/garlic_test.go](./cmd/bostongo/cli/garlic_test.go)</figcaption>

</figure>

## The `cli.App` Type

<figure id="cmd/walker/cli.App">

<go doc="./cmd/walker/cli.App"></go>

<figcaption>[cmd/walker/cli/app.go](./cmd/walker/cli/app.go)</figcaption>

</figure>

## The `fs.FS` Interface

<figure id="io/fs.FS">

<go doc="io/fs.FS"></go>

<figcaption><godoc>io/fs#FS</godoc></figcaption>

</figure>

## The `Commander` Interface

<figure id="commander">

<go sym="./cmd/bostongo/cli.Commander"></go>

<figcaption>[cmd/bostongo/cli/ifaces.go](./cmd/bostongo/cli/ifaces.go)</figcaption>

</figure>

### Default File System

<figure id="os.DirFS">

<go doc="os.DirFS"></go>

<figcaption><godoc>os#DirFS</godoc></figcaption>

</figure>

## The `cli.App#Main` Method

<figure id="cmd/walker/cli.App.Main">

<go sym="./cmd/walker/cli.App.Main"></go>

<figcaption>[cmd/walker/cli/app.go](./cmd/walker/cli/app.go)</figcaption>

</figure>

## The `cli.App#flags` Method

<figure id="cmd/walker/cli.App.flags">

<go sym="./cmd/walker/cli.App.flags"></go>

<figcaption>[cmd/walker/cli/app.go](./cmd/walker/cli/app.go)</figcaption>

</figure>

## Running the `walker` Command

<figure id="help">

<go run="cmd/walker/main.go -h" exit="-1"></go>

<figcaption>The <code>walker</code> command help output.</figcaption>

</figure>

<figure id="running">

<go run="cmd/walker/main.go -dirs testdata"></go>

<figcaption>Running the <code>walker</code> command.</figcaption>

</figure>

## Testing the `walker` Command

<figure id="cmd/walker/cli.App.Test">

<code src="cmd/walker/cli/app_test.go#test"></code>

<figcaption>[cmd/walker/cli/app_test.go](./cmd/walker/cli/app_test.go)</figcaption>

</figure>

<figure id="test-v">

<go src="cmd/walker/cli" test="-v"></go>

<figcaption>Running the <code>walker</code> command tests.</figcaption>

</figure>

## Globals Avoided!

- ~~Command Line Arguments~~
- ~~Current Working Directory~~
- ~~I/O~~
- ~~File System~~
- Environment Variables
