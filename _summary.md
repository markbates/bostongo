# Summing Up

## Avoid Globals

### I/O

<figure id="iox.IO">

<go doc="github.com/markbates/iox.IO"></go>

<figcaption><godoc>github.com/markbates/iox#IO</godoc></figcaption>

</figure>

<figure id="cmd/bostongo/cli.SettableIO">

<go sym="./cmd/bostongo/cli.SettableIO"></go>

<figcaption>[cmd/bostongo/cli/ifaces.go](./cmd/bostongo/cli/ifaces.go)</figcaption>

</figure>

### Environment Variables

<figure id="env.doc">

<go sym="./cmd/server/cli.Env"></go>

<figcaption>[cmd/server/cli/env.go](./cmd/server/cli/env.go)</figcaption>

</figure>

### The File System

<figure id="io/fs.FS">

<go doc="io/fs.FS"></go>

<figcaption><godoc>io/fs#FS</godoc></figcaption>

</figure>

<figure id="garlic.SettableFS">

<go sym="github.com/markbates/garlic.SettableFS"></go>

<figcaption><godoc>github.com/markbates/garlic#SettableFS</godoc></figcaption>

</figure>

### Current Working Directory and Arguments

<figure id="commander">

<go sym="./cmd/bostongo/cli.Commander"></go>

<figcaption>[cmd/bostongo/cli/ifaces.go](./cmd/bostongo/cli/ifaces.go)</figcaption>

</figure>

## Escape the `main` Package

<figure id="cmd/walker/main.go#main">

<code src="cmd/walker/main.go#main"></code>

<figcaption>[cmd/walker/main.go](./cmd/walker/main.go)</figcaption>

</figure>

## Consider the Garlic Pattern

<figure id="main">

<code src="cmd/bostongo/main.go#main"></code>

<figcaption>[cmd/bostongo/main.go](./cmd/bostongo/main.go)</figcaption>

</figure>
