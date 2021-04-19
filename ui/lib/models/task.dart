import 'dart:convert';

class Task {
  String id;
  String name;
  String status;
  String createdAt;
  String updatedAt;

  Task({
    required this.id,
    required this.name,
    required this.status,
    required this.createdAt,
    required this.updatedAt,
  });

  factory Task.fromMap(Map<String, dynamic> json) => Task(
    id: json["id"],
    name: json["name"] ?? "",
    status: json["status"],
    createdAt: json["created_at"] ?? "",
    updatedAt: json["updated_at"] ?? ""
  );

  Map<String, dynamic> toMap() => {
    "id": id,
    "name": name,
    "status": status,
    "createdAt": createdAt,
    "updatedAt": updatedAt,
  };

  String toString() => json.encode(toMap());
}

class Tasks {
  List<Task> tasks;

  Tasks({ required this.tasks });

  factory Tasks.fromList(List<Object> tasks) => Tasks(
    tasks: List<Task>.from(tasks.map((x) => Task.fromMap(x as Map<String, dynamic>)))
  );

  List<Map<String, dynamic>> toList() {
    return tasks.map((task) => task.toMap()).toList();
  }

  String toString() => json.encode(toList());
}
