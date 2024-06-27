package main

func main() {
	nc, _ := nats.Connect("0.0.0.0:4222")
	defer.nc.Close()
	nc.Publish("k",[]byte("hello"))
}
