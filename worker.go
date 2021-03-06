package scraper

type Worker struct {
}

func newWorker() *Worker {
	return &Worker{}
}

func (w *Worker) execute(j Job) []byte {
	return extractBytesFrom(string(j))
}



func (w *Worker) work(jobQueue JobQueue, res chan<- []byte) {
	for job := range jobQueue {

		//you have one job. Do it!
		ans := w.execute(job)

		//sends the result to the channel.
		// The result is received from the other
		//end, freeing up the worker.
		res <- ans
	}

}
