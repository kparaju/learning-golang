package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

/**
 * Take a textfile and count the number of occurences of 
 * each letter of the alphabet
 */
func main() {
  var fileName string
  if (len(os.Args) < 2) {
    fmt.Println("Please specify file to read.")
    os.Exit(1)
  }
  // Difference between = and := is that = is for vars that are
  // already init'd. := init'd and assigns. Pretty cool.
  fileName = os.Args[1]
  content, err := ioutil.ReadFile(fileName)
  if err != nil {
    os.Exit(1)
  }
  strContent := string(content)
  freq := freqAnalysis(strContent)
  for i := 0; i < 26; i++ {
    fmt.Printf("%c,%v\n", i + 'A', freq[i])
  }
}

/**
 * Take a string and returns an array of size 26 with the occurence
 * of each letter of the alphabet
 */
func freqAnalysis(content string) [26]int {
  var counts [26]int;
  for i := 0; i < len(content); i++ {
    var c uint8 = content[i];
    // Convert it to uppercase
    if (c >= 'a' && c <= 'z') {
      c = c - ('a' - 'A')
    }
    if (c >= 'A' && c <= 'Z') {
      c = c - 'A';
      counts[c]++
    }
  }
  return counts
}
