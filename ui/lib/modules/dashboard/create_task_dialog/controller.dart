import 'package:get/get.dart';
import 'package:flutter/widgets.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import '../../../common/request.dart';

class CreateTaskController extends GetxController {
  RxString name = ''.obs;
  RxString originLibraryType = ''.obs;
  RxString originLibraryPath = ''.obs;
  RxString targetLibraryType = ''.obs;
  RxString targetLibraryPath = ''.obs;
  RxString bucketName = ''.obs;
  RxString credential = ''.obs;
  RxString location = ''.obs;

  final autoValidateMode = AutovalidateMode.disabled.obs;

  void onSubmit(getTasks) {
    createTask().then((value) => getTasks());
  }

  Future<QueryResult> createTask() {
    String _query = '''
      mutation {
        createTask(input: {
          name: "$name",
          type: copyDir,
          src: {
            type: $originLibraryType,
            path: "$targetLibraryPath",
          },
          dst: {
            type: $targetLibraryType,
            path: "$targetLibraryPath",
            options: {
              name: "$bucketName",
              credential: "$credential",
              location: "$location",

            }
          },
          options: {
            recursive: true,
          },
        }) { id }
      }
    ''';

    return queryGraphQL(QueryOptions(document: gql(_query))).then((result) {
      Get.back();

      return result;
    });
  }
}
