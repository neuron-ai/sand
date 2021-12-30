// Copyright 2021 Neuron-AI GitHub Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ==============================================================================

package colour

import "fmt"

// ~ Excuse the sea of comments - they come up as docstrings in VS Code. Don't move this line down, or it'll edit the docstrings.

const (
	BLACK       string = "\033[30m"   // Colour code for black. Place at start of fmt.Println call to turn text black.
	BLUE        string = "\033[34m"   // Colour code for blue. Place at start of fmt.Println call to turn text blue.
	CYAN        string = "\033[36m"   // Colour code for cyan. Place at start of fmt.Println call to turn text cyan.
	GREEN       string = "\033[32m"   // Colour code for green. Place at start of fmt.Println call to turn text green.
	GREY        string = "\033[1;30m" // Colour code for grey. Place at start of fmt.Println call to turn text grey.
	LIGHTBLUE   string = "\033[1;34m" // Colour code for pale blue. Place at start of fmt.Println call to turn text light blue.
	LIGHTCYAN   string = "\033[1;36m" // Colour code for pale cyan. Place at start of fmt.Println call to turn text light cyan.
	LIGHTGREY   string = "\033[1;37m" // Colour code for light grey. Place at start of fmt.Println call to turn text light grey.
	LIGHTPURPLE string = "\033[1;35m" // Colour code for pale purple. Place at start of fmt.Println call to turn text light purple.
	LIGHTRED    string = "\033[1;31m" // Colour code for pale red. Place at start of fmt.Println call to turn text light red.
	LIGHTYELLOW string = "\033[1;33m" // Colour code for yellow. Place at start of fmt.Println call to turn text yellow.
	PURPLE      string = "\033[35m"   // Colour code for purple. Place at start of fmt.Println call to turn text purple.
	RED         string = "\033[31m"   // Colour code for red. Place at start of fmt.Println call to turn text red.
	RESET       string = "\033[0m"    // Reset colour code. Place at start of fmt.Println call to reset text colour.
	WHITE       string = "\033[1;37m" // Colour code for white. Place at start of fmt.Println call to turn text white.
	YELLOW      string = "\033[33m"   // Colour code for yellow. Place at start of fmt.Println call to turn text yellow.

	BGBLACK  string = "\033[40m" // Colour code for black background. Place at start of fmt.Println call to turn background black.
	BGBLUE   string = "\033[44m" // Colour code for blue background. Place at start of fmt.Println call to turn background blue.
	BGCYAN   string = "\033[46m" // Colour code for cyan background. Place at start of fmt.Println call to turn background cyan.
	BGGREEN  string = "\033[42m" // Colour code for green background. Place at start of fmt.Println call to turn background green.
	BGPURPLE string = "\033[45m" // Colour code for purple background. Place at start of fmt.Println call to turn background purple.
	BGRED    string = "\033[41m" // Colour code for red background. Place at start of fmt.Println call to turn background red.
	BGWHITE  string = "\033[47m" // Colour code for white background. Place at start of fmt.Println call to turn background white.
	BGYELLOW string = "\033[43m" // Colour code for yellow background. Place at start of fmt.Println call to turn background yellow.
)

// PrintCol prints in colour with the chosen background and
// foreground colours.
func PrintCol(bg string, fore string, args ...interface{}) {
	fmt.Println(bg, fore, args, RESET)
}
