import 'dart:convert';

class Task {
  String id;
  String name;
  String status;
  String createdAt;
  String updatedAt;

  Task({
    this.id,
    this.name,
    this.status,
    this.createdAt,
    this.updatedAt,
  });

  factory Task.fromMap(Map<String, dynamic> json) => Task(
        id: json["id"],
        name: json["name"] ?? "",
        status: json["status"],
        createdAt: json["created_at"] ?? "",
        updatedAt: json["updated_at"] ?? "",
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

  Tasks({this.tasks});

  factory Tasks.fromList(List<Object> tasks) =>
      Tasks(tasks: List<Task>.from(tasks.map((x) => Task.fromMap(x))));

  List<Map<String, dynamic>> toList() {
    return tasks.map((task) => task.toMap()).toList();
  }

  String toString() => json.encode(toList());
}

enum StorageType {
  Fs,
  Qingstor,
}

class StorageOption {
  String key;
  String value;

  StorageOption({
    this.key,
    this.value,
  });

  factory StorageOption.fromMap(Map<String, dynamic> json) => StorageOption(
        key: json["key"],
        value: json["value"],
      );
}

class Storage {
  StorageType type;
  List<StorageOption> options;

  Storage({
    this.type,
    this.options,
  });

  factory Storage.fromMap(Map<String, dynamic> json) => Storage(
        type: json["type"],
        options: List<StorageOption>.from((json["options"] ?? [])
            .map((option) => StorageOption.fromMap(option))),
      );
}

class TaskDetail {
  String id;
  String name;
  List<Storage> storages;

  TaskDetail({
    this.id,
    this.name,
    this.storages,
  });

  factory TaskDetail.fromMap(Map<String, dynamic> json) => TaskDetail(
        id: json["id"],
        name: json["name"],
        storages: List<Storage>.from((json["storages"] ?? [])
            .map((storage) => Storage.fromMap(storage))),
      );
}
