package utils

import "fmt"

//makes text bold
func Bold(text string) string {

	return fmt.Sprintf("**%s**", text)
}


// makes text italic
func Italic(text string) string {

	return fmt.Sprintf("*%s*", text)
}


// Formats /quote command
func FormatQuote(quote string) string {
	
	title := "Quote of the day 💭:"
	signature := Italic("-- your lovely Anatoliy")
	quote = Bold(quote)

	return fmt.Sprintf("%s\n\n\"%s\"\n\n%s", title, quote, signature)
}

// Formats /on-esports command
func FormatOnEsports(quote string) string {

	signature := Italic("-- your lovely ON Esports 💙")
	quote = Bold(quote)

	return fmt.Sprintf("%s\n\n%s", quote, signature)
}