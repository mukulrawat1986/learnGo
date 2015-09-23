package main

import (
    "fmt"
    "github.com/mukulrawat1986/learnGo/poetry"
)

func main(){

    p := poetry.Poem{{"And from my pillow, looking forth by light", "Of moon or favouring stars, I could behold","The antechapel where the statue stood", "Of Newton with his prism and silent face,", "The marble index of a mind for ever", "Voyaging through strange seas of Thought, alone."}}

    fmt.Printf("%s\n", p)
}

