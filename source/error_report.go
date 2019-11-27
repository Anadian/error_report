/**
* @file error_report.go
* @brief Provides interfaces, structs, and functions for working with error reports.
* @author Anadian
* @copyright MITlicensetm(2019,Canosw)
*/

// error_report provides interfaces, structs, and functions for working with error reports.
package error_report;

//# Dependencies
import(
	//## Internal
	//## Standard
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
type ErrorReport_interface interface{
	NoError() bool
	IsError() bool
	CodeEqual( code int64 ) bool
	//GetString() string
	WrapsAnother() bool
	GetWrapped() ErrorReport_struct
	Wrap( error_report ErrorReport_struct )
	//ToJSON() []byte
	//FromJSON( json_data []byte )
	//MatchJSONSchema( schema_json_data [] )
}
//## Structs
type ErrorReport_struct struct{
	Code int64
	Data map[string]interface{}
	Wrapped *ErrorReport_struct
}
//### Methods
/**
* @fn NoError
* @brief Returns true if the error report doesn't signify any errors.
* @struct error_report ErrorReport_struct
* @return bool
* @retval true The preceding `error_report` struct signifies no errors.
* @retval false There are errors signified by the parent value.
*/

// NoError returns true if the error report doesn't signify any errors.
func (error_report ErrorReport_struct) NoError() bool{
	var _return bool;
	/* Variables */
	/* Parametres */
	/* Function */
	if( error_report.Code == 0 ){
		_return = true;
	} else{
		_return = false;
	}
	/* Return */
	return _return;
}
/**
* @fn IsError
* @brief Returns true if the error report does signify an error.
* @struct error_report ErrorReport_struct
* @return bool
* @retval true The preceding `error_report` struct signifies an error.
* @retval false There aren't any errors signified by the parent value.
*/

// IsError returns true if the error report does signify an error.
func (error_report ErrorReport_struct) IsError() bool{
	var _return bool;
	/* Variables */
	/* Parametres */
	/* Function */
	if( error_report.Code == 0 ){
		_return = false;
	} else{
		_return = true;
	}
	/* Return */
	return _return;
}

///**
//* @fn CodeEqual
//* @brief Returns true if the `code` properties of the reports match.
//* @struct error_report *ErrorReport_struct
//* @param other_report ErrorReport_struct [in] The other report to be checked against.
//* @return bool
//* @retval true The two reports' `code` properties are equal.
//* @retval false The two reports' `code` properties are not equal.
//*/
//// CodeEqual returns true if the `code` properties of the reports match.
//func (error_report *ErrorReport_struct) CodeEqual( other_report ErrorReport_struct ) bool{
//	var _return bool = nil;
//	/* Variables */
//	
//	/* Parametres */
//	/* Function */
//	
//	/* Return */
//	return _return;
//}

/**
* @fn CodeEqual
* @brief Returns true if the report's `Code` property matches the `code` parametre given.
* @struct error_report ErrorReport_struct
* @param code int64 [in] The code value to check against the error report.
* @return bool
* @retval true The codes are equal.
* @retval false The codes are not equal.
*/

// CodeEqual returns true if the report's `code` property matches the `code` parametre given.
func (error_report ErrorReport_struct) CodeEqual( code int64 ) bool{
	var _return bool;
	/* Variables */
	/* Parametres */
	/* Function */
	if( error_report.Code == code ){
		_return = true;
	} else{
		_return = false;
	}
	/* Return */
	return _return;
}

/**
* @fn WrapsAnother
* @brief Returns true if this `error_report` wraps another error report.
* @struct error_report ErrorReport_struct
* @return bool
* @retval true `error_report` wraps another error report.
* @retval false `error_report does not wrap another error report.
*/

// WrapsAnother returns true if this `error_report` wraps another error report.
func (error_report ErrorReport_struct) WrapsAnother() bool{
	var _return bool;
	/* Variables */
	/* Parametres */
	/* Function */
	if( error_report.Wrapped != nil ){
		_return = true;
	} else{
		_return = false;
	}
	/* Return */
	return _return;
}

/**
* @fn GetWrapped
* @brief Returns the error report being wrapped by the current error report.
* @struct error_report ErrorReport_struct
* @return ErrorReport_struct
* @retval <wrapped ErrorReport_struct> Success
* @retval nil Failed
*/

// GetWrapped returns the error report being wrapped by the current error report.
func (error_report ErrorReport_struct) GetWrapped() ErrorReport_struct{
	var _return ErrorReport_struct;
	/* Variables */
	/* Parametres */
	/* Function */
	if( error_report.Wrapped != nil ){
		_return = *(error_report.Wrapped);
	}
	/* Return */
	return _return;
}

/**
* @fn Wrap
* @brief Wraps the given error report inside of this error report.
* @struct error_report *ErrorReport_struct
* @param wrap_report *ErrorReport_struct [in] The a pointer to an ErrorReport_struct to be wrapped.
* @return 
*/

// Wrap wraps the given error report inside of this error report.
func (error_report *ErrorReport_struct) Wrap( wrap_report ErrorReport_struct ) {
	/* Variables */
	/* Parametres */
	/* Function */
	error_report.Wrapped = &wrap_report;
	/* Return */
}

//# Global Variables
var(
	//## Exported Variables
	//## Private Variables
);

//# Exported Functions
/**
* @fn New
* @brief Creates a new Error Report.
* @param code int64 [in] The code, a unique, static, signed 64-bit-integer value, to be used for the new error report.
* @param data map[string]interface{} [in] A string-keyed map to be used for the `data` property of the new error report.
* @param wrapped_report *ErrorReport_struct [in] A pointer to a pre-existing error report, if any, which the new error report should wrap.
* @return ErrorReport_struct
* @retval <The new ErrorReport_struct> Success
* @retval nil Error
*/

// New creates a new Error Report.
func New( code int64, data map[string]interface{}, wrapped_report *ErrorReport_struct ) ErrorReport_struct{
	var _return ErrorReport_struct;
	/* Variables */
	/* Parametres */
	/* Function */
	if( code < -9007199254740991 ){
		_return.Code = -9007199254740991;
	} else if( code > 9007199254740991 ){
		_return.Code = 9007199254740991;
	} else{
		_return.Code = code;
	}
	_return.Data = data;
	_return.Wrapped = wrapped_report;
	/* Return */
	return _return;
}

//# Private Functions

