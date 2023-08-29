# The `web` Library

<figure id="web.App">

<go doc="github.com/markbates/bostongo/web.App"></go>

<figcaption>[cmd/server/cli/app.go](./cmd/server/cli/app.go)</figcaption>

</figure>

### The `ServeHTTP`` Method

<figure id="web.App.ServeHTTP">

<go sym="./web.App.ServeHTTP"></go>

<figcaption>[web/app.go](./web/app.go)</figcaption>

</figure>

### The `walk` Method

<figure id="web.App.walk">

<go sym="./web.App.walk"></go>

<figcaption>[web/app.go](./web/app.go)</figcaption>

</figure>

# The Server CLI

## Directory Tree

<figure id="walker.tree">

<cmd exec='tree cmd/server -I testdata'></cmd>

<figcaption>Directory structure of the `cmd/server` command.</figcaption>

</figure>

## The `main` Function

<figure id="main">

<code src="cmd/server/main.go#main"></code>

<figcaption>[cmd/server/main.go](./cmd/server/main.go)</figcaption>

</figure>

## The Server `cli.App` Type

<figure id="cmd/server/cli.App">

<go doc="./cmd/server/cli.App"></go>

<figcaption>[cmd/server/cli/app.go](./cmd/server/cli/app.go)</figcaption>

</figure>

## The `Env` Type

<figure id="env.doc">

<go sym="./cmd/server/cli.Env"></go>

<figcaption>[cmd/server/cli/env.go](./cmd/server/cli/env.go)</figcaption>

</figure>

### The `Getenv` Method

<figure id="env.getenv">

<go sym="./cmd/server/cli.Env.Getenv"></go>

<figcaption>[cmd/server/cli/env.go](./cmd/server/cli/env.go)</figcaption>

</figure>

### Using the `Env` Type

<figure id="env.use">

<code src="cmd/server/cli/app.go#port"></code>

<figcaption>[cmd/server/cli/app.go](./cmd/server/cli/app.go)</figcaption>

</figure>

## The `cli.App#Main` Function

<figure id="cmd/server/cli.App.Main">

<go sym="./cmd/server/cli.App.Main"></go>

<figcaption>[cmd/server/cli/app.go](./cmd/server/cli/app.go)</figcaption>

</figure>

## Testing the `server` Command

<figure id="test-v">

<go src="./cmd/server/cli" test="-v"></go>

<figcaption>Running the <code>server</code> command tests.</figcaption>

</figure>

## Globals Avoided!

- ~~Command Line Arguments~~
- ~~Current Working Directory~~
- ~~I/O~~
- ~~File System~~
- ~~Environment Variables~~
