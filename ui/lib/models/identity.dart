import 'dart:convert';

class Credential {
  String protocol;
  List<String> args;

  Credential({
    required this.protocol,
    required this.args,
  });

  factory Credential.fromMap(Map<String, dynamic> json) => Credential(
        protocol: json["protocol"] ?? "",
        args: List<String>.from(json["args"]),
      );

  Map<String, dynamic> toMap() => {
        "protocol": protocol,
        "args": args.toList(),
      };
}

class Endpoint {
  String protocol;
  String host;
  int port;

  Endpoint({
    required this.protocol,
    required this.host,
    required this.port,
  });

  factory Endpoint.fromMap(Map<String, dynamic> json) => Endpoint(
        protocol: json["protocol"] ?? "",
        host: json["host"] ?? "",
        port: json["port"] ?? "",
      );

  Map<String, dynamic> toMap() => {
        "protocol": protocol,
        "host": host,
        "port": port,
      };
}

class Identity {
  String name;
  String type;
  Credential credential;
  Endpoint endpoint;

  Identity({
    required this.name,
    required this.type,
    required this.credential,
    required this.endpoint,
  });

  factory Identity.fromMap(Map<String, dynamic> json) => Identity(
        name: json["name"] ?? "",
        type: json["type"] ?? "",
        credential: Credential.fromMap(json["credential"] ?? ""),
        endpoint: Endpoint.fromMap(json["endpoint"] ?? ""),
      );

  Map<String, dynamic> toMap() => {
        "name": name,
        "type": type,
        "credential": credential.toMap(),
        "endpoint": endpoint.toMap(),
      };
}

class Identities {
  List<Identity> identities;

  Identities({
    required this.identities,
  });

  factory Identities.fromList(List<Object> _identities) {
    return Identities(
      identities: List<Identity>.from(
        _identities.map(
          (identity) => Identity.fromMap(identity as Map<String, dynamic>),
        ),
      ),
    );
  }

  List<Map<String, dynamic>> toList() {
    return identities.map((identity) => identity.toMap()).toList();
  }

  String toString() => json.encode(toList());

  int length() => identities.length;
}
