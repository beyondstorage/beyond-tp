import 'dart:convert';
class Agent {
  String name;
  String id;
  bool isOnline;
  String ip;
  double? networkSpeed;
  int? taskNumber;
  Agent({
    required this.name,
    required this.id,
    required this.ip,
    required this.isOnline,
    this.taskNumber,
    this.networkSpeed
  });

  factory Agent.fromMap(Map<String, dynamic> json) => Agent(
        name: json["name"] ?? "",
        id: json["id"] ?? "",
        isOnline: json["isOnline"] ?? false,
        ip: json["ip"] ?? "",
        networkSpeed: json["networkSpeed"] ?? null,
        taskNumber: json["taskNumber"] ?? null,
      );

  Map<String, dynamic> toMap() => {
        "name": name,
        "id": id,
        "isOnline": isOnline,
        "ip": ip,
        "networkSpeed": networkSpeed,
        "taskNumber": taskNumber
      };
}

class Agents {
  List<Agent> agents;

  Agents({
    required this.agents,
  });

  factory Agents.fromList(List<Object> _agents) {
    return Agents(
      agents: List<Agent>.from(
        _agents.map(
          (agents) => Agent.fromMap(agents as Map<String, dynamic>),
        ),
      ),
    );
  }

  List<Map<String, dynamic>> toList() {
    return agents.map((agent) => agent.toMap()).toList();
  }

  String toString() => json.encode(toList());

  int length() => agents.length;
}