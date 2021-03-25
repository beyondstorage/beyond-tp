import 'package:get/get.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import '../../models/task.dart';
import '../../common/request.dart';

class DashboardController extends GetxController {
  RxBool loading = false.obs;
  Rx<Tasks> tasks = Tasks.fromList([]).obs;

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

    queryGraphQL(
      QueryOptions(document: gql(query))
    ).then((result) {
      loading(false);

      if (result.data != null) {
        tasks(Tasks.fromList(result?.data["tasks"] ?? []));
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

    return queryGraphQL(
      QueryOptions(document: gql(_query))
    ).then((result) {
      getTasks();

      return result;
    });
  }

  void onPressedNew() {
    print("click new button");
  }
}