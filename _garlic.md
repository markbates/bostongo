# The Garlic Pattern

## The Problem

- CLI toolchain versioning
- Extending CLI toolchains

## The Solution

- User runs `<command x>` in their project
- Look for a local version of `<command x>`
- If found, shell out to local version
- If not found continue using the `<command x>` binary

## The `garlic.Garlic` Type

<figure id="garlic.garlic">

<go doc="github.com/markbates/garlic.Garlic"></go>

<figcaption><godoc>github.com/markbates/garlic#Garlic</godoc></figcaption>

</figure>

## Garlic Commander

<figure id="garlic.commander">

<go sym="github.com/markbates/garlic.Commander"></go>

<figcaption><godoc>github.com/markbates/garlic#Commander</godoc></figcaption>

</figure>

## The `garlic.Garlic#Main` Method

<figure id="garlic.garlic.Main">

<go sym="github.com/markbates/garlic.Garlic.Main"></go>

<figcaption><godoc>github.com/markbates/garlic#Garlic.Main</godoc></figcaption>

</figure>

## The `main` Function

<figure id="main">

<code src="cmd/bostongo/main.go#main"></code>

<figcaption>[cmd/bostongo/main.go](./cmd/bostongo/main.go)</figcaption>

</figure>

## Testing Garlic

<figure id="garlic-works">

<code src="cmd/bostongo/cli/garlic_test.go#garlic-works"></code>

<figcaption>[cmd/bostongo/cli/garlic_test.go](./cmd/bostongo/cli/garlic_test.go)</figcaption>

</figure>

## Running the Tests

<figure id="testing-garlic">

<go src="cmd/bostongo/cli" test="-v -run Garlic_Works"></go>

<figcaption>Running the <code>garlic</code> tests.</figcaption>

</figure>

## Final Folder Structure

<figure id="tree">

<cmd exec="tree -I testdata -I assets -I *.md"></cmd>

<figcaption>Final folder structure.</figcaption>

</figure>
