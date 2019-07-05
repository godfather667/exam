# exam
Program exam provides functions to the answer three given problems.

1. triangleID: Determine type of triangle.
2. fiveList: Find fifth Item from end of a list.
3. equalList: Compare Two Lists for Equality.

Units tests are included for each function.

# Executing Progam

* Download the progam zip file from https://github.com/godfather667/exam
* Rename **exam-master.zip** to **exam.zip**
* Unpack the zip file into a standard golang dir structure: **../src directory**
* Change directory to **../src/exam** 

* _go run exam_
  * Output:  **Exam program**

* _go test exam_
  * Output:  **ok  exam  .002s**

# Test Coverage

* Issue CoverageTool:
  Command: **go tool cover -html=cover.out -o cover.html**
* The **cover.html** file shows which instructions have been test in green.
* Untested code is shown in red.
* The test coverage percentage for exam is **95.8%**

