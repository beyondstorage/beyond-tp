import 'package:get/get.dart';
import 'package:flutter/widgets.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import '../../../common/request.dart';
import '../../../models/task.dart';

class CreateTaskController extends GetxController {
  RxInt step = 1.obs;
  RxBool isEditingName = false.obs;
  RxBool isCreatingIdentity = false.obs;

  RxString name = ''.obs;
  RxString srcType = ''.obs;
  RxString srcPath = '/'.obs;
  RxString srcBucketName = ''.obs;
  RxString srcCredential = ''.obs;
  RxString srcEndpoint = 'https:qingstor.com'.obs;

  RxString dstType = ''.obs;
  RxString dstPath = '/'.obs;
  RxString dstBucketName = ''.obs;
  RxString dstCredential = ''.obs;
  RxString dstEndpoint = 'https:qingstor.com'.obs;

  final autoValidateMode = AutovalidateMode.disabled.obs;

  void closeDialog() {
    Get.back();
    Get.delete<CreateTaskController>();
  }

  void onSubmit(getTasks) {
    createTask().then((value) => getTasks()).then((value) => closeDialog());
  }

  String get src {
    if (srcType.value == 'Fs') {
      return '''
        {
          type: $srcType,
          options: [{
            key: "work_dir",
            value: "$srcPath",
          }],
        }
      ''';
    }

    return '''
      {
        type: $srcType,
        options: [
          {
            key: "work_dir",
            value: "$srcPath",
          },
          {
            key: "credential",
            value: "$srcCredential",
          },
          {
            key: "endpoint",
            value: "$srcEndpoint",
          },
          {
            key: "bucket_name",
            value: "$srcBucketName",
          },
        ],
      }
    ''';
  }

  String get dst {
    if (dstType.value == 'Fs') {
      return '''
        {
          type: $dstType,
          options: [{
            key: "work_dir",
            value: "$dstPath",
          }],
        }
      ''';
    }

    return '''
      {
        type: $dstType,
        options: [
          {
            key: "work_dir",
            value: "$dstPath",
          },
          {
            key: "credential",
            value: "$dstCredential",
          },
          {
            key: "endpoint",
            value: "$dstEndpoint",
          },
          {
            key: "bucket_name",
            value: "$dstBucketName",
          },
        ],
      }
    ''';
  }

  final String staffQuery = '''
    query {
      staffs {
        id
      }
    }
  ''';

  Future<Staffs> getStaffs() {
    return queryGraphQL(QueryOptions(document: gql(staffQuery))).then((result) {
      if (result.data != null) {
        return Staffs.fromList(result.data["staffs"] ?? []);
      } else {
        return Staffs.fromList([]);
      }
    });
  }

  Future<QueryResult> createTask() {
    return getStaffs().then((staffs) {
      String _query = '''
      mutation {
        createTask(input: {
          name: "$name",
          type: copyDir,
          storages: [$src, $dst],
          options: {
            key: "recursive",
            value: "true",
          },
          staffs: ${staffs.toList()},
        }) { id }
      }
    ''';

      return queryGraphQL(QueryOptions(document: gql(staffQuery)))
          .then((result) {
        return queryGraphQL(QueryOptions(document: gql(_query))).then((result) {
          return result;
        });
      });
    });
  }
}
