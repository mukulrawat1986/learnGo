package poetry

import (
    "testing"
)

func TestNumStanza(t *testing.T){
    
    p := Poem{}
    if p.NumStanzas() != 0{
        t.Fatalf("Empty poem is not empty!")
    }


    p = Poem{{
        "And from my pillow, looking forth by light",
        "Of moon or favouring stars, I could behold",
        "The antechapel where the statue stood", 
        "Of Newton with his prism and silent face,", 
        "The marble index of a mind for ever", 
        "Voyaging through strange seas of Thought, alone."}}

    if p.NumStanzas() != 1 {
        t. Fatalf("Unexpected stanza count %d", p.NumStanzas())
    }
}

func TestNumLines(t *testing.T){
    
    p := Poem{}
    if p.NumLines() != 0{
        t.Fatalf("Empty poem is not empty!")
    }

    p = Poem{{
        "And from my pillow, looking forth by light",
        "Of moon or favouring stars, I could behold",
        "The antechapel where the statue stood", 
        "Of Newton with his prism and silent face,",
        "The marble index of a mind for ever", 
        "Voyaging through strange seas of Thought, alone."}}

    if p.NumLines() != 6{
        t.Fatalf("Unexpected Line count %d", p.NumLines())
    }
}


func TestStats(t *testing.T){
    p := Poem{}

    v, c, puncs := p.Stats()

    if v != 0 || c != 0 || puncs != 0 {
        t.Fatalf("Bad number of vowels, consonants or punctuations")
   }

    p = Poem{{"Hello"}}

    v, c, puncs = p.Stats()

    if v != 2 || c != 3 ||  puncs != 0 {
        t.Fatalf("Bad number of vowels, consonants or punctuations vowels : %d consonants: %d puncs: %d", v, c, puncs)
    }

    p = Poem{{"Hello, World!"}}

    v, c, puncs = p.Stats()
    
    if v != 3 || c != 7 || puncs != 3{
        t.Fatalf("Bad number vowels: %d consonants: %d punctuations: %d ", v, c, puncs)
    }

}
