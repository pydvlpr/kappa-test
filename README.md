# kappa-test
My Go implementation of the kappa-test for cryptoanalysis

The program is the result of a task in my studies.

# Background

The kappa-test is also known as Friedman test, named by his inventor William Friedman.

In cryptography, coincidence counting is the technique of putting two texts side-by-side and counting the number of times that identical letters appear in the same position in both texts. This count, either as a ratio of the total or normalized by dividing by the expected count for a random source model, is known as the index of coincidence, or IC for short. (wikipedia)

An index around 0.0762 indicates monoalphabetic encryption. Significantly smaller index says, the text is not monoalphabetisch encrypted.

# Usage

go run kappa-test.go FILENAME

# Testfiles

testfile.txt: Short text about Hucklebarry Fin
cesar.txt: text about Huck (german), usage of step 7
vigenere.txt: text about Huck (german), key: Secret

# Infos 
Learn about cryptography and cryptoanalysis at CryptTool (https://www.cryptool.org/)
