package main

func main() {
	listener := UpstreamListener{addr: "tcp://:5050", multicore: true}
	listener.Start()
}
