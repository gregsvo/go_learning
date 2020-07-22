package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name + "!"

}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case "fr":
		prefix = frenchHelloPrefix
	case "esp":
		prefix = spanishHelloPrefix

	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Greg", "en"))
}
