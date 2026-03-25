module example.com/hello

go 1.26.1

replace example.com/messages => ../messages

replace example.com/greetings => ../greetings

require example.com/greetings v1.1.0

require example.com/messages v1.1.0

require d03x.dev/server v0.0.0-00010101000000-000000000000

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace d03x.dev/server => ../server
