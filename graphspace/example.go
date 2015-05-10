package main

const Example = `graph G {
   layout = neato;

	run -- intr;
	intr -- runbl;
	runbl -- run;
	run -- kernel;
	kernel -- zombie;
	kernel -- sleep;
	kernel -- runmem;
	sleep -- swap;
	swap -- runswap;
	runswap -- new;
	runswap -- runmem;
	new -- runmem;
	sleep -- runmem;
}
`