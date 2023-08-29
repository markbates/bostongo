# Combining Commands

## Directory Tree

<figure id="bostongo.tree">

<cmd exec='tree cmd/bostongo -I testdata'></cmd>

<figcaption>Directory structure of the `cmd/bostongo` command.</figcaption>

</figure>

## The `cli.App` Type

<figure id="cmd/bostongo/cli.App">

<go doc="./cmd/bostongo/cli.App"></go>

<figcaption>[cmd/bostongo/cli/app.go](./cmd/bostongo/cli/app.go)</figcaption>

</figure>

## The `cli.Commands` Type

<figure id="cmd/bostongo/cli.Commands">

<go doc="./cmd/bostongo/cli.Commands"></go>

<figcaption>[cmd/bostongo/cli/app.go](./cmd/bostongo/cli/app.go)</figcaption>

</figure>

### The `cli.Commands#Find` Method

<figure id="cmd/bostongo/cli.Commands.Find">

<go sym="./cmd/bostongo/cli.Commands.Find"></go>

<figcaption>[cmd/bostongo/cli/app.go](./cmd/bostongo/cli/app.go)</figcaption>

</figure>

### The `Commander` Interface

<figure id="commander">

<go sym="./cmd/bostongo/cli.Commander"></go>

<figcaption>[cmd/bostongo/cli/ifaces.go](./cmd/bostongo/cli/ifaces.go)</figcaption>

</figure>

## The `cli.App#Main` Method

<figure id="cmd/bostongo/cli.App.Main">

<go sym="./cmd/bostongo/cli.App.Main"></go>

<figcaption>[cmd/bostongo/cli/app.go](./cmd/bostongo/cli/app.go)</figcaption>

</figure>

## The SettableIO Interface

<figure id="cmd/bostongo/cli.SettableIO">

<go sym="./cmd/bostongo/cli.SettableIO"></go>

<figcaption>[cmd/bostongo/cli/ifaces.go](./cmd/bostongo/cli/ifaces.go)</figcaption>

</figure>

## Running the `bostongo` Command

<figure id="running">

<go run="cmd/bostongo/main.go walker -dirs testdata"></go>

<figcaption>Running the <code>bostongo</code> command.</figcaption>

</figure>

## Testing the `bostongo` Command

<figure id="cmd/bostongo/cli.App.Test">

<code src="cmd/bostongo/cli/app_test.go#test"></code>

<figcaption>[cmd/bostongo/cli/app_test.go](./cmd/bostongo/cli/app_test.go)</figcaption>

</figure>

<figure id="test-v">

<go src="cmd/bostongo/cli" test="-v -run App"></go>

<figcaption>Running the <code>bostongo</code> command tests.</figcaption>

</figure>
