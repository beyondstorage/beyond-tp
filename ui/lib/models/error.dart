class Error {
  String message;
  Map<String, String> extensions;

  Error({ this.message, this.extensions });

  factory Error.fromMap(Map<String, dynamic> error) => Error(
    message: error["message"],
    extensions: error["extensions"],
  );

  String toString() => message;
}

class Errors {
  List<Error> errors;

  Errors({ this.errors });

  factory Errors.fromList(List<Object> errors) => Errors(
    errors: List<Error>.from(errors.map((error) => Error.fromMap(error))),
  );
}
