import 'package:get/get.dart';
import 'package:flutter/widgets.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import '../../../common/request.dart';

class CreateIdentityController extends GetxController {
  RxString type = 'Qingstor'.obs;
  RxString name = ''.obs;
  RxString credentialProtocol = 'hamc'.obs;
  RxString credentialAccessKey = ''.obs;
  RxString credentialSecretKey = ''.obs;
  RxString endpointProtocol = ''.obs;
  RxString endpointHost = ''.obs;
  RxString endpointPort = ''.obs;

  final autoValidateMode = AutovalidateMode.disabled.obs;

  void closeDialog() {
    Get.back();
    Get.delete<CreateIdentityController>();
  }

  void onSubmit(getIdentities) {
    createIdentity()
        .then((value) => getIdentities())
        .then((value) => closeDialog());
  }

  String get credential {
    return '''
      {
        protocol: "hamc",
        args: [
          "$credentialAccessKey",
          "$credentialSecretKey",
        ],
      }
    ''';
  }

  String get endpoint {
    return '''
      {
        protocol: "$endpointProtocol",
        host: "$endpointHost",
        port: $endpointPort,
      }
    ''';
  }

  String get mutation {
    return '''
      mutation {
        createIdentity(input: {
          name: "$name",
          type: $type,
          credential: $credential,
          endpoint: $endpoint,
        }) { name }
      }
    ''';
  }

  Future<QueryResult> createIdentity() {
    return queryGraphQL(QueryOptions(document: gql(mutation))).then((result) {
      return result;
    });
  }
}
