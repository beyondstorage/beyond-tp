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

class Service {
  String name;
  String type;
  Credential credential;
  Endpoint endpoint;

  Service({
    required this.name,
    required this.type,
    required this.credential,
    required this.endpoint,
  });

  factory Service.fromMap(Map<String, dynamic> json) => Service(
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

class Services {
  List<Service> services;

  Services({
    required this.services,
  });

  factory Services.fromList(List<Object> _identities) {
    return Services(
      services: List<Service>.from(
        _identities.map(
          (service) => Service.fromMap(service as Map<String, dynamic>),
        ),
      ),
    );
  }

  List<Map<String, dynamic>> toList() {
    return services.map((service) => service.toMap()).toList();
  }

  String toString() => json.encode(toList());

  int length() => services.length;
}
