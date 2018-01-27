package main

func main() {
	r := newReceiver()
	cmdExecute(r)
	r.Wait()
}
