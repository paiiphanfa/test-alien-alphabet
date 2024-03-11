There are 7 symbols that represent the number:
A for 1
B for 5
Z for 10
L for 50
C for 100
D for 500
R for 1000

The usage of the symbols is quite similar to the Roman numerals. When A is placed before B (AB), it is equal to 4 also with other symbols.
◦ A can be placed before B (5) and Z (10) to make 4 and 9.
◦ Z can be placed before L (50) and C (100) to make 40 and 90.
◦ C can be placed before D (500) and R (1000) to make 400 and 900.

This program is created to translate the symbols above. For example:
Example 1:
Input: s = "AAA"
Output: 3
Explanation: AAA = 3.
Example 2:
Input: s = "LBAAA"
Output: 58
Explanation: L = 50, B= 5, AAA = 3.
Example 3:
Input: s = "RCRZCAB"
Output: 1994
Explanation: R = 1000, CR = 900, ZC = 90 and AB = 4.

Code Explanation
	
    alienValues := map[rune]int{
		'R': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'Z': 10,
		'B': 5,
		'A': 1,
	}
    In this part, I used map function to create symbols as keys that reference to values of the symbols themselves.

        	for i := 0; i < len(input); i++ {
	    	char := rune(input[i])
	    	nextChar := ' '

	    	if i < len(input)-1 {
	    		nextChar = rune(input[i+1])
	    	}

	    	switch char {
	    	case 'A':
	    		if nextChar == 'B' || nextChar == 'Z' {
	    			alienNum = append(alienNum, alienValues[nextChar]-alienValues[char])
	    			explanation = append(explanation, fmt.Sprintf("%s%s = %d", string(char), string(nextChar), alienValues[nextChar]-alienValues[char]))
	    			i++
	    		} else {
	    			alienNum = append(alienNum, alienValues[char])
	    			explanation = append(explanation, fmt.Sprintf("%s = %d", string(char), alienValues[char]))
	    		}
	    	case 'Z':
	    		if nextChar == 'L' || nextChar == 'C' {
	    			alienNum = append(alienNum, alienValues[nextChar]-alienValues[char])
	    			explanation = append(explanation, fmt.Sprintf("%s%s = %d", string(char), string(nextChar), alienValues[nextChar]-alienValues[char]))
	    			i++
	    		} else {
	    			alienNum = append(alienNum, alienValues[char])
	    			explanation = append(explanation, fmt.Sprintf("%s = %d", string(char), alienValues[char]))
	    		}
	    	case 'C':
	    		if nextChar == 'D' || nextChar == 'R' {
	    			alienNum = append(alienNum, alienValues[nextChar]-alienValues[char])
	    			explanation = append(explanation, fmt.Sprintf("%s%s = %d", string(char), string(nextChar), alienValues[nextChar]-alienValues[char]))
	    			i++
	    		} else {
	    			alienNum = append(alienNum, alienValues[char])
	    			explanation = append(explanation, fmt.Sprintf("%s = %d", string(char), alienValues[char]))
	    		}
	    	default:
	    		alienNum = append(alienNum, alienValues[char])
	    		explanation = append(explanation, fmt.Sprintf("%s = %d", string(char), alienValues[char]))
	    	}
	    }
    This loop is where I defined the input alphabet with values. I use for loop to loop through the input string and declare "char" as an alphabet of an index. The value "nextChar" is set in order to check the special case.

    To check the value of symbols, I used switch case to check the char value. If it is not in the condition of A, Z, C, which can be placed before some symbol to reduce the value of the behind degit, its value will remain the same.

    After compare an alphabet with the alienValues mapped, I kept it in an array called alienNum. Moreover, I also add explanation array to keep the explanation of each symbol to display in the output later.

    For the special cases like AB, AZ, I checked it in the switch case by using the "nextChar" value and reduce it to be correct. Then, I kept both symbols to explanation.

    	
        for _, value := range alienNum {
		sum += value
	    }

        fmt.Printf("Output: %s = %d\n", input, sum)
	    fmt.Printf("Explanation: %s.\n", strings.Join(explanation, ", "))

    After I got an array that contains the values of the input symbols, I sum all the values in an array to get the final value and print the output value and explanation array (separate each index by ",") to the output.