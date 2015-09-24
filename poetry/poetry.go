package poetry

import (
    "bufio"
    "fmt"
    _"io"
    "os"
)

// Create some types of our own
type Line string
type Stanza []Line
type Poem []Stanza

// Create an empty poem
func NewPoem() Poem{
    return Poem{}
}

// Load a poem from a file
func LoadPoem(filename string) (Poem, error){ 
    f, err := os.Open(filename)
    if err != nil{
        return nil, err
    }
    defer f.Close()

    p := Poem{}

    var s Stanza

    scan := bufio.NewScanner(f)

    for scan.Scan(){
        l := scan.Text()

        // If there is nothing in the line, that means Stanza is finished, so append it to the poem
        // If not then append the line to Stanza after converting the string to a Line
        if l == ""{
            p = append(p, s)
            s = Stanza{}
            continue
        }
        s = append(s, Line(l))
    }
    // At the end of the scanning we will add the last token which will be emptry String since the Stanza finished
    p = append(p, s)

    if scan.Err() != nil{
        return nil, scan.Err()
    }

    return p, nil
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

// Set String() for Stanza to make it implement Stringer interface
func (s Stanza) String() string{
    result := ""
    for _, l := range s{
        result += fmt.Sprintf("%s\n",l)
    }
    return result
}


// set String() for poem to make it implement Stringer interface
func (p Poem) String() string{
    result := ""
    for _, s := range p{
        result += fmt.Sprintf("%s\n", s)
    }
    return result
}
