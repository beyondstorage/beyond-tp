import 'package:get/get.dart';
import 'package:flutter/widgets.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import '../../../common/request.dart';

class CreateTaskController extends GetxController {
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

  Future<QueryResult> createTask() {
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
          staffs: {
            id: "",
          },
        }) { id }
      }
    ''';

    return queryGraphQL(QueryOptions(document: gql(_query))).then((result) {
      return result;
    });
  }
}
