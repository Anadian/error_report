/**
* @file error_report_test.go
* @brief Tests functions in `error_report.go`.
* @author Anadian
* @copyright 	Copyright 2019 Canosw
	Permission is hereby granted, free of charge, to any person obtaining a copy of this 
software and associated documentation files (the "Software"), to deal in the Software 
without restriction, including without limitation the rights to use, copy, modify, 
merge, publish, distribute, sublicense, and/or sell copies of the Software, and to 
permit persons to whom the Software is furnished to do so, subject to the following 
conditions:
	The above copyright notice and this permission notice shall be included in all copies 
or substantial portions of the Software.
	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, 
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A 
PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT 
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF 
CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE 
OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

// error_report tests functions in `error_report.go`.

package error_report;

//# Dependencies
import(
	//## Internal
	//## Standard
	"testing"
	"log"
	"fmt"
	//## External
);

//# Constants
const(
	//## Exported Constants
	//### Errors
	//## Private Constants
);

//# Types
//## Interfaces
//## Structs
//### Methods

//# Global Variables
var(
	//## Exported Variables
	//## Private Variables
);

//# Exported Functions
func Test( t *testing.T ){
	log.SetFlags( log.LstdFlags | log.Lmicroseconds | log.Lshortfile );
	var function_return ErrorReport_struct;
	//## New
	//### (code < -9007199254740991)
	var expected_error_report ErrorReport_struct = ErrorReport_struct{
		-9007199254740991,
		map[string]interface{}{
			"filename": "error_report_test.go",
		},
		nil,
	};
	var error_report ErrorReport_struct = New( -9007199254740992, map[string]interface{}{ "filename": "error_report_test.go" }, nil );
	if( error_report.Code == expected_error_report.Code ){
		log.Printf("Success: error_report (%v) == expected_error_report (%v)\n", error_report, expected_error_report);
	} else{
		t.Fail();
		log.Printf("Failure: error_report (%v) !== expected_error_report (%v)\n", error_report, expected_error_report);
	}
	//### (code > 9007199254740991)
	expected_error_report.Code = 9007199254740991;
	var error_report2 ErrorReport_struct = New( 9007199254740992, map[string]interface{}{ "filename": "error_report_test.go" }, nil );
	if( error_report2.Code == expected_error_report.Code ){
		log.Printf("Success: error_report (%v) == expected_error_report (%v)\n", error_report2, expected_error_report);
	} else{
		t.Fail();
		log.Printf("Failure: error_report (%v) !== expected_error_report (%v)\n", error_report2, expected_error_report);
	}
	//### else
	expected_error_report.Code = 0;
	var error_report3 ErrorReport_struct = New( 0, map[string]interface{}{ "filename": "error_report_test.go" }, nil );
	if( error_report3.Code == expected_error_report.Code ){
		log.Printf("Success: error_report (%v) == expected_error_report (%v)\n", error_report3, expected_error_report);
	} else{
		t.Fail();
		log.Printf("Failure: error_report (%v) !== expected_error_report (%v)\n", error_report3, expected_error_report);
	}
	//## ErrorReport_struct Methods
	var expected_bool bool;
	var actual_bool bool;
	//### NoError
	//#### (error_report.Code == 0)
	expected_bool = true;
	actual_bool = error_report3.NoError();
	if( actual_bool == expected_bool ){
		log.Printf("Success: actual_bool (%t) == expected_bool (%t)\n", actual_bool, expected_bool);
	} else{
		t.Fail();
		log.Printf("Failure: actual_bool (%t) !== expected_bool (%t)\n", actual_bool, expected_bool);
	}
	//#### else
	expected_bool = false;
	actual_bool = error_report2.NoError();
	if( actual_bool == expected_bool ){
		log.Printf("Success: actual_bool (%t) == expected_bool (%t)\n", actual_bool, expected_bool);
	} else{
		t.Fail();
		log.Printf("Failure: actual_bool (%t) !== expected_bool (%t)\n", actual_bool, expected_bool);
	}
	//### IsError
	//#### (error_report.Code == 0)
	expected_bool = false;
	actual_bool = error_report3.IsError();
	if( actual_bool == expected_bool ){
		log.Printf("Success: actual_bool (%t) == expected_bool (%t)\n", actual_bool, expected_bool);
	} else{
		t.Fail();
		log.Printf("Failure: actual_bool (%t) !== expected_bool (%t)\n", actual_bool, expected_bool);
	}
	//#### else
	expected_bool = true;
	actual_bool = error_report2.IsError();
	if( actual_bool == expected_bool ){
		log.Printf("Success: actual_bool (%t) == expected_bool (%t)\n", actual_bool, expected_bool);
	} else{
		t.Fail();
		log.Printf("Failure: actual_bool (%t) !== expected_bool (%t)\n", actual_bool, expected_bool);
	}
	//### CodeEqual
	//#### (error_report.Code == 0)
	expected_bool = true;
	actual_bool = error_report3.CodeEqual( int64(0) );
	if( actual_bool == expected_bool ){
		log.Printf("Success: actual_bool (%t) == expected_bool (%t)\n", actual_bool, expected_bool);
	} else{
		t.Fail();
		log.Printf("Failure: actual_bool (%t) !== expected_bool (%t)\n", actual_bool, expected_bool);
	}
	//#### else
	expected_bool = false;
	actual_bool = error_report2.CodeEqual( int64(0) );
	if( actual_bool == expected_bool ){
		log.Printf("Success: actual_bool (%t) == expected_bool (%t)\n", actual_bool, expected_bool);
	} else{
		t.Fail();
		log.Printf("Failure: actual_bool (%t) !== expected_bool (%t)\n", actual_bool, expected_bool);
	}
	//### Wrap
	error_report2.Wrap( error_report );
	if( error_report2.Wrapped != nil ){
		log.Printf("Success: error_report2 (%v) now wraps error_report (%v).", error_report2, error_report);
	} else{
		t.Fail();
		log.Printf("Failure: error_report2 (%v) doesn't wrap anything.", error_report2);
	}
	//### WrapsAnother
	//#### true
	expected_bool = true;
	actual_bool = error_report2.WrapsAnother();
	if( actual_bool == expected_bool ){
		log.Printf("Success: actual_bool (%t) == expected_bool (%t)\n", actual_bool, expected_bool);
	} else{
		t.Fail();
		log.Printf("Failure: actual_bool (%t) !== expected_bool (%t)\n", actual_bool, expected_bool);
	}
	//#### else
	expected_bool = false;
	actual_bool = error_report3.WrapsAnother();
	if( actual_bool == expected_bool ){
		log.Printf("Success: actual_bool (%t) == expected_bool (%t)\n", actual_bool, expected_bool);
	} else{
		t.Fail();
		log.Printf("Failure: actual_bool (%t) !== expected_bool (%t)\n", actual_bool, expected_bool);
	}
	//### GetWrapped
	//#### true
	var error_report1_copy ErrorReport_struct;
	error_report1_copy = error_report2.GetWrapped();
	if( error_report1_copy.CodeEqual( error_report.Code ) == true ){
		log.Printf("Success: error_report1_copy (%v) is equal to error_report (%v)\n", error_report1_copy, error_report);
	} else{
		t.Fail();
		log.Printf("Failure: error_report1_copy (%v) is not equal to error_report (%v)\n", error_report1_copy, error_report);
	}
	//#### else
	var error_report_copy ErrorReport_struct;
	error_report_copy = error_report3.GetWrapped();
	if( error_report_copy.CodeEqual( error_report.Code ) == false ){
		log.Printf("Success: error_report_copy (%v) is not equal to error_report (%v)\n", error_report_copy, error_report);
	} else{
		t.Fail();
		log.Printf("Failure: error_report_copy (%v) is equal to error_report (%v)\n", error_report_copy, error_report);
	}
	function_return = SubTestGetWrappedBottom();
	if( function_return.IsError() == true ){
		t.Fail();
		log.Printf("Failure: TestGetWrappedBottom returned an error: %v\n", function_return);
	}
}

//# Private Functions

/**
* @fn TestGetWrappedBottom
* @brief Test error_report.GetWrappedBottom
* @return ( return_report ErrorReport_struct ) 
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/

// TestGetWrappedBottom test error_report.GetWrappedBottom
func SubTestGetWrappedBottom() ( return_report ErrorReport_struct ){
	log.SetFlags( log.LstdFlags | log.Lmicroseconds | log.Lshortfile );
	//Variables
	var er1 ErrorReport_struct = New( 1, map[string]interface{}{}, nil );
	var er2 ErrorReport_struct = New( 2, map[string]interface{}{}, &er1 );
	var er3 ErrorReport_struct = New( 3, map[string]interface{}{}, &er2 );
	var function_return ErrorReport_struct;
	//Parametres
	//Function
	if( er3.WrapsAnother() == true ){
		function_return = er3.GetWrappedBottom();
		if( function_return.NoError() == true ){
			log.Printf("%v\n", function_return);
			function_return = function_return.Data["deepest_wrapped"].(ErrorReport_struct);
			if( function_return.CodeEqual( er1.Code ) == true ){
				log.Printf("Success: GetWrappedBottom retrieved a nested report successfully.\n");
				return_report = New( 0, map[string]interface{}{}, &function_return );
			} else{
				return_report = New( 5, map[string]interface{}{ "message": "Error: Returned report didn't match original nested report." }, &function_return );
			}
		} else{
			return_report = New( 4, map[string]interface{}{ "message": fmt.Sprintf("Error: GetWrappedBottom returned an error: %v\n", function_return) }, &function_return );
		}
	} else{
		return_report = New( 3, map[string]interface{}{ "message": "What the?" }, nil );
	}
	//Return
	return return_report;
}


