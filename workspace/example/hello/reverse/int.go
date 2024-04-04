// Note: this file is added manually, NOT cloned from the remote repository.

package reverse

import "strconv"

// Int returns the decimal reversal of the integer i.
func Int(i int) int {
	i, _ = strconv.Atoi(String(strconv.Itoa(i)))
	return i
}
