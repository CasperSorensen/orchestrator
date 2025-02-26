package scheduler

type Scheduler interface {
	SelectCanditateNodes()
	Score()
	Pick()
}
