package models

type CompileJob struct {
	Title      string `bson:"title"`
	Category   string `bson:"category"`
	Difficulty string `bson:"difficulty"`
}
