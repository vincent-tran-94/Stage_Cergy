package main

type LemmatisationStruct struct {
	Terme         string
	Lemmatisation string
}

type Dictionary struct {
	Key   string
	Value int
}

type DictionaryList []Dictionary
