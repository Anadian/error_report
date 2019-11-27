/**
* @file error_report_test.go
* @brief Tests functions in `error_report.go`.
* @author Anadian
* @copyright MITlicensetm(2019,Canosw)
*/

// error_report tests functions in `error_report.go`.

package error_report;

//# Dependencies
import(
	//## Internal
	//## Standard
	"testing"
	"log"
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
}

//# Private Functions

