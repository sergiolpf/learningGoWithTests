package iteration

const repeatCount = 5


func Repeat(text string) string {

	var repeated string

	for count := 0; count < repeatCount ;count++ {
		repeated += text
	}

	return repeated
}



