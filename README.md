# error_report
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![Semantic Versioning 2.0.0](https://img.shields.io/badge/semver-2.0.0-brightgreen?style=flat-square)](https://semver.org/spec/v2.0.0.html)
[![License](https://img.shields.io/github/license/Anadian/error_report)](https://github.com/Anadian/error_report/Documents/LICENSE)
[![GoDoc](https://godoc.org/github.com/Anadian/error_report/source?status.svg)](https://godoc.org/github.com/Anadian/error_report/source)
[![Build Status](https://travis-ci.org/Anadian/error_report.svg?branch=master)](https://travis-ci.org/Anadian/error_report)
[![Coverage Status](https://coveralls.io/repos/github/Anadian/error_report/badge.svg?branch=master)](https://coveralls.io/github/Anadian/error_report?branch=master)

> A true alternative approach to handling errors in GoLang.
# Table of Contents
- [Background](#Background)
- [Install](#Install)
- [Usage](#Usage)
- [API](#API)
- [Contributing](#Contributing)
- [License](#License)
# Background
# Install
```bash
go get github.com/Anadian/error_report/source
```
# Usage
Example
```golang
package main;

import "github.com/Anadian/error_report/source";

func FunctionWhichReturnsAnErrorReport() error_report.ErrorReport_struct{
	var return_report error_report.ErrorReport_struct = error_report.New( 1, map[string]interface{}{
		"function_name": "FunctionWhichReturnsAnErrorReport",
		"line_number": 8,
	},
	nil );
	return return_report;
}

func ChecksAndWrapsAnErrorReport() error_report.ErrorReport_struct{
	var this_function_error_report error_report.ErrorReport_struct;
	var subroutine_error_report error_report.ErrorReport_struct = FunctionWhichReturnsAnErrorReport();
	if( subroutine_error_report.IsError() == true ){ //Checks if an error ocurred.
		if( subroutine_error_report.CodeEqual( 1 ) == true ){ //Checks the code.
			this_function_error_report.Code = 2;
			this_function_error_report.Data = map[string]interface{}{
				"function_name": "ChecksAndWrapsAnErrorReport",
				"line_number": 16,
			}
			this_function_error_report.Wrap( subroutine_error_report );
		}
	}
	return this_function_error_report;
}

func main(){
	error_report := ChecksAndWrapsAnErrorReport();
	/*
		error_report = {
			Code: 2,
			Data: {
				"function_name": "ChecksAndWrapsAnErrorReport",
				"line_number": 16
			},
			Wrapped -> { (a pointer to another ErrorReport_struct object.)
				Code: 1,
				Data: {
					"function_name": "FunctionWhichReturnsAnErrorReport",
					"line_number": 8
				}
				Wrapped -> nil
			}
		}
	*/
}
```
For example usage of the whole API, checkout the [test file](source/error_report_test.go).
# API
# Contributing
Changes are tracked in [CHANGES.md](CHANGES.md).
# License
MIT ©2019 Anadian
SEE LICENSE IN [LICENSE](LICENSE)
