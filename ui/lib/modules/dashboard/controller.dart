import 'package:get/get.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import '../../models/task.dart';
import '../../common/request.dart';

class DashboardController extends GetxController {
  RxBool loading = false.obs;
  Rx<Tasks> tasks = Tasks.fromList([]).obs;
  Rx<TaskDetail> taskDetail = TaskDetail.fromMap({}).obs;
  RxString filters = ''.obs;
  RxBool showDetail = false.obs;
  RxString detailTaskId = ''.obs;

  final String query = r'''
    query {
      tasks {
        id
        name
        status
        created_at
        updated_at
      }
    }
  ''';

  void getTasks() {
    loading(true);

    queryGraphQL(QueryOptions(document: gql(query))).then((result) {
      loading(false);

      if (result.data != null) {
        tasks(Tasks.fromList(result.data["tasks"] ?? []));
      }
    }).catchError((error) {
      loading(false);
    });
  }

  Future<QueryResult> deleteTask(String id) {
    String _query = '''
      mutation {
        deleteTask(input: { id: "$id" }) { id }
      }
    ''';

    return queryGraphQL(QueryOptions(document: gql(_query))).then((result) {
      getTasks();

      return result;
    });
  }

  Future<QueryResult> runTask(String id) {
    String _query = '''
      mutation {
        runTask(id: "$id") { id }
      }
    ''';

    return queryGraphQL(QueryOptions(document: gql(_query))).then((result) {
      getTasks();

      return result;
    });
  }

  Future<QueryResult> getTaskDetail(String id) {
    String _query = '''
      query {
        task(id: "$id") {
          id
          name
          storages {
            type
            options {
              key
              value
            }
          }
        }
      }
    ''';

    return queryGraphQL(QueryOptions(document: gql(_query))).then((result) {
      if (result.data != null) {
        taskDetail(TaskDetail.fromMap(result.data["task"] ?? {}));
      }

      return result;
    });
  }
}
