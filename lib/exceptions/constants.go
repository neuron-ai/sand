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

package exceptions

const (
	EOFError           string = "EOF Error"            // Thrown when end of file is reached in expectance of something else.
	IndentError        string = "Indentation Error"    // Thrown when indentation is inconsistent.
	MissingFileError   string = "Missing File Error"   // Thrown in the case of a missing file.
	MissingModuleError string = "Missing Module Error" // Thrown in the case of a missing module.
	RuntimeError       string = "Runtime Error"        // Thrown in the event of a runtime error.
	SyntaxError        string = "Syntax Error"         // Thrown when syntax is incorrect.
)
