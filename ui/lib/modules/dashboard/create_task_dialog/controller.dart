import 'package:get/get.dart';
import 'package:flutter/widgets.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import '../../../common/request.dart';

class CreateTaskController extends GetxController {
  RxString name = ''.obs;
  RxString srcType = ''.obs;
  RxString srcPath = ''.obs;
  RxString srcWorkDir = '/'.obs;
  RxString srcBucketName = ''.obs;
  RxString srcCredential = ''.obs;

  RxString dstType = ''.obs;
  RxString dstPath = ''.obs;
  RxString dstWorkDir = '/'.obs;
  RxString dstBucketName = ''.obs;
  RxString dstCredential = ''.obs;

  final autoValidateMode = AutovalidateMode.disabled.obs;

  void closeDialog() {
    Get.back();
    Get.delete<CreateTaskController>();
  }

  void onSubmit(getTasks) {
    createTask().then((value) => getTasks()).then((value) => closeDialog());
  }

  String get src {
    if (srcType.value == 'fs') {
      return '''
        {
          type: $srcType,
          path: "$srcPath",
          options: {
            work_dir: "$srcWorkDir",
          }
        }
      ''';
    }

    return '''
      {
        type: $srcType,
        path: "$srcPath",
        options: {
          work_dir: "$srcWorkDir",
          bucket_name: "$srcBucketName",
          credential: "$srcCredential",
        }
      }
    ''';
  }

  String get dst {
    if (dstType.value == 'fs') {
      return '''
        {
          type: $dstType,
          path: "$dstPath",
          options: {
            work_dir: "$dstWorkDir",
          }
        }
      ''';
    }

    return '''
      {
        type: $dstType,
        path: "$dstPath",
        options: {
          work_dir: "$dstWorkDir",
          bucket_name: "$dstBucketName",
          credential: "$dstCredential",
        }
      }
    ''';
  }

  Future<QueryResult> createTask() {
    String _query = '''
      mutation {
        createTask(input: {
          name: "$name",
          type: copyDir,
          src: $src,
          dst: $dst,
          options: {
            recursive: true
          }
        }) { id }
      }
    ''';

    return queryGraphQL(QueryOptions(document: gql(_query))).then((result) {
      return result;
    });
  }
}
