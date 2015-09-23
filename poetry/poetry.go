package poetry

// Create some types of our own
type Line string
type Stanza []Line
type Poem []Stanza

// Create our poem
func NewPoem() Poem{
    return Poem{}
}

// Get the number of Stanzas in a poem
func (p Poem) NumStanzas() int{
    return len(p)
}

// Get the number of line in a Stanza
func (s Stanza) NumLines() int{
    return len(s)
}

// Get the number of lines in a Poem
func (p Poem) NumLines() (count int){
    for _, s := range p{
        count += s.NumLines()
    }
    return 
}

// Get the count of vowels and consonants from the poem
func (p Poem) Stats() (numVowels, numConsonants, numPuncs int){
    for _, s := range p{
        for _, l := range s{
            for _, r := range l{
                switch r{
                case 'a', 'e', 'i','o','u':
                    numVowels += 1
                case ',', ' ', '!':
                    numPuncs += 1
                default:
                    numConsonants += 1
                }
            }
        }
    }
    return
}
