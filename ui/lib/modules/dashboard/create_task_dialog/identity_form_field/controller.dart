import 'package:get/get.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import '../../../../models/identity.dart';
import '../../../../common/request.dart';

class IdentityFormFieldController extends GetxController {
  Rx<Identities> identities = Identities.fromList([]).obs;

  final String query = '''
    query {
      identities(type: Qingstor) {
        name
        type
        credential {
          protocol
          args
        }
        endpoint {
          protocol
          host
          port
        }
      }
    }
  ''';

  Future<QueryResult> getIdentities() {
    return queryGraphQL(QueryOptions(document: gql(query))).then((result) {
      if (result.data != null) {
        identities(Identities.fromList(result.data["identities"] ?? []));
      }

      return result;
    }).catchError((error) {
      return error;
    });
  }
}
