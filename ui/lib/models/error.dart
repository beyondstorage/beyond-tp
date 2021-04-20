class Error {
  String message;
  Map<String, String> extensions;

  Error({ required this.message, required this.extensions });

  factory Error.fromMap(Map<String, dynamic> error) => Error(
    message: error["message"],
    extensions: error["extensions"],
  );

  String toString() => message;
}

class Errors {
  List<Error> errors;

  Errors({ required this.errors });

  factory Errors.fromList(List<Map<String, dynamic>> errors) => Errors(
    errors: List<Error>.from(errors.map((error) => Error.fromMap(error))),
  );
}
