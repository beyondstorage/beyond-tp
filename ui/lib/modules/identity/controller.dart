import 'package:get/get.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import '../../models/identity.dart';
import '../../common/request.dart';

class IdentityController extends GetxController {
  RxBool loading = false.obs;
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

  void getIdentities() {
    loading(true);

    queryGraphQL(QueryOptions(document: gql(query))).then((result) {
      loading(false);

      if (result.data != null) {
        identities(Identities.fromList(result.data["identities"] ?? []));
      }
    }).catchError((error) {
      loading(false);
    });
  }

  Future<QueryResult> deleteIdentity(Identity identity) {
    String _query = '''
      mutation {
        deleteIdentity(input: { name: "${identity.name}", type: ${identity.type} }) { }
      }
    ''';

    return queryGraphQL(QueryOptions(document: gql(_query))).then((result) {
      getIdentities();

      return result;
    });
  }
}
